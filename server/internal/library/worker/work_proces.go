// package worker
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package worker

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
)

type Queue struct {
	topic string
	w     *Worker
}
type Process interface {
	GetTopic() string                                  // 获取消费主题
	Handle(ctx context.Context, p Payload) (err error) // 处理过程的方法
}

func RegisterQueueProcess(p Process) (q *Queue) {
	q = &Queue{
		topic: p.GetTopic(),
		w:     New(WithWorkerGroup(p.GetTopic()), WithWorkerHandler(p.Handle)),
	}
	return
}

func (q *Queue) Push(ctx context.Context, topic string, data []byte, timeout int) (err error) {
	err = q.w.Once(
		WithTaskUid(guid.S()),
		WithTaskGroup(topic),
		WithTaskPayload(data),
		WithTaskTimeout(timeout),
	)
	if err != nil {
		g.Log().Debug(ctx, "Run Queue TaskWorker %s Error: %v", topic, err)
	}
	return nil
}

type Scheduled struct {
	topic string
	w     *Worker
}

func RegisterScheduledProcess(p Process) (q *Queue) {
	q = &Queue{
		topic: p.GetTopic(),
		w:     New(WithWorkerGroup(p.GetTopic()), WithWorkerHandler(p.Handle)),
	}
	return
}
func (s *Scheduled) Cron(ctx context.Context, topic, cronExpr string, data []byte) (entryID string, err error) {
	s.topic = topic
	entryID, err = s.w.Cron(
		1,
		WithTaskUid(guid.S()),
		WithTaskGroup(topic),
		WithTaskExpr(cronExpr),
		WithTaskPayload(data),
	)
	if err != nil {
		g.Log().Debug(ctx, "Run Cron TaskWorker %s Error: %v", topic, err)
	}
	return
}

func (s *Scheduled) Remove(entryID string) (err error) {
	err = s.w.RemoveCron(entryID)
	return
}
