// Package genrouter
// @Link  https://github.com/xiujie/xiujie-admin
// @Copyright  Copyright (c) 2025 XiuJieZhiLian CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujie/xiujie-admin/blob/master/LICENSE
// @AutoGenerate Version
package genrouter

import "xiujieadmin/internal/controller/gen"

func init() {
	LoginRequiredRouter = append(LoginRequiredRouter, gen.TestDemo) // 测试单表
}