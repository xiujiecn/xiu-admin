package gfredis

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/redis/go-redis/v9"
)

func (r *Redis) HMSet(ctx context.Context, key string, mp map[string]interface{}) error {
	mp2 := make(map[string]string)
	for k, v := range mp {
		mp2[k] = gconv.String(v)
	}
	cmd := r.rdb.HMSet(ctx, key, mp2)
	return cmd.Err()
}
func (r *Redis) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	cmd := r.rdb.HGetAll(ctx, key)
	return cmd.Val(), cmd.Err()
}

func (r *Redis) HGet(ctx context.Context, key, field string) (string, error) {
	cmd := r.rdb.HGet(ctx, key, field)
	if cmd.Err() == redis.Nil {
		return "", nil
	}
	return cmd.Val(), cmd.Err()
}

func (r *Redis) HSet(ctx context.Context, key string, mp map[string]interface{}) error {
	mp2 := make(map[string]string)
	for k, v := range mp {
		mp2[k] = gconv.String(v)
	}
	cmd := r.rdb.HSet(ctx, key, mp2)
	return cmd.Err()
}

func (r *Redis) HSetNX(ctx context.Context, key, field string, value interface{}) (bool, error) {
	cmd := r.rdb.HSetNX(ctx, key, field, gconv.String(value))
	return cmd.Val(), cmd.Err()
}
func (r *Redis) HDel(ctx context.Context, key string, field string) (int64, error) {
	cmd := r.rdb.HDel(ctx, key, field)
	return cmd.Val(), cmd.Err()
}
func (r *Redis) HLen(ctx context.Context, key string) (int64, error) {
	cmd := r.rdb.HLen(ctx, key)
	return cmd.Val(), cmd.Err()
}
