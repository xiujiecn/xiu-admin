// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysClientDao is the data access object for the table sys_client.
type SysClientDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns SysClientColumns // columns contains all the column names of Table for convenient usage.
}

// SysClientColumns defines and stores column names for the table sys_client.
type SysClientColumns struct {
	Id            string // id
	ClientId      string // 客户端id
	ClientKey     string // 客户端key
	ClientSecret  string // 客户端秘钥
	GrantType     string // 授权类型
	DeviceType    string // 设备类型
	ActiveTimeout string // token活跃超时时间
	Timeout       string // token固定超时
	Status        string // 状态（0正常 1停用）
	CreatedDept   string // 创建部门
	CreatedBy     string // 创建者
	CreatedAt     string // 创建时间
	UpdatedBy     string // 更新者
	UpdatedAt     string // 更新时间
	DeletedBy     string // 删除人
	DeletedAt     string // 删除时间
}

// sysClientColumns holds the columns for the table sys_client.
var sysClientColumns = SysClientColumns{
	Id:            "id",
	ClientId:      "client_id",
	ClientKey:     "client_key",
	ClientSecret:  "client_secret",
	GrantType:     "grant_type",
	DeviceType:    "device_type",
	ActiveTimeout: "active_timeout",
	Timeout:       "timeout",
	Status:        "status",
	CreatedDept:   "created_dept",
	CreatedBy:     "created_by",
	CreatedAt:     "created_at",
	UpdatedBy:     "updated_by",
	UpdatedAt:     "updated_at",
	DeletedBy:     "deleted_by",
	DeletedAt:     "deleted_at",
}

// NewSysClientDao creates and returns a new DAO object for table data access.
func NewSysClientDao() *SysClientDao {
	return &SysClientDao{
		group:   "default",
		table:   "sys_client",
		columns: sysClientColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysClientDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysClientDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysClientDao) Columns() SysClientColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysClientDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysClientDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysClientDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
