// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/internal/model"
	"server/internal/model/entity"
)

type (
	ISysAuth interface {
		Login(ctx context.Context, captchaID string, captchaValue string, username string, password string) (res *model.LoginUserOut, token string, err error)
		// 生成Token
		GenerateToken(ctx context.Context, user *model.LoginUserOut) (token string, err error)
		// 解析Token
		ParseToken(ctx context.Context, token string) (claims *model.CustomClaims, err error)
		// 根据Token获取当前登录用户信息
		GetCurrentUser(ctx context.Context) (claims *model.CustomClaims, err error)
	}
	ISysCaptcha interface {
		// 生成验证码
		GenerateCaptcha(ctx context.Context) (key string, image string, err error)
		// 验证验证码
		VerifyCaptcha(ctx context.Context, key string, value string) (err error)
	}
	ISysMenu interface {
		// 获取租户菜单列表， 系统租户返回所有菜单，其他租户返回当前租户菜单
		GetTenantMenu(ctx context.Context) (data []*entity.SysMenu, err error)
		// 获取用户动态路由列表
		GetUserMenu(ctx context.Context) (data []*entity.SysMenu, err error)
		// 构建树结构
		MenuTree(ctx context.Context, parentMenu *model.SysMenuTree, menuList []*entity.SysMenu) (data []*model.SysMenuTree, err error)
		// 获取用户动态路由树
		GetMenuTree(ctx context.Context) (data []*model.SysMenuTree, err error)
	}
	ISysRole interface {
		// 获取租户下角色列表
		GetRoleList(ctx context.Context, tenantId int64) (res []*entity.SysRole, err error)
		// 获取角色详情
		GetRoleDetail(ctx context.Context, id int64) (res *entity.SysRole, err error)
		// 获取角色菜单
		GetRoleMenu(ctx context.Context, id int64) (res []*entity.SysMenu, err error)
		// 获取角色列表对应菜单
		GetRoleListMenu(ctx context.Context, ids []int64) (res []*entity.SysRoleMenu, err error)
	}
	ISysTenant interface {
		// 获取租户信息
		GetTenantInfo(ctx context.Context, tenantId string) (data *entity.SysTenant, err error)
		// 获取租户套餐
		GetTenantPackage(ctx context.Context, packageId int64) (data *entity.SysTenantPackage, err error)
	}
	ISysUser interface {
		// 根据用户名获取用户信息
		GetUserByUsername(ctx context.Context, username string) (user *entity.SysUser, err error)
		// 根据用户名和密码获取用户信息
		GetUserByUsernameAndPassword(ctx context.Context, username string, password string) (user *entity.SysUser, err error)
		// 根据用户ID获取用户信息
		GetUserById(ctx context.Context, id int64) (user *entity.SysUser, err error)
	}
)

var (
	localSysAuth    ISysAuth
	localSysCaptcha ISysCaptcha
	localSysMenu    ISysMenu
	localSysRole    ISysRole
	localSysTenant  ISysTenant
	localSysUser    ISysUser
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

func SysMenu() ISysMenu {
	if localSysMenu == nil {
		panic("implement not found for interface ISysMenu, forgot register?")
	}
	return localSysMenu
}

func RegisterSysMenu(i ISysMenu) {
	localSysMenu = i
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
