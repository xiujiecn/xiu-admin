// package handler
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package handler

import "github.com/gogf/gf/v2/database/gdb"

// ForceCache 强制缓存
func ForceCache(m *gdb.Model) *gdb.Model {
	return m.Cache(gdb.CacheOption{Duration: 0, Force: true})
}
