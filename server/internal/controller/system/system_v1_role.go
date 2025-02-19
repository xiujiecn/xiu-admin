package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/model/response"
	"server/internal/service"
)

func (c *ControllerV1) RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	data, total, err := service.SysRole().GetRoleList(ctx, &req.SysRoleListQuery, &req.PageInfo)
	if err != nil {
		return nil, err
	}
	return &v1.RoleListRes{
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
		Data: data,
	}, nil
}
