package system

import (
	"context"
	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
)

type sSysLogininfor struct{}

func NewSysLogininfor() *sSysLogininfor {
	return &sSysLogininfor{}
}

func init() {
	service.RegisterSysLogininfor(NewSysLogininfor())
}

func (s *sSysLogininfor) ListLogininfor(ctx context.Context, query *model.SysLogininforListParam, pageInfo *request.PageInfo) (items []*model.SysLogininforListModel, total int, err error) {
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

func (s *sSysLogininfor) AddLogininfor(ctx context.Context, logininfor *model.SysLogininforAddModel) (id int64, err error) {
	result, err := dao.SysLogininfor.Ctx(ctx).Data(logininfor).Save()
	if err != nil {
		return 0, err
	}
	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
