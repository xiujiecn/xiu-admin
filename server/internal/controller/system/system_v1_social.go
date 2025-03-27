package system

import (
	"context"

	v1 "xiuadmin/api/system/v1"
	"xiuadmin/internal/model/response"
	"xiuadmin/internal/service"
)

func (c *ControllerV1) SysSocialList(ctx context.Context, req *v1.SysSocialListReq) (res *v1.SysSocialListRes, err error) {
	items, total, err := service.SysSocial().List(ctx, &req.SysSocialListParam, &req.PageInfo)
	if err != nil {
		return nil, err
	}
	res = &v1.SysSocialListRes{
		Items: items,
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}
	return res, nil
}
