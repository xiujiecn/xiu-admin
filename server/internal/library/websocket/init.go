// Package websocket
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package websocket

import (
	"context"
	"net/http"
	"xiuadmin/internal/library/contexts"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gorilla/websocket"
)

var (
	mctx          = gctx.GetInitCtx()  // 上下文
	clientManager = NewClientManager() // 管理者
)
var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartWebSocket(ctx context.Context) {
	g.Log().Info(ctx, "启动：WebSocket")
	go clientManager.start()
	go clientManager.ping(ctx)
}

func StopWebSocket(ctx context.Context) {
	g.Log().Info(ctx, "停止：WebSocket")
	clientManager.closeSignal <- struct{}{}
}

func WsPage(r *ghttp.Request) {
	g.Log().Infof(r.GetCtx(), "websocket.WsPage Begin. UserID:%v ", contexts.GetUserId(r.GetCtx()))
	// r.Response.Header().Set("Content-Type", "text/event-stream")
	contexts.SetIsWebSocket(r.GetCtx(), true)
	conn, err := upGrader.Upgrade(r.Response.Writer, r.Request, nil)
	if err != nil {
		r.Response.Write(err.Error())
		return
	}
	currentTime := uint64(gtime.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, currentTime, r.GetCtx())

	client.UserId = uint64(contexts.GetUserId(r.GetCtx()))
	go client.read()
	go client.write()

	// 用户连接事件
	clientManager.Register <- client

	g.Log().Infof(r.GetCtx(), "websocket.WsPage End. UserID:%v", client.UserId)
}
