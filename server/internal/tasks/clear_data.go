package tasks

import (
	"context"
	"time"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	RegisterInnerTask("clear_operation_log_by_days", "ClearOperationLogByDays")
	RegisterInnerTask("test", "Test")
}

func (t Task) ClearOperationLogByDays(days string) {
	ctx := context.Background()
	g.Log().Debugf(ctx, "执行任务：清理超过%v天的操作日志", days)
	err := service.SysOperLog().ClearOperationLogByDays(ctx, gconv.Int(days))
	if err != nil {
		g.Log().Error(ctx, "TaskJob.ClearOperationLogByDays", err)
	}
	time.Sleep(time.Second * 30)
	g.Log().Debugf(ctx, "执行任务：完成")
}

func (t Task) Test() {
	ctx := context.Background()
	// 测试代码
	g.Log().Debugf(ctx, "执行任务：测试")
}
