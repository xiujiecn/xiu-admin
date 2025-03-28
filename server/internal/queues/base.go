// package queues
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package queues

import (
	"context"
	"errors"
	"xiuadmin/internal/library/worker"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	processes = make(map[string]*worker.Process)
	queues    = make(map[string]*worker.Queue)
)

func RegisterQueueProcess(name string, process worker.Process) {
	processes[name] = &process
}

func GetQueue(name string) *worker.Queue {
	if queue, ok := queues[name]; ok {
		return queue
	}
	return nil
}

func Run(ctx context.Context) {
	for name, process := range processes {
		queues[name] = worker.RegisterQueueProcess(*process)
		g.Log().Infof(ctx, "queues[%s] register success", name)
	}
}

func Push(ctx context.Context, topic string, data []byte, timeout int) (err error) {
	queue := GetQueue(topic)
	if queue != nil {
		err = queue.Push(context.Background(), topic, data, timeout)
	} else {
		g.Log().Errorf(ctx, "server/internal/queues/base.go Push queue[%s] not found", topic)
		return errors.New("queue not found")
	}
	return
}
