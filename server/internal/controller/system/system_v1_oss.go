package system

import (
	"context"

	v1 "xiujieadmin/api/system/v1"
	"xiujieadmin/internal/model/response"
	"xiujieadmin/internal/service"
)

func (c *ControllerV1) SysOssList(ctx context.Context, req *v1.SysOssListReq) (res *v1.SysOssListRes, err error) {
	items, total, err := service.SysOss().List(ctx, &req.SysOssListParam, &req.PageInfo)
	if err != nil {
		return nil, err
	}
	res = &v1.SysOssListRes{
		Items: items,
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}
	return
}
