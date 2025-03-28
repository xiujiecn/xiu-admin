---
outline: deep
---

# 消息队列

> 系统默认的队列asynq，可以参考：https://github.com/hibiken/asynq, 需要Redis支持

## 配置文件
- 配置文件：server/manifest/config/config.yaml

```yaml
# Queue配置
queue:
  type: "asynq" # 队列类型
  asynq:
    redis:
      address: 192.168.158.200:6379
      db: 0
      pass: redis123456@@
      user:
    concurrency: 10 # 最大同时执行的任务数量
    logLevel: 4 # 日志级别 DebugLevel 1, InfoLevel 2, WarnLevel 3, ErrorLevel 4, FatalLevel 5
```

## 实现接口
- 为了提供高度的扩展性，消费队列在设计上采用了接口化的思路。只需要实现以下接口，您就可以在任何地方注册和使用消费队列消费功能，从而实现更大的灵活性和可扩展性。

```go
// Process 消费者接口，实现该接口即可加入到消费队列中
type Process interface {
	GetTopic() string                                  // 获取消费主题
	Handle(ctx context.Context, p Payload) (err error) // 处理过程的方法
}
```


## 一个例子

每个被发送到队列的消息应该被定义为一个单独的文件结构。

例如，如果您需要异步记录系统日志，内容大致如下：

- 文件路径：server/internal/queues/sys_opt_log_queue.go

```go 
// package queues
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package queues

import (
	"context"
	"encoding/json"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/library/worker"
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	RegisterQueueProcess(consts.QueueSysOptLog, SysOptLogQueue)
}

type qSysOperLog struct{}

var SysOptLogQueue = &qSysOperLog{}

func (q *qSysOperLog) GetTopic() string {
	return consts.QueueSysOptLog
}

func (q *qSysOperLog) Handle(ctx context.Context, p worker.Payload) error {
	if p.Payload == nil || q.GetTopic() != p.Group {
		return nil
	}
	var data entity.SysOperLog
	if err := json.Unmarshal(p.Payload, &data); err != nil {
		g.Log().Error(ctx, "qSysOperLog.Handle json.Unmarshal err.", err)
		return err
	}
	return service.SysOperLog().RealWrite(ctx, data)
}


```

下面是将消息添加到队列的方式，大概内容如下:

```go
package main

import (
	"fmt"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/queues"
	"xiuadmin/internal/model/entity"
)

func test()  {
	data := &entity.SysLog{
		//...
    }
	queues.Push(context.Background(), consts.QueueSysOptLog, data, 10)
}

```

