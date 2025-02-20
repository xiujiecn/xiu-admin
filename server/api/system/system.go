// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package system

import (
	"context"

	"server/api/system/v1"
)

type ISystemV1 interface {
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error)
	Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
	GetAccessCodes(ctx context.Context, req *v1.GetAccessCodesReq) (res *v1.GetAccessCodesRes, err error)
	GetCaptcha(ctx context.Context, req *v1.GetCaptchaReq) (res *v1.GetCaptchaRes, err error)
	DeptList(ctx context.Context, req *v1.DeptListReq) (res *v1.DeptListRes, err error)
	SysDictTypeList(ctx context.Context, req *v1.SysDictTypeListReq) (res *v1.SysDictTypeListRes, err error)
	SysDictDataList(ctx context.Context, req *v1.SysDictDataListReq) (res *v1.SysDictDataListRes, err error)
	MenuAll(ctx context.Context, req *v1.MenuAllReq) (res *v1.MenuAllRes, err error)
	MenuList(ctx context.Context, req *v1.MenuListReq) (res *v1.MenuListRes, err error)
	PostList(ctx context.Context, req *v1.PostListReq) (res *v1.PostListRes, err error)
	RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error)
	UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error)
	UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error)
	AddUser(ctx context.Context, req *v1.AddUserReq) (res *v1.AddUserRes, err error)
}
