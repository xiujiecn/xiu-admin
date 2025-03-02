package system

import (
	"context"

	v1 "xiujieadmin/api/system/v1"
	"xiujieadmin/internal/model/response"
	"xiujieadmin/internal/service"
)

func (c *ControllerV1) RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	data, total, err := service.SysRole().List(ctx, &req.SysRoleListParam)
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
	role, err := service.SysRole().Add(ctx, &req.SysRoleAddParam)
	if err != nil {
		return nil, err
	}
	return &v1.RoleAddRes{
		SysRoleAddModel: *role,
	}, nil
}
func (c *ControllerV1) RoleEdit(ctx context.Context, req *v1.RoleEditReq) (res *v1.RoleEditRes, err error) {
	role, err := service.SysRole().Edit(ctx, &req.SysRoleEditParam)
	if err != nil {
		return nil, err
	}
	return &v1.RoleEditRes{
		SysRoleEditModel: *role,
	}, nil
}

func (c *ControllerV1) RoleDelete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error) {
	role, err := service.SysRole().Delete(ctx, &req.SysRoleDeleteParam)
	if err != nil {
		return nil, err
	}
	return &v1.RoleDeleteRes{
		SysRoleDeleteModel: *role,
	}, nil
}

func (c *ControllerV1) RoleView(ctx context.Context, req *v1.RoleViewReq) (res *v1.RoleViewRes, err error) {
	record, err := service.SysRole().View(ctx, &req.SysRoleViewParam)
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
