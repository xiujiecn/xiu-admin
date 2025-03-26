package bcache

import (
	"context"
)

var (
	// 用户退出清理缓存绑定关系
	userLogoutClearCacheMap = map[string]func(ctx context.Context, userId int64) error{}
	// 用户变动清理缓存绑定关系
	userChangeClearCacheMap = map[string]func(ctx context.Context, userId int64) error{}
)
