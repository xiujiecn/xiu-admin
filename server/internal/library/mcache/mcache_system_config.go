// Package mcache
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package mcache

import (
	"context"
	"fmt"
	"time"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/event"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/request"
	"xiuadmin/internal/service"
)

func init() {
	event.EventsInstance().Register([]string{consts.EventKeyDBSysConfigUpdate, consts.EventKeyDBSysConfigCreate, consts.EventKeyDBSysConfigDelete}, func(ctx context.Context, eventKey string, args ...interface{}) {
		if len(args) < 1 {
			return
		}
		// tenantId := args[0].(string)
		Instance().Remove(ctx, getKey(ctx))
	})
}
func getKey(ctx context.Context) string {
	return fmt.Sprintf(consts.MemCacheSystemConfig, contexts.GetTenantId(ctx))
}

func GetSystemConfig(ctx context.Context, key string, defaultValue string) (string, error) {
	mapV, err := Instance().Get(ctx, getKey(ctx))
	if err != nil {
		data, _, err := service.SysConfig().List(ctx, &model.SysConfigListParam{
			PageInfo: request.PageInfo{
				Page:     1,
				PageSize: 9999,
			},
		})
		if err != nil {
			return defaultValue, err
		}
		mapConfig := make(map[string]string)
		for _, item := range data {
			mapConfig[item.ConfigKey] = item.ConfigValue
		}
		Instance().Set(ctx, getKey(ctx), mapConfig, time.Hour*24)
		return mapConfig[key], nil
	}
	mapConfig := mapV.MapStrStr()
	if v, ok := mapConfig[key]; ok {
		return v, nil
	}
	return defaultValue, nil
}

func SetSystemConfig(ctx context.Context, key string, value string) error {
	mapV, err := Instance().Get(ctx, getKey(ctx))
	if err != nil {
		return err
	}
	mapConfig := mapV.MapStrStr()
	mapConfig[key] = value
	return Instance().Set(ctx, getKey(ctx), mapConfig, time.Hour*24)
}

func RemoveSystemConfig(ctx context.Context, key string) error {
	mapV, err := Instance().Get(ctx, getKey(ctx))
	if err != nil {
		return err
	}
	mapConfig := mapV.MapStrStr()
	delete(mapConfig, key)
	return Instance().Set(ctx, getKey(ctx), mapConfig, time.Hour*24)
}
