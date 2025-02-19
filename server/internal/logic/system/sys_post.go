package system

import (
	"context"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/request"
	"server/internal/service"
)

type sSysPost struct {
}

func NewSysPost() *sSysPost {
	return &sSysPost{}
}

func init() {
	service.RegisterSysPost(NewSysPost())
}

func (l *sSysPost) GetPostList(ctx context.Context, query model.SysPostListQuery, pageInfo request.PageInfo) (items []*model.SysPost, total int, err error) {

	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, 0, err
	}
	tenantId := claims.BaseClaims.TenantId

	m := dao.SysPost.Ctx(ctx).Where(dao.SysPost.Columns().TenantId, tenantId)
	if query.PostCode != "" {
		m = m.WhereLike(dao.SysPost.Columns().PostCode, "%"+query.PostCode+"%")
	}
	if query.PostName != "" {
		m = m.WhereLike(dao.SysPost.Columns().PostName, "%"+query.PostName+"%")
	}
	total, err = m.Count()
	if err != nil {
		return nil, 0, err
	}
	err = m.Page(pageInfo.Page, pageInfo.PageSize).Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return
}
