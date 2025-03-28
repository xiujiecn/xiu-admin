// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"errors"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/do"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysLogininfor struct{}

func NewSysLogininfor() *sSysLogininfor {
	return &sSysLogininfor{}
}

func init() {
	service.RegisterSysLogininfor(NewSysLogininfor())
}

func (s *sSysLogininfor) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.SysLogininfor.Ctx(ctx), option...)
}

func (s *sSysLogininfor) List(ctx context.Context, param *model.SysLogininforListParam) (items []*model.SysLogininforListModel, total int, err error) {
	m := s.Model(ctx)

	if param.Ipaddr != "" {
		m = m.WhereLike(dao.SysLogininfor.Columns().Ipaddr, "%"+param.Ipaddr+"%")
	}
	if param.UserName != "" {
		m = m.WhereLike(dao.SysLogininfor.Columns().UserName, "%"+param.UserName+"%")
	}
	if param.Status != "" {
		m = m.Where(dao.SysLogininfor.Columns().Status, param.Status)
	}
	if len(param.LoginTime) == 2 {
		startTime := gtime.NewFromStr(param.LoginTime[0])
		endTime := gtime.NewFromStr(param.LoginTime[1])
		m = m.WhereBetween(dao.SysLogininfor.Columns().LoginTime, startTime, endTime.EndOfDay())
	}
	total, err = m.Count()
	if err != nil {
		return nil, 0, err
	}
	err = m.Page(param.Page, param.PageSize).OrderDesc(dao.SysLogininfor.Columns().InfoId).Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (s *sSysLogininfor) AddLogininfor(ctx context.Context, logininfor *model.SysLogininforAddModel) (id int64, err error) {
	data := do.SysLogininfor{}
	gconv.Struct(logininfor, &data)
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

func (s *sSysLogininfor) Delete(ctx context.Context, param *model.SysLogininforDeleteParam) (output *model.SysLogininforDeleteModel, err error) {
	m := s.Model(ctx)
	output = &model.SysLogininforDeleteModel{
		InfoIds: param.InfoIds,
	}
	if len(param.InfoIds) == 0 {
		return nil, errors.New("请选择要删除的配置")
	}
	_, err = m.WhereIn(dao.SysLogininfor.Columns().InfoId, param.InfoIds).Delete()
	if err != nil {
		return
	}
	return
}
