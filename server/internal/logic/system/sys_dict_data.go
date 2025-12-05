// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/event"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/do"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysDictData struct{}

func NewSysDictData() *sSysDictData {
	return &sSysDictData{}
}

func init() {
	service.RegisterSysDictData(NewSysDictData())
}

func (l *sSysDictData) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: false,
			FilterAuth:   false,
		})
	}
	return handler.Model(dao.SysDictData.Ctx(ctx), option...)
}

func (l *sSysDictData) ModelQuery(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: false,
			FilterAuth:   false,
		})
	}
	return handler.Model(service.MemoryDB().DB(ctx).Ctx(ctx).Model(dao.SysDictData.Table()), option...)
}

func (s *sSysDictData) List(ctx context.Context, param *model.SysDictDataListParam) (items []model.SysDictDataListModel, total int, err error) {
	typeData := &model.SysDictTypeViewModel{}
	// 获取字典类型
	if param.DictType == "" && param.DictId != 0 {
		typeData, err = service.SysDictType().View(ctx, &model.SysDictTypeViewParam{
			DictId: param.DictId,
		})
		if err != nil {
			return nil, 0, err
		}
		if typeData == nil {
			return nil, 0, gerror.New("字典类型不存在")
		}
		param.DictType = typeData.DictType
	}
	if param.DictType == "" {
		return nil, 0, gerror.New("字典类型不能为空")
	}
	// 获取字典数据
	items = make([]model.SysDictDataListModel, 0)
	m := s.ModelQuery(ctx).Where(dao.SysDictData.Columns().DictType, param.DictType)
	if param.DictValue != "" {
		m = m.Where(dao.SysDictData.Columns().DictValue, param.DictValue)
	}
	total, err = m.Count()
	if err != nil {
		return nil, 0, err
	}
	err = m.Order(dao.SysDictData.Columns().DictSort, "ASC").Page(param.Page, param.PageSize).Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return
}

func (s *sSysDictData) View(ctx context.Context, param *model.SysDictDataViewParam) (data *model.SysDictDataViewModel, err error) {
	if param.DictCode == 0 {
		return nil, gerror.New("字典编码不能为空")
	}

	m := s.ModelQuery(ctx)
	if param.DictCode != 0 {
		m = m.Where(dao.SysDictData.Columns().DictCode, param.DictCode)
	}
	err = m.Scan(&data)
	if err != nil {
		return nil, err
	}
	return
}

func (s *sSysDictData) Add(ctx context.Context, param *model.SysDictDataAddParam) (output *model.SysDictDataAddModel, err error) {
	m := s.Model(ctx)
	data := do.SysDictData{}
	gconv.Struct(param, &data)
	// 通过 DictType 查询字典类型列表
	listParam := &model.SysDictTypeListParam{
		DictType: param.DictType,
	}
	typeList, _, err := service.SysDictType().List(ctx, listParam)
	if err != nil {
		return nil, err
	}
	if len(typeList) == 0 {
		return nil, gerror.New("字典类型不存在")
	}
	isSys := typeList[0].IsSys
	if isSys == "0" && !contexts.IsSuperAdmin(ctx) {
		return nil, gerror.New("非超级管理员不能添加系统内置字典数据")
	}
	data.TenantId = contexts.GetTenantId(ctx)
	data.CreatedAt = gtime.Now()
	data.CreatedBy = contexts.GetUserId(ctx)
	data.CreatedDept = contexts.GetDeptId(ctx)
	m = m.Data(data).OmitNil()
	output = &model.SysDictDataAddModel{}
	output.DictCode, err = m.InsertAndGetId()
	if err != nil {
		return nil, err
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysDictDataCreate, output.DictCode)
	return
}

func (s *sSysDictData) Edit(ctx context.Context, param *model.SysDictDataEditParam) (output *model.SysDictDataEditModel, err error) {
	if param.DictCode == 0 {
		return nil, gerror.New("字典编码不能为空")
	}
	m := s.Model(ctx)
	data := do.SysDictData{}
	gconv.Struct(param, &data)
	data.UpdatedAt = gtime.Now()
	data.UpdatedBy = contexts.GetUserId(ctx)
	m = m.Data(data).OmitNil().Where(dao.SysDictData.Columns().DictCode, param.DictCode)
	output = &model.SysDictDataEditModel{
		DictCode: param.DictCode,
	}
	_, err = m.Update()
	if err != nil {
		return nil, err
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysDictDataUpdate, param.DictCode)
	return
}

func (s *sSysDictData) Delete(ctx context.Context, param *model.SysDictDataDeleteParam) (output *model.SysDictDataDeleteModel, err error) {
	if len(param.DictCodes) == 0 {
		return nil, gerror.New("字典编码不能为空")
	}

	// 检查要删除的字典数据所属的字典类型是否为系统内置类型
	if !contexts.IsSuperAdmin(ctx) {
		// 查询字典数据所属的字典类型
		var dictTypes []string
		err = dao.SysDictData.Ctx(ctx).
			Fields(dao.SysDictData.Columns().DictType).
			WhereIn(dao.SysDictData.Columns().DictCode, param.DictCodes).
			Distinct().
			Scan(&dictTypes)
		if err != nil {
			return nil, err
		}

		// 检查这些字典类型是否有系统内置类型
		var systemTypeCount int
		systemTypeCount, err = dao.SysDictType.Ctx(ctx).
			WhereIn(dao.SysDictType.Columns().DictType, dictTypes).
			Where(dao.SysDictType.Columns().IsSys, "0").
			Count()
		if err != nil {
			return nil, err
		}

		if systemTypeCount > 0 {
			return nil, gerror.New("系统内置字典数据不能删除，请联系管理员")
		}
	}

	m := s.Model(ctx)
	m = m.WhereIn(dao.SysDictData.Columns().DictCode, param.DictCodes)
	_, err = m.Delete()
	if err != nil {
		return nil, err
	}
	output = &model.SysDictDataDeleteModel{
		DictCodes: param.DictCodes,
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysDictDataDelete, param.DictCodes)
	return
}

func (s *sSysDictData) GetDictLabel(ctx context.Context, dictType string, dictCode string) string {
	dictData, _, err := s.List(ctx, &model.SysDictDataListParam{
		DictType:  dictType,
		DictValue: dictCode,
	})
	if err != nil {
		return ""
	}
	if len(dictData) > 0 {
		return dictData[0].DictLabel
	}
	return ""
}

func (s *sSysDictData) GetDictListByTypes(ctx context.Context, dictTypes []string) (dictDataList []model.SysDictDataListModel, err error) {
	dictDataList = make([]model.SysDictDataListModel, 0)
	m := s.ModelQuery(ctx).WhereIn(dao.SysDictData.Columns().DictType, dictTypes)
	err = m.Scan(&dictDataList)
	return
}
