// Package cmd
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package cmd

import (
	"context"
	"errors"
	"xiuadmin/internal/library/cache"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmode"
)

// 系统初始化
func SystemInit(ctx context.Context) {
	// 检查数据库连接
	if err := CheckConnDB(ctx); err != nil {
		panic(err)
	}
	// 设置gf运行模式
	SetGFMode(ctx)
	// 默认上海时区
	if err := gtime.SetTimeZone("Asia/Shanghai"); err != nil {
		g.Log().Fatalf(ctx, "时区设置异常 err: %+v", err)
		return
	}
	// 设置缓存适配器
	cache.SetAdapter(ctx)
}

// CheckConnDB 检查数据库连接
func CheckConnDB(ctx context.Context) error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("连接到数据库失败")
	}
	return nil
}

// 设置gf运行模式
func SetGFMode(ctx context.Context) {
	gfmode := g.Cfg().MustGet(ctx, "system.gfmode").String()
	if len(gfmode) == 0 {
		gfmode = gmode.NOT_SET
	}
	var modes = []string{gmode.DEVELOP, gmode.TESTING, gmode.STAGING, gmode.PRODUCT}
	for _, mode := range modes {
		if mode == gfmode {
			gmode.Set(mode)
			break
		}
	}
}

func DBHandle(ctx context.Context) {

}
