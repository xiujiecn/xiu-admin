// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysGenTableDao is the data access object for the table sys_gen_table.
type SysGenTableDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns SysGenTableColumns // columns contains all the column names of Table for convenient usage.
}

// SysGenTableColumns defines and stores column names for the table sys_gen_table.
type SysGenTableColumns struct {
	TableId       string // 生成ID
	GenType       string // 生成类型
	GenTemplate   string // 生成模板
	VarName       string // 实体命名
	Options       string // 配置选项
	DbName        string // 数据库名称
	TableName     string // 主表名称
	TableComment  string // 主表注释
	DaoName       string // 主表dao模型
	MasterColumns string // 主表字段
	AddonName     string // 插件名称
	Status        string // 生成状态（0成功 1未开始）
	CreatedDept   string // 创建部门
	CreatedBy     string // 创建者
	CreatedAt     string // 创建时间
	UpdatedBy     string // 更新者
	UpdatedAt     string // 更新时间
}

// sysGenTableColumns holds the columns for the table sys_gen_table.
var sysGenTableColumns = SysGenTableColumns{
	TableId:       "table_id",
	GenType:       "gen_type",
	GenTemplate:   "gen_template",
	VarName:       "var_name",
	Options:       "options",
	DbName:        "db_name",
	TableName:     "table_name",
	TableComment:  "table_comment",
	DaoName:       "dao_name",
	MasterColumns: "master_columns",
	AddonName:     "addon_name",
	Status:        "status",
	CreatedDept:   "created_dept",
	CreatedBy:     "created_by",
	CreatedAt:     "created_at",
	UpdatedBy:     "updated_by",
	UpdatedAt:     "updated_at",
}

// NewSysGenTableDao creates and returns a new DAO object for table data access.
func NewSysGenTableDao() *SysGenTableDao {
	return &SysGenTableDao{
		group:   "default",
		table:   "sys_gen_table",
		columns: sysGenTableColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysGenTableDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysGenTableDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysGenTableDao) Columns() SysGenTableColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysGenTableDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysGenTableDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysGenTableDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
