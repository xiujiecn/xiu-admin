package router

import (
	"context"
	"xiujieadmin/internal/controller/gen_codes"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

// GenCodes 代码生成路由
func GenCodes(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/genCodes", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware().Auth,
		)
		group.Bind(
			gen_codes.NewV1(),
		)
	})
}
