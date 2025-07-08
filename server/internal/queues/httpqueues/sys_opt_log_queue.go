// package queues
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package httpqueues

import (
	"context"
	"encoding/json"
	"xiuadmin/internal/cmd/inithttp"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/service"

	"xiuadmin/utility/queue"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	// memqueue.Register(gctx.New(), SysOptLogQueue)
	inithttp.RegisterHttpInitFunc("init "+consts.QueueSysOptLog, func(ctx context.Context) error {
		queue.NewConsumer(gctx.New(), SysOptLogQueue, &queue.Config{
			Workers: 100,
		})
		queue.NewProducer(gctx.New(), queue.QueueTypeMemory, consts.QueueSysOptLog, &queue.Config{})
		return nil
	})
}

type qSysOperLog struct{}

var SysOptLogQueue = &qSysOperLog{}

func (q *qSysOperLog) GetTopic() string {
	return consts.QueueSysOptLog
}

func (q *qSysOperLog) GetType() queue.QueueType {
	return queue.QueueTypeMemory
}

func (q *qSysOperLog) Handle(ctx context.Context, p queue.Payload) error {
	var data entity.SysOperLog
	if err := json.Unmarshal(p.Data, &data); err != nil {
		g.Log().Error(ctx, "qSysOperLog.Handle json.Unmarshal err.", err)
		return err
	}
	return service.SysOperLog().RealWrite(ctx, data)
}
