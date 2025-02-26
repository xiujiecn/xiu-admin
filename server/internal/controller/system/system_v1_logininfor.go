package system

import (
	"context"

	v1 "xiujieadmin/api/system/v1"
	"xiujieadmin/internal/model/response"
	"xiujieadmin/internal/service"
)

func (c *ControllerV1) ListLogininfor(ctx context.Context, req *v1.ListLogininforReq) (res *v1.ListLogininforRes, err error) {
	items, total, err := service.SysLogininfor().ListLogininfor(ctx, &req.SysLogininforListParam, &req.PageInfo)
	if err != nil {
		return nil, err
	}
	return &v1.ListLogininforRes{
		Items: items,
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}, nil
}
