// package xgen
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package xgen

import (
	"context"
	gendao "xiuadmin/internal/library/xgen/gen_dao"
)

func Init(ctx context.Context) (err error) {
	err = gendao.Init(ctx)
	if err != nil {
		return
	}
	return
}
