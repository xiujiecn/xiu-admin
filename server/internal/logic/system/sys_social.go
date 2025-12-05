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

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sSysSocial struct {
}

func NewSysSocial() *sSysSocial {
	return &sSysSocial{}
}

func init() {
	service.RegisterSysSocial(NewSysSocial())
}

// 查询社会化关系表
func (s *sSysSocial) List(ctx context.Context, query *model.SysSocialListParam, page *request.PageInfo) (items []*model.SysSocialListModel, total int, err error) {
	if query.UserId == 0 {
		query.UserId = contexts.GetUserId(ctx)
	}

	db := dao.SysSocial.Ctx(ctx)
	if query.UserId != 0 {
		db = db.Where(dao.SysSocial.Columns().UserId, query.UserId)
	}
	if query.Source != "" {
		db = db.Where(dao.SysSocial.Columns().Source, query.Source)
	}
	if query.OpenId != "" {
		db = db.Where(dao.SysSocial.Columns().OpenId, query.OpenId)
	}
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

// 删除社会化关系表
func (s *sSysSocial) Delete(ctx context.Context, id int64) (err error) {
	userId := contexts.GetUserId(ctx)
	_, err = dao.SysSocial.Ctx(ctx).WherePri(id).Where(dao.SysSocial.Columns().UserId, userId).Delete()
	return err
}

func (s *sSysSocial) Create(ctx context.Context, social *model.SysSocialSaveParam) (err error) {
	g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 查询是否存在
		exist, err := dao.SysSocial.Ctx(ctx).Where(dao.SysSocial.Columns().OpenId, social.OpenId).Where(dao.SysSocial.Columns().Source, social.Source).Where(dao.SysSocial.Columns().UserId, social.UserId).Exist()
		if err != nil {
			return err
		}
		if exist {
			// 绑定信息已存在
			g.Log().Errorf(ctx, "绑定信息已存在, openId: %s, source: %s, userId: %d", social.OpenId, social.Source, social.UserId)
			return gerror.New("绑定信息已存在")
		}

		_, err = dao.SysSocial.Ctx(ctx).Data(social).OmitEmpty().Save()
		return err
	})
	return err
}
