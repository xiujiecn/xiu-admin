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

// 删除绑定关系
func (c *ControllerV1) SysSocialDelete(ctx context.Context, req *v1.SysSocialDeleteReq) (res *v1.SysSocialDeleteRes, err error) {
	err = service.SysSocial().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.SysSocialDeleteRes{}, nil
}
