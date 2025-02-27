package v1

import (
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/model/response"

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
	UserInfo
}

type UserListReq struct {
	g.Meta `path:"/user/list" method:"get" tags:"系统" summary:"获取用户列表"`
	request.PageInfo
	model.UserListParam
}

type UserListRes struct {
	response.PageResult
	Data []*model.SysUserListModel `json:"items" dc:"用户列表"`
}

type AddUserReq struct {
	g.Meta `path:"/user/add" method:"post" tags:"系统" summary:"新增用户"`
	model.SysUserAddModel
}

type AddUserRes struct {
	Data *model.SysUserViewModel `json:"data" dc:"用户信息"`
}

type UserProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" tags:"系统" summary:"获取当前用户信息"`
}

type UserProfileRes struct {
	User *model.UserProfileModel `json:"user" dc:"用户信息"`
}

type UpdateCurrentUserReq struct {
	g.Meta `path:"/user/profile/update" method:"post" tags:"系统" summary:"更新当前用户信息"`
	model.UpdateCurrentUserModel
}

type UpdateCurrentUserRes struct {
	UserInfo
}

type UpdateCurrentUserPasswordReq struct {
	g.Meta `path:"/user/profile/password" method:"post" tags:"系统" summary:"更新当前用户密码"`
	model.UpdateCurrentUserPasswordModel
}

type UpdateCurrentUserPasswordRes struct {
}

type UpdateUserReq struct {
	g.Meta `path:"/user/update" method:"post" tags:"系统" summary:"更新用户信息"`
	model.SysUserUpdateModel
}

type UpdateUserRes struct {
	model.SysUserViewModel
}

type DeleteUserReq struct {
	g.Meta  `path:"/user/delete" method:"post" tags:"系统" summary:"删除用户"`
	UserId  int64   `json:"userId" dc:"用户ID"`
	UserIds []int64 `json:"userIds" dc:"用户ID列表"`
}

type DeleteUserRes struct {
}

type GetUserReq struct {
	g.Meta `path:"/user/view" method:"get" tags:"系统" summary:"获取用户信息"`
	UserId int64 `json:"userId" dc:"用户ID"`
}

type GetUserRes struct {
	model.SysUserViewModel
}

type ResetPasswordReq struct {
	g.Meta   `path:"/user/resetPassword" method:"post" tags:"系统" summary:"重置用户密码"`
	UserId   int64  `json:"userId" dc:"用户ID"`
	Password string `json:"password" dc:"新密码"`
}

type ResetPasswordRes struct {
}

type GetUserListByDeptIdReq struct {
	g.Meta `path:"/user/listByDeptId" method:"get" tags:"系统" summary:"获取部门用户列表"`
	DeptId int64 `json:"deptId" dc:"部门ID"`
	request.PageInfo
}

type GetUserListByDeptIdRes struct {
	response.PageResult
	Items []*model.SysUserMiniModel `json:"items" dc:"用户列表"`
}
