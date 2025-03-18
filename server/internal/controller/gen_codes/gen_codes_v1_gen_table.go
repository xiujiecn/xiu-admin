package gen_codes

import (
	"context"

	v1 "xiujieadmin/api/gen_codes/v1"
	"xiujieadmin/internal/model/response"
	"xiujieadmin/internal/service"
)

func (c *ControllerV1) SysGenTableList(ctx context.Context, req *v1.SysGenTableListReq) (res *v1.SysGenTableListRes, err error) {
	items, total, err := service.SysGenTable().List(ctx, req.SysGenTableListParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysGenTableListRes{
		Items: items,
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}
	return res, nil
}

func (c *ControllerV1) SysGenTableView(ctx context.Context, req *v1.SysGenTableViewReq) (res *v1.SysGenTableViewRes, err error) {
	item, err := service.SysGenTable().View(ctx, req.SysGenTableViewParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysGenTableViewRes{
		SysGenTableViewModel: item,
	}
	return res, nil
}

func (c *ControllerV1) SysGenTableAdd(ctx context.Context, req *v1.SysGenTableAddReq) (res *v1.SysGenTableAddRes, err error) {
	output, err := service.SysGenTable().Add(ctx, req.SysGenTableAddParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysGenTableAddRes{
		SysGenTableAddModel: output,
	}
	return res, nil
}
