// package model
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package model

import "github.com/gogf/gf/v2/frame/g"

// 上下文结构
type Context struct {
	User        *Identity // 上下文用户信息
	IsWebSocket bool      // 是否是WebSocket
	Data        g.Map     // 自定KV变量，业务模块根据需要设置，不固定
}

type Role struct {
	RoleId    int64  // 角色ID
	DataScope string // 数据范围
}

// 通用身份模型
type Identity struct {
	BaseClaims
	AccessCodeList []string // 访问码列表
}
