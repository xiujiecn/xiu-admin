package system

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "xiuadmin/api/system/v1"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/request"
	"xiuadmin/internal/model/response"
	"xiuadmin/internal/service"
)

func (c *ControllerV1) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	// 获取token
	userId := contexts.GetUserId(ctx)
	if userId == 0 {
		return nil, gerror.NewCode(gcode.CodeSecurityReason)
	}
	// 获取用户信息
	user, err := service.SysUser().GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	menus, err := service.SysMenu().GetUserMenuTree(ctx)
	if err != nil {
		return nil, err
	}
	homePath := ""
	firstMenu := "404"
	if len(menus) > 0 {
		for _, menu := range menus {
			firstMenu = menu.Path
			if len(menu.Children) == 0 {
				break
			}
			for _, menu2 := range menu.Children {
				firstMenu = firstMenu + "/" + menu2.Path
				if len(menu2.Children) == 0 {
					break
				}
				for _, menu3 := range menu2.Children {
					firstMenu = firstMenu + "/" + menu3.Path
					if len(menu3.Children) == 0 {
						break
					}
					for _, menu4 := range menu3.Children {
						firstMenu = firstMenu + "/" + menu4.Path
						if len(menu4.Children) == 0 {
							break
						}
						for _, menu5 := range menu4.Children {
							firstMenu = firstMenu + "/" + menu5.Path
							if len(menu5.Children) == 0 {
								break
							}
							for _, menu6 := range menu5.Children {
								firstMenu = firstMenu + "/" + menu6.Path
								if len(menu6.Children) == 0 {
									break
								}
							}
						}
					}
				}
			}
		}
	}
	g.Log().Debugf(ctx, "firstMenu: %s", firstMenu)
	if homePath == "" {
		homePath = "/" + firstMenu
	}

	return &v1.UserInfoRes{
		UserInfo: v1.UserInfo{
			Id:       user.UserId,
			Username: user.UserName,
			Nickname: user.NickName,
			Avatar:   user.Avatar,
			Email:    user.Email,
			HomePath: homePath,
		},
	}, nil
}
func (c *ControllerV1) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	// 获取用户列表
	users, total, err := service.SysUser().List(ctx, &req.PageInfo, &req.UserListParam)
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
	// 如果调用方传入 userId == 0，则不返回错误，直接返回空的用户视图结构
	if req.UserId == 0 {
		return &v1.GetUserRes{
			SysUserViewModel: model.SysUserViewModel{},
		}, nil
	}

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
func (c *ControllerV1) GetUserListByDeptId(ctx context.Context, req *v1.GetUserListByDeptIdReq) (res *v1.GetUserListByDeptIdRes, err error) {
	users, total, err := service.SysUser().List(ctx, &request.PageInfo{
		Page:     req.Page,
		PageSize: req.PageSize,
	}, &model.UserListParam{
		DeptId: req.DeptId,
	})
	if err != nil {
		return nil, err
	}
	miniUsers := make([]*model.SysUserMiniModel, 0)
	gconv.Structs(users, &miniUsers)
	return &v1.GetUserListByDeptIdRes{
		Items: miniUsers,
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}, nil
}

// UserRegister 用户注册
func (c *ControllerV1) UserRegister(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	err = service.SysUser().Register(ctx, &req.SysUserRegisterModel)
	if err != nil {
		return nil, err
	}
	return &v1.UserRegisterRes{}, nil
}

// UserStatus 更新用户状态
func (c *ControllerV1) UserStatus(ctx context.Context, req *v1.UserStatusReq) (res *v1.UserStatusRes, err error) {
	err = service.SysUser().Status(ctx, &req.SysUserStatusParam)
	if err != nil {
		return nil, err
	}
	return &v1.UserStatusRes{}, nil
}
