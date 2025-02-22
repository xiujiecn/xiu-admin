package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/model/response"
	"server/internal/service"
)

func (c *ControllerV1) SysTenantList(ctx context.Context, req *v1.SysTenantListReq) (res *v1.SysTenantListRes, err error) {
	items, total, err := service.SysTenant().List(ctx, &req.SysTenantListQuery, &req.PageInfo)
	if err != nil {
		return nil, err
	}
	res = &v1.SysTenantListRes{
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
		Items: items,
	}
	return
}

func (c *ControllerV1) SysTenantPackageList(ctx context.Context, req *v1.SysTenantPackageListReq) (res *v1.SysTenantPackageListRes, err error) {
	items, total, err := service.SysTenant().TenantPackageList(ctx, &req.SysTenantPackageListQuery, &req.PageInfo)
	if err != nil {
		return nil, err
	}
	res = &v1.SysTenantPackageListRes{
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
		Items: items,
	}
	return
}
