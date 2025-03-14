package router

import (
	"context"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// InitRouter 初始化路由
func InitRouter(ctx context.Context, s *ghttp.Server) {
	uploadPath := g.Cfg().MustGet(ctx, "system.upload.local.path").String()
	if uploadPath != "" {
		s.AddStaticPath("/upload", uploadPath)
	}
	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware().ResponseHandler,
			service.Middleware().Ctx,
			service.Middleware().CORS,
		)
		System(ctx, group)
		Monitor(ctx, group)
		Common(ctx, group)
	})
	//操作日志
	s.BindHookHandler("/*", ghttp.HookAfterOutput, func(r *ghttp.Request) {
		service.Middleware().OperationLog(r)
	})
}
