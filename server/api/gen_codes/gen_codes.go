// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package gen_codes

import (
	"context"

	"xiujieadmin/api/gen_codes/v1"
)

type IGenCodesV1 interface {
	SysGenTableList(ctx context.Context, req *v1.SysGenTableListReq) (res *v1.SysGenTableListRes, err error)
	SysGenTableView(ctx context.Context, req *v1.SysGenTableViewReq) (res *v1.SysGenTableViewRes, err error)
	SysGenTableAdd(ctx context.Context, req *v1.SysGenTableAddReq) (res *v1.SysGenTableAddRes, err error)
	SysGenTableEdit(ctx context.Context, req *v1.SysGenTableEditReq) (res *v1.SysGenTableEditRes, err error)
	SysGenTableDelete(ctx context.Context, req *v1.SysGenTableDeleteReq) (res *v1.SysGenTableDeleteRes, err error)
	SysGenTableSelects(ctx context.Context, req *v1.SysGenTableSelectsReq) (res *v1.SysGenTableSelectsRes, err error)
	SysGenTableTableSelect(ctx context.Context, req *v1.SysGenTableTableSelectReq) (res *v1.SysGenTableTableSelectRes, err error)
	SysGenTableColumnSelect(ctx context.Context, req *v1.SysGenTableColumnSelectReq) (res *v1.SysGenTableColumnSelectRes, err error)
	SysGenTableColumnList(ctx context.Context, req *v1.SysGenTableColumnListReq) (res *v1.SysGenTableColumnListRes, err error)
	SysGenTablePreview(ctx context.Context, req *v1.SysGenTablePreviewReq) (res *v1.SysGenTablePreviewRes, err error)
	SysGenTableBuild(ctx context.Context, req *v1.SysGenTableBuildReq) (res *v1.SysGenTableBuildRes, err error)
}
