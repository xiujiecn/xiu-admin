package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/model/response"
	"server/internal/service"
)

func (c *ControllerV1) DeptList(ctx context.Context, req *v1.DeptListReq) (res *v1.DeptListRes, err error) {
	items, total, err := service.SysDept().GetDeptList(ctx, req.DeptListQuery)
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
