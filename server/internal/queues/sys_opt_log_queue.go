// package queues
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
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
