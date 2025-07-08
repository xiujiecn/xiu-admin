package queue

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
)

type tRedisBroadcast struct {
	p         Process
	RedisName string
	isRunning bool
	grpool    *grpool.Pool
	clientId  string
	conn      gredis.Conn
}

var (
	redisBroadcastRwLock sync.RWMutex
	redisBroadcastQueues = make(map[string]*tRedisBroadcast)
)

func RegisterRedisBroadcast(ctx context.Context, topic string, p Process, config *Config) *tRedisBroadcast {
	redisBroadcastRwLock.Lock()
	defer redisBroadcastRwLock.Unlock()
	q, ok := redisBroadcastQueues[topic+"__"+config.ClientId]
	if ok {
		q.p = p
		q.RedisName = config.RedisName
		q.isRunning = true
		q.clientId = config.ClientId
	} else {
		q = &tRedisBroadcast{
			p:         p,
			RedisName: config.RedisName,
			isRunning: true,
			clientId:  config.ClientId,
		}
		redisBroadcastQueues[topic+"__"+config.ClientId] = q
	}
	workers := 1
	if config != nil && config.Workers > 0 {
		workers = config.Workers
	}
	if p != nil {
		q.p = p
		q.grpool = grpool.New(workers)

		q.conn, _ = q.getConn(ctx)
		go q.Handle(ctx)
	}
	return q
}

func UnregisterRedisBroadcast(ctx context.Context, topic string, clientId string) {
	redisBroadcastRwLock.Lock()
	defer redisBroadcastRwLock.Unlock()
	q, ok := redisBroadcastQueues[topic+"__"+clientId]
	if !ok {
		return
	}
	q.isRunning = false
	delete(redisBroadcastQueues, topic+"__"+clientId)
}

func (q *tRedisBroadcast) GetType() QueueType {
	return QueueTypeRedis
}

func (q *tRedisBroadcast) GetTopic() string {
	return q.p.GetTopic()
}

func (q *tRedisBroadcast) Push(ctx context.Context, topic string, data interface{}, expire int64) (err error) {
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
	num, err := g.Redis(q.RedisName).Publish(ctx, topic, payloadJson)
	if err != nil {
		return err
	}
	g.Log().Debugf(ctx, "redisqueue Push %s %d", topic, num)
	return
}

func (q *tRedisBroadcast) getConn(ctx context.Context) (conn gredis.Conn, err error) {
	conn, err = g.Redis(q.RedisName).Conn(ctx)
	if err != nil || conn == nil {
		g.Log().Errorf(ctx, "tRedisBroadcast.Handle Conn %s error: %v", q.p.GetTopic(), err)
		return
	}
	g.Log().Debugf(ctx, "tRedisBroadcast.Handle 获取连接 %s", q.p.GetTopic())
	subs, err := conn.Subscribe(ctx, q.p.GetTopic())
	if err != nil {
		g.Log().Errorf(ctx, "tRedisBroadcast.Handle Subscribe %s error: %v", q.p.GetTopic(), err)
		conn.Close(ctx)
		conn = nil
		return
	}
	g.Log().Debugf(ctx, "tRedisBroadcast.Handle Subscribe %s 成功, sub:%+v", q.p.GetTopic(), subs)
	return
}

func (q *tRedisBroadcast) Handle(ctx context.Context) (err error) {
	g.Log().Debugf(ctx, "tRedisBroadcast.Handle 开始处理 %s, group_id: %s", q.p.GetTopic(), q.clientId)
	topic := q.p.GetTopic()

	for {
		if !q.isRunning {
			break
		}
		if q.conn == nil {
			q.conn, err = q.getConn(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "tRedisBroadcast.Handle getConn %s error: %v", topic, err)
				time.Sleep(time.Second)
				continue
			}
		}
		data, err := q.conn.ReceiveMessage(ctx)
		if err != nil {
			g.Log().Errorf(ctx, "tRedisBroadcast.Handle ReceiveMessage %s error: %v", topic, err)
			q.conn.Close(ctx)
			q.conn = nil
			continue
		}
		var payload Payload
		err = json.Unmarshal([]byte(data.Payload), &payload)
		if err != nil {
			g.Log().Errorf(ctx, "tRedisBroadcast.Handle unmarshal error: %v", data)
			continue
		}
		payload.ClientId = q.clientId
		subCtx := gctx.New()
		gtrace.WithTraceID(subCtx, payload.TraceId)
		err = q.grpool.Add(subCtx, func(ctx context.Context) {
			err = q.p.Handle(subCtx, payload)
			if err != nil {
				g.Log().Errorf(ctx, "tRedisBroadcast.Handle process error: %v, topic: %s, trace_id: %s, data: %v",
					err, q.p.GetTopic(), payload.TraceId, payload.Data)
				return
			}
		})
	}
	redisBroadcastRwLock.Lock()
	delete(redisBroadcastQueues, topic)
	redisBroadcastRwLock.Unlock()
	if q.conn != nil {
		q.conn.Close(ctx)
	}
	return
}
