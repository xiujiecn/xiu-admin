package tasks

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"go.opentelemetry.io/otel"
)

// Task 任务结构体 内置任务
type Task struct {
	ID         string        //任务ID
	TaskType   string        //任务类型
	MethodName string        //方法名
	Params     []interface{} //参数
	Explain    string        //任务描述
}

type TaskOptions struct {
	Uid        string
	Payload    []byte
	Expr       *string         // 仅周期任务
	In         *time.Duration  // 仅一次性任务 当前时间的延迟
	At         *time.Time      // 仅一次性任务 指定时间
	Now        *bool           // 仅一次性任务 当前时间
	Replace    bool            // 仅一次性任务 是否替换已存在的任务
	Ctx        context.Context // 仅一次性任务
	MaxRetry   int
	Timeout    int // 限制并发堵塞超时时间
	Concurrent int // 0不限制并发 1限制并发
}
type taskInfo struct {
	entryID int
	option  TaskOptions
	run     bool
}

type Tasks struct {
	cron    *cron.Cron
	lock    sync.Mutex
	taskMap map[string]*taskInfo
}

var (
	instance *Tasks
	once     sync.Once
)

func TasksInstance(ctx context.Context) *Tasks {
	once.Do(func() {
		instance, _ = NewTasks(ctx) // Initialize only once
	})
	return instance
}

// Once 注册一个任务以运行一次。
func (tk *Tasks) Once(ctx context.Context, taskOptions TaskOptions) error {
	// return tk.worker.Once(options...)
	info := &taskInfo{
		entryID: 0,
		option:  taskOptions,
		run:     false,
	}
	spec := ""
	if taskOptions.Expr != nil && *taskOptions.Expr != "" {
		spec = *taskOptions.Expr
	} else if taskOptions.In != nil {
		spec = fmt.Sprintf("@every %s", *taskOptions.In)
	} else if taskOptions.At != nil {
		spec = fmt.Sprintf("@every %s", time.Until(*taskOptions.At).String())
	} else if taskOptions.Now != nil && *taskOptions.Now {
		go func() {
			tk.process(ctx, info)
		}()
		g.Log().Infof(ctx, "Tasks.Once end. info: %+v", info)
		return nil
	} else {
		return errors.New("expr is required")
	}
	tk.lock.Lock()
	tk.taskMap[taskOptions.Uid] = info
	tk.lock.Unlock()
	entryID, _ := tk.cron.AddFunc(spec, func() {
		tk.process(ctx, info)
		tk.lock.Lock()
		delete(tk.taskMap, taskOptions.Uid)
		tk.lock.Unlock()
	})
	tk.lock.Lock()
	info, ok := tk.taskMap[taskOptions.Uid]
	if ok {
		info.entryID = int(entryID)
	}
	tk.lock.Unlock()
	g.Log().Infof(ctx, "Tasks.Once end. info: %+v", info)
	return nil
}

// 注册一个任务以由cron表达式运行。
func (tk *Tasks) Cron(ctx context.Context, taskOptions TaskOptions) (id string, err error) {
	spec := ""
	if taskOptions.Expr != nil && *taskOptions.Expr != "" {
		spec = *taskOptions.Expr
	} else {
		return "", errors.New("expr is required")
	}
	tk.lock.Lock()
	tk.taskMap[taskOptions.Uid] = &taskInfo{
		entryID: 0,
		option:  taskOptions,
		run:     false,
	}
	tk.lock.Unlock()
	entryID, _ := tk.cron.AddFunc(spec, func() {
		tk.lock.Lock()
		info, ok := tk.taskMap[taskOptions.Uid]
		g.Log().Infof(ctx, "Tasks.Cron AddFunc func begin. uid: %s, spec: %s, ok: %v", taskOptions.Uid, spec, ok)
		if !ok {
			tk.lock.Unlock()
			return
		}
		tk.lock.Unlock()
		start := time.Now()
		for {
			if taskOptions.Timeout > 0 && time.Since(start) > time.Second*time.Duration(taskOptions.Timeout) {
				g.Log().Infof(ctx, "Tasks.Cron AddFunc func timeout. uid: %s, spec: %s, timeout: %d, start: %s", taskOptions.Uid, spec, taskOptions.Timeout, start.Format("2006-01-02 15:04:05"))
				return
			}
			tk.lock.Lock()
			if info.run && info.option.Concurrent == 1 {
				tk.lock.Unlock()
				time.Sleep(time.Second * 1)
				continue
			}
			info.run = true
			tk.lock.Unlock()
			break
		}
		tk.process(ctx, info)
		tk.lock.Lock()
		info.run = false
		tk.lock.Unlock()
	})
	tk.lock.Lock()
	info, ok := tk.taskMap[taskOptions.Uid]
	if ok {
		info.entryID = int(entryID)
	}
	tk.lock.Unlock()
	nextTime := tk.cron.Entry(entryID).Next
	nextTimeStr := nextTime.Format("2006-01-02 15:04:05")
	sub := nextTime.Sub(time.Now())
	g.Log().Infof(ctx, "Tasks.Cron end.	entryID: %d, uid: %s, spec: %s, nextTime:%s, sub:%v", info.entryID, info.option.Uid, spec, nextTimeStr, sub)
	return gconv.String(info.entryID), nil
}

