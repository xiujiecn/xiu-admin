package v1

import "github.com/gogf/gf/v2/frame/g"

type UserInfo struct {
	Id       int64  `json:"id" dc:"用户ID"`
	Username string `json:"username" dc:"用户名"`
	Nickname string `json:"realName" dc:"昵称"`
	Avatar   string `json:"avatar" dc:"头像"`
	Email    string `json:"email" dc:"邮箱"`
}

type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"系统" summary:"获取用户信息"`
}

type UserInfoRes struct {
	Data *UserInfo `json:"data" dc:"用户信息"`
}
