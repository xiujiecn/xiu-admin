// Package worker
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package worker

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type Payload struct {
	Group   string `json:"group"`
	Uid     string `json:"uid"`
	Payload []byte `json:"payload"`
}

type WorkerOptions struct {
	group          string                                     //任务处理器的组名
	redisAddr      string                                     //redis地址
	redisUser      string                                     //redis用户名
	redisPassword  string                                     //redis密码
	redisDB        int                                        //redis数据库
	handler        func(ctx context.Context, p Payload) error //任务处理函数
	maxRetry       int                                        //任务最大重试次数
	clearArchived  int                                        //清除已归档任务的时间间隔
	timeout        int                                        //任务超时时间
	redisPeriodKey string                                     //redis周期任务键
}

func GetDefaultWorkerOptions(ops *WorkerOptions) *WorkerOptions {
	if ops != nil {
		return ops
	}
	addr := g.Cfg().MustGet(context.Background(), "queue.asynq.redis.address").String()
	pass := g.Cfg().MustGet(context.Background(), "queue.asynq.redis.pass").String()
	db := g.Cfg().MustGet(context.Background(), "queue.asynq.redis.db").Int()
	user := g.Cfg().MustGet(context.Background(), "queue.asynq.redis.user").String()
	if user == "" {
		user = "default"
	}
	return &WorkerOptions{
		group:         "task",
		redisAddr:     addr,
		redisPassword: pass,
		redisDB:       db,
		redisUser:     user,
	}
}

func WithWorkerGroup(group string) func(*WorkerOptions) {
	return func(o *WorkerOptions) {
		GetDefaultWorkerOptions(o).group = group
	}
}

func WithWorkerRedisAddr(addr string) func(*WorkerOptions) {
	return func(o *WorkerOptions) {
		GetDefaultWorkerOptions(o).redisAddr = addr
	}
}

func WithWorkerRedisUser(user string) func(*WorkerOptions) {
	return func(o *WorkerOptions) {
		GetDefaultWorkerOptions(o).redisUser = user
	}
}

func WithWorkerRedisPassword(pass string) func(*WorkerOptions) {
	return func(o *WorkerOptions) {
		GetDefaultWorkerOptions(o).redisPassword = pass
	}
}

func WithWorkerRedisDB(db int) func(*WorkerOptions) {
	return func(o *WorkerOptions) {
		GetDefaultWorkerOptions(o).redisDB = db
	}
}

func WithWorkerHandler(handler func(ctx context.Context, p Payload) error) func(*WorkerOptions) {
	return func(o *WorkerOptions) {
		GetDefaultWorkerOptions(o).handler = handler
	}
}

func WithWorkerMaxRetry(maxRetry int) func(*WorkerOptions) {
	return func(o *WorkerOptions) {
		GetDefaultWorkerOptions(o).maxRetry = maxRetry
	}
}

func WithWorkerClearArchived(clearArchived int) func(*WorkerOptions) {
	return func(o *WorkerOptions) {
		GetDefaultWorkerOptions(o).clearArchived = clearArchived
	}
}

func WithWorkerTimeout(timeout int) func(*WorkerOptions) {
	return func(o *WorkerOptions) {
		GetDefaultWorkerOptions(o).timeout = timeout
	}
}

func WithWorkerRedisPeriodKey(redisPeriodKey string) func(*WorkerOptions) {
	return func(o *WorkerOptions) {
		GetDefaultWorkerOptions(o).redisPeriodKey = redisPeriodKey
	}
}

type TaskOptions struct {
	uid       string
	group     string
	payload   []byte
	expr      string          // 仅周期任务
	in        *time.Duration  // 仅一次性任务 当前时间的延迟
	at        *time.Time      // 仅一次性任务 指定时间
	now       bool            // 仅一次性任务 当前时间
	retention int             // 仅一次性任务 任务保留时间 在这个时间内任务会被保留为已完成状态，过了这个时间后，任务将被自动删除。
	replace   bool            // 仅一次性任务 是否替换已存在的任务
	ctx       context.Context // 仅一次性任务
	maxRetry  int
	timeout   int
}

func GetDefaultTaskOptions(opt *TaskOptions) *TaskOptions {
	if opt != nil {
		return opt
	}
	return &TaskOptions{
		group:    "group",
		maxRetry: 3,
		timeout:  120,
	}
}

func WithTaskUid(uid string) func(*TaskOptions) {

	return func(o *TaskOptions) {
		GetDefaultTaskOptions(o).uid = uid
	}
}

func WithTaskGroup(group string) func(*TaskOptions) {
	return func(o *TaskOptions) {
		GetDefaultTaskOptions(o).group = group
	}
}

func WithTaskPayload(payload []byte) func(*TaskOptions) {
	return func(o *TaskOptions) {
		GetDefaultTaskOptions(o).payload = payload
	}
}

func WithTaskExpr(expr string) func(*TaskOptions) {
	return func(o *TaskOptions) {
		GetDefaultTaskOptions(o).expr = expr
	}
}

func WithTaskIn(in *time.Duration) func(*TaskOptions) {
	return func(o *TaskOptions) {
		GetDefaultTaskOptions(o).in = in
	}
}

func WithTaskAt(at *time.Time) func(*TaskOptions) {
	return func(o *TaskOptions) {
		GetDefaultTaskOptions(o).at = at
	}
}

func WithTaskNow(now bool) func(*TaskOptions) {
	return func(o *TaskOptions) {
		GetDefaultTaskOptions(o).now = now
	}
}

func WithTaskRetention(retention int) func(*TaskOptions) {
	return func(o *TaskOptions) {
		GetDefaultTaskOptions(o).retention = retention
	}
}

func WithTaskReplace(replace bool) func(*TaskOptions) {
	return func(o *TaskOptions) {
		GetDefaultTaskOptions(o).replace = replace
	}
}

func WithTaskCtx(ctx context.Context) func(*TaskOptions) {
	return func(o *TaskOptions) {
		GetDefaultTaskOptions(o).ctx = ctx
	}
}

func WithTaskMaxRetry(maxRetry int) func(*TaskOptions) {
	return func(o *TaskOptions) {
		GetDefaultTaskOptions(o).maxRetry = maxRetry
	}
}

func WithTaskTimeout(timeout int) func(*TaskOptions) {
	return func(o *TaskOptions) {
		GetDefaultTaskOptions(o).timeout = timeout
	}
}
