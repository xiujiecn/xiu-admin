package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/model/response"
	"server/internal/service"
)

func (c *ControllerV1) DeptList(ctx context.Context, req *v1.DeptListReq) (res *v1.DeptListRes, err error) {
	items, total, err := service.SysDept().GetDeptList(ctx, req.DeptListParam)
	if err != nil {
		return nil, err
	}
	return &v1.DeptListRes{
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
		Data: items,
	}, nil
}

func (c *ControllerV1) DeptTree(ctx context.Context, req *v1.DeptTreeReq) (res *v1.DeptTreeRes, err error) {
	items, err := service.SysDept().GetDeptTree(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.DeptTreeRes{
		Data: items,
	}, nil
}
