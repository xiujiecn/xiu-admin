// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysRoleDao is the data access object for the table sys_role.
type SysRoleDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns SysRoleColumns // columns contains all the column names of Table for convenient usage.
}

// SysRoleColumns defines and stores column names for the table sys_role.
type SysRoleColumns struct {
	RoleId            string // 角色ID
	TenantId          string // 租户编号
	RoleName          string // 角色名称
	RoleKey           string // 角色权限字符串
	RoleSort          string // 显示顺序
	DataScope         string // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	MenuCheckStrictly string // 菜单树选择项是否关联显示
	DeptCheckStrictly string // 部门树选择项是否关联显示
	Status            string // 角色状态（0正常 1停用）
	CreatedDept       string // 创建部门
	CreatedBy         string // 创建者
	CreatedAt         string // 创建时间
	UpdatedBy         string // 更新者
	UpdatedAt         string // 更新时间
	DeletedBy         string // 删除人
	DeletedAt         string // 删除时间
	Remark            string // 备注
}

// sysRoleColumns holds the columns for the table sys_role.
var sysRoleColumns = SysRoleColumns{
	RoleId:            "role_id",
	TenantId:          "tenant_id",
	RoleName:          "role_name",
	RoleKey:           "role_key",
	RoleSort:          "role_sort",
	DataScope:         "data_scope",
	MenuCheckStrictly: "menu_check_strictly",
	DeptCheckStrictly: "dept_check_strictly",
	Status:            "status",
	CreatedDept:       "created_dept",
	CreatedBy:         "created_by",
	CreatedAt:         "created_at",
	UpdatedBy:         "updated_by",
	UpdatedAt:         "updated_at",
	DeletedBy:         "deleted_by",
	DeletedAt:         "deleted_at",
	Remark:            "remark",
}

// NewSysRoleDao creates and returns a new DAO object for table data access.
func NewSysRoleDao() *SysRoleDao {
	return &SysRoleDao{
		group:   "default",
		table:   "sys_role",
		columns: sysRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysRoleDao) Columns() SysRoleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
