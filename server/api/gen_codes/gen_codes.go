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
}