// 从任务队列中删除一个任务。
func (tk *Tasks) Remove(uid string) error {
	tk.lock.Lock()
	info, ok := tk.taskMap[uid]
	if ok {
		tk.cron.Remove(cron.EntryID(info.entryID))
		delete(tk.taskMap, uid)
	}
	tk.lock.Unlock()
	return nil
}

func (tk *Tasks) RemoveCron(entryID string) error {
	tk.lock.Lock()
	for _, info := range tk.taskMap {
		if gconv.String(info.entryID) == entryID {
			tk.cron.Remove(cron.EntryID(info.entryID))
			delete(tk.taskMap, info.option.Uid)
			break
		}
	}
	tk.lock.Unlock()
	return nil
}

func (tk Tasks) GetTask(uid string) (task *Task, err error) {
	tk.lock.Lock()
	info, ok := tk.taskMap[uid]
	if !ok {
		return nil, errors.New("task not found")
	}
	tk.lock.Unlock()

	taskRun := &Task{}

	err = json.Unmarshal(info.option.Payload, &taskRun)
	if err != nil {
		return nil, err
	}
	return taskRun, nil
}

func NewTasks(ctx context.Context) (tk *Tasks, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = errors.Errorf("%v", e)
		}
	}()
	// logger := &TaskLog{ctx: ctx}
	c := cron.New(cron.WithLocation(time.Local),
		cron.WithSeconds(),
		// cron.WithChain(cron.SkipIfStillRunning(logger)),
	)
	// w := worker.New(
	// 	worker.WithWorkerHandler(func(ctx context.Context, p worker.Payload) error {
	// 		return process(task{
	// 			ctx:     ctx,
	// 			payload: p,
	// 		})
	// 	}),
	// 	worker.WithWorkerRedisPeriodKey("TaskSchedule"),
	// )
	// if w.Error != nil {
	// 	err = errors.WithMessage(w.Error, "initialize worker failed")
	// 	return
	// }
	c.Start()
	tk = &Tasks{
		cron:    c,
		taskMap: make(map[string]*taskInfo),
	}
	g.Log().Debug(ctx, "NewTasks initialize tasks success")
	return
}

func (tk *Tasks) process(ctx context.Context, t *taskInfo) (err error) {
	g.Log().Debug(ctx, "Tasks.process payload", t.option.Payload)
	tr := otel.Tracer("task")
	_, span := tr.Start(ctx, "Task")
	defer span.End()
	var taskRun Task
	err = json.Unmarshal(t.option.Payload, &taskRun)
	if err != nil {
		return err
	}
	g.Log().Debug(ctx, "Tasks.process task", taskRun)
	err = CallMethod(ctx, &taskRun)
	if err != nil {
		fmt.Println("CallMethod err:", err.Error())
		return err
	}
	g.Log().Debug(ctx, "Tasks.process end. task", taskRun)
	return
}

