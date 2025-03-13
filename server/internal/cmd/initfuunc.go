package cmd

import (
	"context"
	"xiujieadmin/internal/queues"
)

func InitSystemDeferFunc(ctx context.Context) error {
	queues.Run(ctx)
	return nil
}
