package system

import (
	"context"

	v1 "xiujieadmin/api/system/v1"
	"xiujieadmin/internal/model/response"
	"xiujieadmin/internal/service"
)

func (c *ControllerV1) GetConfigList(ctx context.Context, req *v1.GetConfigListReq) (res *v1.GetConfigListRes, err error) {

	items, total, err := service.SysConfig().List(ctx, &req.SysConfigListParam)
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
func (c *ControllerV1) AddConfig(ctx context.Context, req *v1.AddConfigReq) (res *v1.AddConfigRes, err error) {
	output, err := service.SysConfig().Add(ctx, &req.SysConfigAddParam)
	if err != nil {
		return nil, err
	}
	res = &v1.AddConfigRes{
		SysConfigAddModel: *output,
	}
	return
}
func (c *ControllerV1) EditConfig(ctx context.Context, req *v1.EditConfigReq) (res *v1.EditConfigRes, err error) {
	output, err := service.SysConfig().Edit(ctx, &req.SysConfigEditParam)
	if err != nil {
		return nil, err
	}
	res = &v1.EditConfigRes{
		SysConfigEditModel: *output,
	}
	return
}
func (c *ControllerV1) DeleteConfig(ctx context.Context, req *v1.DeleteConfigReq) (res *v1.DeleteConfigRes, err error) {
	output, err := service.SysConfig().Delete(ctx, &req.SysConfigDeleteParam)
	if err != nil {
		return nil, err
	}
	res = &v1.DeleteConfigRes{
		SysConfigDeleteModel: *output,
	}
	return
}
func (c *ControllerV1) GetConfigById(ctx context.Context, req *v1.GetConfigByIdReq) (res *v1.GetConfigByIdRes, err error) {
	output, err := service.SysConfig().View(ctx, &req.SysConfigViewParam)
	if err != nil {
		return nil, err
	}
	res = &v1.GetConfigByIdRes{
		SysConfigViewModel: *output,
	}
	return
}
