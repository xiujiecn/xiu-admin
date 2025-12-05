package gfredis

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/redis/go-redis/v9"
)

func (r *Redis) Subscribe(ctx context.Context, channels ...string) (*redis.PubSub, error) {
	cmd := r.rdb.Subscribe(ctx, channels...)
	return cmd, nil
}

func (r *Redis) Publish(ctx context.Context, channel string, message interface{}) (int64, error) {
	cmd := r.rdb.Publish(ctx, channel, gconv.String(message))
	return cmd.Val(), cmd.Err()
}
