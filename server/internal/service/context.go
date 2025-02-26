// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"xiujieadmin/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IContext interface {
		// Init初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
		Init(r *ghttp.Request, customCtx *model.Context)
		// Get 获取上下文变量；
		Get(ctx context.Context) *model.Context
		// Set 设置上下文变量
		SetUser(ctx context.Context, user *model.Identity)
		// GetUser 获取上下文变量中的用户信息
		GetUser(ctx context.Context) *model.Identity
		// GetUserId 获取上下文变量中的用户ID
		GetUserId(ctx context.Context) int64
		// GetUserName 获取上下文变量中的用户名
		GetUserName(ctx context.Context) string
		// GetNickName 获取上下文变量中的用户昵称
		GetNickName(ctx context.Context) string
		// GetDeptId 获取上下文变量中的部门ID
		GetDeptId(ctx context.Context) int64
		// GetTenantId 获取上下文变量中的租户ID
		GetTenantId(ctx context.Context) string
		// GetAuthorityIds 获取上下文变量中的角色ID
		GetAuthorityIds(ctx context.Context) []int64
		// IsSystemTenant 判断是否是系统租户
		IsSystemTenant(ctx context.Context) bool
		// IsSuperAdmin 判断是否是超级管理员角色
		IsSuperAdmin(ctx context.Context) bool
	}
)

var (
	localContext IContext
)

func Context() IContext {
	if localContext == nil {
		panic("implement not found for interface IContext, forgot register?")
	}
	return localContext
}

func RegisterContext(i IContext) {
	localContext = i
}
