package gfredis

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/redis/go-redis/v9"
)

func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	cmd := r.rdb.Set(ctx, key, gconv.String(value), expiration)
	return cmd.Err()
}

func (r *Redis) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	cmd := r.rdb.SetNX(ctx, key, gconv.String(value), expiration)
	return cmd.Val(), cmd.Err()
}

func (r *Redis) Get(ctx context.Context, key string) (*g.Var, error) {
	cmd := r.rdb.Get(ctx, key)
	if cmd.Err() == redis.Nil {
		return g.NewVar(nil), nil
	}
	return g.NewVar(cmd.Val()), cmd.Err()
}

func (r *Redis) UpdateExpire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	cmd := r.rdb.Expire(ctx, key, expiration)
	return cmd.Val(), cmd.Err()
}

func (r *Redis) Remove(ctx context.Context, key string) (bool, error) {
	cmd := r.rdb.Del(ctx, key)
	return cmd.Val() > 0, cmd.Err()
}
func (r *Redis) Del(ctx context.Context, key string) (int64, error) {
	cmd := r.rdb.Del(ctx, key)
	return cmd.Val(), cmd.Err()
}
