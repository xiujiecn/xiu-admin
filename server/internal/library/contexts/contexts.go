// Package contexts
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package contexts

import (
	"context"
	"slices"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type ContextKey struct {
	name string
}

var contextKey = &ContextKey{
	name: consts.ContextKey,
}

// Init初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(contextKey, customCtx)
}

// Set设置本地上下文
func Set(ctx context.Context, customCtx *model.Context) context.Context {
	ctx = context.WithValue(ctx, contextKey, customCtx)
	return ctx
}

// Get 获取上下文变量；
func Get(ctx context.Context) *model.Context {
	customCtx, ok := ctx.Value(contextKey).(*model.Context)
	if !ok {
		return nil
	}
	return customCtx
}

// Set 设置上下文变量
func SetUser(ctx context.Context, user *model.Identity) {
	customCtx := Get(ctx)
	customCtx.User = user
}

// GetUser 获取上下文变量中的用户信息
func GetUser(ctx context.Context) *model.Identity {
	customCtx := Get(ctx)
	if customCtx == nil {
		return nil
	}
	return customCtx.User
}

// GetUserId 获取上下文变量中的用户ID
func GetUserId(ctx context.Context) int64 {
	user := GetUser(ctx)
	if user == nil {
		return 0
	}
	return user.ID
}

// GetUserUuid 获取上下文变量中的用户UUID
func GetUserUuid(ctx context.Context) string {
	user := GetUser(ctx)
	if user == nil {
		return ""
	}
	return user.UUID
}

// GetUserName 获取上下文变量中的用户名
func GetUserName(ctx context.Context) string {
	user := GetUser(ctx)
	if user == nil {
		return ""
	}
	return user.Username
}

// GetNickName 获取上下文变量中的用户昵称
func GetNickName(ctx context.Context) string {
	user := GetUser(ctx)
	if user == nil {
		return ""
	}
	return user.NickName
}

// GetDeptId 获取上下文变量中的部门ID
func GetDeptId(ctx context.Context) int64 {
	user := GetUser(ctx)
	if user == nil {
		return 0
	}
	return user.DeptId
}

// GetTenantId 获取上下文变量中的租户ID
func GetTenantId(ctx context.Context) string {
	user := GetUser(ctx)
	if user == nil {
		return ""
	}
	return user.TenantId
}

// GetRoles 获取上下文变量中的角色列表
func GetRoles(ctx context.Context) []model.Role {
	user := GetUser(ctx)
	if user == nil {
		return nil
	}
	return user.Roles
}

// GetLoginAt 获取上下文变量中的登录时间
func GetLoginAt(ctx context.Context) int64 {
	user := GetUser(ctx)
	if user == nil {
		return 0
	}
	return user.LoginAt
}

// IsSystemTenant 判断是否是系统租户
func IsSystemTenant(ctx context.Context) bool {
	tenantId := GetTenantId(ctx)
	return tenantId == consts.DefaultSystemTenantCode
}

// IsSuperAdmin 判断是否是超级管理员角色
func IsSuperAdmin(ctx context.Context) bool {
	roles := GetRoles(ctx)
	return slices.ContainsFunc(roles, func(role model.Role) bool {
		return role.RoleId == consts.SuperAdminRoleId
	})
}
func IsWebSocket(ctx context.Context) bool {
	customCtx := Get(ctx)
	if customCtx == nil {
		return false
	}
	return customCtx.IsWebSocket
}
func SetIsWebSocket(ctx context.Context, isWebSocket bool) {
	customCtx := Get(ctx)
	if customCtx == nil {
		return
	}
	customCtx.IsWebSocket = isWebSocket
}

// GetAccessCodeList 获取上下文变量中的访问码列表
func GetAccessCodeList(ctx context.Context) []string {
	user := GetUser(ctx)
	if user == nil {
		return nil
	}
	if len(user.AccessCodeList) == 0 {
		return nil
	}
	return user.AccessCodeList
}
