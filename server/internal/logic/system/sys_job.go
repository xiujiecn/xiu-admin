package system

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"xiujieadmin/internal/consts"
	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/library/contexts"
	"xiujieadmin/internal/library/worker"
	"xiujieadmin/internal/library/xgorm/handler"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/entity"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/service"
	"xiujieadmin/internal/tasks"
)

func init() {
	service.RegisterSysJob(NewSysJob())
}

type sSysJob struct{}

func NewSysJob() *sSysJob {
	return &sSysJob{}
}

func (s *sSysJob) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: false,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.SysJob.Ctx(ctx), option...)
}

func (c *sSysJob) List(ctx context.Context, query *model.SysJobListParam, pageInfo *request.PageInfo) (Data []*model.SysJobListModel, total int, err error) {
	m := c.Model(ctx)
	if query.JobName != "" {
		m = m.WhereLike(dao.SysJob.Columns().JobName, "%"+strings.Trim(query.JobName, " ")+"%")
	}
	if query.JobGroup != "" {
		m = m.Where(dao.SysJob.Columns().JobGroup, query.JobGroup)
	}
	if query.Status != "" {
		m = m.Where(dao.SysJob.Columns().Status, query.Status)
	}
	t, err := m.Count()
	if err != nil {
		return nil, 0, err
	}
	Data = make([]*model.SysJobListModel, 0)
	if err := m.Page(pageInfo.Page, pageInfo.PageSize).Order(dao.SysJob.Columns().JobId, "DESC").Scan(&Data); err != nil {
		return nil, 0, err
	}
	return Data, t, nil
}

func (c *sSysJob) View(ctx context.Context, jobId int64) (Data *model.SysJobViewModel, err error) {
	m := c.Model(ctx)
	Data = &model.SysJobViewModel{}
	if err := m.Where(dao.SysJob.Columns().JobId, jobId).Scan(&Data); err != nil {
		return nil, err
	}

	return Data, nil
}

func (c *sSysJob) Add(ctx context.Context, jobAdd *model.SysJobAddModel) (LastInsertId int64, err error) {
	claims := contexts.GetUser(ctx)
	if claims == nil {
		return 0, gerror.NewCode(gcode.CodeMissingParameter, "用户信息不存在")
	}

	//获取task目录下是否绑定对应的方法
	_, exist := tasks.TasksInstance().CheckFuncName(jobAdd.InvokeTarget)
	if !exist {
		errInfo := fmt.Sprintf("没有绑定对应的方法:%s", jobAdd.InvokeTarget)
		return 0, gerror.New(errInfo)
	}

	jobAdd.CreatedDept = claims.BaseClaims.DeptId
	jobAdd.CreatedAt = gtime.Now()
	jobAdd.CreatedBy = claims.BaseClaims.ID
	jobAdd.UpdatedAt = gtime.Now()
	jobAdd.UpdatedBy = claims.BaseClaims.ID

	result, err := c.Model(ctx).Data(jobAdd).Insert()
	if err != nil {
		return 0, err
	}
	lastInserId, _ := result.LastInsertId()
	if lastInserId > 0 && jobAdd.Status == consts.SysJobStatusNormal {
		job := &model.SysJobViewModel{}
		if err := c.Model(ctx).Where(dao.SysJob.Columns().JobId, lastInserId).Scan(&job); err != nil {
			return lastInserId, err
		}
		//任务启动
		err = c.taskRun(ctx, &job.SysJob)
		if err != nil {
			return lastInserId, err
		}
	}

	return result.LastInsertId()
}

func (c *sSysJob) Update(ctx context.Context, jobUpdate *model.SysJobUpdateModel) (RowsAffected int64, err error) {
	claims := contexts.GetUser(ctx)
	if claims == nil {
		return 0, gerror.NewCode(gcode.CodeMissingParameter, "用户信息不存在")
	}
	//获取task目录下是否绑定对应的方法
	_, exist := tasks.TasksInstance().CheckFuncName(jobUpdate.InvokeTarget)
	if !exist {
		errInfo := fmt.Sprintf("没有绑定对应的方法:%s", jobUpdate.InvokeTarget)
		return 0, gerror.New(errInfo)
	}

	jobUpdate.UpdatedAt = gtime.Now()
	jobUpdate.UpdatedBy = claims.BaseClaims.ID
	result, err := c.Model(ctx).Where(dao.SysJob.Columns().JobId, jobUpdate.JobId).Update(jobUpdate)

	if err != nil {
		return 0, err
	}
	row, err := result.RowsAffected()
	if row > 0 && jobUpdate.Status == consts.SysJobStatusNormal {
		job := &model.SysJobViewModel{}
		if err := c.Model(ctx).Where(dao.SysJob.Columns().JobId, jobUpdate.JobId).Scan(&job); err != nil {
			return 0, err
		}

		//任务启动
		err = c.taskRestart(ctx, &job.SysJob)
		if err != nil {
			return row, err
		}
	}
	return row, err
}

