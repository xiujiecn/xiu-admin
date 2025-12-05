// Package websocket
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package websocket

import (
	"fmt"

	"github.com/gogf/gf/v2/util/gconv"
)

const (
	Error = "error"
	Login = "login"
	Join  = "join"
	Quit  = "quit"
	IsApp = "is_app"
	Ping  = "ping"
	Joins = "joins"
	Quits = "quits"
)

// ProcessData 处理数据
func ProcessData(client *Client, message []byte) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("websocket.ProcessData 处理数据 stop", r, "message=", message)
		}
	}()
	request := &request{}
	err := gconv.Struct(message, request)
	if err != nil {
		fmt.Println("websocket.ProcessData 数据解析失败：", err)
		return
	}
	fmt.Println("websocket.ProcessData request", request)
	switch request.Event {
	case Login:
		LoginController(client, request)
		break
	case Join:
		JoinController(client, request)
		break
	case Quit:
		QuitController(client, request)
		break
	case IsApp:
		IsAppController(client)
		break
	case Ping:
		PingController(client)
		break
	case Joins:
		JoinsController(client, request)
		break
	case Quits:
		QuitsController(client, request)
		break
	}
}
