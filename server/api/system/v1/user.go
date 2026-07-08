// Package v1
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package v1

import (
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/request"
	"xiuadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type UserInfo struct {
	Id       int64  `json:"id" dc:"用户ID"`
	Username string `json:"userName" dc:"用户名"`
	Nickname string `json:"realName" dc:"昵称"`
	Avatar   string `json:"avatar" dc:"头像"`
	Email    string `json:"email" dc:"邮箱"`
	HomePath string `json:"homePath" dc:"首页路径"`
}

type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"系统-用户管理" summary:"获取当前用户信息" x-check-permission:"cpc:current:user"`
}

type UserInfoRes struct {
	UserInfo
}

type UserListReq struct {
	g.Meta `path:"/user/list" method:"get" tags:"系统-用户管理" summary:"获取用户列表" x-check-permission:"cpm:system:user:list"`
	request.PageInfo
	model.UserListParam
}

type UserListRes struct {
	response.PageResult
	Data []*model.SysUserListModel `json:"items" dc:"用户列表"`
}

type AddUserReq struct {
	g.Meta `path:"/user/add" method:"post" tags:"系统-用户管理" summary:"新增用户" x-check-permission:"cpm:system:user:add"`
	model.SysUserAddModel
}

type AddUserRes struct {
	Data *model.SysUserViewModel `json:"data" dc:"用户信息"`
}

type UserProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" tags:"系统-用户管理" summary:"获取当前用户信息" x-check-permission:"cpc:current:user"`
}

type UserProfileRes struct {
	User *model.UserProfileModel `json:"user" dc:"用户信息"`
}

type UpdateCurrentUserReq struct {
	g.Meta `path:"/user/profile/update" method:"post" tags:"系统-用户管理" summary:"更新当前用户信息" x-check-permission:"cpc:current:user"`
	model.UpdateCurrentUserModel
}

type UpdateCurrentUserRes struct {
	UserInfo
}

type UpdateCurrentUserPasswordReq struct {
	g.Meta `path:"/user/profile/password" method:"post" tags:"系统-用户管理" summary:"更新当前用户密码" x-check-permission:"cpc:current:user"`
	model.UpdateCurrentUserPasswordModel
}

type UpdateCurrentUserPasswordRes struct {
}

type UpdateCurrentUserContactCodeReq struct {
	g.Meta `path:"/user/profile/contact/code" method:"post" tags:"系统-用户管理" summary:"发送当前用户邮箱/手机修改验证码" x-check-permission:"cpc:current:user"`
	model.UpdateCurrentUserContactCodeModel
}

type UpdateCurrentUserContactCodeRes struct {
}

type UpdateCurrentUserContactReq struct {
	g.Meta `path:"/user/profile/contact/update" method:"post" tags:"系统-用户管理" summary:"修改当前用户邮箱/手机" x-check-permission:"cpc:current:user"`
	model.UpdateCurrentUserContactModel
}

type UpdateCurrentUserContactRes struct {
}

type UpdateUserReq struct {
	g.Meta `path:"/user/update" method:"post" tags:"系统-用户管理" summary:"更新用户信息" x-check-permission:"cpm:system:user:edit"`
	model.SysUserUpdateModel
}

type UpdateUserRes struct {
	model.SysUserViewModel
}

type DeleteUserReq struct {
	g.Meta  `path:"/user/delete" method:"post" tags:"系统-用户管理" summary:"删除用户" x-check-permission:"cpm:system:user:remove"`
	UserId  int64   `json:"userId" dc:"用户ID"`
	UserIds []int64 `json:"userIds" dc:"用户ID列表"`
}

type DeleteUserRes struct {
}

type GetUserReq struct {
	g.Meta `path:"/user/view" method:"get" tags:"系统-用户管理" summary:"获取用户信息" x-check-permission:"cpm:system:user:query"`
	UserId int64 `json:"userId" dc:"用户ID"`
}

type GetUserRes struct {
	model.SysUserViewModel
}

type ResetPasswordReq struct {
	g.Meta   `path:"/user/resetPassword" method:"post" tags:"系统-用户管理" summary:"重置用户密码" x-check-permission:"cpm:system:user:resetPwd"`
	UserId   int64  `json:"userId" dc:"用户ID"`
	Password string `json:"password" dc:"新密码"`
}

type ResetPasswordRes struct {
}

type GetUserListByDeptIdReq struct {
	g.Meta `path:"/user/listByDeptId" method:"get" tags:"系统-用户管理" summary:"获取部门用户列表" x-check-permission:"cpm:system:user:list"`
	DeptId int64 `json:"deptId" dc:"部门ID"`
	request.PageInfo
}

type GetUserListByDeptIdRes struct {
	response.PageResult
	Items []*model.SysUserMiniModel `json:"items" dc:"用户列表"`
}

// 用户注册
type UserRegisterReq struct {
	g.Meta `path:"/user/register" method:"post" tags:"系统-用户管理" summary:"用户注册" x-check-permission:"cpm:system:user:register"`
	model.SysUserRegisterModel
}

type UserRegisterRes struct {
}

// UserRegisterCodeReq 发送注册验证码
type UserRegisterCodeReq struct {
	g.Meta `path:"/user/register/code" method:"post" tags:"系统-用户管理" summary:"发送注册验证码"`
	model.SysUserRegisterCodeModel
}

type UserRegisterCodeRes struct {
}

// StatusReq 更新用户状态
type UserStatusReq struct {
	g.Meta `path:"/user/status" method:"post" tags:"系统-用户管理" summary:"更新用户状态" x-check-permission:"cpm:system:user:edit"`
	model.SysUserStatusParam
}

type UserStatusRes struct {
}
