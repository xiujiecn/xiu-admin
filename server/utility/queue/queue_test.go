package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

type testMemoryConsumer struct {
}

func (c *testMemoryConsumer) GetTopic() string {
	return "test"
}

func (c *testMemoryConsumer) GetType() QueueType {
	return QueueTypeMemory
}

func (c *testMemoryConsumer) Handle(ctx context.Context, p Payload) error {
	fmt.Println("testMemoryConsumer", p)
	return nil
}

type testRedisConsumer struct {
	QueueType QueueType
}

func (c *testRedisConsumer) GetTopic() string {
	return "test"
}

func (c *testRedisConsumer) GetType() QueueType {
	return c.QueueType
}

func (c *testRedisConsumer) Handle(ctx context.Context, p Payload) error {
	fmt.Printf("testRedisConsumer %+v \n", p)
	var data string
	json.Unmarshal(p.Data, &data)
	fmt.Printf("testRedisConsumer %+v \n", data)

	time.Sleep(1 * time.Second)
	return nil
}

func TestQueue(t *testing.T) {
	ctx := context.Background()
	NewProducer(ctx, QueueTypeMemory, "test", nil)
	NewConsumer(ctx, &testMemoryConsumer{}, nil)
	Push(ctx, "test", "content1")
	Push(ctx, "test", "content2")
	Push(ctx, "test", "content3")
	time.Sleep(1 * time.Second)
}

func TestRedisQueue(t *testing.T) {
	ctx := gctx.New()

	NewProducer(ctx, QueueTypeRedis, "test", &Config{
		RedisName: "default",
	})
	NewConsumer(ctx, &testRedisConsumer{QueueType: QueueTypeRedis}, &Config{
		RedisName: "default",
	})
	time.Sleep(2 * time.Second)
	Push(ctx, "test", "content1")
	Push(ctx, "test", "content2")
	Push(ctx, "test", "content3")
	time.Sleep(5 * time.Second)
}

func TestRedisBroadcast(t *testing.T) {
	ctx := gctx.New()
	NewProducer(ctx, QueueTypeRedisBroadcast, "test", &Config{
		RedisName: "default",
	})
	consumerHandle := &testRedisConsumer{QueueType: QueueTypeRedisBroadcast}
	NewConsumer(ctx, consumerHandle, &Config{
		RedisName: "default",
		ClientId:  "group1",
	})
	NewConsumer(ctx, consumerHandle, &Config{
		RedisName: "default",
		ClientId:  "group2",
	})
	Push(ctx, "test", "content1")
	Push(ctx, "test", "content2")
	Push(ctx, "test", "content3")
	time.Sleep(50 * time.Second)
	UnregisterConsumer(ctx, "test", "group1")
	UnregisterConsumer(ctx, "test", "group2")
	UnregisterProducer(ctx, "test")
	time.Sleep(5 * time.Second)
}
