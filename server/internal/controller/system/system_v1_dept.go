package system

import (
	"context"

	v1 "xiuadmin/api/system/v1"
	"xiuadmin/internal/model/response"
	"xiuadmin/internal/service"
)

func (c *ControllerV1) DeptList(ctx context.Context, req *v1.DeptListReq) (res *v1.DeptListRes, err error) {
	items, total, err := service.SysDept().GetDeptList(ctx, req.SysDeptListParam)
	if err != nil {
		return nil, err
	}
	return &v1.DeptListRes{
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
		Data: items,
	}, nil
}

func (c *ControllerV1) DeptTree(ctx context.Context, req *v1.DeptTreeReq) (res *v1.DeptTreeRes, err error) {
	items, err := service.SysDept().GetDeptTree(ctx, req.SysDeptTreeParam)
	if err != nil {
		return nil, err
	}
	return &v1.DeptTreeRes{
		Data: items,
	}, nil
}
func (c *ControllerV1) DeptAdd(ctx context.Context, req *v1.DeptAddReq) (res *v1.DeptAddRes, err error) {
	deptId, err := service.SysDept().AddDept(ctx, &req.SysDeptAddModel)
	if err != nil {
		return nil, err
	}
	return &v1.DeptAddRes{
		DeptId: deptId,
	}, nil
}
func (c *ControllerV1) DeptEdit(ctx context.Context, req *v1.DeptEditReq) (res *v1.DeptEditRes, err error) {
	deptId, err := service.SysDept().EditDept(ctx, &req.SysDeptEditModel)
	if err != nil {
		return nil, err
	}
	return &v1.DeptEditRes{
		DeptId: deptId,
	}, nil
}
func (c *ControllerV1) DeptDelete(ctx context.Context, req *v1.DeptDeleteReq) (res *v1.DeptDeleteRes, err error) {
	deptId, err := service.SysDept().DeleteDept(ctx, &req.SysDeptDeleteModel)
	if err != nil {
		return nil, err
	}
	return &v1.DeptDeleteRes{
		DeptId: deptId,
	}, nil
}
func (c *ControllerV1) DeptView(ctx context.Context, req *v1.DeptViewReq) (res *v1.DeptViewRes, err error) {
	dept, err := service.SysDept().GetDeptById(ctx, req.SysDeptViewModel.DeptId)
	if err != nil {
		return nil, err
	}
	return &v1.DeptViewRes{
		SysDeptViewModel: *dept,
	}, nil
}
