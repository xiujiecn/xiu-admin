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

// RegisterUserChangeClearCache 注册用户变动清理缓存绑定关系
func RegisterUserChangeClearCache(key string, f func(ctx context.Context, userId int64) error) {
	userChangeClearCacheMap[key] = f
}

// UserChangeClearCache 用户变动清理缓存
func UserChangeClearCache(ctx context.Context, userId int64) error {
	for _, f := range userChangeClearCacheMap {
		f(ctx, userId)
	}
	return nil
}

// RegisterUserLogoutClearCache 注册用户退出清理缓存绑定关系
func RegisterUserLogoutClearCache(key string, f func(ctx context.Context, userId int64) error) {
	userLogoutClearCacheMap[key] = f
}

// UserLogoutClearCache 用户退出清理缓存
func UserLogoutClearCache(ctx context.Context, userId int64) error {
	for _, f := range userLogoutClearCacheMap {
		f(ctx, userId)
	}
	return nil
}
