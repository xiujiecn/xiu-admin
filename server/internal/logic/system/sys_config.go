// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"errors"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/event"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/do"
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysConfig struct {
}

func NewSysConfig() *sSysConfig {
	return &sSysConfig{}
}

func init() {
	service.RegisterSysConfig(NewSysConfig())
}

func (l *sSysConfig) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   false,
		})
	}
	return handler.Model(dao.SysConfig.Ctx(ctx), option...)
}

func (l *sSysConfig) ModelQuery(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   false,
		})
	}
	return handler.Model(service.MemoryDB().DB(ctx).Ctx(ctx).Model(dao.SysConfig.Table()), option...)
}

func (s *sSysConfig) List(ctx context.Context, param *model.SysConfigListParam) (items []*model.SysConfigListModel, total int, err error) {
	m := s.ModelQuery(ctx)

	if param.ConfigName != "" {
		m = m.WhereLike(dao.SysConfig.Columns().ConfigName, "%"+param.ConfigName+"%")
	}

	if param.ConfigKey != "" {
		m = m.WhereLike(dao.SysConfig.Columns().ConfigKey, "%"+param.ConfigKey+"%")
	}

	if param.ConfigType != "" {
		m = m.Where(dao.SysConfig.Columns().ConfigType, param.ConfigType)
	}

	if param.ConfigValue != "" {
		m = m.WhereLike(dao.SysConfig.Columns().ConfigValue, "%"+param.ConfigValue+"%")
	}

	if len(param.CreatedAt) > 0 {
		startTime := gtime.NewFromStr(param.CreatedAt[0])
		endTime := gtime.NewFromStr(param.CreatedAt[1])
		m = m.WhereBetween(dao.SysConfig.Columns().CreatedAt, startTime, endTime.EndOfDay())
	}

	total, err = m.Count()
	if err != nil {
		return nil, 0, err
	}

	err = m.Page(param.Page, param.PageSize).Order(dao.SysConfig.Columns().ConfigId, "ASC").Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return
}

func (s *sSysConfig) Add(ctx context.Context, param *model.SysConfigAddParam) (output *model.SysConfigAddModel, err error) {
	if param == nil || param.ConfigName == "" {
		return nil, errors.New("配置名称不能为空")
	}
	m := s.Model(ctx)

	data := &do.SysConfig{}
	gconv.Struct(param, data)
	data.TenantId = contexts.GetTenantId(ctx)
	data.CreatedDept = contexts.GetDeptId(ctx)
	data.CreatedBy = contexts.GetUserId(ctx)
	data.CreatedAt = gtime.Now()
	data.UpdatedBy = contexts.GetUserId(ctx)
	data.UpdatedAt = gtime.Now()

	lastInsertId, err := m.Data(data).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	output = &model.SysConfigAddModel{
		ConfigId: lastInsertId,
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysConfigCreate, lastInsertId, data.TenantId, data.ConfigKey)
	return
}

func (s *sSysConfig) Edit(ctx context.Context, param *model.SysConfigEditParam) (output *model.SysConfigEditModel, err error) {
	if param == nil || param.ConfigId == 0 {
		return nil, errors.New("配置ID不能为空")
	}

	data := &do.SysConfig{}
	gconv.Struct(param, data)
	data.UpdatedBy = contexts.GetUserId(ctx)
	data.UpdatedAt = gtime.Now()
	data.ConfigId = nil

	m := s.Model(ctx)
	_, err = m.Data(data).Where(dao.SysConfig.Columns().ConfigId, param.ConfigId).Update()
	if err != nil {
		return nil, err
	}

	output = &model.SysConfigEditModel{
		ConfigId: param.ConfigId,
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysConfigUpdate, param.ConfigId, contexts.GetTenantId(ctx), data.ConfigKey)
	return
}

func (s *sSysConfig) Delete(ctx context.Context, param *model.SysConfigDeleteParam) (output *model.SysConfigDeleteModel, err error) {
	m := s.Model(ctx)

	if len(param.ConfigIds) == 0 {
		return nil, errors.New("请选择要删除的配置")
	}
	m = m.WhereIn(dao.SysConfig.Columns().ConfigId, param.ConfigIds)
	data := make([]entity.SysConfig, 0)
	err = m.Scan(&data)
	if err != nil {
		return nil, err
	}
	for _, v := range data {
		if v.ConfigType == "Y" {
			return nil, errors.New("系统内置配置不能删除")
		}
	}

	_, err = m.Delete()
	if err != nil {
		return nil, err
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysConfigDelete, param.ConfigIds, contexts.GetTenantId(ctx))
	return &model.SysConfigDeleteModel{
		ConfigIds: param.ConfigIds,
	}, nil
}

func (s *sSysConfig) View(ctx context.Context, param *model.SysConfigViewParam) (output *model.SysConfigViewModel, err error) {
	m := s.ModelQuery(ctx)

	err = m.Where(dao.SysConfig.Columns().ConfigId, param.ConfigId).Scan(&output)
	if err != nil {
		return nil, err
	}

	return
}

func (s *sSysConfig) GetConfigByKey(ctx context.Context, configKey string) (config *entity.SysConfig, err error) {
	m := s.ModelQuery(ctx)
	err = m.Where(dao.SysConfig.Columns().ConfigKey, configKey).Scan(&config)
	if err != nil {
		return nil, err
	}
	return
}
