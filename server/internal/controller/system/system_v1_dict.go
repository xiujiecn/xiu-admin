package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/model/response"
	"server/internal/service"
)

func (c *ControllerV1) SysDictTypeList(ctx context.Context, req *v1.SysDictTypeListReq) (res *v1.SysDictTypeListRes, err error) {
	items, total, err := service.SysDict().GetDictTypeList(ctx, &req.SysDictTypeListQuery, &req.PageInfo)
	if err != nil {
		return nil, err
	}
	return &v1.SysDictTypeListRes{
		Items: items,
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}, nil
}

func (c *ControllerV1) SysDictDataList(ctx context.Context, req *v1.SysDictDataListReq) (res *v1.SysDictDataListRes, err error) {
	items, total, err := service.SysDict().GetDictDataList(ctx, &req.SysDictDataListQuery, &req.PageInfo)
	if err != nil {
		return nil, err
	}
	return &v1.SysDictDataListRes{
		Items: items,
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}, nil
}