func (c *sSysJob) UpdateStatus(ctx context.Context, jobUpdate *model.SysJobUpdateStatusModel) (RowsAffected int64, err error) {
	claims := contexts.GetUser(ctx)
	if claims == nil {
		return 0, gerror.NewCode(gcode.CodeMissingParameter, "用户信息不存在")
	}

	exists := &model.SysJobViewModel{}
	if err := c.Model(ctx).Where(dao.SysJob.Columns().JobId, jobUpdate.JobId).Scan(&exists); err != nil {
		return 0, err
	}

	if jobUpdate.Status == consts.SysJobStatusNormal {
		//任务启动
		err = c.taskRun(ctx, &exists.SysJob)
		if err != nil {
			return 0, err
		}
	}
	if jobUpdate.Status == consts.SysJobStatusDisable {
		//任务停止
		if err = c.taskStop(ctx, &exists.SysJob); err != nil {
			return 0, err
		}
	}
	jobUpdate.UpdatedAt = gtime.Now()
	jobUpdate.UpdatedBy = claims.BaseClaims.ID
	result, err := c.Model(ctx).Where(dao.SysJob.Columns().JobId, jobUpdate.JobId).Update(jobUpdate)

	if err != nil {
		return 0, err
	}
	row, err := result.RowsAffected()
	return row, err
}

// 任务重启
func (c *sSysJob) taskRestart(ctx context.Context, job *entity.SysJob) (err error) {
	//任务停止
	if err = c.taskStop(ctx, job); err != nil {
		return
	}

	//任务启动
	err = c.taskRun(ctx, job)
	if err != nil {
		return
	}
	return
}

// 任务启动
func (c *sSysJob) taskRun(ctx context.Context, job *entity.SysJob) error {
	//获取task目录下是否绑定对应的方法
	actualName, exist := tasks.TasksInstance().CheckFuncName(job.InvokeTarget)
	if !exist {
		errInfo := fmt.Sprintf("没有绑定对应的方法:%s", job.InvokeTarget)
		return gerror.New(errInfo)
	}

	//传参解析
	paramArr, err := tasks.TasksInstance().ParseParameters(job.JobParams)
	if err != nil {
		return err
	}

	taskData := tasks.Task{
		ID:         fmt.Sprintf("%s-job-%d", job.InvokeTarget, job.JobId),
		TaskType:   "Type-" + gconv.String(job.MisfirePolicy),
		MethodName: actualName,
		Params:     paramArr,
		Explain:    job.JobName,
	}
	runPayload, _ := json.Marshal(taskData)
	if job.MisfirePolicy == consts.SysJobMisfirePolicyMulty {
		entryID, err := tasks.TasksInstance().Cron(
			job.Concurrent,
			worker.WithTaskCtx(context.Background()),
			worker.WithTaskUid(taskData.ID),           // 任务ID
			worker.WithTaskGroup(taskData.MethodName), // 任务组
			worker.WithTaskExpr(job.CronExpression),
			worker.WithTaskTimeout(10),
			worker.WithTaskReplace(true),
			worker.WithTaskPayload(runPayload),
		)
		if err != nil {
			g.Log().Debug(ctx, taskData.MethodName, taskData.Explain, "启动任务失败")
			return err
		}
		g.Log().Debug(ctx, taskData.MethodName, taskData.Explain, "启动任务成功", "entryID", entryID)
	} else {
		err := tasks.TasksInstance().Once(
			worker.WithTaskCtx(context.Background()),
			worker.WithTaskUid(taskData.ID),           // 任务ID
			worker.WithTaskGroup(taskData.MethodName), // 任务组
			worker.WithTaskTimeout(10),
			worker.WithTaskNow(true),
			worker.WithTaskReplace(true),
			worker.WithTaskPayload(runPayload),
		)
		if err != nil {
			g.Log().Debug(ctx, taskData.MethodName, taskData.Explain, "启动任务失败")
			return err
		}
	}

	return nil
}

