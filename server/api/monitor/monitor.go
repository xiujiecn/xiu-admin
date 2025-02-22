// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package monitor

import (
	"context"

	"server/api/monitor/v1"
)

type IMonitorV1 interface {
	RedisInfo(ctx context.Context, req *v1.RedisInfoReq) (res *v1.RedisInfoRes, err error)
}
