package system

import (
	"context"

	v1 "xiujieadmin/api/system/v1"
	"xiujieadmin/internal/model/response"
	"xiujieadmin/internal/service"
)

func (c *ControllerV1) SysTenantList(ctx context.Context, req *v1.SysTenantListReq) (res *v1.SysTenantListRes, err error) {
	items, total, err := service.SysTenant().List(ctx, req.SysTenantListParam)
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
	items, total, err := service.SysTenantPackage().List(ctx, &req.SysTenantPackageListParam)
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
func (c *ControllerV1) SysTenantAdd(ctx context.Context, req *v1.SysTenantAddReq) (res *v1.SysTenantAddRes, err error) {
	data, err := service.SysTenant().Add(ctx, req.SysTenantAddParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysTenantAddRes{
		SysTenantAddModel: data,
	}, nil
}
func (c *ControllerV1) SysTenantEdit(ctx context.Context, req *v1.SysTenantEditReq) (res *v1.SysTenantEditRes, err error) {
	data, err := service.SysTenant().Edit(ctx, req.SysTenantEditParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysTenantEditRes{
		SysTenantEditModel: data,
	}, nil
}
func (c *ControllerV1) SysTenantDelete(ctx context.Context, req *v1.SysTenantDeleteReq) (res *v1.SysTenantDeleteRes, err error) {
	data, err := service.SysTenant().Delete(ctx, req.SysTenantDeleteParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysTenantDeleteRes{
		SysTenantDeleteModel: data,
	}, nil
}
func (c *ControllerV1) SysTenantStatus(ctx context.Context, req *v1.SysTenantStatusReq) (res *v1.SysTenantStatusRes, err error) {
	data, err := service.SysTenant().Status(ctx, req.SysTenantStatusParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysTenantStatusRes{
		SysTenantStatusModel: data,
	}, nil
}
func (c *ControllerV1) SysTenantView(ctx context.Context, req *v1.SysTenantViewReq) (res *v1.SysTenantViewRes, err error) {
	data, err := service.SysTenant().View(ctx, req.SysTenantViewParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysTenantViewRes{
		SysTenantViewModel: data,
	}, nil
}
func (c *ControllerV1) SysTenantPackageView(ctx context.Context, req *v1.SysTenantPackageViewReq) (res *v1.SysTenantPackageViewRes, err error) {
	data, err := service.SysTenantPackage().View(ctx, req.SysTenantPackageViewParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysTenantPackageViewRes{
		SysTenantPackageViewModel: data,
	}, nil
}
func (c *ControllerV1) SysTenantPackageStatus(ctx context.Context, req *v1.SysTenantPackageStatusReq) (res *v1.SysTenantPackageStatusRes, err error) {
	data, err := service.SysTenantPackage().Status(ctx, req.SysTenantPackageStatusParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysTenantPackageStatusRes{
		SysTenantPackageStatusModel: data,
	}, nil
}
func (c *ControllerV1) SysTenantPackageAdd(ctx context.Context, req *v1.SysTenantPackageAddReq) (res *v1.SysTenantPackageAddRes, err error) {
	data, err := service.SysTenantPackage().Add(ctx, req.SysTenantPackageAddParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysTenantPackageAddRes{
		SysTenantPackageAddModel: data,
	}, nil
}
func (c *ControllerV1) SysTenantPackageEdit(ctx context.Context, req *v1.SysTenantPackageEditReq) (res *v1.SysTenantPackageEditRes, err error) {
	data, err := service.SysTenantPackage().Edit(ctx, req.SysTenantPackageEditParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysTenantPackageEditRes{
		SysTenantPackageEditModel: data,
	}, nil
}
func (c *ControllerV1) SysTenantPackageDelete(ctx context.Context, req *v1.SysTenantPackageDeleteReq) (res *v1.SysTenantPackageDeleteRes, err error) {
	data, err := service.SysTenantPackage().Delete(ctx, req.SysTenantPackageDeleteParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysTenantPackageDeleteRes{
		SysTenantPackageDeleteModel: data,
	}, nil
}
