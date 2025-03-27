// Package cmd
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package cmd

import (
	"context"
	"xiuadmin/internal/library/xgen"
	"xiuadmin/internal/queues"
	"xiuadmin/internal/service"
	"xiuadmin/internal/tasks"

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
