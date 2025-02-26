package system

import (
	"context"

	v1 "xiujieadmin/api/system/v1"
	"xiujieadmin/internal/model/response"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) SysNoticeList(ctx context.Context, req *v1.SysNoticeListReq) (res *v1.SysNoticeListRes, err error) {
	g.Log().Infof(ctx, "SystemControl.SysNoticeList %+v", req)
	items, total, err := service.SysNotice().GetNoticeList(ctx, &req.SysNoticeListParam, &req.PageInfo)
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
