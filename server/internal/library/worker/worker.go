// package worker
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package worker

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"xiuadmin/utility/uuid32"

	"github.com/dromara/carbon/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/xiujiecn/asynq"
)

type Worker struct {
	ops       *WorkerOptions
	client    *asynq.Client
	inspector *asynq.Inspector
	scheduler *asynq.Scheduler
	svr       *asynq.Server
	Error     error
}

type periodTask struct {
	Expr      string `json:"expr"` // asynq
	Group     string `json:"group"`
	Uid       string `json:"uid"`
	Payload   []byte `json:"payload"`
	Next      int64  `json:"next"`      // next schedule unix timestamp
	Processed int64  `json:"processed"` // run times
	MaxRetry  int    `json:"maxRetry"`
	Timeout   int    `json:"timeout"`
}

func (p *periodTask) String() (str string) {
	bs, _ := json.Marshal(p)
	str = string(bs)
	return
}

func (p *periodTask) FromString(str string) {
	err := json.Unmarshal([]byte(str), p)
	if err != nil {
		return
	}
	return
}

func aggregate(group string, tasks []*asynq.Task) *asynq.Task {
	// g.Log().Infof(context.Background(), "server/internal/library/worker/worker.go aggregate group %s, tasks %d", group, len(tasks))
	list := make([][]byte, 0)
	for _, t := range tasks {
		list = append(list, t.Payload())
	}
	payload, _ := json.Marshal(list)
	return asynq.NewTask(group+".aggregate", payload, asynq.TaskID(uuid32.S()))
}

func New(options ...func(*WorkerOptions)) *Worker {
	ops := GetDefaultWorkerOptions(nil)
	for _, op := range options {
		op(ops)
	}

	rs := asynq.RedisClientOpt{
		Addr:     ops.redisAddr,
		Password: ops.redisPassword,
		DB:       ops.redisDB,
	}

	w := &Worker{}
	if ops.redisAddr == "" {
		w.Error = errors.New("redisAddr is nil")
		return w
	}
	w.ops = ops
	w.client = asynq.NewClient(rs)
	w.inspector = asynq.NewInspector(rs)
	if ops.redisPeriodKey != "" {
		w.scheduler = asynq.NewScheduler(rs, &asynq.SchedulerOpts{
			Location: time.Now().Location(),
			LogLevel: asynq.LogLevel(g.Cfg().MustGet(context.Background(), "queue.asynq.logLevel").Int()), // 日志级别 DebugLevel 1, InfoLevel 2, WarnLevel 3, ErrorLevel 4, FatalLevel 5
		})
		go func() {
			if err := w.scheduler.Start(); err != nil {
				g.Log().Error(context.Background(), "server/internal/library/worker/worker.go New 启动调度器失败", err)
				panic(err)
			}
			g.Log().Info(context.Background(), "server/internal/library/worker/worker.go New 启动调度器成功", "group", w.ops.group)
		}()
	}
	conf := asynq.Config{
		Concurrency: g.Cfg().MustGet(context.Background(), "queue.asynq.concurrency").Int(),              // 最大同时执行的任务数量
		Queues:      map[string]int{ops.group: 10},                                                       // 队列名称和优先级
		LogLevel:    asynq.LogLevel(g.Cfg().MustGet(context.Background(), "queue.asynq.logLevel").Int()), // 日志级别 DebugLevel 1, InfoLevel 2, WarnLevel 3, ErrorLevel 4, FatalLevel 5
	}
	if w.ops.aggregator != nil {
		conf.GroupAggregator = asynq.GroupAggregatorFunc(aggregate)
		conf.GroupGracePeriod = time.Duration(w.ops.aggregator.GroupGracePeriod) * time.Second
		conf.GroupMaxDelay = time.Duration(w.ops.aggregator.GroupMaxDelay) * time.Second
		conf.GroupMaxSize = w.ops.aggregator.GroupMaxSize
		g.Log().Infof(context.Background(), "server/internal/library/worker/worker.go New 聚合器配置 group %s, gracePeriod %d, maxDelay %d, maxSize %d",
			w.ops.group, w.ops.aggregator.GroupGracePeriod, w.ops.aggregator.GroupMaxDelay, w.ops.aggregator.GroupMaxSize)
	}
	svr := asynq.NewServer(rs, conf)
	h := &taskHandlerBase{w: w}

	mux := asynq.NewServeMux()
	if w.ops.aggregator != nil {
		mux.HandleFunc(w.ops.group+".aggregate", h.ProcessTask)
	} else {
		mux.HandleFunc(w.ops.group+".once", h.ProcessTask)
		mux.HandleFunc(w.ops.group+".cron", h.ProcessTask)
	}

	if err := svr.Start(mux); err != nil {
		panic(err)
	}
	w.svr = svr
	if w.ops.clearArchived > 0 {
		go func() {
			time.Sleep(time.Duration(w.ops.clearArchived) * time.Second)
			w.clearArchived()
		}()
	}
	g.Log().Debug(context.Background(), "server/internal/library/worker/worker.go New", "group", w.ops.group)
	return w
}

