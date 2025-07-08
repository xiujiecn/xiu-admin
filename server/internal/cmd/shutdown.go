package cmd

import (
	"context"
	"os"
	"sync"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/library/event"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
)

var (
	serverCloseSignal = make(chan struct{}, 1)
	serverWg          = sync.WaitGroup{}
	once              sync.Once
)

// 等待接收服务关闭信号
func waitServerClose(ctx context.Context) {
	signalListen(ctx, handleSignalClose)
	<-serverCloseSignal
}

// 等待所有服务关闭
func waitAllServerClose(ctx context.Context) {
	serverWg.Wait()
}

// 处理服务关闭信号
func handleSignalClose(sig os.Signal) {
	serverCloseEvent(gctx.GetInitCtx())
	serverCloseSignal <- struct{}{}
}

// 监听信号
func signalListen(ctx context.Context, handler ...gproc.SigHandler) {
	g.Go(ctx, func(ctx context.Context) {
		// syscall.SIGINT,
		// syscall.SIGQUIT,
		// syscall.SIGKILL,
		// syscall.SIGTERM,
		// syscall.SIGABRT.
		gproc.AddSigHandlerShutdown(handler...)
		gproc.Listen()
	}, func(ctx context.Context, exception error) {
		g.Log().Fatalf(ctx, "信号监听异常 err: %+v", exception)
	})
}

// serverCloseEvent 关闭事件
func serverCloseEvent(ctx context.Context) {
	once.Do(func() {
		event.EventsInstance().Emit(ctx, consts.EventKeyServerClose, nil)
	})
}
