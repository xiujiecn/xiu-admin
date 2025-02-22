package router

import (
	"context"
	"server/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

// InitRouter 初始化路由
func InitRouter(ctx context.Context, s *ghttp.Server) {
	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().ResponseHandler)
		System(ctx, group)
		Monitor(ctx, group)
	})
}
