// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserOnlineDao is the data access object for the table sys_user_online.
type SysUserOnlineDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns SysUserOnlineColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserOnlineColumns defines and stores column names for the table sys_user_online.
type SysUserOnlineColumns struct {
	OnlineId      string // 访问ID
	TenantId      string // 租户编号
	Uuid          string // UUID
	UserName      string // 用户账号
	ClientKey     string // 客户端
	DeviceType    string // 设备类型
	Ipaddr        string // 登录IP地址
	LoginLocation string // 登录地点
	Browser       string // 浏览器类型
	Os            string // 操作系统
	Token         string // Token
	LoginTime     string // 访问时间
	ExpireTime    string // 过期时间
	DeletedAt     string // 删除时间
}

// sysUserOnlineColumns holds the columns for the table sys_user_online.
var sysUserOnlineColumns = SysUserOnlineColumns{
	OnlineId:      "online_id",
	TenantId:      "tenant_id",
	Uuid:          "uuid",
	UserName:      "user_name",
	ClientKey:     "client_key",
	DeviceType:    "device_type",
	Ipaddr:        "ipaddr",
	LoginLocation: "login_location",
	Browser:       "browser",
	Os:            "os",
	Token:         "token",
	LoginTime:     "login_time",
	ExpireTime:    "expire_time",
	DeletedAt:     "deleted_at",
}

// NewSysUserOnlineDao creates and returns a new DAO object for table data access.
func NewSysUserOnlineDao() *SysUserOnlineDao {
	return &SysUserOnlineDao{
		group:   "default",
		table:   "sys_user_online",
		columns: sysUserOnlineColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysUserOnlineDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysUserOnlineDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysUserOnlineDao) Columns() SysUserOnlineColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysUserOnlineDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysUserOnlineDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysUserOnlineDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
