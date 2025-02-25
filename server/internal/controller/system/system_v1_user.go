package system

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "server/api/system/v1"
	"server/internal/model/response"
	"server/internal/service"
)

func (c *ControllerV1) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	// 获取token
	authorization := g.RequestFromCtx(ctx).Header.Get("Authorization")
	if authorization == "" {
		return nil, gerror.NewCode(gcode.CodeNotImplemented)
	}
	token := strings.TrimPrefix(authorization, "Bearer ")
	// 解析token
	claims, err := service.SysAuth().ParseToken(ctx, token)
	if err != nil {
		return nil, err
	}
	// 获取用户信息
	user, err := service.SysUser().GetUserById(ctx, claims.BaseClaims.ID)
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoRes{
		UserInfo: v1.UserInfo{
			Id:       user.UserId,
			Username: user.UserName,
			Nickname: user.NickName,
			Avatar:   user.Avatar,
			Email:    user.Email,
		},
	}, nil
}
func (c *ControllerV1) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	// 获取当前用户信息
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	req.TenantId = claims.BaseClaims.TenantId
	// 获取用户列表
	users, total, err := service.SysUser().GetUserList(ctx, req.PageInfo, req.UserListQuery)
	if err != nil {
		return nil, err
	}
	return &v1.UserListRes{
		Data: users,
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}, nil
}

func (c *ControllerV1) AddUser(ctx context.Context, req *v1.AddUserReq) (res *v1.AddUserRes, err error) {
	data, err := service.SysUser().AddUser(ctx, &req.SysUserAddModel)
	if err != nil {
		return nil, err
	}
	return &v1.AddUserRes{
		Data: data,
	}, nil
}

func (c *ControllerV1) UserProfile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error) {
	user, err := service.SysUser().Profile(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.UserProfileRes{
		User: user,
	}, nil
}

func (c *ControllerV1) UpdateCurrentUser(ctx context.Context, req *v1.UpdateCurrentUserReq) (res *v1.UpdateCurrentUserRes, err error) {
	user, err := service.SysUser().UpdateCurrentUser(ctx, &req.UpdateCurrentUserModel)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateCurrentUserRes{
		UserInfo: v1.UserInfo{
			Id:       user.UserId,
			Username: user.UserName,
			Nickname: user.NickName,
			Avatar:   user.Avatar,
			Email:    user.Email,
		},
	}, nil
}

func (c *ControllerV1) UpdateCurrentUserPassword(ctx context.Context, req *v1.UpdateCurrentUserPasswordReq) (res *v1.UpdateCurrentUserPasswordRes, err error) {
	err = service.SysUser().UpdateCurrentUserPassword(ctx, &req.UpdateCurrentUserPasswordModel)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateCurrentUserPasswordRes{}, nil
}

func (c *ControllerV1) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	err = service.SysUser().UpdateUser(ctx, &req.SysUserUpdateModel)
	if err != nil {
		return nil, err
	}
	user, err := service.SysUser().GetUserById(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateUserRes{
		SysUserViewModel: *user,
	}, nil
}

func (c *ControllerV1) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	if req.UserId != 0 {
		err = service.SysUser().DeleteUser(ctx, []int64{req.UserId})
	} else {
		err = service.SysUser().DeleteUser(ctx, req.UserIds)
	}
	if err != nil {
		return nil, err
	}
	return &v1.DeleteUserRes{}, nil
}
func (c *ControllerV1) GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error) {
	user, err := service.SysUser().GetUserById(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserRes{
		SysUserViewModel: *user,
	}, nil
}
func (c *ControllerV1) ResetPassword(ctx context.Context, req *v1.ResetPasswordReq) (res *v1.ResetPasswordRes, err error) {
	err = service.SysUser().ResetPassword(ctx, req.UserId, req.Password)
	if err != nil {
		return nil, err
	}
	return &v1.ResetPasswordRes{}, nil
}
