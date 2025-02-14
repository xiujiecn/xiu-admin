// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysOssDao is the data access object for the table sys_oss.
type SysOssDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of the current DAO.
	columns SysOssColumns // columns contains all the column names of Table for convenient usage.
}

// SysOssColumns defines and stores column names for the table sys_oss.
type SysOssColumns struct {
	OssId        string // 对象存储主键
	TenantId     string // 租户编号
	FileName     string // 文件名
	OriginalName string // 原名
	FileSuffix   string // 文件后缀名
	Url          string // URL地址
	CreatedDept  string // 创建部门
	CreatedAt    string // 创建时间
	CreatedBy    string // 创建者
	UpdatedAt    string // 更新时间
	UpdatedBy    string // 更新者
	Service      string // 服务商
}

// sysOssColumns holds the columns for the table sys_oss.
var sysOssColumns = SysOssColumns{
	OssId:        "oss_id",
	TenantId:     "tenant_id",
	FileName:     "file_name",
	OriginalName: "original_name",
	FileSuffix:   "file_suffix",
	Url:          "url",
	CreatedDept:  "created_dept",
	CreatedAt:    "created_at",
	CreatedBy:    "created_by",
	UpdatedAt:    "updated_at",
	UpdatedBy:    "updated_by",
	Service:      "service",
}

// NewSysOssDao creates and returns a new DAO object for table data access.
func NewSysOssDao() *SysOssDao {
	return &SysOssDao{
		group:   "default",
		table:   "sys_oss",
		columns: sysOssColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysOssDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysOssDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysOssDao) Columns() SysOssColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysOssDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysOssDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysOssDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
