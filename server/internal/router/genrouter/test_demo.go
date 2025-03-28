// Package genrouter
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
// @AutoGenerate Version
package genrouter

import "xiuadmin/internal/controller/gen"

func init() {
	LoginRequiredRouter = append(LoginRequiredRouter, gen.TestDemo) // 测试单表
}
