package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta       `path:"/auth/login" method:"post" tags:"系统" summary:"登录"`
	Username     string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password     string `json:"password" v:"required#密码不能为空" dc:"密码"`
	CaptchaID    string `json:"captchaID" dc:"验证码ID"`
	CaptchaValue string `json:"captchaValue" v:"required#验证码不能为空" dc:"验证码"`
}

type LoginRes struct {
	ID          int64  `json:"id" dc:"用户ID"`
	Username    string `json:"username" dc:"用户名"`
	NickName    string `json:"nickName" dc:"昵称"`
	Avatar      string `json:"avatar" dc:"头像"`
	AccessToken string `json:"accessToken" dc:"访问令牌"`
}

type RefreshTokenReq struct {
	g.Meta          `path:"/auth/refresh" method:"post" tags:"系统" summary:"刷新访问令牌"`
	WithCredentials bool `json:"withCredentials" dc:"是否携带凭证"`
}

type RefreshTokenRes struct {
	Data   string `json:"data" dc:"新的访问令牌"`
	Status int    `json:"status" dc:"状态码"`
}

type LogoutReq struct {
	g.Meta          `path:"/auth/logout" method:"post" tags:"系统" summary:"退出登录"`
	WithCredentials bool `json:"withCredentials" dc:"是否携带凭证"`
}

type LogoutRes struct {
	Data   string `json:"data" dc:"退出登录"`
	Status int    `json:"status" dc:"状态码"`
}

type GetAccessCodesReq struct {
	g.Meta `path:"/auth/codes" method:"get" tags:"系统" summary:"获取用户权限码"`
}

type GetAccessCodesRes struct {
	Data   []string `json:"data" dc:"用户权限码"`
	Status int      `json:"status" dc:"状态码"`
}
