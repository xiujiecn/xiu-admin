package router

import (
	"context"

	"xiujieadmin/internal/controller/system"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

// System 系统默认功能的路由，不含业务属性的
func System(ctx context.Context, group *ghttp.RouterGroup) {

	group.Group("/system", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware().Auth,
		)
		group.Bind(
			system.NewV1(),
		)
	})
}
