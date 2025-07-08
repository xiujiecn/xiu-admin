// Package websocket
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package websocket

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// LoginController 用户登录
func LoginController(client *Client, req *request) {
	l := &WsLoginDataReq{}
	gconv.Struct(req.Data, &l)
	userId := gconv.Uint64(l.UserID)
	client.UserId = userId
	// 用户登录
	login := &login{
		UserId: userId,
		Client: client,
	}
	clientManager.Login <- login

	client.SendMsg(&WResponse{
		Event: Login,
		Data:  "success",
	})
}

func IsAppController(client *Client) {
	client.isApp = true
}

// JoinController 加入
func JoinController(client *Client, req *request) {
	name := gconv.String(req.Data["name"])

	if !client.tags.Contains(name) {
		client.tags.Append(name)
		if callback, ok := clientManager.TagCallbackMap[name]; ok {
			if callback != nil {
				callback(client)
			}
		}
		AddTagCount(name)
	}
	client.SendMsg(&WResponse{
		Event: Join,
		Data:  client.tags.Slice(),
	})
}

// QuitController 退出
func QuitController(client *Client, req *request) {
	name := gconv.String(req.Data["name"])
	if client.tags.Contains(name) {
		client.tags.RemoveValue(name)
		DelTagCount(name)
	}
	client.SendMsg(&WResponse{
		Event: Quit,
		Data:  client.tags.Slice(),
	})
}

func PingController(client *Client) {
	currentTime := uint64(gtime.Now().Unix())
	client.Heartbeat(currentTime)
}

// JoinsController 加入
func JoinsController(client *Client, req *request) {
	for _, name := range gconv.Strings(req.Data["names"]) {
		if !client.tags.Contains(name) {
			client.tags.Append(name)
			if callback, ok := clientManager.TagCallbackMap[name]; ok {
				if callback != nil {
					callback(client)
				}
			}
			AddTagCount(name)
		}
	}
	client.SendMsg(&WResponse{
		Event: Joins,
		Data:  client.tags.Slice(),
	})
}

// QuitController 退出
func QuitsController(client *Client, req *request) {
	for _, name := range gconv.Strings(req.Data["names"]) {
		if client.tags.Contains(name) {
			client.tags.RemoveValue(name)
			DelTagCount(name)
		}
	}
	client.SendMsg(&WResponse{
		Event: Quits,
		Data:  client.tags.Slice(),
	})
}
