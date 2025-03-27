// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"errors"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/do"
	"xiuadmin/internal/model/request"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
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

func (s *sSysUserOnline) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.SysUserOnline.Ctx(ctx), option...)
}

func (s *sSysUserOnline) Add(ctx context.Context, userOnline *model.SysUserOnlineAddModel) (err error) {
	data := do.SysUserOnline{}
	gconv.Struct(userOnline, &data)
	_, err = s.Model(ctx).Data(data).OmitEmpty().Save()
	return
}

func (s *sSysUserOnline) List(ctx context.Context, query *model.SysUserOnlineListParam, page *request.PageInfo) (items []*model.SysUserOnlineListModel, total int, err error) {
	tenantId := contexts.GetTenantId(ctx)
	db := s.Model(ctx).Where(dao.SysUserOnline.Columns().TenantId, tenantId)
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
	if query.Token != "" {
		db = db.Where(dao.SysUserOnline.Columns().Token, query.Token)
	}
	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}
	err = db.Page(page.Page, page.PageSize).Order(dao.SysUserOnline.Columns().LoginTime, "desc").Scan(&items)
	return
}

func (s *sSysUserOnline) Delete(ctx context.Context, ids []int64) (err error) {
	if len(ids) == 0 {
		return errors.New("ids不能为空")
	}
	var data = make([]*model.SysUserOnlineViewModel, 0)
	err = s.Model(ctx).WhereIn(dao.SysUserOnline.Columns().OnlineId, ids).Scan(&data)
	if err != nil {
		return err
	}
	for _, v := range data {
		service.SysAuth().DeleteToken(ctx, v.Token)
	}
	_, err = s.Model(ctx).WhereIn(dao.SysUserOnline.Columns().OnlineId, ids).Delete()
	if err != nil {
		return err
	}
	return
}

func (s *sSysUserOnline) DeleteByToken(ctx context.Context, token string) (err error) {
	var data *model.SysUserOnlineViewModel
	err = s.Model(ctx).Where(dao.SysUserOnline.Columns().Token, token).Scan(&data)
	if err != nil {
		return err
	}
	if data != nil && data.Token != "" {
		service.SysAuth().DeleteToken(ctx, data.Token)
	}
	_, err = s.Model(ctx).Where(dao.SysUserOnline.Columns().Token, token).Delete()
	return
}
