package router

import (
	"context"
	"server/internal/controller/monitor"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Monitor 监控功能的路由，不含业务属性的
func Monitor(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/monitor", func(group *ghttp.RouterGroup) {
		group.Bind(
			monitor.NewV1(),
		)
	})
}
