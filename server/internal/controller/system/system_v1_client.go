package system

import (
	"context"

	v1 "xiujieadmin/api/system/v1"
	"xiujieadmin/internal/model/response"
	"xiujieadmin/internal/service"
)

func (c *ControllerV1) SysClientList(ctx context.Context, req *v1.SysClientListReq) (res *v1.SysClientListRes, err error) {
	items, total, err := service.SysClient().List(ctx, &req.SysClientListParam, &req.PageInfo)
	if err != nil {
		return nil, err
	}
	res = &v1.SysClientListRes{
		Items: items,
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}
	return
}
