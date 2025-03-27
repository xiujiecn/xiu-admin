// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"errors"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/event"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/do"
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/model/request"
	"xiuadmin/internal/service"
	"xiuadmin/utility"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
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

func (l *sSysUser) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.SysUser.Ctx(ctx), option...)
}

// 根据用户名获取用户信息，不验证当前租户
func (l *sSysUser) GetUserByUsername(ctx context.Context, username string, tenantId string) (user *entity.SysUser, err error) {
	var data *entity.SysUser
	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserName, username).Where(dao.SysUser.Columns().TenantId, tenantId).Scan(&data)
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

// 根据邮箱获取用户信息，不验证当前租户
func (l *sSysUser) GetUserByEmail(ctx context.Context, email string, tenantId string) (user *entity.SysUser, err error) {
	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Email, email).Where(dao.SysUser.Columns().TenantId, tenantId).Scan(&user)
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

// 根据手机号获取用户信息,不验证当前租户
func (l *sSysUser) GetUserByPhone(ctx context.Context, phone string, tenantId string) (user *entity.SysUser, err error) {
	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Phonenumber, phone).Where(dao.SysUser.Columns().TenantId, tenantId).Scan(&user)
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
func (l *sSysUser) GetUserByUsernameAndPassword(ctx context.Context, tenantId string, username string, password string) (user *entity.SysUser, err error) {
	// 判断用户名是否存在
	user, err = l.GetUserByUsername(ctx, username, tenantId)
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
	err = l.Model(ctx).WithAll().Where(dao.SysUser.Columns().UserId, id).Scan(&user)
	if err != nil {
		return nil, err
	}
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
		mini := &model.SysRoleMiniModel{}
		gconv.Struct(role, mini)
		user.Roles = append(user.Roles, mini)
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
		mini := &model.SysPostMiniModel{}
		gconv.Struct(post, mini)
		user.Posts = append(user.Posts, mini)
	}
	return
}

// 获取用户列表
func (l *sSysUser) List(ctx context.Context, page *request.PageInfo, query *model.UserListParam) (items []*model.SysUserListModel, total int, err error) {
	deptIds := make([]int64, 0)
	if query.DeptId != 0 {
		deptIds = append(deptIds, query.DeptId)
		subDeptIds, err := service.SysDept().GetDeptIdsByParentId(ctx, query.DeptId)
		if err != nil {
			return nil, 0, err
		}
		deptIds = append(deptIds, subDeptIds...)
	}
	g.Log().Infof(ctx, "sSysUser.List 获取部门列表: deptIds:%+v", deptIds)
	// 实现分页查询逻辑
	var users []*model.SysUserListModel
	m := l.Model(ctx)
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
	if len(query.CreatedAt) == 2 {
		begin, end := gtime.NewFromStr(query.CreatedAt[0]), gtime.NewFromStr(query.CreatedAt[1])
		m = m.WhereBetween(dao.SysUser.Columns().CreatedAt, begin, end.EndOfDay())
	}
	total, err = m.Count()
	if err != nil {
		return nil, 0, err
	}
	err = m.WithAll().Page(page.Page, page.PageSize).Scan(&users)
	if err != nil {
		return nil, 0, err
	}
	// g.Log().Infof(ctx, "获取用户列表成功: total:%d, users:%+v", total, users)
	return users, total, nil
}

