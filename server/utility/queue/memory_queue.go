package queue

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
)

type tMemoryQueue struct {
	q      *gqueue.Queue
	p      Process
	grpool *grpool.Pool
}

var (
	memRwLock sync.RWMutex
	memQueues = make(map[string]*tMemoryQueue)
)

func RegisterMemoryQueue(ctx context.Context, topic string, p Process, config *Config) *tMemoryQueue {
	memRwLock.Lock()
	defer memRwLock.Unlock()
	q, ok := memQueues[topic]
	if !ok {
		q = &tMemoryQueue{
			q: gqueue.New(),
			p: p,
		}
		memQueues[topic] = q
	}
	workers := 1
	if config != nil && config.Workers > 0 {
		workers = config.Workers
	}
	if p != nil {
		q.p = p
		q.grpool = grpool.New(workers)
		go q.Handle(ctx)
	}
	g.Log().Infof(ctx, "RegisterMemoryQueue %s success, workers: %d, process: %v", topic, workers, p)
	return q
}

func UnregisterMemoryQueue(ctx context.Context, topic string) {
	memRwLock.Lock()
	defer memRwLock.Unlock()
	q, ok := memQueues[topic]
	if !ok {
		return
	}
	q.q.Close()
	delete(memQueues, topic)
}

func (q *tMemoryQueue) GetType() QueueType {
	return QueueTypeMemory
}

func (q *tMemoryQueue) GetTopic() string {
	return q.p.GetTopic()
}

func (q *tMemoryQueue) Push(ctx context.Context, topic string, data interface{}, expire int64) (err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	payload := Payload{
		Ts:      time.Now().UnixMilli(),
		TraceId: gctx.CtxId(ctx),
		Topic:   topic,
		Data:    json.RawMessage(jsonData),
		Expire:  expire,
	}
	q.q.Push(payload)
	return
}

func (q *tMemoryQueue) Handle(ctx context.Context) (err error) {
	if q.grpool == nil {
		g.Log().Errorf(ctx, "grpool is nil")
		return errors.New("grpool is nil")
	}
	p := q.p
	if p == nil {
		g.Log().Errorf(ctx, "process is nil")
		return errors.New("process is nil")
	}

	for {

		data := q.q.Pop()
		if data == nil {
			break
		}
		payload, ok := data.(Payload)
		if !ok {
			g.Log().Errorf(ctx, "memqueue process unmarshal error: %v", data)
			continue
		}
		subCtx := gctx.New()
		gtrace.WithTraceID(subCtx, payload.TraceId)
		q.grpool.Add(ctx, func(ctx context.Context) {
			defer func() {
				if r := recover(); r != nil {
					g.Log().Errorf(ctx, "memqueue process panic: %v, topic: %s, trace_id: %s, hex(data): %v",
						r, q.p.GetTopic(), payload.TraceId, hex.EncodeToString(payload.Data))
				}
			}()
			err := p.Handle(subCtx, payload)
			if err != nil {
				g.Log().Errorf(ctx, "memqueue process error: %v, topic: %s, trace_id: %s, hex(data): %v",
					err, q.p.GetTopic(), payload.TraceId, hex.EncodeToString(payload.Data))
			}
		})
	}
	return
}
