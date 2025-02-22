package system

import (
	"context"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/request"
	"server/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
)

type sSysLogininfor struct{}

func NewSysLogininfor() *sSysLogininfor {
	return &sSysLogininfor{}
}

func (s *sSysLogininfor) ListLogininfor(ctx context.Context, query *model.SysLogininforListQuery, pageInfo *request.PageInfo) (items []*model.SysLogininforListModel, total int, err error) {
	// 获取当前用户租户编码
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, 0, err
	}
	tenantId := claims.TenantId
	m := dao.SysLogininfor.Ctx(ctx).Where(dao.SysLogininfor.Columns().TenantId, tenantId)
	if query.Ipaddr != "" {
		m = m.WhereLike(dao.SysLogininfor.Columns().Ipaddr, "%"+query.Ipaddr+"%")
	}
	if query.UserName != "" {
		m = m.WhereLike(dao.SysLogininfor.Columns().UserName, "%"+query.UserName+"%")
	}
	if query.Status != "" {
		m = m.Where(dao.SysLogininfor.Columns().Status, query.Status)
	}
	if len(query.LoginTime) == 2 {
		startTime := gtime.NewFromStr(query.LoginTime[0])
		endTime := gtime.NewFromStr(query.LoginTime[1])
		m = m.WhereBetween(dao.SysLogininfor.Columns().LoginTime, startTime, endTime.EndOfDay())
	}
	total, err = m.Count()
	if err != nil {
		return nil, 0, err
	}
	err = m.Page(pageInfo.Page, pageInfo.PageSize).Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