// 新增用户
func (l *sSysUser) AddUser(ctx context.Context, req *model.SysUserAddModel) (data *model.SysUserViewModel, err error) {
	// 获取当前登录用户的租户
	if req.TenantId == "" {
		req.TenantId = contexts.GetTenantId(ctx)
	}
	// 判断部门是否存在
	dept, err := service.SysDept().GetDeptById(ctx, req.DeptId)
	if err != nil {
		return nil, err
	}
	if dept == nil || dept.TenantId != req.TenantId {
		return nil, gerror.NewCode(consts.CodeDeptNotFound, "部门不存在")
	}

	// 判断用户名是否存在
	var user *entity.SysUser
	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserName, req.UserName).Unscoped().Scan(&user)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
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

	dataInsert := do.SysUser{}
	gconv.Struct(req, &dataInsert)
	dataInsert.Password = password
	dataInsert.Salt = salt
	dataInsert.CreatedDept = contexts.GetDeptId(ctx)
	dataInsert.CreatedBy = contexts.GetUserId(ctx)
	dataInsert.CreatedAt = gtime.Now()
	dataInsert.UpdatedBy = contexts.GetUserId(ctx)
	dataInsert.UpdatedAt = gtime.Now()
	dataInsert.TenantId = req.TenantId
	result, err := dao.SysUser.Ctx(ctx).Data(dataInsert).Insert()
	if err != nil {
		return nil, err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	data, err = l.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	// 新增用户角色
	d := make([]map[string]any, 0)
	if len(req.RoleIds) > 0 {
		for _, roleId := range req.RoleIds {
			d = append(d, map[string]any{dao.SysUserRole.Columns().UserId: userId, dao.SysUserRole.Columns().RoleId: roleId})
		}
		_, err = dao.SysUserRole.Ctx(ctx).Data(d).Insert()
		if err != nil {
			return nil, err
		}
	}
	// 新增用户岗位
	if len(req.PostIds) > 0 {
		d = make([]map[string]any, 0)
		for _, postId := range req.PostIds {
			d = append(d, map[string]any{dao.SysUserPost.Columns().UserId: userId, dao.SysUserPost.Columns().PostId: postId})
		}
		_, err = dao.SysUserPost.Ctx(ctx).Data(d).Insert()
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func (l *sSysUser) Profile(ctx context.Context) (user *model.UserProfileModel, err error) {

	u, err := l.GetUserById(ctx, contexts.GetUserId(ctx))
	if err != nil {
		return nil, err
	}
	return &model.UserProfileModel{
		SysUserViewModel: *u,
	}, nil
}

func (l *sSysUser) UpdateCurrentUser(ctx context.Context, req *model.UpdateCurrentUserModel) (user *model.SysUserViewModel, err error) {
	userId := contexts.GetUserId(ctx)
	_, err = l.Model(ctx).Where(dao.SysUser.Columns().UserId, userId).Data(req).OmitEmpty().Update()
	if err != nil {
		return nil, err
	}
	user, err = l.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyUserUpdate, userId)
	return user, nil
}

func (l *sSysUser) UpdateCurrentUserPassword(ctx context.Context, req *model.UpdateCurrentUserPasswordModel) (err error) {
	userId := contexts.GetUserId(ctx)
	var user *entity.SysUser
	err = l.Model(ctx).Where(dao.SysUser.Columns().UserId, userId).Scan(&user)
	if err != nil {
		return err
	}
	if user == nil {
		return gerror.NewCode(consts.CodeUserNotFound, "用户不存在")
	}

	oldPassword := utility.PasswordEncrypt(req.OldPassword, user.Salt)
	if oldPassword != user.Password {
		return gerror.NewCode(consts.CodeUserPasswordError, "旧密码错误")
	}
	salt := utility.RandomString(5)
	password := utility.PasswordEncrypt(req.NewPassword, salt)
	_, err = l.Model(ctx).Where(dao.SysUser.Columns().UserId, userId).Data(map[string]any{
		dao.SysUser.Columns().Password:  password,
		dao.SysUser.Columns().Salt:      salt,
		dao.SysUser.Columns().UpdatedBy: userId,
		dao.SysUser.Columns().UpdatedAt: gtime.Now(),
	}).OmitEmpty().Update()
	if err != nil {
		return err
	}
	return nil
}

func (l *sSysUser) UpdateUser(ctx context.Context, req *model.SysUserUpdateModel) (err error) {
	currUserId := contexts.GetUserId(ctx)
	req.UpdatedAt = gtime.Now()
	req.UpdatedBy = &currUserId
	_, err = l.Model(ctx).Where(dao.SysUser.Columns().UserId, req.UserId).Data(req).OmitNil().Update()
	if err != nil {
		return err
	}
	// 用户角色
	urList := make([]*entity.SysUserRole, 0)
	err = dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns().UserId, req.UserId).Scan(&urList)
	if err != nil {
		return err
	}
	// 计算删除和新增的角色
	delRoleIds := make([]int64, 0)
	addRoleIds := make([]int64, 0)
	hasRoleId := false
	for _, ur := range urList {
		hasRoleId = false
		for _, roleId := range req.RoleIds {
			if ur.RoleId == roleId {
				hasRoleId = true
				break
			}
		}
		if !hasRoleId {
			delRoleIds = append(delRoleIds, ur.RoleId)
		}
	}
	for _, roleId := range req.RoleIds {
		hasRoleId = false
		for _, ur := range urList {
			if ur.RoleId == roleId {
				hasRoleId = true
				break
			}
		}
		if !hasRoleId {
			addRoleIds = append(addRoleIds, roleId)
		}
	}
	// 删除多余的角色
	if len(delRoleIds) > 0 {
		_, err = dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns().UserId, req.UserId).WhereIn(dao.SysUserRole.Columns().RoleId, delRoleIds).Delete()
		if err != nil {
			return err
		}
	}
	// 新增角色
	d := make([]map[string]any, 0)
	if len(addRoleIds) > 0 {
		for _, roleId := range addRoleIds {
			d = append(d, map[string]any{dao.SysUserRole.Columns().UserId: req.UserId, dao.SysUserRole.Columns().RoleId: roleId})
		}
		_, err = dao.SysUserRole.Ctx(ctx).Data(d).Insert()
		if err != nil {
			return err
		}
	}
	// 查询用户岗位
	upList := make([]*entity.SysUserPost, 0)
	err = dao.SysUserPost.Ctx(ctx).Where(dao.SysUserPost.Columns().UserId, req.UserId).Scan(&upList)
	if err != nil {
		return err
	}
	// 计算删除和新增的岗位
	delPostIds := make([]int64, 0)
	addPostIds := make([]int64, 0)
	hasPostId := false
	for _, up := range upList {
		hasPostId = false
		for _, postId := range req.PostIds {
			if up.PostId == postId {
				hasPostId = true
				break
			}
		}
		if !hasPostId {
			delPostIds = append(delPostIds, up.PostId)
		}
	}
	for _, postId := range req.PostIds {
		hasPostId = false
		for _, up := range upList {
			if up.PostId == postId {
				hasPostId = true
				break
			}
		}
		if !hasPostId {
			addPostIds = append(addPostIds, postId)
		}
	}
	// 删除多余岗位
	if len(delPostIds) > 0 {
		_, err = dao.SysUserPost.Ctx(ctx).Where(dao.SysUserPost.Columns().UserId, req.UserId).WhereIn(dao.SysUserPost.Columns().PostId, delPostIds).Delete()
		if err != nil {
			return err
		}
	}
	// 新增岗位
	d = make([]map[string]any, 0)
	if len(addPostIds) > 0 {
		for _, postId := range addPostIds {
			d = append(d, map[string]any{dao.SysUserPost.Columns().UserId: req.UserId, dao.SysUserPost.Columns().PostId: postId})
		}
		_, err = dao.SysUserPost.Ctx(ctx).Data(d).Insert()
		if err != nil {
			return err
		}
	}

	event.EventsInstance().Emit(ctx, consts.EventKeyUserUpdate, req.UserId)
	return nil
}

