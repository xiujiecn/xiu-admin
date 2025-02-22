package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/model/response"
	"server/internal/service"
)

func (c *ControllerV1) SysOssList(ctx context.Context, req *v1.SysOssListReq) (res *v1.SysOssListRes, err error) {
	items, total, err := service.SysOss().List(ctx, &req.SysOssListQuery, &req.PageInfo)
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
