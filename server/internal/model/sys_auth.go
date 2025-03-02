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
type LoginParams struct {
	CaptchaID    string `json:"captchaId" dc:"验证码ID"`
	CaptchaValue string `json:"captchaValue" dc:"验证码值"`
	Username     string `json:"username" dc:"用户名"`
	Password     string `json:"password" dc:"密码"`
	TenantId     string `json:"tenantId" dc:"租户ID"`
}

type LoginResult struct {
	Token string
}
