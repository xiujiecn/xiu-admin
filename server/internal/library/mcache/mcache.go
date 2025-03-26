// 内存缓存管理
package mcache

import (
	"context"
	"sync"
	"time"

	"github.com/gogf/gf/v2/os/gcache"
)

// cache 缓存驱动
var (
	mcache *gcache.Cache
	mu     sync.Mutex
	// 用户退出清理缓存绑定关系
	userLogoutClearCacheMap = map[string]func(ctx context.Context, userId int64) error{}
	// 用户变动清理缓存绑定关系
	userChangeClearCacheMap = map[string]func(ctx context.Context, userId int64) error{}
)

// Instance 缓存实例
func Instance() *gcache.Cache {
	if mcache == nil {
		mu.Lock()
		if mcache == nil {
			mcache = gcache.NewWithAdapter(gcache.NewAdapterMemory())
		}
		mu.Unlock()
	}
	return mcache
}

func Set(ctx context.Context, key string, value interface{}, timeout time.Duration) error {
	return Instance().Set(ctx, key, value, timeout)
}

func Get(ctx context.Context, key string) (interface{}, error) {
	return Instance().Get(ctx, key)
}

func Remove(ctx context.Context, key string) error {
	_, err := Instance().Remove(ctx, key)
	return err
}

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
