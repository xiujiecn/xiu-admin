// package genrouter
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package genrouter

import (
	"context"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	NoLoginRouter       []interface{} // 无需登录
	LoginRequiredRouter []interface{} // 需要登录
)

// Register 注册通过代码生成的后台路由
func Register(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/gen", func(group *ghttp.RouterGroup) {
		if len(NoLoginRouter) > 0 {
			group.Bind(NoLoginRouter...)
		}
		group.Middleware(service.Middleware().Auth)
		if len(LoginRequiredRouter) > 0 {
			group.Bind(LoginRequiredRouter...)
		}
	})
}
