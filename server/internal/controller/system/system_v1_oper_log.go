package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/model/response"
	"server/internal/service"
)

func (c *ControllerV1) GetOperLogList(ctx context.Context, req *v1.GetOperLogListReq) (res *v1.GetOperLogListRes, err error) {
	items, total, err := service.SysOperLog().GetOperLogList(ctx, &req.SysOperLogListQuery, &req.PageInfo)
	if err != nil {
		return nil, err
	}

	return &v1.GetOperLogListRes{
		Items: items,
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}, nil
}
