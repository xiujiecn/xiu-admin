package model

import "github.com/gogf/gf/v2/frame/g"

// 上下文结构
type Context struct {
	User *Identity // 上下文用户信息
	Data g.Map     // 自定KV变量，业务模块根据需要设置，不固定
}

// 通用身份模型
type Identity struct {
	UUID         string
	ID           int64   // 用户ID
	Username     string  // 用户名
	NickName     string  // 昵称
	DeptId       int64   // 部门ID
	TenantId     string  // 租户编号
	AuthorityIds []int64 // 角色ID
}
