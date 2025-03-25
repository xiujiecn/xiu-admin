// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"xiujieadmin/internal/library/xgorm/handler"
	"xiujieadmin/internal/model"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISysGenTable interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// 列表
		List(ctx context.Context, param *model.SysGenTableListParam) (output []*model.SysGenTableListModel, total int, err error)
		// 详情
		View(ctx context.Context, param *model.SysGenTableViewParam) (output *model.SysGenTableViewModel, err error)
		// 新增
		Add(ctx context.Context, param *model.SysGenTableAddParam) (output *model.SysGenTableAddModel, err error)
		// 修改
		Edit(ctx context.Context, param *model.SysGenTableEditParam) (output *model.SysGenTableEditModel, err error)
		Delete(ctx context.Context, param *model.SysGenTableDeleteParam) (output *model.SysGenTableDeleteModel, err error)
		// 获取选择项
		Selects(ctx context.Context) (output *model.SelectsModel, err error)
		// 获取表选择项
		TableSelect(ctx context.Context, param *model.GenCodesTableSelectParam) (output []*model.GenCodesTableSelectModel, err error)
		// 获取字段选择项
		ColumnSelect(ctx context.Context, param *model.GenCodesColumnSelectParam) (output []*model.GenCodesColumnSelectModel, err error)
		// 获取字段列表
		ColumnList(ctx context.Context, param *model.GenCodesColumnListParam) (output []*model.GenCodesColumnListModel, err error)
		// 预览
		Preview(ctx context.Context, param *model.GenCodesPreviewParam) (output *model.GenCodesPreviewModel, err error)
		// 构建
		Build(ctx context.Context, param *model.GenCodesBuildParam) (output *model.GenCodesBuildModel, err error)
	}
)

var (
	localSysGenTable ISysGenTable
)

func SysGenTable() ISysGenTable {
	if localSysGenTable == nil {
		panic("implement not found for interface ISysGenTable, forgot register?")
	}
	return localSysGenTable
}

func RegisterSysGenTable(i ISysGenTable) {
	localSysGenTable = i
}
