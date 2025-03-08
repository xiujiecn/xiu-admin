package system

import (
	"context"
	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/do"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
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
	data := do.SysUserOnline{}
	gconv.Struct(userOnline, &data)
	_, err = dao.SysUserOnline.Ctx(ctx).Data(data).OmitEmpty().Save()
	return
}

func (s *sSysUserOnline) List(ctx context.Context, query *model.SysUserOnlineListParam, page *request.PageInfo) (items []*model.SysUserOnlineListModel, total int, err error) {
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
