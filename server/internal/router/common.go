// package router
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package router

import (
	"context"
	"xiuadmin/internal/controller/common"
	"xiuadmin/internal/service"

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
