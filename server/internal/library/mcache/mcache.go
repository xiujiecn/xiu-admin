// Package mcache
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
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
