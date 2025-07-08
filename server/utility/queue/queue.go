package queue

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
)

var ErrNotFoundTopic = errors.New("not found topic")

type Payload struct {
	Ts       int64           `json:"ts"`        // 毫秒时间戳
	TraceId  string          `json:"trace_id"`  // 链路ID
	Topic    string          `json:"topic"`     // 消费主题 队列名
	Data     json.RawMessage `json:"data"`      // 数据
	Expire   int64           `json:"expire"`    // 过期时长毫秒 0:不过期 业务内自行处理
	ClientId string          `json:"client_id"` // 客户端ID
}

type Config struct {
	RedisName string `json:"redis_name"` // redis名称
	Workers   int    `json:"workers"`    // 工作线程数
	ClientId  string `json:"client_id"`  // 客户端ID
}

type Process interface {
	GetTopic() string                                  // 获取消费主题
	GetType() QueueType                                // 获取类型
	Handle(ctx context.Context, p Payload) (err error) // 处理过程的方法
}

type QueueType string

const (
	QueueTypeMemory         QueueType = "memory_queue"
	QueueTypeRedis          QueueType = "redis_queue"
	QueueTypeRedisBroadcast QueueType = "redis_broadcast"
)

type sProducer interface {
	GetType() QueueType                                                                 // 获取类型
	GetTopic() string                                                                   // 获取主题
	Push(ctx context.Context, topic string, data interface{}, expire int64) (err error) // 推送消息
}

type sConsumer interface {
	GetType() QueueType
	GetTopic() string                       // 获取消费主题
	Handle(ctx context.Context) (err error) // 处理过程的方法
}

var (
	rwLock   sync.RWMutex
	producer = make(map[string]sProducer)
	consumer = make(map[string]sConsumer)
)

func NewProducer(ctx context.Context, t QueueType, topic string, config *Config) error {
	if _, ok := producer[topic]; ok {
		return errors.New("producer already exists")
	}
	var p sProducer
	switch t {
	case QueueTypeMemory:
		p = RegisterMemoryQueue(ctx, topic, nil, config)
	case QueueTypeRedis:
		p = RegisterRedisQueue(ctx, topic, nil, config)
	case QueueTypeRedisBroadcast:
		p = RegisterRedisBroadcast(ctx, topic, nil, config)
	}
	rwLock.Lock()
	producer[topic] = p
	rwLock.Unlock()
	return nil
}

func NewConsumer(ctx context.Context, p Process, config *Config) error {
	if _, ok := consumer[p.GetTopic()+"__"+config.ClientId]; ok {
		return errors.New("consumer already exists")
	}
	var c sConsumer
	switch p.GetType() {
	case QueueTypeMemory:
		c = RegisterMemoryQueue(ctx, p.GetTopic(), p, config)
	case QueueTypeRedis:
		c = RegisterRedisQueue(ctx, p.GetTopic(), p, config)
	case QueueTypeRedisBroadcast:
		c = RegisterRedisBroadcast(ctx, p.GetTopic(), p, config)
	}
	rwLock.Lock()
	consumer[p.GetTopic()+"__"+config.ClientId] = c
	rwLock.Unlock()
	return nil
}

func UnregisterProducer(ctx context.Context, topic string) error {
	rwLock.Lock()
	defer rwLock.Unlock()
	p, ok := producer[topic]
	if !ok {
		return errors.New("producer not found")
	}
	switch p.GetType() {
	case QueueTypeMemory:
		UnregisterMemoryQueue(ctx, topic)
	case QueueTypeRedis:
		UnregisterRedisQueue(ctx, topic)
	case QueueTypeRedisBroadcast:
		UnregisterRedisBroadcast(ctx, topic, "")
	}
	delete(producer, topic)
	return nil
}

func UnregisterConsumer(ctx context.Context, topic string, groupId string) error {
	rwLock.Lock()
	defer rwLock.Unlock()
	c, ok := consumer[topic+"__"+groupId]
	if !ok {
		return ErrNotFoundTopic
	}
	switch c.GetType() {
	case QueueTypeMemory:
		UnregisterMemoryQueue(ctx, topic)
	case QueueTypeRedis:
		UnregisterRedisQueue(ctx, topic)
	case QueueTypeRedisBroadcast:
		UnregisterRedisBroadcast(ctx, topic, groupId)
	}
	delete(consumer, topic+"__"+groupId)
	return nil
}

func ProducerTopicExists(topic string) bool {
	rwLock.RLock()
	_, ok := producer[topic]
	rwLock.RUnlock()
	return ok
}

func ConsumerTopicExists(topic string) bool {
	rwLock.RLock()
	_, ok := consumer[topic]
	rwLock.RUnlock()
	return ok
}

func Push(ctx context.Context, topic string, data interface{}) (err error) {
	rwLock.RLock()
	p, ok := producer[topic]
	rwLock.RUnlock()
	if !ok {
		g.Log().Errorf(ctx, "topic not found. topic: %s", topic)
		return ErrNotFoundTopic
	}
	return p.Push(ctx, topic, data, 0)
}

func PushWithExpire(ctx context.Context, topic string, data interface{}, expire int64) (err error) {
	rwLock.RLock()
	p, ok := producer[topic]
	rwLock.RUnlock()
	if !ok {
		g.Log().Errorf(ctx, "topic not found. topic: %s", topic)
		return ErrNotFoundTopic
	}
	return p.Push(ctx, topic, data, expire)
}
