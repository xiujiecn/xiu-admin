package system

import (
	"context"
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

func (l *sSysDept) GetDeptList(ctx context.Context, query model.DeptListQuery) (items []*model.SysDept, total int, err error) {
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
