package gfredis

import (
	"context"
)

func (r *Redis) ZPopByScore(ctx context.Context, key string, min, max string, limit int64) (res []string, err error) {
	var luaScript = `
	local members = redis.call('ZRANGEBYSCORE', KEYS[1], ARGV[1], ARGV[2], 'LIMIT', 0, ARGV[3])
	if #members > 0 then
		redis.call('ZREM', KEYS[1], unpack(members))
	end
	return members
	`
	cmd := r.rdb.Eval(ctx, luaScript, []string{key}, min, max, limit)
	if err = cmd.Err(); err != nil {
		return nil, err
	}
	res, err = cmd.StringSlice()
	return
}
