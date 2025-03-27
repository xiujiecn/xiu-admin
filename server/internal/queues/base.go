// package queues
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package queues

import (
	"context"
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
