// package router
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package router

import (
	"context"
	"xiuadmin/internal/controller/info"
	"xiuadmin/internal/service"

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
			service.Middleware().CORS,
		)
		group.Bind(
			info.NewV1(),
		)
	})
	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware().Ctx,
			service.Middleware().ResponseHandler,
			service.Middleware().CORS,
		)
		System(ctx, group)
		Monitor(ctx, group)
		Common(ctx, group)
		GenCodes(ctx, group)
		Gen(ctx, group)
	})
	//操作日志
	s.BindHookHandler("/*", ghttp.HookAfterOutput, func(r *ghttp.Request) {
		service.Middleware().OperationLog(r)
	})
}
