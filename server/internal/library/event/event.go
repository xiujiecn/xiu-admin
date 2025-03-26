package event

import (
	"context"
	"sync"
)

type EventCallbackFunc func(ctx context.Context, eventKey string, args ...interface{})

type sEvent struct {
	sync.Mutex
	eventList map[string][]EventCallbackFunc // 事件列表 [事件Key]:回调函数
}

var (
	instance *sEvent
	once     sync.Once
)

// EventsInstance 事件实例
func EventsInstance() *sEvent {
	once.Do(func() {
		instance = &sEvent{
			eventList: make(map[string][]EventCallbackFunc),
		}
	})
	return instance
}

// Register 注册事件
func (e *sEvent) Register(eventKey interface{}, callback EventCallbackFunc) {
	e.Lock()
	defer e.Unlock()
	eventKeyStr, ok1 := eventKey.(string)
	eventKeyList, ok2 := eventKey.([]string)
	if !ok1 && !ok2 {
		return
	}
	if ok1 {
		e.eventList[eventKeyStr] = append(e.eventList[eventKeyStr], callback)
	} else {
		for _, eventKey := range eventKeyList {
			e.eventList[eventKey] = append(e.eventList[eventKey], callback)
		}
	}
}

// Call 回调事件
func (e *sEvent) Call(ctx context.Context, eventKey string, args ...interface{}) {
	if events, ok := e.eventList[eventKey]; ok {
		for _, f := range events {
			f(ctx, eventKey, args...)
		}
	}
}

// Emit 触发事件
func (e *sEvent) Emit(ctx context.Context, eventKey string, args ...interface{}) {
	e.Call(ctx, eventKey, args...)
}

// Remove 删除事件
func (e *sEvent) Remove(eventKey string) {
	delete(e.eventList, eventKey)
}

// Clear 清空事件列表
func (e *sEvent) Clear() {
	e.eventList = make(map[string][]EventCallbackFunc)
}
