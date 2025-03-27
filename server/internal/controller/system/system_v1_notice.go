package system

import (
	"context"

	v1 "xiuadmin/api/system/v1"
	"xiuadmin/internal/model/response"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) SysNoticeList(ctx context.Context, req *v1.SysNoticeListReq) (res *v1.SysNoticeListRes, err error) {
	g.Log().Infof(ctx, "SystemControl.SysNoticeList %+v", req)
	items, total, err := service.SysNotice().List(ctx, req.SysNoticeListParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysNoticeListRes{
		Items: items,
		PageResult: response.PageResult{
			Total:    total,
			Page:     req.PageInfo.Page,
			PageSize: req.PageInfo.PageSize,
		},
	}
	return
}
func (c *ControllerV1) SysNoticeAdd(ctx context.Context, req *v1.SysNoticeAddReq) (res *v1.SysNoticeAddRes, err error) {
	err = service.SysNotice().Add(ctx, req.SysNoticeAddParam)
	if err != nil {
		return nil, err
	}
	return
}
func (c *ControllerV1) SysNoticeEdit(ctx context.Context, req *v1.SysNoticeEditReq) (res *v1.SysNoticeEditRes, err error) {
	err = service.SysNotice().Edit(ctx, req.SysNoticeEditParam)
	if err != nil {
		return nil, err
	}
	return
}
func (c *ControllerV1) SysNoticeDelete(ctx context.Context, req *v1.SysNoticeDeleteReq) (res *v1.SysNoticeDeleteRes, err error) {
	err = service.SysNotice().Delete(ctx, req.SysNoticeDeleteParam)
	if err != nil {
		return nil, err
	}
	return
}
func (c *ControllerV1) SysNoticeView(ctx context.Context, req *v1.SysNoticeViewReq) (res *v1.SysNoticeViewRes, err error) {
	data, err := service.SysNotice().View(ctx, req.SysNoticeViewParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysNoticeViewRes{
		SysNoticeViewModel: data,
	}
	return
}
