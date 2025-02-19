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
		Data: &v1.UserInfo{
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
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
