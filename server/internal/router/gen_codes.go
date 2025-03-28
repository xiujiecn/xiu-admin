// package router
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package router

import (
	"context"
	"xiuadmin/internal/controller/gen_codes"
	"xiuadmin/internal/service"

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
