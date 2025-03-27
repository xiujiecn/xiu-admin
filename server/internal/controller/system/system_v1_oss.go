package system

import (
	"context"

	v1 "xiuadmin/api/system/v1"
	"xiuadmin/internal/model/response"
	"xiuadmin/internal/service"
)

func (c *ControllerV1) SysOssList(ctx context.Context, req *v1.SysOssListReq) (res *v1.SysOssListRes, err error) {
	items, total, err := service.SysOss().List(ctx, &req.SysOssListParam, &req.PageInfo)
	if err != nil {
		return nil, err
	}
	res = &v1.SysOssListRes{
		Items: items,
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}
	return
}
func (c *ControllerV1) SysOssView(ctx context.Context, req *v1.SysOssViewReq) (res *v1.SysOssViewRes, err error) {
	info, err := service.SysOss().View(ctx, req.SysOssViewParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysOssViewRes{
		SysOssViewModel: info,
	}
	return
}
func (c *ControllerV1) SysOssDelete(ctx context.Context, req *v1.SysOssDeleteReq) (res *v1.SysOssDeleteRes, err error) {
	_, err = service.SysOss().Delete(ctx, req.SysOssDeleteParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysOssDeleteRes{}
	return
}
