// package gen_codes
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package gen_codes

import (
	"context"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/xgen"
	genconsts "xiuadmin/internal/library/xgen/gen_consts"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/do"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysGenTable struct {
}

func SysGenTableNew() *sSysGenTable {
	return &sSysGenTable{}
}

func init() {
	service.RegisterSysGenTable(SysGenTableNew())
}
func (l *sSysGenTable) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: false,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.SysGenTable.Ctx(ctx), option...)
}

// 列表
func (s *sSysGenTable) List(ctx context.Context, param *model.SysGenTableListParam) (output []*model.SysGenTableListModel, total int, err error) {
	m := s.Model(ctx)
	if param.GenType != 0 {
		m = m.Where(dao.SysGenTable.Columns().GenType, param.GenType)
	}
	if param.VarName != "" {
		m = m.Where(dao.SysGenTable.Columns().VarName, param.VarName)
	}
	if param.Status != "" {
		m = m.Where(dao.SysGenTable.Columns().Status, param.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	output = make([]*model.SysGenTableListModel, 0)
	err = m.Page(param.Page, param.PageSize).OrderDesc(dao.SysGenTable.Columns().TableId).Scan(&output)
	if err != nil {
		return
	}
	return
}

// 详情
func (s *sSysGenTable) View(ctx context.Context, param *model.SysGenTableViewParam) (output *model.SysGenTableViewModel, err error) {
	if param.TableId == 0 {
		err = gerror.New("参数错误")
		return
	}
	m := s.Model(ctx)
	err = m.Where(dao.SysGenTable.Columns().TableId, param.TableId).Scan(&output)
	if err != nil {
		return
	}
	return
}

// 新增
func (s *sSysGenTable) Add(ctx context.Context, param *model.SysGenTableAddParam) (output *model.SysGenTableAddModel, err error) {
	data := do.SysGenTable{}
	gconv.Struct(param, &data)
	data.Options = "{}"
	data.MasterColumns = "[]"
	data.Status = "1"
	data.CreatedDept = contexts.GetDeptId(ctx)
	data.CreatedBy = contexts.GetUserId(ctx)
	data.CreatedAt = gtime.Now()
	m := s.Model(ctx)
	id, err := m.Data(data).OmitNil().InsertAndGetId()
	if err != nil {
		return
	}
	output = &model.SysGenTableAddModel{
		TableId: id,
	}
	return
}

// 修改
func (s *sSysGenTable) Edit(ctx context.Context, param *model.SysGenTableEditParam) (output *model.SysGenTableEditModel, err error) {
	if param == nil || param.TableId == 0 {
		err = gerror.New("参数错误")
		return
	}
	data := do.SysGenTable{}
	gconv.Struct(param, &data)
	if param.Options != nil {
		data.Options = gjson.New(*param.Options).MustToJsonString()
	}
	if param.MasterColumns != nil {
		data.MasterColumns = gjson.New(*param.MasterColumns).MustToJsonString()
	}

	data.UpdatedBy = contexts.GetUserId(ctx)
	data.UpdatedAt = gtime.Now()
	m := s.Model(ctx)
	_, err = m.Data(data).OmitNil().Where(dao.SysGenTable.Columns().TableId, param.TableId).Update()
	if err != nil {
		return
	}
	output = &model.SysGenTableEditModel{
		TableId: param.TableId,
	}
	return
}

func (s *sSysGenTable) Delete(ctx context.Context, param *model.SysGenTableDeleteParam) (output *model.SysGenTableDeleteModel, err error) {
	if len(param.TableIds) == 0 {
		err = gerror.New("参数错误")
		return
	}
	m := s.Model(ctx)
	_, err = m.WhereIn(dao.SysGenTable.Columns().TableId, param.TableIds).Delete()
	if err != nil {
		return
	}
	output = &model.SysGenTableDeleteModel{
		TableIds: param.TableIds,
	}
	return
}

// 获取选择项
func (s *sSysGenTable) Selects(ctx context.Context) (output *model.SelectsModel, err error) {
	output, err = xgen.GenCodesSelects(ctx)
	if err != nil {
		return
	}
	return
}

// 获取表选择项
func (s *sSysGenTable) TableSelect(ctx context.Context, param *model.GenCodesTableSelectParam) (output []*model.GenCodesTableSelectModel, err error) {
	output, err = xgen.GenCodesTableSelect(ctx, param)
	if err != nil {
		return
	}
	return
}

// 获取字段选择项
func (s *sSysGenTable) ColumnSelect(ctx context.Context, param *model.GenCodesColumnSelectParam) (output []*model.GenCodesColumnSelectModel, err error) {
	output, err = xgen.GenCodesColumnSelect(ctx, param)
	if err != nil {
		return
	}
	return
}

// 获取字段列表
func (s *sSysGenTable) ColumnList(ctx context.Context, param *model.GenCodesColumnListParam) (output []*model.GenCodesColumnListModel, err error) {
	output, err = xgen.GenCodesColumnList(ctx, param)
	if err != nil {
		return
	}
	return
}

// 预览
func (s *sSysGenTable) Preview(ctx context.Context, param *model.GenCodesPreviewParam) (output *model.GenCodesPreviewModel, err error) {
	output, err = xgen.GenCodesPreview(ctx, param)
	if err != nil {
		return
	}
	return
}

// 构建
func (s *sSysGenTable) Build(ctx context.Context, param *model.GenCodesBuildParam) (output *model.GenCodesBuildModel, err error) {
	editParam := &model.SysGenTableEditParam{}
	gconv.Struct(param, editParam)
	status := genconsts.GenCodesStatusOk
	editParam.Status = &status
	_, err = s.Edit(ctx, editParam)
	if err != nil {
		return
	}
	output, err = xgen.GenCodesBuild(ctx, param)
	if err != nil {
		status = genconsts.GenCodesStatusFail
		editParam.Status = &status
		_, err = s.Edit(ctx, &model.SysGenTableEditParam{
			TableId: editParam.TableId,
			Status:  &status,
		})
		if err != nil {
			return
		}
		return
	}
	return
}
