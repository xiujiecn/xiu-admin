// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"errors"
	"fmt"
	"strings"
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
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysDictType struct{}

func NewSysDictType() *sSysDictType {
	return &sSysDictType{}
}

func init() {
	service.RegisterSysDictType(NewSysDictType())
}

func (l *sSysDictType) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	model := dao.SysDictType.Ctx(ctx)
	model = model.Where("is_sys = '0' OR tenant_id = ?", contexts.GetTenantId(ctx))
	return model
}

func (l *sSysDictType) ModelQuery(ctx context.Context, option ...*handler.Option) *gdb.Model {
	model := service.MemoryDB().DB(ctx).Ctx(ctx).Model(dao.SysDictType.Table())
	model = model.Where("is_sys = '0' OR tenant_id = ?", contexts.GetTenantId(ctx))
	return model
}

func (s *sSysDictType) List(ctx context.Context, param *model.SysDictTypeListParam) (items []model.SysDictTypeListModel, total int, err error) {
	m := s.ModelQuery(ctx)
	if param.DictName != "" {
		m = m.WhereLike(dao.SysDictType.Columns().DictName, "%"+param.DictName+"%")
	}
	if param.DictType != "" {
		m = m.WhereLike(dao.SysDictType.Columns().DictType, "%"+param.DictType+"%")
	}
	if param.IsSys != "" {
		m = m.Where(dao.SysDictType.Columns().IsSys, param.IsSys)
	}
	if len(param.CreatedAt) == 2 {
		createdAt := gtime.NewFromStr(param.CreatedAt[0])
		createdAtEnd := gtime.NewFromStr(param.CreatedAt[1])
		m = m.WhereBetween(dao.SysDictType.Columns().CreatedAt, createdAt, createdAtEnd.EndOfDay())
	}
	total, err = m.Count()
	if err != nil {
		return nil, 0, err
	}
	err = m.Order(dao.SysDictType.Columns().DictId, "DESC").Page(param.Page, param.PageSize).Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return
}

func (s *sSysDictType) View(ctx context.Context, param *model.SysDictTypeViewParam) (data *model.SysDictTypeViewModel, err error) {
	m := s.ModelQuery(ctx)

	if param.DictId != 0 {
		m = m.Where(dao.SysDictType.Columns().DictId, param.DictId)
	} else if param.DictType != "" {
		m = m.Where(dao.SysDictType.Columns().DictType, param.DictType)
	} else {
		return nil, errors.New("字典类型ID不能为空")
	}
	err = m.Scan(&data)
	if err != nil {
		return nil, err
	}
	return
}

func (s *sSysDictType) Add(ctx context.Context, param *model.SysDictTypeAddParam) (output *model.SysDictTypeAddModel, err error) {
	// 如果设置为系统字典类型，需要检查权限
	if param.IsSys == "0" {
		if !contexts.IsSuperAdmin(ctx) {
			return nil, errors.New("非超级管理员不能添加系统字典类型")
		}
	}
	m := s.Model(ctx)
	data := do.SysDictType{}
	gconv.Struct(param, &data)
	data.TenantId = contexts.GetTenantId(ctx)
	data.CreatedAt = gtime.Now()
	data.CreatedBy = contexts.GetUserId(ctx)
	data.CreatedDept = contexts.GetDeptId(ctx)
	m = m.Data(data).OmitNil()
	output = &model.SysDictTypeAddModel{}
	output.DictId, err = m.InsertAndGetId()
	if err != nil {
		return nil, err
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysDictTypeCreate, output.DictId)
	return
}

func (s *sSysDictType) Edit(ctx context.Context, param *model.SysDictTypeEditParam) (output *model.SysDictTypeEditModel, err error) {
	if param.DictId == 0 {
		g.Log().Errorf(ctx, "字典类型ID不能为空,param:%+v", param)
		return nil, errors.New("字典类型ID不能为空")
	}

	// 如果设置为系统字典类型，需要检查权限
	if param.IsSys == "0" {
		if !contexts.IsSuperAdmin(ctx) {
			return nil, errors.New("非超级管理员不能修改为系统字典类型")
		}
	}

	m := s.Model(ctx)
	data := do.SysDictType{}
	gconv.Struct(param, &data)
	data.UpdatedAt = gtime.Now()
	data.UpdatedBy = contexts.GetUserId(ctx)
	m = m.Data(data).OmitNil().Where(dao.SysDictType.Columns().DictId, param.DictId)
	output = &model.SysDictTypeEditModel{
		DictId: param.DictId,
	}
	_, err = m.Update()
	if err != nil {
		return nil, err
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysDictTypeUpdate, param.DictId)
	return
}

func (s *sSysDictType) Delete(ctx context.Context, param *model.SysDictTypeDeleteParam) (output *model.SysDictTypeDeleteModel, err error) {
	if len(param.DictIds) == 0 {
		return nil, errors.New("字典类型ID不能为空")
	}

	// 检查是否有系统内置字典类型
	var systemDictTypes []entity.SysDictType
	err = s.Model(ctx).WhereIn(dao.SysDictType.Columns().DictId, param.DictIds).
		Where(dao.SysDictType.Columns().IsSys, "0").
		Scan(&systemDictTypes)
	if err != nil {
		return nil, err
	}

	// 如果有系统内置字典类型，检查当前用户权限
	if len(systemDictTypes) > 0 {
		if !contexts.IsSuperAdmin(ctx) {
			var dictNames []string
			for _, dict := range systemDictTypes {
				dictNames = append(dictNames, dict.DictName)
			}
			return nil, fmt.Errorf("系统内置字典类型 [%s] 只有超级管理员可以删除", strings.Join(dictNames, ", "))
		}
	}

	m := s.Model(ctx)
	m = m.WhereIn(dao.SysDictType.Columns().DictId, param.DictIds)
	_, err = m.Delete()
	if err != nil {
		return nil, err
	}
	output = &model.SysDictTypeDeleteModel{
		DictIds: param.DictIds,
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysDictTypeDelete, param.DictIds)
	return
}
