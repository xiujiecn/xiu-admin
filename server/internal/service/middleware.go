// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// 响应处理中间件
		ResponseHandler(r *ghttp.Request)
		// Ctx 自定义上下文对象
		Ctx(r *ghttp.Request)
		// CORS 跨域处理
		CORS(r *ghttp.Request)
		// Auth 认证处理
		Auth(r *ghttp.Request)
		// IsExceptLogin 判断是否需要验证登录
		IsExceptLogin(ctx context.Context, path string) bool
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
