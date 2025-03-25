package cmd

import (
	"context"
	"xiujieadmin/internal/queues"
	"xiujieadmin/internal/tasks"
)

func InitSystemDeferFunc(ctx context.Context) error {
	tasks.TasksInstance()
	queues.Run(ctx)
	InitMonitor()
	return nil
}
