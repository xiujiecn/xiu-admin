// package router
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package router

import (
	"context"
	"xiuadmin/internal/controller/monitor"
	"xiuadmin/internal/library/websocket"
	"xiuadmin/internal/service"

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
