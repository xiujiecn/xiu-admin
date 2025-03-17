package router

import (
	"context"
	"xiujieadmin/internal/controller/monitor"
	"xiujieadmin/internal/library/websocket"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Monitor 监控功能的路由，不含业务属性的
func Monitor(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/monitor", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware().Auth,
		)
		group.Bind(
			monitor.NewV1(),
		)
		group.ALL("/ws", websocket.WsPage)
	})
}
