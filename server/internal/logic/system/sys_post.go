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

func (l *sSysPost) GetPostList(ctx context.Context, query model.SysPostListParam, pageInfo request.PageInfo) (items []*model.SysPostListModel, total int, err error) {

	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, 0, err
	}
	tenantId := claims.BaseClaims.TenantId

	deptIds := make([]int64, 0)
	if query.DeptId != 0 {
		deptIds = append(deptIds, query.DeptId)
	} else if query.BelongDeptId != 0 {
		deptIds = append(deptIds, query.BelongDeptId)
		subDeptIds, err := service.SysDept().GetDeptIdsByParentId(ctx, query.BelongDeptId)
		if err != nil {
			return nil, 0, err
		}
		deptIds = append(deptIds, subDeptIds...)
	}
	m := dao.SysPost.Ctx(ctx).Where(dao.SysPost.Columns().TenantId, tenantId)
	if query.PostCode != "" {
		m = m.WhereLike(dao.SysPost.Columns().PostCode, "%"+query.PostCode+"%")
	}
	if query.PostName != "" {
		m = m.WhereLike(dao.SysPost.Columns().PostName, "%"+query.PostName+"%")
	}
	if len(deptIds) > 0 {
		m = m.WhereIn(dao.SysPost.Columns().DeptId, deptIds)
	}

	total, err = m.Count()
	if err != nil {
		return nil, 0, err
	}
	err = m.WithAll().Page(pageInfo.Page, pageInfo.PageSize).OrderAsc(dao.SysPost.Columns().PostSort).Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return
}
