package gen_codes

import (
	"context"

	v1 "xiuadmin/api/gen_codes/v1"
	"xiuadmin/internal/model/response"
	"xiuadmin/internal/service"
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
func (c *ControllerV1) SysGenTableDelete(ctx context.Context, req *v1.SysGenTableDeleteReq) (res *v1.SysGenTableDeleteRes, err error) {
	output, err := service.SysGenTable().Delete(ctx, req.SysGenTableDeleteParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysGenTableDeleteRes{
		SysGenTableDeleteModel: output,
	}
	return res, nil
}
func (c *ControllerV1) SysGenTableSelects(ctx context.Context, req *v1.SysGenTableSelectsReq) (res *v1.SysGenTableSelectsRes, err error) {
	output, err := service.SysGenTable().Selects(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.SysGenTableSelectsRes{
		SelectsModel: output,
	}
	return res, nil
}
func (c *ControllerV1) SysGenTableTableSelect(ctx context.Context, req *v1.SysGenTableTableSelectReq) (res *v1.SysGenTableTableSelectRes, err error) {
	output, err := service.SysGenTable().TableSelect(ctx, req.GenCodesTableSelectParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysGenTableTableSelectRes{
		Items: output,
	}
	return res, nil
}
func (c *ControllerV1) SysGenTableColumnSelect(ctx context.Context, req *v1.SysGenTableColumnSelectReq) (res *v1.SysGenTableColumnSelectRes, err error) {
	output, err := service.SysGenTable().ColumnSelect(ctx, req.GenCodesColumnSelectParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysGenTableColumnSelectRes{
		Items: output,
	}
	return res, nil
}
func (c *ControllerV1) SysGenTablePreview(ctx context.Context, req *v1.SysGenTablePreviewReq) (res *v1.SysGenTablePreviewRes, err error) {
	output, err := service.SysGenTable().Preview(ctx, req.GenCodesPreviewParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysGenTablePreviewRes{
		GenCodesPreviewModel: output,
	}
	return res, nil
}
func (c *ControllerV1) SysGenTableBuild(ctx context.Context, req *v1.SysGenTableBuildReq) (res *v1.SysGenTableBuildRes, err error) {
	output, err := service.SysGenTable().Build(ctx, req.GenCodesBuildParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysGenTableBuildRes{
		GenCodesBuildModel: output,
	}
	return res, nil
}
func (c *ControllerV1) SysGenTableColumnList(ctx context.Context, req *v1.SysGenTableColumnListReq) (res *v1.SysGenTableColumnListRes, err error) {
	output, err := service.SysGenTable().ColumnList(ctx, req.GenCodesColumnListParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysGenTableColumnListRes{
		Items: output,
	}
	return res, nil
}
func (c *ControllerV1) SysGenTableEdit(ctx context.Context, req *v1.SysGenTableEditReq) (res *v1.SysGenTableEditRes, err error) {
	output, err := service.SysGenTable().Edit(ctx, req.SysGenTableEditParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysGenTableEditRes{
		SysGenTableEditModel: output,
	}
	return res, nil
}
