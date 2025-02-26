package system

import (
	"context"
	"fmt"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/service"
)

type sSysDept struct {
}

func NewSysDept() *sSysDept {
	return &sSysDept{}
}

func init() {
	service.RegisterSysDept(NewSysDept())
}

func (l *sSysDept) GetDeptList(ctx context.Context, query model.DeptListParam) (items []*model.SysDept, total int, err error) {
	// 获取当前用户租户
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, 0, err
	}
	tenantId := claims.BaseClaims.TenantId
	err = dao.SysDept.Ctx(ctx).Where(dao.SysDept.Columns().TenantId, tenantId).Page(1, 5000).Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	total = len(items)
	return
}

func (l *sSysDept) GetDeptById(ctx context.Context, id int64) (dept *model.SysDept, err error) {
	err = dao.SysDept.Ctx(ctx).Where(dao.SysDept.Columns().DeptId, id).Scan(&dept)
	if err != nil {
		return nil, err
	}
	return dept, nil
}

// 构建树结构
func (l *sSysDept) DeptTree(ctx context.Context, parentDept *model.SysDeptTreeModel, deptList []*model.SysDept, ancestors string) (data []*model.SysDeptTreeModel, err error) {
	data = make([]*model.SysDeptTreeModel, 0)
	for _, dept := range deptList {
		pId := int64(0)
		if parentDept != nil {
			pId = parentDept.DeptId
		}
		if dept.ParentId == pId {
			item := &model.SysDeptTreeModel{
				DeptId:    dept.DeptId,
				ParentId:  dept.ParentId,
				DeptName:  dept.DeptName,
				Key:       fmt.Sprintf("%d", dept.DeptId),
				Children:  nil,
				Ancestors: ancestors,
			}
			subDept, err := l.DeptTree(ctx, item, deptList, ancestors+","+fmt.Sprintf("%d", dept.DeptId))
			if err != nil {
				return nil, err
			}
			item.Children = subDept
			data = append(data, item)
		}
	}
	return
}

func (l *sSysDept) GetDeptTree(ctx context.Context) (items []*model.SysDeptTreeModel, err error) {
	depts, _, err := l.GetDeptList(ctx, model.DeptListParam{})
	if err != nil {
		return nil, err
	}
	items, err = l.DeptTree(ctx, nil, depts, "0")
	if err != nil {
		return nil, err
	}
	return items, nil
}

// 递归构建结构
func (l *sSysDept) RecursionDeptIds(ctx context.Context, parentId int64, deptList []*model.SysDept, data *[]int64) (err error) {
	for _, dept := range deptList {
		if dept.ParentId == parentId {
			*data = append(*data, dept.DeptId)
			err := l.RecursionDeptIds(ctx, dept.DeptId, deptList, data)
			if err != nil {
				return err
			}
		}
	}
	return
}

// 根据父部门id获取部门列表
func (l *sSysDept) GetDeptIdsByParentId(ctx context.Context, parentId int64) (ids []int64, err error) {
	depts, _, err := l.GetDeptList(ctx, model.DeptListParam{})
	if err != nil {
		return nil, err
	}
	ids = make([]int64, 0)
	err = l.RecursionDeptIds(ctx, parentId, depts, &ids)
	if err != nil {
		return nil, err
	}
	// g.Log().Infof(ctx, "sSysDept.GetDeptIdsByParentId 获取机构列表: ids:%+v", ids)
	return ids, nil
}
