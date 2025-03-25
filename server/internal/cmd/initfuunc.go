package cmd

import (
	"context"
	"xiujieadmin/internal/library/xgen"
	"xiujieadmin/internal/queues"
	"xiujieadmin/internal/service"
	"xiujieadmin/internal/tasks"

	"github.com/gogf/gf/v2/util/gmode"
)

func InitSystemDeferFunc(ctx context.Context) error {
	tasks.TasksInstance()
	queues.Run(ctx)
	InitMonitor()
	service.SysJob().InitRegister()
	// 初始化生成代码配置
	if !gmode.IsProduct() {
		xgen.Init(ctx)
	}
	return nil
}
