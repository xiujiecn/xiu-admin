// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysTenantPackageDao is the data access object for the table sys_tenant_package.
type SysTenantPackageDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  SysTenantPackageColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// SysTenantPackageColumns defines and stores column names for the table sys_tenant_package.
type SysTenantPackageColumns struct {
	PackageId         string // 租户套餐id
	PackageName       string // 套餐名称
	MenuIds           string // 关联菜单id
	Remark            string // 备注
	MenuCheckStrictly string // 菜单树选择项是否关联显示
	Status            string // 状态（0正常 1停用）
	CreatedDept       string // 创建部门
	CreatedBy         string // 创建者
	CreatedAt         string // 创建时间
	UpdatedBy         string // 更新者
	UpdatedAt         string // 更新时间
	DeletedBy         string // 删除人
	DeletedAt         string // 删除时间
}

// sysTenantPackageColumns holds the columns for the table sys_tenant_package.
var sysTenantPackageColumns = SysTenantPackageColumns{
	PackageId:         "package_id",
	PackageName:       "package_name",
	MenuIds:           "menu_ids",
	Remark:            "remark",
	MenuCheckStrictly: "menu_check_strictly",
	Status:            "status",
	CreatedDept:       "created_dept",
	CreatedBy:         "created_by",
	CreatedAt:         "created_at",
	UpdatedBy:         "updated_by",
	UpdatedAt:         "updated_at",
	DeletedBy:         "deleted_by",
	DeletedAt:         "deleted_at",
}

// NewSysTenantPackageDao creates and returns a new DAO object for table data access.
func NewSysTenantPackageDao(handlers ...gdb.ModelHandler) *SysTenantPackageDao {
	return &SysTenantPackageDao{
		group:    "default",
		table:    "sys_tenant_package",
		columns:  sysTenantPackageColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysTenantPackageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysTenantPackageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysTenantPackageDao) Columns() SysTenantPackageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysTenantPackageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysTenantPackageDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysTenantPackageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
