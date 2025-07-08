// Package cmd
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package cmd

import (
	"context"
	"errors"
	"xiuadmin/internal/library/xgen"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gmode"
)

var (
	initFuncList map[string]func(ctx context.Context) error
)

func RegisterInitFunc(name string, fn func(ctx context.Context) error) {
	if initFuncList == nil {
		initFuncList = make(map[string]func(ctx context.Context) error)
	}
	initFuncList[name] = fn
}

func InitSystemDeferFunc(ctx context.Context) error {
	for name, fn := range initFuncList {
		if fn == nil {
			g.Log().Errorf(ctx, "InitSystemDeferFunc %s fn is nil.", name)
			return errors.New("func is nil")
		}
		err := fn(ctx)
		if err != nil {
			g.Log().Errorf(ctx, "InitSystemDeferFunc %s err: %v", name, err)
			return err
		}
		g.Log().Infof(ctx, "InitSystemDeferFunc %s success.", name)
	}
	// 初始化生成代码配置
	if !gmode.IsProduct() {
		xgen.Init(ctx)
	}
	return nil
}
