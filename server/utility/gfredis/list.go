package gfredis

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/redis/go-redis/v9"
)

func (r *Redis) LRange(ctx context.Context, key string, start, stop int64) (*g.Var, error) {
	cmd := r.rdb.LRange(ctx, key, start, stop)
	if cmd.Err() == redis.Nil {
		return g.NewVar(nil), nil
	}
	return g.NewVar(cmd.Val()), cmd.Err()
}
func (r *Redis) RPush(ctx context.Context, key string, value interface{}) (int64, error) {
	cmd := r.rdb.RPush(ctx, key, gconv.String(value))
	return cmd.Val(), cmd.Err()
}
func (r *Redis) LLen(ctx context.Context, key string) (int64, error) {
	cmd := r.rdb.LLen(ctx, key)
	return cmd.Val(), cmd.Err()
}
func (r *Redis) LTrim(ctx context.Context, key string, start, stop int64) error {
	cmd := r.rdb.LTrim(ctx, key, start, stop)
	return cmd.Err()
}
func (r *Redis) BLPop(ctx context.Context, timeout time.Duration, keys ...string) (*g.Var, error) {
	cmd := r.rdb.BLPop(ctx, timeout, keys...)
	if cmd.Err() == redis.Nil {
		return g.NewVar(nil), nil
	}
	return g.NewVar(cmd.Val()), cmd.Err()
}
