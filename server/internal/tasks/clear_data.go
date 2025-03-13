package tasks

import (
	"context"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	RegisterInnerTask("clear_data", "ClearOperationLogByDays")
}

func (t Task) ClearOperationLogByDays(days string) {
	ctx := context.Background()
	glog.Debugf(ctx, "执行任务：清理超过%v天的操作日志", days)
	err := service.SysOperLog().ClearOperationLogByDays(ctx, gconv.Int(days))
	if err != nil {
		glog.Error(ctx, "TaskJob.ClearOperationLogByDays", err)
	}
}
