// package worker
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package worker

import (
	"context"

	"xiuadmin/utility/uuid32"

	"github.com/gogf/gf/v2/frame/g"
)

type Queue struct {
	topic string
	w     *Worker
}

type Process interface {
	GetTopic() string                                  // 获取消费主题
	Handle(ctx context.Context, p Payload) (err error) // 处理过程的方法
	GetAggregator() *AggregatorOptions                 // 获取聚合器配置，返回nil不使用聚合
}

func RegisterQueueProcess(p Process) (q *Queue) {
	q = &Queue{
		topic: p.GetTopic(),
		w:     New(WithWorkerGroup(p.GetTopic()), WithWorkerHandler(p.Handle), WithWorkerAggregator(p.GetAggregator())),
	}
	return
}

func (q *Queue) Push(ctx context.Context, topic string, data []byte, timeout int) (err error) {
	err = q.w.Once(
		WithTaskUid(uuid32.New()),
		WithTaskGroup(topic),
		WithTaskPayload(data),
		WithTaskTimeout(timeout),
	)
	if err != nil {
		g.Log().Debug(ctx, "Run Queue TaskWorker %s Error: %v", topic, err)
	}
	return nil
}

func (q *Queue) Stop() {
	if q.w != nil {
		q.w.Stop()
	}
}

type Scheduled struct {
	topic string
	w     *Worker
}

func RegisterScheduledProcess(p Process) (q *Queue) {
	q = &Queue{
		topic: p.GetTopic(),
		w:     New(WithWorkerGroup(p.GetTopic()), WithWorkerHandler(p.Handle), WithWorkerAggregator(p.GetAggregator())),
	}
	return
}
func (s *Scheduled) Cron(ctx context.Context, topic, cronExpr string, data []byte) (entryID string, err error) {
	s.topic = topic
	entryID, err = s.w.Cron(
		1,
		WithTaskUid(uuid32.S()),
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
