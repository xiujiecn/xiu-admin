// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "server/api/system/v1"
	"server/internal/model"
	"server/internal/model/entity"
	"server/internal/model/request"
)

type (
	ISysAuth interface {
		Login(ctx context.Context, captchaID string, captchaValue string, username string, password string) (res *model.LoginUserOut, token string, err error)
		// 生成Token
		GenerateToken(ctx context.Context, user *model.LoginUserOut) (claims *model.CustomClaims, token string, err error)
		// 解析Token
		ParseToken(ctx context.Context, token string) (claims *model.CustomClaims, err error)
		// 删除Token
		DeleteToken(ctx context.Context, token string) (err error)
		// 根据Token获取当前登录用户信息
		GetCurrentUser(ctx context.Context) (claims *model.CustomClaims, err error)
	}
	ISysCaptcha interface {
		// 生成验证码
		GenerateCaptcha(ctx context.Context) (key string, image string, err error)
		// 验证验证码
		VerifyCaptcha(ctx context.Context, key string, value string) (err error)
	}
	ISysClient interface {
		List(ctx context.Context, query *model.SysClientListParam, pageInfo *request.PageInfo) (items []*model.SysClientListModel, total int, err error)
	}
	ISysConfig interface {
		GetConfigList(ctx context.Context, query *model.SysConfigListParam, page *request.PageInfo) (items []*model.SysConfig, total int, err error)
	}
	ISysDept interface {
		GetDeptList(ctx context.Context, query model.DeptListParam) (items []*model.SysDept, total int, err error)
		GetDeptById(ctx context.Context, id int64) (dept *model.SysDept, err error)
		// 构建树结构
		DeptTree(ctx context.Context, parentDept *model.SysDeptTreeModel, deptList []*model.SysDept, ancestors string) (data []*model.SysDeptTreeModel, err error)
		GetDeptTree(ctx context.Context) (items []*model.SysDeptTreeModel, err error)
		// 递归构建结构
		RecursionDeptIds(ctx context.Context, parentId int64, deptList []*model.SysDept, data *[]int64) (err error)
		// 根据父部门id获取部门列表
		GetDeptIdsByParentId(ctx context.Context, parentId int64) (ids []int64, err error)
	}
	ISysDict interface {
		GetDictTypeList(ctx context.Context, req *model.SysDictTypeListParam, pageInfo *request.PageInfo) (items []model.SysDictType, total int, err error)
		GetDictDataList(ctx context.Context, query *model.SysDictDataListParam, pageInfo *request.PageInfo) (items []model.SysDictData, total int, err error)
	}
	ISysLogininfor interface {
		ListLogininfor(ctx context.Context, query *model.SysLogininforListParam, pageInfo *request.PageInfo) (items []*model.SysLogininforListModel, total int, err error)
		AddLogininfor(ctx context.Context, logininfor *model.SysLogininforAddModel) (id int64, err error)
	}
	ISysMenu interface {
		// 获取租户菜单列表， 系统租户返回所有菜单，其他租户返回当前租户菜单
		GetTenantMenu(ctx context.Context, query *model.SysMenuListParam) (data []*model.SysMenuListModel, total int, err error)
		// 构建树结构
		MenuTree(ctx context.Context, parentMenu *model.SysMenuTree, menuList []*entity.SysMenu) (data []*model.SysMenuTree, err error)
		// 获取用户动态路由列表
		GetUserMenu(ctx context.Context) (data []*entity.SysMenu, err error)
		// 构建用户动态路由树
		BuildUserMenuTree(ctx context.Context, parentMenu *v1.RouteMenu, menuList []*entity.SysMenu, allPath string) (data v1.MenuAllRes, err error)
		// 获取用户动态路由树
		GetUserMenuTree(ctx context.Context) (data v1.MenuAllRes, err error)
		GetSysMenuView(ctx context.Context, menuId int64) (data *model.SysMenuViewModel, err error)
		UpdateSysMenu(ctx context.Context, menu *model.SysMenuUpdateModel) (data *model.SysMenuViewModel, err error)
		AddSysMenu(ctx context.Context, menu *model.SysMenuAddModel) (data *model.SysMenuViewModel, err error)
		DeleteSysMenu(ctx context.Context, menuId int64) (err error)
	}
	ISysNotice interface {
		GetNoticeList(ctx context.Context, query *model.SysNoticeListParam, page *request.PageInfo) (items []*model.SysNotice, total int, err error)
	}
	ISysOperLog interface {
		GetOperLogList(ctx context.Context, query *model.SysOperLogListParam, page *request.PageInfo) (items []*model.SysOperLogListModel, total int, err error)
	}
	ISysOss interface {
		List(ctx context.Context, query *model.SysOssListParam, pageInfo *request.PageInfo) (items []*model.SysOssListModel, total int, err error)
	}
	ISysPost interface {
		GetPostList(ctx context.Context, query model.SysPostListParam, pageInfo request.PageInfo) (items []*model.SysPostListModel, total int, err error)
	}
	ISysRole interface {
		// 获取租户下角色列表
		GetRoleList(ctx context.Context, model *model.SysRoleListParam, pageInfo *request.PageInfo) (res []*model.SysRoleListModel, total int, err error)
		// 获取角色详情
		GetRoleView(ctx context.Context, id int64) (res *model.SysRoleViewModel, err error)
		// 获取角色菜单
		GetRoleMenu(ctx context.Context, id int64) (res []*entity.SysMenu, err error)
		// 获取角色列表对应菜单
		GetRoleListMenu(ctx context.Context, ids []int64) (res []*entity.SysRoleMenu, err error)
	}
	ISysSocial interface {
		List(ctx context.Context, query *model.SysSocialListParam, page *request.PageInfo) (items []*model.SysSocialListModel, total int, err error)
	}
	ISysTenant interface {
		// 获取租户信息
		GetTenantInfo(ctx context.Context, tenantId string) (data *entity.SysTenant, err error)
		// 获取租户套餐
		GetTenantPackage(ctx context.Context, packageId int64) (data *entity.SysTenantPackage, err error)
		// 获取租户列表
		List(ctx context.Context, query *model.SysTenantListParam, page *request.PageInfo) (data []*model.SysTenantListModel, total int, err error)
		TenantPackageList(ctx context.Context, query *model.SysTenantPackageListParam, page *request.PageInfo) (data []*model.SysTenantPackageListModel, total int, err error)
	}
	ISysUser interface {
		// 根据用户名获取用户信息
		GetUserByUsername(ctx context.Context, username string) (user *entity.SysUser, err error)
		// 根据邮箱获取用户信息
		GetUserByEmail(ctx context.Context, email string) (user *entity.SysUser, err error)
		// 根据手机号获取用户信息
		GetUserByPhone(ctx context.Context, phone string) (user *entity.SysUser, err error)
		// 根据用户名和密码获取用户信息
		GetUserByUsernameAndPassword(ctx context.Context, username string, password string) (user *entity.SysUser, err error)
		// 根据用户ID获取用户信息
		GetUserById(ctx context.Context, id int64) (user *model.SysUserViewModel, err error)
		// 获取用户列表
		GetUserList(ctx context.Context, page request.PageInfo, query model.UserListParam) (items []*model.SysUserListModel, total int, err error)
		// 新增用户
		AddUser(ctx context.Context, req *model.SysUserAddModel) (data *model.SysUserViewModel, err error)
		Profile(ctx context.Context) (user *model.UserProfileModel, err error)
		UpdateCurrentUser(ctx context.Context, req *model.UpdateCurrentUserModel) (user *model.SysUserViewModel, err error)
		UpdateCurrentUserPassword(ctx context.Context, req *model.UpdateCurrentUserPasswordModel) (err error)
		UpdateUser(ctx context.Context, req *model.SysUserUpdateModel) (err error)
		DeleteUser(ctx context.Context, userIds []int64) (err error)
		ResetPassword(ctx context.Context, userId int64, password string) (err error)
	}
	ISysUserOnline interface {
		Add(ctx context.Context, userOnline *model.SysUserOnlineAddModel) (err error)
		List(ctx context.Context, query *model.SysUserOnlineListParam, page *request.PageInfo) (items []*model.SysUserOnlineListModel, total int, err error)
		Delete(ctx context.Context, id int64) (err error)
	}
)

