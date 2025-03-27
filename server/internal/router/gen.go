// package router
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package router

import (
	"context"
	"xiuadmin/internal/router/genrouter"

	"github.com/gogf/gf/v2/net/ghttp"
)

// 生成代码注册
func Gen(ctx context.Context, group *ghttp.RouterGroup) {
	genrouter.Register(ctx, group)
}
