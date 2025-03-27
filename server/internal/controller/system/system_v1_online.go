package system

import (
	"context"
	"errors"

	v1 "xiuadmin/api/system/v1"
	"xiuadmin/internal/model/response"
	"xiuadmin/internal/service"
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
	if len(req.Ids) == 0 {
		return nil, errors.New("ids不能为空")
	}
	err = service.SysUserOnline().Delete(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	return &v1.SysUserOnlineDeleteRes{}, nil
}
