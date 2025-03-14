package system

import (
	"context"

	v1 "xiujieadmin/api/system/v1"
	"xiujieadmin/internal/model/response"
	"xiujieadmin/internal/service"
)

func (c *ControllerV1) SysClientList(ctx context.Context, req *v1.SysClientListReq) (res *v1.SysClientListRes, err error) {
	items, total, err := service.SysClient().List(ctx, &req.SysClientListParam)
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
func (c *ControllerV1) SysClientView(ctx context.Context, req *v1.SysClientViewReq) (res *v1.SysClientViewRes, err error) {
	info, err := service.SysClient().View(ctx, req.SysClientViewParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysClientViewRes{
		SysClientViewModel: info,
	}
	return
}
func (c *ControllerV1) SysClientAdd(ctx context.Context, req *v1.SysClientAddReq) (res *v1.SysClientAddRes, err error) {
	output, err := service.SysClient().Add(ctx, req.SysClientAddParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysClientAddRes{
		SysClientAddModel: output,
	}
	return
}
func (c *ControllerV1) SysClientEdit(ctx context.Context, req *v1.SysClientEditReq) (res *v1.SysClientEditRes, err error) {
	output, err := service.SysClient().Edit(ctx, req.SysClientEditParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysClientEditRes{
		SysClientEditModel: output,
	}
	return
}
func (c *ControllerV1) SysClientDelete(ctx context.Context, req *v1.SysClientDeleteReq) (res *v1.SysClientDeleteRes, err error) {
	output, err := service.SysClient().Delete(ctx, req.SysClientDeleteParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysClientDeleteRes{
		SysClientDeleteModel: output,
	}
	return
}
func (c *ControllerV1) SysClientStatus(ctx context.Context, req *v1.SysClientStatusReq) (res *v1.SysClientStatusRes, err error) {
	output, err := service.SysClient().Status(ctx, req.SysClientStatusParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysClientStatusRes{
		SysClientStatusModel: output,
	}
	return
}