var (
	localSysAuth       ISysAuth
	localSysCaptcha    ISysCaptcha
	localSysClient     ISysClient
	localSysConfig     ISysConfig
	localSysDept       ISysDept
	localSysDict       ISysDict
	localSysLogininfor ISysLogininfor
	localSysMenu       ISysMenu
	localSysNotice     ISysNotice
	localSysOperLog    ISysOperLog
	localSysOss        ISysOss
	localSysPost       ISysPost
	localSysRole       ISysRole
	localSysSocial     ISysSocial
	localSysTenant     ISysTenant
	localSysUser       ISysUser
	localSysUserOnline ISysUserOnline
)

func SysAuth() ISysAuth {
	if localSysAuth == nil {
		panic("implement not found for interface ISysAuth, forgot register?")
	}
	return localSysAuth
}

func RegisterSysAuth(i ISysAuth) {
	localSysAuth = i
}

func SysCaptcha() ISysCaptcha {
	if localSysCaptcha == nil {
		panic("implement not found for interface ISysCaptcha, forgot register?")
	}
	return localSysCaptcha
}

func RegisterSysCaptcha(i ISysCaptcha) {
	localSysCaptcha = i
}

func SysClient() ISysClient {
	if localSysClient == nil {
		panic("implement not found for interface ISysClient, forgot register?")
	}
	return localSysClient
}