// CallMethod 调用任务的实际方法
func CallMethod(ctx context.Context, task *Task) (err error) {
	// 获取TaskModel的反射值
	taskValue := reflect.ValueOf(task)
	// 准备要调用的方法的参数
	var args []reflect.Value
	for _, param := range task.Params {
		if !reflect.ValueOf(param).IsValid() {
			err = errors.New("invalid parameter")
			return
		}
		args = append(args, reflect.ValueOf(param))
	}
	// 查找并执行方法
	method := taskValue.MethodByName(task.MethodName)
	if method.IsValid() {
		if method.Type().NumIn() == 0 {
			method.Call(nil)
			return
		} else if method.Type().NumIn() != len(args) {
			err = errors.New("incorrect number of parameters")
			return
		}
		g.Log().Debug(context.Background(), "执行的任务：", task.MethodName)
		for i := 0; i < method.Type().NumIn(); i++ {
			switch method.Type().In(i).Kind() {
			case reflect.Float32:
				args[i] = reflect.ValueOf(gconv.Float32(task.Params[i]))
			case reflect.Float64:
				args[i] = reflect.ValueOf(gconv.Float64(task.Params[i]))
			case reflect.Uint64:
				args[i] = reflect.ValueOf(gconv.Uint64(task.Params[i]))
			case reflect.Int64:
				args[i] = reflect.ValueOf(gconv.Int64(task.Params[i]))
			case reflect.Int:
				args[i] = reflect.ValueOf(gconv.Int(task.Params[i]))
			case reflect.Int32:
				args[i] = reflect.ValueOf(gconv.Int32(task.Params[i]))
			case reflect.Int16:
				args[i] = reflect.ValueOf(gconv.Int16(task.Params[i]))
			case reflect.Int8:
				args[i] = reflect.ValueOf(gconv.Int8(task.Params[i]))
			case reflect.Uint:
				args[i] = reflect.ValueOf(gconv.Uint(task.Params[i]))
			case reflect.Uint32:
				args[i] = reflect.ValueOf(gconv.Uint32(task.Params[i]))
			case reflect.Uint16:
				args[i] = reflect.ValueOf(gconv.Uint16(task.Params[i]))
			case reflect.Uint8:
				args[i] = reflect.ValueOf(gconv.Uint8(task.Params[i]))
			case reflect.String:
				args[i] = reflect.ValueOf(gconv.String(task.Params[i]))
			case reflect.Bool:
				args[i] = reflect.ValueOf(gconv.Bool(task.Params[i]))
			default:
				args[i] = reflect.ValueOf(task.Params[i])
			}
		}
		if len(task.Params) > 0 {
			method.Call(args)
		} else {
			method.Call(nil)
		}
	} else {
		errInfo := fmt.Sprintf("Method not found: %s", task.MethodName)
		err = errors.New(errInfo)
	}
	return
}

// UnmarshalTask 解析TaskJob，使用自定义JSON解析来处理类型问题
func UnmarshalTask(data []byte) (*Task, error) {
	var task Task
	dec := json.NewDecoder(bytes.NewReader(data))
	// 使用json.Decoder逐个解析Token，自定义解析逻辑
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		// 根据Token类型处理
		switch v := t.(type) {
		case float64:
			// 如果是数字 ，检查是否可以转换为整数
			if v == float64(int64(v)) {
				// 如果可以转换为整数，则转换为int类型
				task.Params = append(task.Params, int64(v))
			} else {
				// 如果不能转换为整数，则保留为float64类型
				task.Params = append(task.Params, v)
			}
		case string, bool, nil:
			// 对于字符串、布尔和空值，直接添加到Params中
			task.Params = append(task.Params, v)
			// 可以添加更多的类型处理
		}
	}
	return &task, nil
}

// CheckFuncName 检查方法名是否存在
func (tk *Tasks) CheckFuncName(funcName string) (actualfuncName string, exists bool) {
	exist := GetInnerTaskName(funcName)
	g.Log().Debug(context.Background(), "CheckFuncName", exist)
	if exist != "" {
		return exist, true
	}
	return
}

// ParseParameters 解析参数
func (tk *Tasks) ParseParameters(parseData string) (params []interface{}, err error) {
	parts := strings.Split(parseData, "|")
	params = make([]interface{}, len(parts))
	for i, part := range parts {
		params[i] = part
	}
	return
}
