package gfredis

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/redis/go-redis/v9"
)

func (r *Redis) Do(ctx context.Context, args ...interface{}) (*g.Var, error) {
	cmd := r.rdb.Do(ctx, args...)
	if cmd.Err() == redis.Nil {
		return g.NewVar(nil), nil
	}
	return g.NewVar(cmd.Val()), cmd.Err()
}
