package model

import (
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type BaseClaims struct {
	UUID         uuid.UUID
	ID           int64   // 用户ID
	Username     string  // 用户名
	NickName     string  // 昵称
	DeptId       int64   // 部门ID
	TenantId     string  // 租户编号
	AuthorityIds []int64 // 角色ID
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}
