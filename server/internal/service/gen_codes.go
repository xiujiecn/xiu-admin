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
