package queues

import (
	"context"
	"encoding/json"
	"xiujieadmin/internal/consts"
	"xiujieadmin/internal/library/worker"
	"xiujieadmin/internal/model/entity"
	"xiujieadmin/internal/service"

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