func (l *sSysUser) DeleteUser(ctx context.Context, userIds []int64) (err error) {
	if len(userIds) == 0 {
		return errors.New("用户ID列表不能为空")
	}
	currUserId := contexts.GetUserId(ctx)
	_, err = l.Model(ctx).WhereIn(dao.SysUser.Columns().UserId, userIds).Data(map[string]any{
		dao.SysUser.Columns().DeletedAt: gtime.Now(),
		dao.SysUser.Columns().DeletedBy: currUserId,
	}).OmitEmpty().Update()
	if err != nil {
		return err
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyUserDelete, userIds)
	return nil
}

func (l *sSysUser) ResetPassword(ctx context.Context, userId int64, password string) (err error) {
	currUserId := contexts.GetUserId(ctx)
	salt := utility.RandomString(5)
	password = utility.PasswordEncrypt(password, salt)
	_, err = l.Model(ctx).Where(dao.SysUser.Columns().UserId, userId).Data(map[string]any{
		dao.SysUser.Columns().Password:  password,
		dao.SysUser.Columns().Salt:      salt,
		dao.SysUser.Columns().UpdatedBy: currUserId,
		dao.SysUser.Columns().UpdatedAt: gtime.Now(),
	}).OmitEmpty().Update()
	if err != nil {
		return err
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyUserUpdate, userId)
	return nil
}

// 获取用户角色ID列表
func (l *sSysUser) GetUserRoleIds(ctx context.Context, userId int64) (roleIds []int64, err error) {
	urList := make([]*entity.SysUserRole, 0)
	err = dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns().UserId, userId).Scan(&urList)
	if err != nil {
		return nil, err
	}
	for _, ur := range urList {
		roleIds = append(roleIds, ur.RoleId)
	}
	return roleIds, nil
}

// 获取用户岗位ID列表
func (l *sSysUser) GetUserPostIds(ctx context.Context, userId int64) (postIds []int64, err error) {
	upList := make([]*entity.SysUserPost, 0)
	err = dao.SysUserPost.Ctx(ctx).Where(dao.SysUserPost.Columns().UserId, userId).Scan(&upList)
	if err != nil {
		return nil, err
	}
	for _, up := range upList {
		postIds = append(postIds, up.PostId)
	}
	return postIds, nil
}
