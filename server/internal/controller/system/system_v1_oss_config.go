package system

import (
	"context"

	v1 "xiujieadmin/api/system/v1"
	"xiujieadmin/internal/model/response"
	"xiujieadmin/internal/service"
)

func (c *ControllerV1) SysOssConfigList(ctx context.Context, req *v1.SysOssConfigListReq) (res *v1.SysOssConfigListRes, err error) {
	items, total, err := service.SysOssConfig().List(ctx, req.SysOssConfigListParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysOssConfigListRes{
		Items: items,
		PageResult: response.PageResult{
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	}, nil
}
func (c *ControllerV1) SysOssConfigView(ctx context.Context, req *v1.SysOssConfigViewReq) (res *v1.SysOssConfigViewRes, err error) {
	item, err := service.SysOssConfig().View(ctx, req.SysOssConfigViewParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysOssConfigViewRes{
		SysOssConfigViewModel: item,
	}, nil
}
func (c *ControllerV1) SysOssConfigAdd(ctx context.Context, req *v1.SysOssConfigAddReq) (res *v1.SysOssConfigAddRes, err error) {
	item, err := service.SysOssConfig().Add(ctx, req.SysOssConfigAddParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysOssConfigAddRes{
		SysOssConfigAddModel: item,
	}, nil
}
func (c *ControllerV1) SysOssConfigEdit(ctx context.Context, req *v1.SysOssConfigEditReq) (res *v1.SysOssConfigEditRes, err error) {
	item, err := service.SysOssConfig().Edit(ctx, req.SysOssConfigEditParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysOssConfigEditRes{
		SysOssConfigEditModel: item,
	}, nil
}

func (c *ControllerV1) SysOssConfigDelete(ctx context.Context, req *v1.SysOssConfigDeleteReq) (res *v1.SysOssConfigDeleteRes, err error) {
	item, err := service.SysOssConfig().Delete(ctx, req.SysOssConfigDeleteParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysOssConfigDeleteRes{
		SysOssConfigDeleteModel: item,
	}, nil
}
