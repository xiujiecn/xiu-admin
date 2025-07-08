// // package system_test
// // @Link  https://github.com/xiujiecn/xiu-admin
// // @Copyright  Copyright (c) 2025 LiXiujie
// // @Author  Lxj <li@xiujie.cn>
// // @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system_test

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"testing"
// 	"time"
// 	"xiuadmin/internal/consts"
// 	"xiuadmin/internal/dao"
// 	"xiuadmin/internal/library/worker"
// 	"xiuadmin/internal/model"
// 	"xiuadmin/internal/tasks"

// 	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
// 	"github.com/gogf/gf/v2/frame/g"
// 	"github.com/gogf/gf/v2/os/gctx"
// 	"github.com/gogf/gf/v2/util/gconv"
// )

// func init() {
// 	// 数据库只读，不执行
// 	// g.DB().SetDryRun(true)
// 	g.DB().SetDebug(true)
// 	testing.Init()
// }

// func TestTask_runOnce(t *testing.T) {
// 	// 设置上下文用户身份为用户
// 	ctx := context.WithValue(gctx.New(), consts.ContextKey, &model.Context{
// 		// 为了测试只设置了hook中需要用到的数据
// 		User: &model.Identity{
// 			BaseClaims: model.BaseClaims{
// 				ID:       1,
// 				TenantId: "000000",
// 				DeptId:   1,
// 			},
// 		},
// 	})

// 	exists := &model.SysJobViewModel{}
// 	dao.SysJob.Ctx(ctx).Where(dao.SysJob.Columns().JobId, 4).Scan(&exists)
// 	paramArr, _ := tasks.TasksInstance().ParseParameters(exists.JobParams)
// 	taskData := tasks.Task{
// 		ID:         fmt.Sprintf("%s-job-%d", exists.InvokeTarget, exists.JobId),
// 		TaskType:   "Type-" + gconv.String(exists.MisfirePolicy),
// 		MethodName: exists.InvokeTarget,
// 		Params:     paramArr,
// 		Explain:    exists.JobName,
// 	}
// 	runPayload, _ := json.Marshal(taskData)

// 	tasks.TasksInstance().Once(
// 		worker.WithTaskCtx(context.Background()),
// 		worker.WithTaskUid(taskData.ID),           // 任务ID
// 		worker.WithTaskGroup(taskData.MethodName), // 任务组
// 		worker.WithTaskTimeout(10),
// 		worker.WithTaskNow(true),
// 		worker.WithTaskReplace(true),
// 		worker.WithTaskPayload(runPayload),
// 	)
// }

// func TestTask_schedule(t *testing.T) {
// 	// 设置上下文用户身份为用户
// 	ctx := context.WithValue(gctx.New(), consts.ContextKey, &model.Context{
// 		// 为了测试只设置了hook中需要用到的数据
// 		User: &model.Identity{
// 			BaseClaims: model.BaseClaims{
// 				ID:       1,
// 				TenantId: "000000",
// 				DeptId:   1,
// 			},
// 		},
// 	})

// 	exists := &model.SysJobViewModel{}
// 	dao.SysJob.Ctx(ctx).Where(dao.SysJob.Columns().JobId, 4).Scan(&exists)
// 	paramArr, _ := tasks.TasksInstance().ParseParameters(exists.JobParams)
// 	//获取task目录下是否绑定对应的方法
// 	actualName, exist := tasks.TasksInstance().CheckFuncName(exists.InvokeTarget)
// 	if !exist {
// 		errInfo := fmt.Sprintf("没有绑定对应的方法:%s", exists.InvokeTarget)
// 		g.Log().Error(ctx, errInfo)
// 	}

// 	//传参解析
// 	paramArr, err := tasks.TasksInstance().ParseParameters(exists.JobParams)
// 	if err != nil {
// 		g.Log().Error(ctx, err)
// 	}

// 	taskData := tasks.Task{
// 		ID:         fmt.Sprintf("%s-job-%d", exists.InvokeTarget, exists.JobId),
// 		TaskType:   "Type-" + gconv.String(exists.MisfirePolicy),
// 		MethodName: actualName,
// 		Params:     paramArr,
// 		Explain:    exists.JobName,
// 	}
// 	runPayload, _ := json.Marshal(taskData)
// 	_, err2 := tasks.TasksInstance().Cron(
// 		1,
// 		worker.WithTaskCtx(context.Background()),
// 		worker.WithTaskUid(taskData.ID), // 任务ID
// 		worker.WithTaskGroup(taskData.MethodName), // 任务组
// 		worker.WithTaskExpr("0/2 * * * * *"),
// 		//worker.WithTaskExpr("@every 2s"),
// 		worker.WithTaskTimeout(10),
// 		worker.WithTaskReplace(true),
// 		worker.WithTaskPayload(runPayload),
// 	)
// 	if err2 != nil {
// 		g.Log().Debug(ctx, taskData.MethodName, taskData.Explain, "启动任务失败")
// 	}

// 	time.Sleep(10 * time.Minute)
// }

// func TestTask_remove(t *testing.T) {
// 	// 设置上下文用户身份为用户
// 	ctx := context.WithValue(gctx.New(), consts.ContextKey, &model.Context{
// 		// 为了测试只设置了hook中需要用到的数据
// 		User: &model.Identity{
// 			BaseClaims: model.BaseClaims{
// 				ID:       1,
// 				TenantId: "000000",
// 				DeptId:   1,
// 			},
// 		},
// 	})

// 	exists := &model.SysJobViewModel{}
// 	dao.SysJob.Ctx(ctx).Where(dao.SysJob.Columns().JobId, 4).Scan(&exists)
// 	taskJobId := fmt.Sprintf("%s-job-%d", exists.InvokeTarget, exists.JobId)
// 	g.Log().Debug(ctx, taskJobId)
// 	err := tasks.TasksInstance().Remove(taskJobId)
// 	if err != nil {
// 		g.Log().Debug(ctx, exists.JobName, "启动停止失败")
// 		g.Log().Error(ctx, err)
// 	}
// }

// func TestTask_getTask(t *testing.T) {
// 	ctx := context.WithValue(gctx.New(), consts.ContextKey, &model.Context{
// 		// 为了测试只设置了hook中需要用到的数据
// 		User: &model.Identity{
// 			BaseClaims: model.BaseClaims{
// 				ID:       1,
// 				TenantId: "000000",
// 				DeptId:   1,
// 			},
// 		},
// 	})

// 	exists := &model.SysJobViewModel{}
// 	dao.SysJob.Ctx(ctx).Where(dao.SysJob.Columns().JobId, 4).Scan(&exists)
// 	taskJobId := fmt.Sprintf("%s-job-%d", exists.InvokeTarget, exists.JobId)
// 	task, err := tasks.TasksInstance().GetTask(taskJobId)
// 	if err != nil {
// 		g.Log().Debug(ctx, exists.JobName, "获取任务失败")
// 		g.Log().Error(ctx, err)
// 	}
// 	fmt.Println("task", task)
// }
