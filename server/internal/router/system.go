package router

import (
	"context"

	"server/internal/controller/system"

	"github.com/gogf/gf/v2/net/ghttp"
)

// System 系统默认功能的路由，不含业务属性的
func System(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.Bind(
			system.NewV1(),
		)
	})
}
