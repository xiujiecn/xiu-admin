package info

import (
	"context"

	v1 "xiuadmin/api/info/v1"
	"xiuadmin/internal/service"
)

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	info := service.Info().GetInfo()
	return &v1.InfoRes{
		InfoModel: info,
	}, nil
}
