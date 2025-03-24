package system_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"xiujieadmin/internal/consts"
	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/library/worker"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/tasks"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	// 数据库只读，不执行
	// g.DB().SetDryRun(true)
	g.DB().SetDebug(true)
	testing.Init()
}

func TestTask_runOnce(t *testing.T) {
	// 设置上下文用户身份为用户
	ctx := context.WithValue(gctx.New(), consts.ContextKey, &model.Context{
		// 为了测试只设置了hook中需要用到的数据
		User: &model.Identity{
			BaseClaims: model.BaseClaims{
				ID:       1,
				TenantId: "000000",
				DeptId:   1,
			},
		},
	})

	exists := &model.SysJobViewModel{}
	dao.SysJob.Ctx(ctx).Where(dao.SysJob.Columns().JobId, 4).Scan(&exists)
	paramArr, _ := tasks.TasksInstance().ParseParameters(exists.JobParams)
	taskData := tasks.Task{
		ID:         fmt.Sprintf("%s-job-%d", exists.InvokeTarget, exists.JobId),
		TaskType:   "Type-" + gconv.String(exists.MisfirePolicy),
		MethodName: exists.InvokeTarget,
		Params:     paramArr,
		Explain:    exists.JobName,
	}
	runPayload, _ := json.Marshal(taskData)

	tasks.TasksInstance().Once(
		worker.WithTaskCtx(context.Background()),
		worker.WithTaskUid(taskData.ID),           // 任务ID
		worker.WithTaskGroup(taskData.MethodName), // 任务组
		worker.WithTaskTimeout(10),
		worker.WithTaskNow(true),
		worker.WithTaskReplace(true),
		worker.WithTaskPayload(runPayload),
	)
	fmt.Println("UpdateUserPassword success")
}
