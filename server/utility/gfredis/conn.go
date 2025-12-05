package gfredis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func (r *Redis) Conn(ctx context.Context) (*redis.Conn, error) {
	return r.rdb.Conn(), nil
}
