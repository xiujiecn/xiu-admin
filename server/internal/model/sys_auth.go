package model

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

type BaseClaims struct {
	UUID     string
	ID       int64  // 用户ID
	Username string // 用户名
	NickName string // 昵称
	DeptId   int64  // 部门ID
	TenantId string // 租户编号
	Roles    []Role // 角色列表
	LoginAt  int64  // 登录时间戳(秒)
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}
