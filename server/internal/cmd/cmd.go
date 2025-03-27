// Package cmd
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"

	"xiuadmin/internal/library/websocket"
	"xiuadmin/internal/router"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			err = InitSystemDeferFunc(ctx)
			if err != nil {
				g.Log().Fatalf(ctx, "系统初始化异常 err: %+v", err)
				return err
			}
			s := g.Server()
			s.SetDumpRouterMap(false)
			router.InitRouter(ctx, s)
			websocket.StartWebSocket(ctx)
			s.Run()
			return nil
		},
	}
)
