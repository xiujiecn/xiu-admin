package system

import (
	"context"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/request"
	"server/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
)

type sSysUserOnline struct {
}

func NewSysUserOnline() *sSysUserOnline {
	return &sSysUserOnline{}
}

func init() {
	service.RegisterSysUserOnline(NewSysUserOnline())
}

func (s *sSysUserOnline) Add(ctx context.Context, userOnline *model.SysUserOnlineAddModel) (err error) {
	_, err = dao.SysUserOnline.Ctx(ctx).Data(userOnline).OmitEmpty().Save()
	return
}

func (s *sSysUserOnline) List(ctx context.Context, query *model.SysUserOnlineListQuery, page *request.PageInfo) (items []*model.SysUserOnlineListModel, total int, err error) {
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, 0, err
	}
	tenantId := claims.TenantId
	db := dao.SysUserOnline.Ctx(ctx).Where(dao.SysUserOnline.Columns().TenantId, tenantId)
	if query.UserName != "" {
		db = db.Where(dao.SysUserOnline.Columns().UserName, query.UserName)
	}
	if query.ClientKey != "" {
		db = db.WhereLike(dao.SysUserOnline.Columns().ClientKey, "%"+query.ClientKey+"%")
	}
	if query.DeviceType != "" {
		db = db.Where(dao.SysUserOnline.Columns().DeviceType, query.DeviceType)
	}
	if query.Ipaddr != "" {
		db = db.Where(dao.SysUserOnline.Columns().Ipaddr, query.Ipaddr)
	}
	if len(query.LoginTime) == 2 {
		startTime := gtime.NewFromStr(query.LoginTime[0])
		endTime := gtime.NewFromStr(query.LoginTime[1])
		db = db.WhereBetween(dao.SysUserOnline.Columns().LoginTime, startTime, endTime.EndOfDay())
	}
	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}
	err = db.Page(page.Page, page.PageSize).Order(dao.SysUserOnline.Columns().LoginTime, "desc").Scan(&items)
	return
}

func (s *sSysUserOnline) Delete(ctx context.Context, id int64) (err error) {
	_, err = dao.SysUserOnline.Ctx(ctx).Where(dao.SysUserOnline.Columns().OnlineId, id).Delete()
	return
}
