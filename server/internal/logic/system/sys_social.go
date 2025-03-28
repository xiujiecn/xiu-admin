// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/request"
	"xiuadmin/internal/service"
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
