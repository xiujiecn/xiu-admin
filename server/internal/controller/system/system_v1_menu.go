package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/service"
)

func (c *ControllerV1) MenuAll(ctx context.Context, req *v1.MenuAllReq) (res *v1.MenuAllRes, err error) {
	data, err := service.SysMenu().GetUserMenuTree(ctx)
	return &data, err
}
func (c *ControllerV1) MenuList(ctx context.Context, req *v1.MenuListReq) (res *v1.MenuListRes, err error) {
	data, total, err := service.SysMenu().GetTenantMenu(ctx, &req.SysMenuListQuery)
	return &v1.MenuListRes{
		Data:  data,
		Total: total,
	}, err
}
func (c *ControllerV1) MenuView(ctx context.Context, req *v1.MenuViewReq) (res *v1.MenuViewRes, err error) {
	data, err := service.SysMenu().GetSysMenuView(ctx, req.MenuId)
	if err != nil {
		return nil, err
	}
	return &v1.MenuViewRes{
		SysMenuViewModel: *data,
	}, nil
}

func (c *ControllerV1) MenuAdd(ctx context.Context, req *v1.MenuAddReq) (res *v1.MenuAddRes, err error) {
	data, err := service.SysMenu().AddSysMenu(ctx, &req.SysMenuAddModel)
	if err != nil {
		return nil, err
	}
	return &v1.MenuAddRes{
		SysMenuViewModel: *data,
	}, nil
}

func (c *ControllerV1) MenuUpdate(ctx context.Context, req *v1.MenuUpdateReq) (res *v1.MenuUpdateRes, err error) {
	data, err := service.SysMenu().UpdateSysMenu(ctx, &req.SysMenuUpdateModel)
	if err != nil {
		return nil, err
	}
	return &v1.MenuUpdateRes{
		SysMenuViewModel: *data,
	}, nil
}
func (c *ControllerV1) MenuDelete(ctx context.Context, req *v1.MenuDeleteReq) (res *v1.MenuDeleteRes, err error) {
	err = service.SysMenu().DeleteSysMenu(ctx, req.MenuId)
	if err != nil {
		return nil, err
	}
	return &v1.MenuDeleteRes{}, nil
}
