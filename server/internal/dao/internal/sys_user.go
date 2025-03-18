// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserDao is the data access object for the table sys_user.
type SysUserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysUserColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysUserColumns defines and stores column names for the table sys_user.
type SysUserColumns struct {
	UserId      string // 用户ID
	TenantId    string // 租户编号
	DeptId      string // 部门ID
	UserName    string // 用户账号
	NickName    string // 用户昵称
	UserType    string // 用户类型（sys_user系统用户）
	Email       string // 用户邮箱
	Phonenumber string // 手机号码
	Sex         string // 用户性别（0男 1女 2未知）
	Avatar      string // 头像地址
	Salt        string // 加密盐
	Password    string // 密码
	Status      string // 帐号状态（0正常 1停用）
	LoginIp     string // 最后登录IP
	LoginDate   string // 最后登录时间
	CreatedDept string // 创建部门
	CreatedBy   string // 创建者
	CreatedAt   string // 创建时间
	UpdatedBy   string // 更新者
	UpdatedAt   string // 更新时间
	DeletedBy   string // 删除人
	DeletedAt   string // 删除时间
	Remark      string // 备注
}

// sysUserColumns holds the columns for the table sys_user.
var sysUserColumns = SysUserColumns{
	UserId:      "user_id",
	TenantId:    "tenant_id",
	DeptId:      "dept_id",
	UserName:    "user_name",
	NickName:    "nick_name",
	UserType:    "user_type",
	Email:       "email",
	Phonenumber: "phonenumber",
	Sex:         "sex",
	Avatar:      "avatar",
	Salt:        "salt",
	Password:    "password",
	Status:      "status",
	LoginIp:     "login_ip",
	LoginDate:   "login_date",
	CreatedDept: "created_dept",
	CreatedBy:   "created_by",
	CreatedAt:   "created_at",
	UpdatedBy:   "updated_by",
	UpdatedAt:   "updated_at",
	DeletedBy:   "deleted_by",
	DeletedAt:   "deleted_at",
	Remark:      "remark",
}

// NewSysUserDao creates and returns a new DAO object for table data access.
func NewSysUserDao(handlers ...gdb.ModelHandler) *SysUserDao {
	return &SysUserDao{
		group:    "default",
		table:    "sys_user",
		columns:  sysUserColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysUserDao) Columns() SysUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysUserDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