func (c *sSysJob) taskStop(ctx context.Context, job *entity.SysJob) error {
	//获取task目录下是否绑定对应的方法
	_, exist := tasks.TasksInstance().CheckFuncName(job.InvokeTarget)
	if !exist {
		errInfo := fmt.Sprintf("没有绑定对应的方法:%s", job.InvokeTarget)
		return gerror.New(errInfo)
	}

	taskJobId := fmt.Sprintf("%s-job-%d", job.InvokeTarget, job.JobId)

	err := tasks.TasksInstance().Remove(taskJobId)
	if err != nil {
		g.Log().Debug(ctx, job.JobName, "启动停止失败")
		return err
	}
	return nil
}

func (c *sSysJob) Delete(ctx context.Context, jobDelete *model.SysJobDeleteModel) (RowsAffected int64, err error) {
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return 0, err
	}

	exists := &[]*model.SysJobViewModel{}
	c.Model(ctx).WhereIn(dao.SysJob.Columns().JobId, jobDelete.JobIds).Where(dao.SysJob.Columns().Status, consts.SysJobStatusNormal).Scan(exists)
	if len(*exists) > 0 {
		return 0, gerror.New("存在正常运行的任务，请先停止后再删除")
	}

	jobDelete.DeletedAt = gtime.Now()
	jobDelete.DeletedBy = claims.BaseClaims.ID
	result, err := c.Model(ctx).WhereIn(dao.SysJob.Columns().JobId, jobDelete.JobIds).Update(map[string]interface{}{
		dao.SysJob.Columns().DeletedAt: jobDelete.DeletedAt,
		dao.SysJob.Columns().DeletedBy: jobDelete.DeletedBy})

	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (c *sSysJob) Exec(ctx context.Context, jobId int64) error {
	exists := &model.SysJobViewModel{}
	err := c.Model(ctx).Where(dao.SysJob.Columns().JobId, jobId).Scan(&exists)
	if err != nil {
		return err
	}

	actualName, exist := tasks.TasksInstance().CheckFuncName(exists.InvokeTarget)
	if !exist {
		errInfo := fmt.Sprintf("没有绑定对应的方法:%s", exists.InvokeTarget)
		return gerror.New(errInfo)
	}

	//传参解析
	paramArr, err := tasks.TasksInstance().ParseParameters(exists.JobParams)
	if err != nil {
		g.Log().Error(ctx, "sSysJob.JobRun worker.ParseParameters", err)
		return err
	}

	taskData := tasks.Task{
		ID:         fmt.Sprintf("%s-job-%d", exists.InvokeTarget, exists.JobId),
		TaskType:   "Type-" + gconv.String(exists.MisfirePolicy),
		MethodName: actualName,
		Params:     paramArr,
		Explain:    exists.JobName,
	}
	runPayload, _ := json.Marshal(taskData)

	err = tasks.TasksInstance().Once(
		worker.WithTaskCtx(context.Background()),
		worker.WithTaskUid(taskData.ID),           // 任务ID
		worker.WithTaskGroup(taskData.MethodName), // 任务组
		worker.WithTaskTimeout(10),
		worker.WithTaskNow(true),
		worker.WithTaskReplace(true),
		worker.WithTaskPayload(runPayload),
	)

	if err != nil {
		errInfo := fmt.Sprintf(exists.InvokeTarget, taskData.Explain, "启动任务失败")
		return errors.New(errInfo)
	}
	return nil
}

// 查询所有的状态正常的任务列表，并进行初始注册
func (c *sSysJob) InitRegister() error {
	g.Log().Debug(context.Background(), "初始化任务注册")
	jobs := &[]*model.SysJobViewModel{}
	err := c.Model(context.Background()).Where(dao.SysJob.Columns().Status, consts.SysJobStatusNormal).Scan(jobs)
	if err != nil {
		return err
	}
	for _, job := range *jobs {
		err = c.taskRun(context.Background(), &job.SysJob)
		if err != nil {
			g.Log().Error(context.Background(), "初始化注册任务 %s 失败", job.JobName, err)
			continue
		}
		g.Log().Debug(context.Background(), "初始化注册任务 %s 成功", job.JobName)
	}
	g.Log().Debug(context.Background(), "初始化任务注册完成")
	return nil
}
