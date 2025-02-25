package system

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/entity"
	"server/internal/model/request"
	"server/internal/service"
	"server/utility"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"gorm.io/gorm"
)

type sSysUser struct {
}

func NewSysUser() *sSysUser {
	return &sSysUser{}
}

func init() {
	service.RegisterSysUser(NewSysUser())
}

// 根据用户名获取用户信息
func (l *sSysUser) GetUserByUsername(ctx context.Context, username string) (user *entity.SysUser, err error) {
	var data *entity.SysUser
	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserName, username).Scan(&data)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gerror.NewCode(consts.CodeUserNotFound, "账号不存在")
		}
		return nil, err
	}
	if data == nil {
		return nil, gerror.NewCode(consts.CodeUserNotFound, "账号不存在")
	}
	if data.Status == consts.SysUserStatusDisable {
		return nil, gerror.NewCode(consts.CodeUserDisabled, "账号已禁用")
	}
	if data.Status == consts.SysUserStatusLocked {
		return nil, gerror.NewCode(consts.CodeUserLocked, "账号已锁定")
	}
	if data.Status == consts.SysUserStatusExpired {
		return nil, gerror.NewCode(consts.CodeUserExpired, "账号已过期")
	}
	if data.Status == consts.SysUserStatusDeleted {
		return nil, gerror.NewCode(consts.CodeUserDeleted, "账号已删除")
	}
	return data, nil
}

// 根据邮箱获取用户信息
func (l *sSysUser) GetUserByEmail(ctx context.Context, email string) (user *entity.SysUser, err error) {
	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Email, email).Scan(&user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gerror.NewCode(consts.CodeUserNotFound, "账号不存在")
		}
		return nil, err
	}
	if user == nil {
		return nil, gerror.NewCode(consts.CodeUserNotFound, "账号不存在")
	}
	if user.Status == consts.SysUserStatusDisable {
		return nil, gerror.NewCode(consts.CodeUserDisabled, "账号已禁用")
	}
	if user.Status == consts.SysUserStatusLocked {
		return nil, gerror.NewCode(consts.CodeUserLocked, "账号已锁定")
	}
	if user.Status == consts.SysUserStatusExpired {
		return nil, gerror.NewCode(consts.CodeUserExpired, "账号已过期")
	}
	if user.Status == consts.SysUserStatusDeleted {
		return nil, gerror.NewCode(consts.CodeUserDeleted, "账号已删除")
	}
	return user, nil
}

// 根据手机号获取用户信息
func (l *sSysUser) GetUserByPhone(ctx context.Context, phone string) (user *entity.SysUser, err error) {
	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Phonenumber, phone).Scan(&user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gerror.NewCode(consts.CodeUserNotFound, "账号不存在")
		}
		return nil, err
	}
	if user == nil {
		return nil, gerror.NewCode(consts.CodeUserNotFound, "账号不存在")
	}
	if user.Status == consts.SysUserStatusDisable {
		return nil, gerror.NewCode(consts.CodeUserDisabled, "账号已禁用")
	}
	if user.Status == consts.SysUserStatusLocked {
		return nil, gerror.NewCode(consts.CodeUserLocked, "账号已锁定")
	}
	if user.Status == consts.SysUserStatusExpired {
		return nil, gerror.NewCode(consts.CodeUserExpired, "账号已过期")
	}

	return
}

// 根据用户名和密码获取用户信息
func (l *sSysUser) GetUserByUsernameAndPassword(ctx context.Context, username string, password string) (user *entity.SysUser, err error) {
	// 判断用户名是否存在
	user, err = l.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	// 判断密码是否正确
	password = utility.PasswordEncrypt(password, user.Salt)
	if password != user.Password {
		g.Log().Errorf(ctx, "密码错误: username:%s, password:%s, user:%+v", username, password, user)
		return nil, gerror.NewCode(consts.CodeUserPasswordError, "密码错误")
	}
	return user, nil
}

