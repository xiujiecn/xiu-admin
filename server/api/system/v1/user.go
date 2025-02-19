package v1

import (
	"server/internal/model"
	"server/internal/model/request"
	"server/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type UserInfo struct {
	Id       int64  `json:"id" dc:"用户ID"`
	Username string `json:"userName" dc:"用户名"`
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

type UserListReq struct {
	g.Meta `path:"/user/list" method:"get" tags:"系统" summary:"获取用户列表"`
	request.PageInfo
	model.UserListQuery
}

type UserListRes struct {
	response.PageResult
	Data []*model.SysUser `json:"items" dc:"用户列表"`
}

type AddUserReq struct {
	g.Meta `path:"/user/add" method:"post" tags:"系统" summary:"新增用户"`
	model.AddUser
}

type AddUserRes struct {
	Data *model.SysUser `json:"data" dc:"用户信息"`
}
