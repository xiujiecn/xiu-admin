# 定时任务

> 在实际的项目开发中，定时任务几乎成为不可或缺的一部分。XiuAdmin为定时任务提供一个方便的后台操作界面，让您能够轻松地进行在线启停、修改和立即执行等操作。这样的设计可以极大地改善您在使用定时任务过程中的体验，让整个过程更加顺畅、高效。


## 实现接口
- 为了提供高度的扩展性，定时任务在设计上采用了注册回调的思路。只需要实现接口，注册回调函数，从而实现更大的灵活性和可扩展性。


## 一个例子

定时任务的文件结构可以根据具体需要进行调整，以下是一个常见的参考结构：

- 文件路径：server/internal/tasks/clear_data.go

```go 
package tasks

import (
	"context"
	"time"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	RegisterInnerTask("clear_data", "ClearOperationLogByDays")
	RegisterInnerTask("test", "Test")
}

func (t Task) ClearOperationLogByDays(days string) {
	ctx := context.Background()
	glog.Debugf(ctx, "执行任务：清理超过%v天的操作日志", days)
	err := service.SysOperLog().ClearOperationLogByDays(ctx, gconv.Int(days))
	if err != nil {
		glog.Error(ctx, "TaskJob.ClearOperationLogByDays", err)
	}
	time.Sleep(time.Second * 30)
	glog.Debugf(ctx, "执行任务：完成")
}

func (t Task) Test() {
	ctx := context.Background()
	// 测试代码
	glog.Debugf(ctx, "执行任务：测试")
}

```

继续在后台 系统监控 -> 任务调度中心 -> 添加任务，填写的任务名称需要和上面的名称保持一致，再进行简单的策略配置以后，一个后台可控的定时任务就添加好了！


## 更多

定时任务源码路径：`server/internal/library/worker`

内部机制有`asynq`实现，可以参考：https://github.com/hibiken/asynq

秒级定时任务，修改了asynq源码，可以参考：https://github.com/xiujiecn/asynq
