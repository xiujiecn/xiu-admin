package system

import (
	"context"
	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/library/contexts"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/service"
)

type sSysSocial struct {
}

func NewSysSocial() *sSysSocial {
	return &sSysSocial{}
}

func init() {
	service.RegisterSysSocial(NewSysSocial())
}

func (s *sSysSocial) List(ctx context.Context, query *model.SysSocialListParam, page *request.PageInfo) (items []*model.SysSocialListModel, total int, err error) {
	if query.UserId == 0 {
		query.UserId = contexts.GetUserId(ctx)
	}

	db := dao.SysSocial.Ctx(ctx).Where(dao.SysSocial.Columns().UserId, query.UserId)
	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}
	err = db.Page(page.Page, page.PageSize).Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