func (w *Worker) Once(options ...func(*TaskOptions)) (err error) {
	ops := GetDefaultTaskOptions(nil)
	for _, op := range options {
		op(ops)
	}
	if ops.uid == "" {
		err = errors.New("uid is nil")
		return
	}
	t := asynq.NewTask(ops.group+".once", ops.payload, asynq.TaskID(ops.uid), asynq.Group(ops.group))
	taskOpts := []asynq.Option{
		asynq.Queue(w.ops.group),
		asynq.MaxRetry(w.ops.maxRetry),
		asynq.Timeout(time.Duration(ops.timeout) * time.Second),
	}
	if ops.maxRetry > 0 {
		taskOpts = append(taskOpts, asynq.MaxRetry(ops.maxRetry))
	}
	if ops.retention > 0 {
		taskOpts = append(taskOpts, asynq.Retention(time.Duration(ops.retention)*time.Second))
	} else {
		taskOpts = append(taskOpts, asynq.Retention(time.Duration(w.ops.timeout)*time.Second))
	}
	if ops.in != nil {
		taskOpts = append(taskOpts, asynq.ProcessIn(*ops.in))
	} else if ops.at != nil {
		taskOpts = append(taskOpts, asynq.ProcessAt(*ops.at))
	} else if ops.now {
		taskOpts = append(taskOpts, asynq.ProcessIn(time.Millisecond))
	}
	_, err = w.client.Enqueue(t, taskOpts...)
	if ops.replace && errors.Is(err, asynq.ErrTaskIDConflict) {
		// 错误表示由于给定的任务ID已经存在，因此无法将该任务加入队列。这可能是因为任务已经存在，或者任务ID冲突。
		err = w.inspector.DeleteTask(w.ops.group, ops.uid)
		if err != nil {
			return
		}
		_, err = w.client.Enqueue(t, taskOpts...)
	}
	return
}

// Cron 注册一个cron任务, 如果concurrent为0,则不限制并发,否则限制并发
func (w *Worker) Cron(concurrent int, options ...func(*TaskOptions)) (entryID string, err error) {
	ops := GetDefaultTaskOptions(nil)
	for _, op := range options {
		op(ops)
	}
	if ops.expr == "" {
		err = errors.New("expr is nil")
		return
	}
	if concurrent == 0 {
		// 不限制并发, 不传递任务id，任务ID随机，可并发
		entryID, err = w.scheduler.Register(ops.expr, asynq.NewTask(ops.group+".cron", ops.payload), asynq.Queue(w.ops.group))
	} else {
		// 限制并发, 传递任务id，任务ID固定，不可并发
		entryID, err = w.scheduler.Register(ops.expr, asynq.NewTask(ops.group+".cron", ops.payload, asynq.TaskID(ops.uid)), asynq.Queue(w.ops.group))
	}
	if err != nil {
		return
	}
	g.Log().Debug(context.Background(), "server/internal/library/worker/worker.go Cron", "group", w.ops.group, "entryID", entryID, "expr", ops.expr, "concurrent", concurrent)
	return
}

func (w *Worker) Remove(uid string) (err error) {
	entries, err := w.inspector.SchedulerEntries()
	if err != nil {
		return
	}
	for _, entry := range entries {
		payload := string(entry.Task.Payload())
		g.Log().Debug(context.Background(), "server/internal/library/worker/worker.go Remove", "group", w.ops.group, "entryID", entry.ID, "payload", payload)
		if strings.Contains(string(entry.Task.Payload()), "\""+uid+"\"") {
			g.Log().Debug(context.Background(), "strings.Contains(string(entry.Task.Payload()), uid) ")
			err = w.scheduler.Unregister(entry.ID)
		}
	}
	return
}

func (w *Worker) GetTaskInfo(uid string) (task *asynq.TaskInfo, err error) {
	task, err = w.inspector.GetTaskInfo(w.ops.group, uid)
	if err != nil {
		return
	}
	return task, nil
}

func (w *Worker) RemoveCron(entryID string) (err error) {
	err = w.scheduler.Unregister(entryID)
	return
}

// 获取带有默认超时的上下文
func (w *Worker) getDefaultTimeoutCtx() context.Context {
	c, _ := context.WithTimeout(context.Background(), time.Duration(w.ops.timeout)*time.Second)
	return c
}

// 清除已归档任务
func (w *Worker) clearArchived() {
	list, err := w.inspector.ListArchivedTasks(w.ops.group, asynq.Page(1), asynq.PageSize(100))
	if err != nil {
		return
	}
	// ctx := w.getDefaultTimeoutCtx()
	for _, item := range list {
		last := carbon.CreateFromStdTime(item.LastFailedAt)
		if !last.IsZero() && item.Retried < item.MaxRetry {
			continue
		}
		uid := item.ID
		var flag bool
		if strings.HasSuffix(item.Type, ".cron") {
			if carbon.Now().Gt(last.AddMinutes(5)) {
				flag = true
			}
		} else {
			if carbon.Now().Gt(last.AddMinutes(5)) {
				flag = true
			}
		}
		if flag {
			err := w.inspector.DeleteTask(w.ops.group, uid)
			if err != nil {
				return
			}
		}
	}
}

type taskHandlerBase struct {
	w *Worker
}

func (h *taskHandlerBase) ProcessTask(ctx context.Context, task *asynq.Task) (err error) {
	var isAggregate bool
	if strings.Contains(task.Type(), ".aggregate") {
		isAggregate = true
	}
	group := strings.TrimSuffix(strings.TrimSuffix(strings.TrimSuffix(task.Type(), ".aggregate"), ".once"), ".cron")
	if h.w.ops.handler == nil {
		g.Log().Errorf(ctx, "server/internal/library/worker/worker.go ProcessTask handler is nil, group: %s, uid: %s, payload: %s", group, task.ResultWriter().TaskID(), string(task.Payload()))
		return errors.New("handler is nil")
	}
	err = h.w.ops.handler(ctx, Payload{
		IsAggregate: isAggregate,
		Group:       group,
		Uid:         task.ResultWriter().TaskID(),
		Payload:     task.Payload(),
	})

	return
}

func (w *Worker) Stop() {
	if w.scheduler != nil {
		w.scheduler.Shutdown()
	}
	if w.client != nil {
		w.client.Close()
	}
	if w.inspector != nil {
		w.inspector.Close()
	}
	if w.svr != nil {
		w.svr.Shutdown()
	}
}
