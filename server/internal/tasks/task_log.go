package tasks

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type TaskLog struct {
	ctx context.Context
}

func (l *TaskLog) Info(msg string, keysAndValues ...interface{}) {
	g.Log().Infof(l.ctx, "%s, %v", msg, keysAndValues)
}

func (l *TaskLog) Error(err error, msg string, keysAndValues ...interface{}) {
	g.Log().Errorf(l.ctx, "%s, %v", msg, keysAndValues)
}
