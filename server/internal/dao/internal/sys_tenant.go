// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysTenantDao is the data access object for the table sys_tenant.
type SysTenantDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysTenantColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysTenantColumns defines and stores column names for the table sys_tenant.
type SysTenantColumns struct {
	Id              string // id
	TenantId        string // 租户编号
	ContactUserName string // 联系人
	ContactPhone    string // 联系电话
	CompanyName     string // 企业名称
	LicenseNumber   string // 统一社会信用代码
	Address         string // 地址
	Intro           string // 企业简介
	Domain          string // 域名
	Remark          string // 备注
	PackageId       string // 租户套餐编号
	ExpireTime      string // 过期时间
	AccountCount    string // 用户数量（-1不限制）
	Status          string // 租户状态（0正常 1停用）
	CreatedDept     string // 创建部门
	CreatedBy       string // 创建者
	CreatedAt       string // 创建时间
	UpdatedBy       string // 更新者
	UpdatedAt       string // 更新时间
	DeletedBy       string // 删除人
	DeletedAt       string // 删除时间
}

// sysTenantColumns holds the columns for the table sys_tenant.
var sysTenantColumns = SysTenantColumns{
	Id:              "id",
	TenantId:        "tenant_id",
	ContactUserName: "contact_user_name",
	ContactPhone:    "contact_phone",
	CompanyName:     "company_name",
	LicenseNumber:   "license_number",
	Address:         "address",
	Intro:           "intro",
	Domain:          "domain",
	Remark:          "remark",
	PackageId:       "package_id",
	ExpireTime:      "expire_time",
	AccountCount:    "account_count",
	Status:          "status",
	CreatedDept:     "created_dept",
	CreatedBy:       "created_by",
	CreatedAt:       "created_at",
	UpdatedBy:       "updated_by",
	UpdatedAt:       "updated_at",
	DeletedBy:       "deleted_by",
	DeletedAt:       "deleted_at",
}

// NewSysTenantDao creates and returns a new DAO object for table data access.
func NewSysTenantDao(handlers ...gdb.ModelHandler) *SysTenantDao {
	return &SysTenantDao{
		group:    "default",
		table:    "sys_tenant",
		columns:  sysTenantColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysTenantDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysTenantDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysTenantDao) Columns() SysTenantColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysTenantDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysTenantDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysTenantDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
