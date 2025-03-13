package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"

	"xiujieadmin/internal/router"
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
			router.InitRouter(ctx, s)
			s.Run()
			return nil
		},
	}
)
