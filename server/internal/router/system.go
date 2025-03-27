// package router
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package router

import (
	"context"

	"xiuadmin/internal/controller/system"
	"xiuadmin/internal/service"

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
