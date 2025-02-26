package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/model/response"
	"server/internal/service"
)

func (c *ControllerV1) GetConfigList(ctx context.Context, req *v1.GetConfigListReq) (res *v1.GetConfigListRes, err error) {

	items, total, err := service.SysConfig().GetConfigList(ctx, &req.SysConfigListParam, &req.PageInfo)
	res = &v1.GetConfigListRes{
		Data: items,
		PageResult: response.PageResult{
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	}
	return
}
