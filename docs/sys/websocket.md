---
outline: deep
---

# WebSocket服务器


> XiuAdmin提供了一个WebSocket服务器，随`HTTP服务`启停。集成了许多常用功能，如JWT身份认证、路由消息处理器、一对一消息/群组消息/广播消息、在线用户管理、心跳保持等，大大简化和规范了WebSocket服务器的开发流程。
- [Websocket客户端](websocket-client.md)

###  一个基本的消息收发例子
- 这是一个基本的消息接收并进行处理的简单例子

#### 1. 注册`标签`订阅,处理接口
- 文件路径：server/internal/cmd/monitor.go
```go
package cmd

import (
	"context"
	"fmt"
	"time"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/library/websocket"
	"xiuadmin/internal/service"
)

func InitMonitor() {
	go func() {
		for {
			RunMonitor()
			time.Sleep(time.Second * 5)
		}
	}()
	websocket.RegisterTagCallback(consts.WSTagMonitorServer, func(client *websocket.Client) {
		RunMonitor()
	})
}

```

#### 2.广播给订阅指定`标签`消息
- 以下是功能案例中的一个简单演示，实现了消息处理接口，并将收到的消息原样发送给客户端
- 文件路径：server/addons/hgexample/controller/websocket/handler/index.go
```go
	websocket.SendToTag(consts.WSTagMonitorServer, &websocket.WResponse{
		Event: fmt.Sprintf(consts.WSEventMonitorServer, consts.MonitorServerHost),
		Data:  hostInfo,
	})
```

- 到此，你已了解了WebSocket消息接收并进行处理的基本流程


### 常用方法
- websocket服务器还提供了一些常用的方法，下面只对部分进行说明
```go
func test() {
    websocket.SendToAll(response *WResponse)                        //发送全部客户端
    websocket.SendToClientID(clientId string, response *WResponse)  //发送单个客户端
    websocket.SendToUser(userId int64, response *WResponse)         //发送单个用户
    websocket.SendToTag(tag string, response *WResponse)            //发送某个标签、群组
}
```


### 其他
- WebSocket请求URL中需要携带`access_token`参数，授权中间件验证授权信息识别用户.
