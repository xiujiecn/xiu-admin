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
