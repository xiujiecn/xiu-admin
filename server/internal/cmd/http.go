package cmd

import (
	"context"
	"xiuadmin/internal/cmd/inithttp"
	"xiuadmin/internal/library/websocket"
	"xiuadmin/internal/router"
	"xiuadmin/internal/service"
	"xiuadmin/internal/tasks"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"

	_ "xiuadmin/internal/queues/httpqueues"
)

func init() {
	inithttp.RegisterHttpInitFunc("initJobTask", initJobTask)
}

var Http = &gcmd.Command{
	Name:  "http",
	Usage: "http",
	Brief: "HTTP控制台服务",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		g.Log().Infof(ctx, "Http init start")
		serverWg.Add(1)
		s := g.Server()
		go func() {
			waitServerClose(ctx)
			websocket.StopWebSocket(ctx)
			s.Shutdown()
			serverWg.Done()
		}()
		err = inithttp.InitHttp(ctx)
		if err != nil {
			g.Log().Errorf(ctx, "http init error: %v", err)
			return err
		}
		s.SetDumpRouterMap(false)
		router.InitRouter(ctx, s)
		websocket.StartWebSocket(ctx)
		s.Run()
		inithttp.CleanupHttp(ctx)
		return nil
	},
}

func initJobTask(ctx context.Context) error {
	tasks.TasksInstance(ctx)
	service.SysJob().InitRegister()
	return nil
}
