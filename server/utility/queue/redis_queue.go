package queue

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"sync"
	"time"

	"xiuadmin/utility/gfredis"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
)

type tRedisQueue struct {
	p         Process
	RedisName string
	isRunning bool
	grpool    *grpool.Pool
}

var (
	redisRwLock sync.RWMutex
	redisQueues = make(map[string]*tRedisQueue)
)

func RegisterRedisQueue(ctx context.Context, topic string, p Process, config *Config) *tRedisQueue {
	redisRwLock.Lock()
	defer redisRwLock.Unlock()
	q, ok := redisQueues[topic]
	if ok {
		if p != nil {
			q.p = p
		}
		q.RedisName = config.RedisName
		q.isRunning = true
	} else {
		q = &tRedisQueue{
			p:         p,
			RedisName: config.RedisName,
			isRunning: true,
		}
		redisQueues[topic] = q
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
	return q
}

func UnregisterRedisQueue(ctx context.Context, topic string) {
	redisRwLock.Lock()
	defer redisRwLock.Unlock()
	q, ok := redisQueues[topic]
	if !ok {
		return
	}
	q.isRunning = false
}

func (q *tRedisQueue) GetType() QueueType {
	return QueueTypeRedis
}

func (q *tRedisQueue) GetTopic() string {
	return q.p.GetTopic()
}

func (q *tRedisQueue) Push(ctx context.Context, topic string, data interface{}, expire int64) (err error) {
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
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	_, err = gfredis.Instance(q.RedisName).RPush(ctx, topic, payloadJson)
	if err != nil {
		return err
	}
	// g.Log().Debugf(ctx, "redisqueue Push %s %d", topic, num)
	return
}

func (q *tRedisQueue) Handle(ctx context.Context) (err error) {
	topic := q.p.GetTopic()
	redisName := q.RedisName
	for {
		if !q.isRunning {
			break
		}
		data, err := gfredis.Instance(redisName).BLPop(ctx, 1*time.Second, topic)
		if err != nil {
			g.Log().Errorf(ctx, "redisqueue Pop %s error: %v", topic, err)
			time.Sleep(2 * time.Second)
			continue
		}
		arr := data.Strings()
		if len(arr) < 2 {
			continue
		}
		// g.Log().Debugf(ctx, "redisqueue Pop %s %v", topic, arr)
		var payload Payload
		err = json.Unmarshal([]byte(arr[1]), &payload)
		if err != nil {
			g.Log().Errorf(ctx, "redisqueue process unmarshal error: %v, data: %v, topic: %s",
				err, data, topic)
			continue
		}
		if q.p == nil {
			g.Log().Errorf(ctx, "redisqueue process is nil, topic: %s, trace_id: %s, data: %v",
				topic, payload.TraceId, payload.Data)
			continue
		}
		subCtx := gctx.New()
		gtrace.WithTraceID(subCtx, payload.TraceId)
		err = q.grpool.Add(subCtx, func(ctx context.Context) {
			defer func() {
				if r := recover(); r != nil {
					g.Log().Errorf(ctx, "redisqueue process panic: %v, topic: %s, trace_id: %s, len(data): %d, hex(data): %v",
						r, q.p.GetTopic(), payload.TraceId, len(payload.Data), hex.EncodeToString(payload.Data))
				}
			}()
			err = q.p.Handle(subCtx, payload)
			if err != nil {
				g.Log().Errorf(ctx, "redisqueue process error: %v, topic: %s, trace_id: %s, len(data): %d, hex(data): %v",
					err, q.p.GetTopic(), payload.TraceId, len(payload.Data), hex.EncodeToString(payload.Data))
				return
			}
		})
	}
	redisRwLock.Lock()
	delete(redisQueues, topic)
	redisRwLock.Unlock()
	return
}
