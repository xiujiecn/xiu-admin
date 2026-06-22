// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package info

import (
	"context"

	"xiuadmin/api/info/v1"
)

type IInfoV1 interface {
	Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error)
}