// 根据用户ID获取用户信息
func (l *sSysUser) GetUserById(ctx context.Context, id int64) (user *model.SysUserViewModel, err error) {
	err = dao.SysUser.Ctx(ctx).WithAll().Where(dao.SysUser.Columns().UserId, id).Scan(&user)
	// 获取用户角色
	// 获取用户角色
	userRoles := make([]*entity.SysUserRole, 0)
	err = dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns().UserId, id).Scan(&userRoles)
	if err != nil {
		return nil, err
	}
	roleIds := make([]int64, 0)
	for _, userRole := range userRoles {
		roleIds = append(roleIds, userRole.RoleId)
	}
	roles := make([]*entity.SysRole, 0)
	err = dao.SysRole.Ctx(ctx).WhereIn(dao.SysRole.Columns().RoleId, roleIds).Scan(&roles)
	if err != nil {
		return nil, err
	}
	user.Roles = make([]*model.SysRoleMiniModel, 0)
	for _, role := range roles {
		user.Roles = append(user.Roles, &model.SysRoleMiniModel{
			RoleId:   role.RoleId,
			RoleName: role.RoleName,
		})
	}
	// 获取用户岗位
	userPosts := make([]*entity.SysUserPost, 0)
	err = dao.SysUserPost.Ctx(ctx).Where(dao.SysUserPost.Columns().UserId, id).Scan(&userPosts)
	if err != nil {
		return nil, err
	}
	postIds := make([]int64, 0)
	for _, userPost := range userPosts {
		postIds = append(postIds, userPost.PostId)
	}
	posts := make([]*entity.SysPost, 0)
	err = dao.SysPost.Ctx(ctx).WhereIn(dao.SysPost.Columns().PostId, postIds).Scan(&posts)
	if err != nil {
		return nil, err
	}
	user.Posts = make([]*model.SysPostMiniModel, 0)
	for _, post := range posts {
		user.Posts = append(user.Posts, &model.SysPostMiniModel{
			PostId:   post.PostId,
			PostCode: post.PostCode,
			PostName: post.PostName,
			DeptId:   post.DeptId,
		})
	}

	return
}

// 获取用户列表
func (l *sSysUser) GetUserList(ctx context.Context, page request.PageInfo, query model.UserListQuery) (items []*model.SysUserListModel, total int, err error) {
	tenantId := query.TenantId
	if tenantId == "" {
		claims, err := service.SysAuth().GetCurrentUser(ctx)
		if err != nil {
			return nil, 0, err
		}
		tenantId = claims.BaseClaims.TenantId
	}
	deptIds := make([]int64, 0)
	if query.DeptId != 0 {
		deptIds = append(deptIds, query.DeptId)
		subDeptIds, err := service.SysDept().GetDeptIdsByParentId(ctx, query.DeptId)
		if err != nil {
			return nil, 0, err
		}
		deptIds = append(deptIds, subDeptIds...)
	}
	g.Log().Infof(ctx, "sSysUser.GetUserList 获取机构列表: deptIds:%+v", deptIds)
	// 实现分页查询逻辑
	var users []*model.SysUserListModel
	m := dao.SysUser.Ctx(ctx).Page(page.Page, page.PageSize)
	if query.TenantId != "" {
		m = m.Where(dao.SysUser.Columns().TenantId, query.TenantId)
	}
	if query.UserId != 0 {
		m = m.Where(dao.SysUser.Columns().UserId, query.UserId)
	}
	if len(deptIds) > 0 {
		m = m.WhereIn(dao.SysUser.Columns().DeptId, deptIds)
	}
	if query.UserName != "" {
		m = m.WhereLike(dao.SysUser.Columns().UserName, "%"+query.UserName+"%")
	}
	if query.NickName != "" {
		m = m.WhereLike(dao.SysUser.Columns().NickName, "%"+query.NickName+"%")
	}
	if query.Email != "" {
		m = m.WhereLike(dao.SysUser.Columns().Email, "%"+query.Email+"%")
	}
	if query.Phonenumber != "" {
		m = m.WhereLike(dao.SysUser.Columns().Phonenumber, "%"+query.Phonenumber+"%")
	}
	if query.Status != "" {
		m = m.Where(dao.SysUser.Columns().Status, query.Status)
	}
	total, err = m.Count()
	if err != nil {
		return nil, 0, err
	}
	err = m.WithAll().Scan(&users)
	if err != nil {
		return nil, 0, err
	}
	g.Log().Infof(ctx, "获取用户列表成功: total:%d, users:%+v", total, users)
	return users, total, nil
}

