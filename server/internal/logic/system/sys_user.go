// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"
	"regexp"
	"slices"
	"strings"
	"time"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/event"
	"xiuadmin/internal/library/mcache"
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
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"gorm.io/gorm"
)

type sSysUser struct {
}

const (
	registerCodeDailyLimit = 5
	registerCodeTTL        = 6 * time.Hour
)

type registerSmsConfig struct {
	URL           string            `json:"url"`
	Method        string            `json:"method"`
	Headers       map[string]string `json:"headers"`
	BodyTemplate  string            `json:"bodyTemplate"`
	SuccessStatus int               `json:"successStatus"`
}

type registerEmailConfig struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	From         string `json:"from"`
	FromName     string `json:"fromName"`
	Subject      string `json:"subject"`
	BodyTemplate string `json:"bodyTemplate"`
	SSL          bool   `json:"ssl"`
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
	if user.Status == consts.SysUserStatusDeleted {
		return nil, gerror.NewCode(consts.CodeUserDeleted, "账号已删除")
	}

	return
}

// 根据用户名、手机号、邮箱和密码获取用户信息
func (l *sSysUser) GetUserByUsernameAndPassword(ctx context.Context, tenantId string, username string, password string) (user *entity.SysUser, err error) {
	loginName := strings.TrimSpace(username)
	if loginName == "" {
		return nil, gerror.NewCode(consts.CodeUserNotFound, "账号不存在")
	}
	if isEmailContact(loginName) {
		user, err = l.GetUserByEmail(ctx, loginName, tenantId)
	} else if isPhoneContact(loginName) {
		user, err = l.GetUserByPhone(ctx, loginName, tenantId)
	} else {
		user, err = l.GetUserByUsername(ctx, loginName, tenantId)
	}
	if err != nil {
		return nil, err
	}
	// 判断密码是否正确
	password = utility.PasswordEncrypt(password, user.Salt)
	if password != user.Password {
		g.Log().Errorf(ctx, "密码错误: loginName:%s, password:%s, user:%+v", loginName, password, user)
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
	if len(query.DeptIdList) > 0 {
		m = m.WhereIn(dao.SysUser.Columns().DeptId, query.DeptIdList)
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

	// 判断用户名是否存在（在当前租户下）
	var user *entity.SysUser
	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserName, req.UserName).Where(dao.SysUser.Columns().TenantId, req.TenantId).Scan(&user)
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
	companyInfo := service.SysDept().GetParentCompanyInfo(ctx, u.DeptId)
	return &model.UserProfileModel{
		SysUserViewModel: *u,
		CompanyInfo: &model.SysDeptMiniModel{
			DeptId:   companyInfo.DeptId,
			DeptName: companyInfo.DeptName,
			ParentId: companyInfo.ParentId,
		},
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

func (l *sSysUser) validateCurrentUserContact(ctx context.Context, userId int64, tenantId string, contact string) error {
	if isPhoneContact(contact) {
		existsUser, err := l.GetUserByPhone(ctx, contact, tenantId)
		if err != nil {
			return err
		}
		if existsUser != nil && existsUser.UserId != userId {
			return gerror.NewCode(consts.CodePhoneExists, "手机号已存在")
		}
		return nil
	}
	if isEmailContact(contact) {
		existsUser, err := l.GetUserByEmail(ctx, contact, tenantId)
		if err != nil {
			return err
		}
		if existsUser != nil && existsUser.UserId != userId {
			return gerror.NewCode(consts.CodeEmailExists, "邮箱已存在")
		}
		return nil
	}
	return gerror.New("请输入正确的手机号或邮箱")
}

func (l *sSysUser) SendCurrentUserContactCode(ctx context.Context, req *model.UpdateCurrentUserContactCodeModel) (err error) {
	userId := contexts.GetUserId(ctx)
	user := contexts.GetUser(ctx)
	if user == nil || userId == 0 {
		return gerror.NewCode(consts.CodeLoginExpired, "登录已过期")
	}
	req.Contact = strings.TrimSpace(req.Contact)
	if err = l.validateCurrentUserContact(ctx, userId, user.TenantId, req.Contact); err != nil {
		return err
	}
	if err = service.SysCaptcha().VerifyCaptcha(ctx, req.CaptchaID, req.CaptchaValue); err != nil {
		return err
	}
	code := grand.Digits(6)
	if err = mcache.Set(ctx, userContactCodeCacheKey(userId, req.Contact), code, registerCodeTTL); err != nil {
		return err
	}
	if isPhoneContact(req.Contact) {
		err = sendRegisterSmsCode(ctx, req.Contact, code)
	} else {
		err = sendRegisterEmailCode(ctx, req.Contact, code)
	}
	if err != nil {
		_, _ = mcache.Instance().Remove(ctx, userContactCodeCacheKey(userId, req.Contact))
		return err
	}
	return nil
}

func (l *sSysUser) UpdateCurrentUserContact(ctx context.Context, req *model.UpdateCurrentUserContactModel) (err error) {
	userId := contexts.GetUserId(ctx)
	user := contexts.GetUser(ctx)
	if user == nil || userId == 0 {
		return gerror.NewCode(consts.CodeLoginExpired, "登录已过期")
	}
	req.Contact = strings.TrimSpace(req.Contact)
	if err = l.validateCurrentUserContact(ctx, userId, user.TenantId, req.Contact); err != nil {
		return err
	}
	cacheValue, err := mcache.Get(ctx, userContactCodeCacheKey(userId, req.Contact))
	cacheCode := gconv.String(cacheValue)
	if err != nil || cacheCode == "" {
		return gerror.NewCode(consts.CodeCaptchaError, "验证码已失效")
	}
	if cacheCode != req.Code {
		return gerror.NewCode(consts.CodeCaptchaError, "验证码错误")
	}
	data := map[string]any{
		dao.SysUser.Columns().UpdatedBy: userId,
		dao.SysUser.Columns().UpdatedAt: gtime.Now(),
	}
	if isPhoneContact(req.Contact) {
		data[dao.SysUser.Columns().Phonenumber] = req.Contact
	} else {
		data[dao.SysUser.Columns().Email] = req.Contact
	}
	_, err = l.Model(ctx).Where(dao.SysUser.Columns().UserId, userId).Data(data).Update()
	if err != nil {
		return err
	}
	_, _ = mcache.Instance().Remove(ctx, userContactCodeCacheKey(userId, req.Contact))
	event.EventsInstance().Emit(ctx, consts.EventKeyUserUpdate, userId)
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
		g.Log().Warningf(ctx, "sSysUser.DeleteUser 用户ID列表不能为空. userIds: %v", userIds)
		return errors.New("用户ID列表不能为空")
	}
	currUserId := contexts.GetUserId(ctx)
	if slices.Contains(userIds, 1) {
		g.Log().Warningf(ctx, "sSysUser.DeleteUser 不能删除超级管理员. userIds: %v", userIds)
		return errors.New("不能删除超级管理员")
	}
	if slices.Contains(userIds, currUserId) {
		g.Log().Warningf(ctx, "sSysUser.DeleteUser 不能删除当前用户. userIds: %v, currUserId: %d", userIds, currUserId)
		return errors.New("不能删除当前用户")
	}
	isAdminUserIds, err := service.SysTenant().IsTenantAdmin(ctx, userIds)
	if err != nil {
		g.Log().Errorf(ctx, "sSysUser.DeleteUser IsTenantAdmin err: %v, userIds: %v", err, userIds)
		return err
	}
	if len(isAdminUserIds) > 0 {
		g.Log().Warningf(ctx, "sSysUser.DeleteUser 不能删除租户管理员. isAdminUserIds: %v, userIds: %v", isAdminUserIds, userIds)
		return errors.New(fmt.Sprintf("不能删除租户管理员ID: %v", isAdminUserIds))
	}

	_, err = l.Model(ctx).WhereIn(dao.SysUser.Columns().UserId, userIds).Data(map[string]any{
		dao.SysUser.Columns().DeletedAt: gtime.Now(),
		dao.SysUser.Columns().DeletedBy: currUserId,
	}).OmitEmpty().Update()
	if err != nil {
		g.Log().Errorf(ctx, "sSysUser.DeleteUser Update err: %v, userIds: %v", err, userIds)
		return err
	}

	// 删除用户绑定关系
	_, err = dao.SysSocial.Ctx(ctx).WhereIn(dao.SysSocial.Columns().UserId, userIds).Delete()
	if err != nil {
		g.Log().Errorf(ctx, "sSysUser.DeleteUser SysSocial.Ctx(ctx).Delete err: %v, userIds: %v", err, userIds)
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

func (l *sSysUser) UpdateLoginInfo(ctx context.Context, userId int64, loginIp string) (err error) {
	_, err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserId, userId).Data(map[string]any{
		dao.SysUser.Columns().LoginIp:   loginIp,
		dao.SysUser.Columns().LoginDate: gtime.Now(),
	}).Update()
	if err != nil {
		g.Log().Errorf(ctx, "sSysUser.UpdateLoginInfo err: %v, userId: %d, loginIp: %s", err, userId, loginIp)
		return err
	}
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

func normalizeRegisterTenantId(tenantId string) string {
	if tenantId == "" {
		return "000000"
	}
	return tenantId
}

func isPhoneContact(contact string) bool {
	return regexp.MustCompile(`^1[3-9]\d{9}$`).MatchString(contact)
}

func isEmailContact(contact string) bool {
	return regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`).MatchString(contact)
}

func registerCodeCacheKey(tenantId string, contact string) string {
	return fmt.Sprintf("system:register:code:%s:%s", tenantId, strings.ToLower(strings.TrimSpace(contact)))
}

func userContactCodeCacheKey(userId int64, contact string) string {
	return fmt.Sprintf("system:user:contact:code:%d:%s", userId, strings.ToLower(strings.TrimSpace(contact)))
}

func registerCodeDailyCountKey(tenantId string, contact string) string {
	return fmt.Sprintf("system:register:code:daily:%s:%s:%s", tenantId, strings.ToLower(strings.TrimSpace(contact)), time.Now().Format("20060102"))
}

func registerCodeDailyLimitTTL() time.Duration {
	now := time.Now()
	tomorrow := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	return time.Until(tomorrow)
}

func maskRegisterContact(contact string) string {
	if isPhoneContact(contact) && len(contact) == 11 {
		return contact[:3] + "****" + contact[7:]
	}
	if at := strings.Index(contact, "@"); at > 1 {
		return contact[:1] + "****" + contact[at:]
	}
	return contact
}

func renderRegisterTemplate(template string, contact string, code string) string {
	replacer := strings.NewReplacer(
		"{contact}", contact,
		"{code}", code,
		"${contact}", contact,
		"${code}", code,
	)
	return replacer.Replace(template)
}

func getRegisterConfig(ctx context.Context, key string, defaultValue string) string {
	value, err := mcache.GetSystemConfig(ctx, key, defaultValue)
	if err != nil || strings.TrimSpace(value) == "" {
		return g.Cfg().MustGet(ctx, key, defaultValue).String()
	}
	return value
}

func parseRegisterSmsConfig(ctx context.Context) (cfg registerSmsConfig, err error) {
	cfg.URL = g.Cfg().MustGet(ctx, "register.sms.url", "").String()
	cfg.Method = g.Cfg().MustGet(ctx, "register.sms.method", "POST").String()
	cfg.Headers = g.Cfg().MustGet(ctx, "register.sms.headers", map[string]string{}).MapStrStr()
	cfg.BodyTemplate = g.Cfg().MustGet(ctx, "register.sms.bodyTemplate", `{"phone":"{contact}","code":"{code}"}`).String()
	cfg.SuccessStatus = g.Cfg().MustGet(ctx, "register.sms.successStatus", 200).Int()
	return cfg, nil
}

func parseRegisterEmailConfig(ctx context.Context) (cfg registerEmailConfig, err error) {
	cfg.Host = g.Cfg().MustGet(ctx, "register.email.host", "").String()
	cfg.Port = g.Cfg().MustGet(ctx, "register.email.port", 465).Int()
	cfg.Username = g.Cfg().MustGet(ctx, "register.email.username", "").String()
	cfg.Password = g.Cfg().MustGet(ctx, "register.email.password", "").String()
	cfg.From = g.Cfg().MustGet(ctx, "register.email.from", "").String()
	cfg.FromName = g.Cfg().MustGet(ctx, "register.email.fromName", "DTU管理平台").String()
	cfg.Subject = g.Cfg().MustGet(ctx, "register.email.subject", "注册验证码").String()
	cfg.BodyTemplate = g.Cfg().MustGet(ctx, "register.email.bodyTemplate", "您的注册验证码是：{code}，6小时内有效。").String()
	cfg.SSL = g.Cfg().MustGet(ctx, "register.email.ssl", true).Bool()
	return cfg, nil
}

func sendRegisterSmsCode(ctx context.Context, contact string, code string) error {
	cfg, err := parseRegisterSmsConfig(ctx)
	if err != nil {
		return gerror.Wrap(err, "短信验证码配置错误")
	}
	if strings.TrimSpace(cfg.URL) == "" {
		g.Log().Infof(ctx, "注册短信验证码发送接口未配置，跳过真实发送 contact=%s code=%s", maskRegisterContact(contact), code)
		return nil
	}
	method := strings.ToUpper(strings.TrimSpace(cfg.Method))
	if method == "" {
		method = "POST"
	}
	if cfg.SuccessStatus == 0 {
		cfg.SuccessStatus = 200
	}
	body := renderRegisterTemplate(cfg.BodyTemplate, contact, code)
	client := g.Client()
	client.SetHeader("Content-Type", "application/json")
	for k, v := range cfg.Headers {
		client.SetHeader(k, renderRegisterTemplate(v, contact, code))
	}
	var response *gclient.Response
	switch method {
	case "GET":
		response, err = client.Get(ctx, renderRegisterTemplate(cfg.URL, contact, code))
	default:
		response, err = client.Post(ctx, cfg.URL, body)
	}
	if err != nil {
		return gerror.Wrap(err, "调用短信验证码接口失败")
	}
	defer response.Close()
	if response.StatusCode != cfg.SuccessStatus {
		return gerror.Newf("调用短信验证码接口失败，状态码：%d，响应：%s", response.StatusCode, response.ReadAllString())
	}
	return nil
}

func sendRegisterEmailCode(ctx context.Context, contact string, code string) error {
	cfg, err := parseRegisterEmailConfig(ctx)
	if err != nil {
		return gerror.Wrap(err, "邮件验证码配置错误")
	}
	if strings.TrimSpace(cfg.Host) == "" || strings.TrimSpace(cfg.Username) == "" {
		g.Log().Infof(ctx, "注册邮件SMTP未配置，跳过真实发送 contact=%s code=%s", maskRegisterContact(contact), code)
		return nil
	}
	if cfg.Port == 0 {
		cfg.Port = 25
	}
	if cfg.Subject == "" {
		cfg.Subject = "注册验证码"
	}
	if cfg.BodyTemplate == "" {
		cfg.BodyTemplate = "您的注册验证码是：{code}，6小时内有效。"
	}
	from := cfg.From
	if from == "" {
		from = cfg.Username
	}
	body := renderRegisterTemplate(cfg.BodyTemplate, contact, code)
	fromHeader := from
	if cfg.FromName != "" {
		fromHeader = fmt.Sprintf("%s <%s>", cfg.FromName, from)
	}
	message := strings.Join([]string{
		"From: " + fromHeader,
		"To: " + contact,
		"Subject: " + cfg.Subject,
		"MIME-Version: 1.0",
		"Content-Type: text/html; charset=UTF-8",
		"",
		body,
	}, "\r\n")
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
	if cfg.SSL {
		conn, err := tls.Dial("tcp", addr, &tls.Config{ServerName: cfg.Host, MinVersion: tls.VersionTLS12})
		if err != nil {
			return gerror.Wrap(err, "连接SMTP服务器失败")
		}
		defer conn.Close()
		client, err := smtp.NewClient(conn, cfg.Host)
		if err != nil {
			return gerror.Wrap(err, "创建SMTP客户端失败")
		}
		defer client.Close()
		if err = client.Auth(auth); err != nil {
			return gerror.Wrap(err, "SMTP认证失败")
		}
		if err = client.Mail(from); err != nil {
			return err
		}
		if err = client.Rcpt(contact); err != nil {
			return err
		}
		writer, err := client.Data()
		if err != nil {
			return err
		}
		if _, err = writer.Write([]byte(message)); err != nil {
			_ = writer.Close()
			return err
		}
		return writer.Close()
	}
	return smtp.SendMail(addr, auth, from, []string{contact}, []byte(message))
}

func (l *sSysUser) SendRegisterCode(ctx context.Context, param *model.SysUserRegisterCodeModel) (err error) {
	param.TenantId = normalizeRegisterTenantId(param.TenantId)
	param.Contact = strings.TrimSpace(param.Contact)
	if !isPhoneContact(param.Contact) && !isEmailContact(param.Contact) {
		return gerror.New("请输入正确的手机号或邮箱")
	}
	if err = service.SysCaptcha().VerifyCaptcha(ctx, param.CaptchaID, param.CaptchaValue); err != nil {
		return err
	}
	dailyCountKey := registerCodeDailyCountKey(param.TenantId, param.Contact)
	dailyCountValue, err := mcache.Get(ctx, dailyCountKey)
	if err == nil && gconv.Int(dailyCountValue) >= registerCodeDailyLimit {
		return gerror.Newf("当天验证码发送次数已达上限，每个手机号或邮箱每天最多发送%d条", registerCodeDailyLimit)
	}
	code := grand.Digits(6)
	if err = mcache.Set(ctx, registerCodeCacheKey(param.TenantId, param.Contact), code, registerCodeTTL); err != nil {
		return err
	}
	if isPhoneContact(param.Contact) {
		err = sendRegisterSmsCode(ctx, param.Contact, code)
	} else {
		err = sendRegisterEmailCode(ctx, param.Contact, code)
	}
	if err != nil {
		_, _ = mcache.Instance().Remove(ctx, registerCodeCacheKey(param.TenantId, param.Contact))
		return err
	}
	if err = mcache.Set(ctx, dailyCountKey, gconv.Int(dailyCountValue)+1, registerCodeDailyLimitTTL()); err != nil {
		return err
	}
	return nil
}

func (l *sSysUser) Register(ctx context.Context, param *model.SysUserRegisterModel) (err error) {
	param.TenantId = normalizeRegisterTenantId(param.TenantId)
	param.Contact = strings.TrimSpace(param.Contact)
	if param.Contact == "" {
		param.Contact = strings.TrimSpace(param.UserName)
	}
	if param.UserName == "" {
		param.UserName = param.Contact
	}
	if param.ConfirmPassword == "" {
		param.ConfirmPassword = param.Password
	}
	if param.CompanyName == "" {
		param.CompanyName = param.UserName
	}
	if param.Password != param.ConfirmPassword {
		return gerror.New("两次输入的密码不一致")
	}
	if param.Code != "" {
		if !isPhoneContact(param.Contact) && !isEmailContact(param.Contact) {
			return gerror.New("请输入正确的手机号或邮箱")
		}
		cacheValue, err := mcache.Get(ctx, registerCodeCacheKey(param.TenantId, param.Contact))
		cacheCode := gconv.String(cacheValue)
		if err != nil || cacheCode == "" {
			return gerror.NewCode(consts.CodeCaptchaError, "验证码已失效")
		}
		if cacheCode != param.Code {
			return gerror.NewCode(consts.CodeCaptchaError, "验证码错误")
		}
	} else if param.CaptchaID != "" || param.CaptchaValue != "" {
		if err = service.SysCaptcha().VerifyCaptcha(ctx, param.CaptchaID, param.CaptchaValue); err != nil {
			return err
		}
	} else {
		return gerror.NewCode(consts.CodeCaptchaError, "验证码不能为空")
	}

	// 检查是否开启用户注册功能
	registerConfig := &model.SysConfigViewModel{}
	mod := g.DB().Model(dao.SysConfig.Table())
	mod = mod.Where(dao.SysConfig.Columns().ConfigKey, "sys.account.registerUser")
	mod = mod.Where(dao.SysConfig.Columns().TenantId, param.TenantId)
	err = mod.Scan(&registerConfig)
	if err != nil {
		return gerror.New("请联系管理员开启用户注册功能")
	}
	if registerConfig == nil || registerConfig.ConfigValue != "true" {
		return gerror.New("请联系管理员开启用户注册功能")
	}

	// 获取租户信息
	tenant, err := service.SysTenant().View(ctx, &model.SysTenantViewParam{
		TenantId: param.TenantId,
	})
	if err != nil {
		return err
	}
	if tenant == nil {
		return gerror.New("租户无效")
	}

	// 获取租户的默认注册组织ID
	config := &model.SysConfigViewModel{}
	mod = g.DB().Model(dao.SysConfig.Table())
	mod = mod.Where(dao.SysConfig.Columns().ConfigKey, consts.ConfigKeyTenantDefaultRegisterDeptId)
	mod = mod.Where(dao.SysConfig.Columns().TenantId, param.TenantId)
	err = mod.Scan(&config)
	if err != nil {
		return gerror.New("该租户不允许自主注册")
	}
	if config == nil || config.ConfigValue == "" {
		return gerror.New("该租户不允许自主注册")
	}

	// 验证用户名是否存在
	mod1 := g.DB().Model(dao.SysUser.Table())
	mod1 = mod1.Where(dao.SysUser.Columns().UserName, param.UserName)
	mod1 = mod1.Where(dao.SysUser.Columns().TenantId, param.TenantId)
	total, err := mod1.Count()
	if err != nil {
		return err
	}
	if total > 0 {
		return gerror.New("用户名已存在")
	}
	if isPhoneContact(param.Contact) {
		total, err = g.DB().Model(dao.SysUser.Table()).
			Where(dao.SysUser.Columns().Phonenumber, param.Contact).
			Where(dao.SysUser.Columns().TenantId, param.TenantId).
			Count()
		if err != nil {
			return err
		}
		if total > 0 {
			return gerror.NewCode(consts.CodePhoneExists, "手机号已存在")
		}
	} else {
		total, err = g.DB().Model(dao.SysUser.Table()).
			Where(dao.SysUser.Columns().Email, param.Contact).
			Where(dao.SysUser.Columns().TenantId, param.TenantId).
			Count()
		if err != nil {
			return err
		}
		if total > 0 {
			return gerror.NewCode(consts.CodeEmailExists, "邮箱已存在")
		}
	}
	// 随机生成5位字符
	salt := utility.RandomString(5)
	password := utility.PasswordEncrypt(param.Password, salt)

	newUser := &entity.SysUser{
		UserName: param.UserName,
		NickName: param.CompanyName,
		Password: password,
		Salt:     salt,
		TenantId: param.TenantId,
		DeptId:   gconv.Int64(config.ConfigValue),
		Status:   consts.SysDeptStatusNormal,
		Remark:   "注册公司：" + param.CompanyName,
	}
	if isPhoneContact(param.Contact) {
		newUser.Phonenumber = param.Contact
	} else if isEmailContact(param.Contact) {
		newUser.Email = param.Contact
	}
	newUser.CreatedBy = 0
	newUser.CreatedAt = gtime.Now()
	newUser.UpdatedBy = 0
	newUser.UpdatedAt = gtime.Now()

	userId, err := dao.SysUser.Ctx(ctx).Data(newUser).InsertAndGetId()
	if err != nil {
		return err
	}

	//配置角色
	// 新增用户角色
	mod = g.DB().Model(dao.SysConfig.Table())
	mod = mod.Where(dao.SysConfig.Columns().ConfigKey, consts.ConfigKeyTenantDefaultRegisterRoleId)
	mod = mod.Where(dao.SysConfig.Columns().TenantId, param.TenantId)
	err = mod.Scan(&config)
	if err != nil {
		g.Log().Error(ctx, "获取租户的默认角色失败", err)
	}
	if config != nil && config.ConfigValue != "" {
		roleId := gconv.Int64(config.ConfigValue)
		d := map[string]any{dao.SysUserRole.Columns().UserId: userId, dao.SysUserRole.Columns().RoleId: roleId}
		_, err = dao.SysUserRole.Ctx(ctx).Data(d).Insert()
		if err != nil {
			g.Log().Error(ctx, "获取租户的默认角色失败", err)
		}
	}
	if param.Code != "" {
		_, _ = mcache.Instance().Remove(ctx, registerCodeCacheKey(param.TenantId, param.Contact))
	}
	return nil
}

// Status 更新用户状态
func (l *sSysUser) Status(ctx context.Context, param *model.SysUserStatusParam) (err error) {
	_, err = l.Model(ctx).Where(dao.SysUser.Columns().UserId, param.UserId).Data(map[string]any{
		dao.SysUser.Columns().Status: param.Status,
	}).OmitEmpty().Update()
	if err != nil {
		return err
	}
	return nil
}

// 批量查询用户迷你信息
func (l *sSysUser) BatchGetUserMiniInfo(ctx context.Context, userIds []int64) (users []*model.SysUserMiniModel, err error) {
	err = dao.SysUser.Ctx(ctx).WhereIn(dao.SysUser.Columns().UserId, userIds).Scan(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
