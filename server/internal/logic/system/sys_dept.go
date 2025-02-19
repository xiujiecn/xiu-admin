package system

import (
	"context"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/request"
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

func (l *sSysDept) GetDeptList(ctx context.Context, page request.PageInfo, query model.DeptListQuery) (items []*model.SysDept, total int, err error) {
	return
}

func (l *sSysDept) GetDeptById(ctx context.Context, id int64) (dept *model.SysDept, err error) {
	err = dao.SysDept.Ctx(ctx).Where(dao.SysDept.Columns().DeptId, id).Scan(&dept)
	if err != nil {
		return nil, err
	}
	return dept, nil
}
