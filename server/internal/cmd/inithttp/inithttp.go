package inithttp

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	httpInitFunc    map[string]func(ctx context.Context) error
	httpCleanupFunc map[string]func(ctx context.Context) error
	rwLock          sync.RWMutex
)

func RegisterHttpInitFunc(name string, fn func(ctx context.Context) error) {
	rwLock.Lock()
	defer rwLock.Unlock()
	if httpInitFunc == nil {
		httpInitFunc = make(map[string]func(ctx context.Context) error)
	}
	httpInitFunc[name] = fn
}

func RegisterHttpCleanupFunc(name string, fn func(ctx context.Context) error) {
	rwLock.Lock()
	defer rwLock.Unlock()
	if httpCleanupFunc == nil {
		httpCleanupFunc = make(map[string]func(ctx context.Context) error)
	}
	httpCleanupFunc[name] = fn
}

func InitHttp(ctx context.Context) error {
	rwLock.RLock()
	defer rwLock.RUnlock()
	for name, fn := range httpInitFunc {
		if err := fn(ctx); err != nil {
			g.Log().Errorf(ctx, "httpInitFunc %s error: %v", name, err)
			return err
		} else {
			g.Log().Infof(ctx, "httpInitFunc %s success", name)
		}
	}
	return nil
}

func CleanupHttp(ctx context.Context) error {
	rwLock.RLock()
	defer rwLock.RUnlock()
	for name, fn := range httpCleanupFunc {
		if err := fn(ctx); err != nil {
			g.Log().Errorf(ctx, "httpCleanupFunc %s error: %v", name, err)
		} else {
			g.Log().Infof(ctx, "httpCleanupFunc %s success", name)
		}
	}
	return nil
}
