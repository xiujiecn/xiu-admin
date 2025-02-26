package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/model/response"
	"server/internal/service"
)

func (c *ControllerV1) SysUserOnlineList(ctx context.Context, req *v1.SysUserOnlineListReq) (res *v1.SysUserOnlineListRes, err error) {
	items, total, err := service.SysUserOnline().List(ctx, &req.SysUserOnlineListParam, &req.PageInfo)
	if err != nil {
		return nil, err
	}
	return &v1.SysUserOnlineListRes{
		Items: items,
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}, nil
}

func (c *ControllerV1) SysUserOnlineDelete(ctx context.Context, req *v1.SysUserOnlineDeleteReq) (res *v1.SysUserOnlineDeleteRes, err error) {
	err = service.SysUserOnline().Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &v1.SysUserOnlineDeleteRes{}, nil
}