// 新增用户
func (l *sSysUser) AddUser(ctx context.Context, req model.AddUser) (user *entity.SysUser, err error) {
	// 获取当前登录用户的租户
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	req.TenantId = claims.BaseClaims.TenantId
	// 判断部门是否存在
	dept, err := service.SysDept().GetDeptById(ctx, req.DeptId)
	if err != nil {
		return nil, err
	}
	if dept == nil || dept.TenantId != claims.BaseClaims.TenantId {
		return nil, gerror.NewCode(consts.CodeDeptNotFound, "部门不存在")
	}

	// 判断用户名是否存在
	user, err = l.GetUserByUsername(ctx, req.UserName)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, gerror.NewCode(consts.CodeUserExists, "用户名已存在")
	}
	// // 判断邮箱是否存在
	// user, err = l.GetUserByEmail(ctx, req.Email)
	// if err != nil {
	// 	return nil, err
	// }
	// if user != nil {
	// 	return nil, gerror.NewCode(consts.CodeEmailExists, "邮箱已存在")
	// }
	// // 判断手机号是否存在
	// user, err = l.GetUserByPhone(ctx, req.Phone)
	// if err != nil {
	// 	return nil, err
	// }
	// if user != nil {
	// 	return nil, gerror.NewCode(consts.CodePhoneExists, "手机号已存在")
	// }
	// 随机生成5位字符
	salt := utility.RandomString(5)
	password := utility.PasswordEncrypt(req.Password, salt)
	// 插入用户
	user = &entity.SysUser{
		UserName:    req.UserName,
		TenantId:    req.TenantId,
		DeptId:      req.DeptId,
		Password:    password,
		Email:       req.Email,
		Phonenumber: req.Phone,
		Status:      consts.SysUserStatusNormal,
		CreatedAt:   gtime.Now(),
		UpdatedAt:   gtime.Now(),
	}
	_, err = dao.SysUser.Ctx(ctx).Data(user).OmitEmpty().Insert()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (l *sSysUser) Profile(ctx context.Context) (user *model.UserProfileModel, err error) {
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	u, err := l.GetUserById(ctx, claims.BaseClaims.ID)
	if err != nil {
		return nil, err
	}
	return &model.UserProfileModel{
		SysUserViewModel: *u,
	}, nil
}

func (l *sSysUser) UpdateCurrentUser(ctx context.Context, req *model.UpdateCurrentUserModel) (user *model.SysUserViewModel, err error) {
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	_, err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserId, claims.BaseClaims.ID).Data(req).OmitEmpty().Update()
	if err != nil {
		return nil, err
	}
	user, err = l.GetUserById(ctx, claims.BaseClaims.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (l *sSysUser) UpdateCurrentUserPassword(ctx context.Context, req *model.UpdateCurrentUserPasswordModel) (err error) {
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return err
	}
	user, err := l.GetUserById(ctx, claims.BaseClaims.ID)
	if err != nil {
		return err
	}
	oldPassword := utility.PasswordEncrypt(req.OldPassword, user.Salt)
	if oldPassword != user.Password {
		return gerror.NewCode(consts.CodeUserPasswordError, "旧密码错误")
	}
	salt := utility.RandomString(5)
	password := utility.PasswordEncrypt(req.NewPassword, salt)
	_, err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserId, claims.BaseClaims.ID).Data(map[string]any{
		dao.SysUser.Columns().Password: password,
		dao.SysUser.Columns().Salt:     salt,
	}).OmitEmpty().Update()
	if err != nil {
		return err
	}
	return nil
}
