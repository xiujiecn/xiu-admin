package system

import (
	"context"

	v1 "xiujieadmin/api/system/v1"
	"xiujieadmin/internal/model/response"
	"xiujieadmin/internal/service"
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
