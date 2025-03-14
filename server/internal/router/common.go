package router

import (
	"context"
	"xiujieadmin/internal/controller/common"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Common 通用路由
func Common(ctx context.Context, group *ghttp.RouterGroup) {

	group.Group("/common", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware().Auth,
		)
		group.Bind(
			common.NewV1(),
		)
	})
}
