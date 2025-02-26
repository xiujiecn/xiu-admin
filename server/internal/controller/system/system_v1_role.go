package system

import (
	"context"

	v1 "xiujieadmin/api/system/v1"
	"xiujieadmin/internal/model/response"
	"xiujieadmin/internal/service"
)

func (c *ControllerV1) RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	data, total, err := service.SysRole().GetRoleList(ctx, &req.SysRoleListParam, &req.PageInfo)
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
func (c *ControllerV1) RoleAdd(ctx context.Context, req *v1.RoleAddReq) (res *v1.RoleAddRes, err error) {
	err = service.SysRole().AddRole(ctx, &req.SysRoleAddParam)
	if err != nil {
		return nil, err
	}
	return &v1.RoleAddRes{}, nil
}
func (c *ControllerV1) RoleEdit(ctx context.Context, req *v1.RoleEditReq) (res *v1.RoleEditRes, err error) {
	err = service.SysRole().EditRole(ctx, &req.SysRoleEditParam)
	if err != nil {
		return nil, err
	}
	return &v1.RoleEditRes{}, nil
}

func (c *ControllerV1) RoleDelete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error) {
	err = service.SysRole().DeleteRole(ctx, &req.SysRoleDeleteParam)
	if err != nil {
		return nil, err
	}
	return &v1.RoleDeleteRes{}, nil
}

func (c *ControllerV1) RoleView(ctx context.Context, req *v1.RoleViewReq) (res *v1.RoleViewRes, err error) {
	record, err := service.SysRole().GetRoleView(ctx, req.SysRoleViewParam.RoleId)
	if err != nil {
		return nil, err
	}
	return &v1.RoleViewRes{
		SysRoleViewModel: *record,
	}, nil
}
func (c *ControllerV1) RoleDataScopeEdit(ctx context.Context, req *v1.RoleDataScopeEditReq) (res *v1.RoleDataScopeEditRes, err error) {
	err = service.SysRole().EditRoleDataScope(ctx, &req.SysRoleDataScopeEditParam)
	if err != nil {
		return nil, err
	}
	return &v1.RoleDataScopeEditRes{}, nil
}
