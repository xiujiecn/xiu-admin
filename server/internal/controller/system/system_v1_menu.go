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