func RegisterSysClient(i ISysClient) {
	localSysClient = i
}

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}

func SysDept() ISysDept {
	if localSysDept == nil {
		panic("implement not found for interface ISysDept, forgot register?")
	}
	return localSysDept
}

func RegisterSysDept(i ISysDept) {
	localSysDept = i
}

func SysDict() ISysDict {
	if localSysDict == nil {
		panic("implement not found for interface ISysDict, forgot register?")
	}
	return localSysDict
}

func RegisterSysDict(i ISysDict) {
	localSysDict = i
}

func SysLogininfor() ISysLogininfor {
	if localSysLogininfor == nil {
		panic("implement not found for interface ISysLogininfor, forgot register?")
	}
	return localSysLogininfor
}

func RegisterSysLogininfor(i ISysLogininfor) {
	localSysLogininfor = i
}

func SysMenu() ISysMenu {
	if localSysMenu == nil {
		panic("implement not found for interface ISysMenu, forgot register?")
	}
	return localSysMenu
}

func RegisterSysMenu(i ISysMenu) {
	localSysMenu = i
}

func SysNotice() ISysNotice {
	if localSysNotice == nil {
		panic("implement not found for interface ISysNotice, forgot register?")
	}
	return localSysNotice
}

func RegisterSysNotice(i ISysNotice) {
	localSysNotice = i
}

func SysOperLog() ISysOperLog {
	if localSysOperLog == nil {
		panic("implement not found for interface ISysOperLog, forgot register?")
	}
	return localSysOperLog
}

func RegisterSysOperLog(i ISysOperLog) {
	localSysOperLog = i
}

func SysOss() ISysOss {
	if localSysOss == nil {
		panic("implement not found for interface ISysOss, forgot register?")
	}
	return localSysOss
}

func RegisterSysOss(i ISysOss) {
	localSysOss = i
}

func SysPost() ISysPost {
	if localSysPost == nil {
		panic("implement not found for interface ISysPost, forgot register?")
	}
	return localSysPost
}

func RegisterSysPost(i ISysPost) {
	localSysPost = i
}

func SysRole() ISysRole {
	if localSysRole == nil {
		panic("implement not found for interface ISysRole, forgot register?")
	}
	return localSysRole
}

func RegisterSysRole(i ISysRole) {
	localSysRole = i
}

func SysSocial() ISysSocial {
	if localSysSocial == nil {
		panic("implement not found for interface ISysSocial, forgot register?")
	}
	return localSysSocial
}

func RegisterSysSocial(i ISysSocial) {
	localSysSocial = i
}

func SysTenant() ISysTenant {
	if localSysTenant == nil {
		panic("implement not found for interface ISysTenant, forgot register?")
	}
	return localSysTenant
}

func RegisterSysTenant(i ISysTenant) {
	localSysTenant = i
}

func SysUser() ISysUser {
	if localSysUser == nil {
		panic("implement not found for interface ISysUser, forgot register?")
	}
	return localSysUser
}

func RegisterSysUser(i ISysUser) {
	localSysUser = i
}

func SysUserOnline() ISysUserOnline {
	if localSysUserOnline == nil {
		panic("implement not found for interface ISysUserOnline, forgot register?")
	}
	return localSysUserOnline
}

func RegisterSysUserOnline(i ISysUserOnline) {
	localSysUserOnline = i
}
