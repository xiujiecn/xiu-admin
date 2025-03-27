package system

import (
	"context"

	v1 "xiuadmin/api/system/v1"
	"xiuadmin/internal/model/response"
	"xiuadmin/internal/service"
)

func (c *ControllerV1) ListLogininfor(ctx context.Context, req *v1.ListLogininforReq) (res *v1.ListLogininforRes, err error) {
	items, total, err := service.SysLogininfor().List(ctx, &req.SysLogininforListParam)
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
func (c *ControllerV1) DeleteLogininfor(ctx context.Context, req *v1.DeleteLogininforReq) (res *v1.DeleteLogininforRes, err error) {
	out, err := service.SysLogininfor().Delete(ctx, &req.SysLogininforDeleteParam)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteLogininforRes{
		SysLogininforDeleteModel: *out,
	}, nil
}
