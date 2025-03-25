package router

import (
	"context"
	"xiujieadmin/internal/router/genrouter"

	"github.com/gogf/gf/v2/net/ghttp"
)

// 生成代码注册
func Gen(ctx context.Context, group *ghttp.RouterGroup) {
	genrouter.Register(ctx, group)
}
