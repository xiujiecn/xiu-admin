// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TestTreeDao is the data access object for the table test_tree.
type TestTreeDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TestTreeColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TestTreeColumns defines and stores column names for the table test_tree.
type TestTreeColumns struct {
	Id          string // 主键
	TenantId    string // 租户编号
	ParentId    string // 父id
	DeptId      string // 部门id
	UserId      string // 用户id
	TreeName    string // 值
	Level       string // 关系树等级
	Tree        string // 关系树
	Version     string // 版本
	CreatedDept string // 创建部门
	CreatedAt   string // 创建时间
	CreatedBy   string // 创建者
	UpdatedAt   string // 更新时间
	UpdatedBy   string // 更新者
	DeletedBy   string // 删除人
	DeletedAt   string // 删除时间
}

// testTreeColumns holds the columns for the table test_tree.
var testTreeColumns = TestTreeColumns{
	Id:          "id",
	TenantId:    "tenant_id",
	ParentId:    "parent_id",
	DeptId:      "dept_id",
	UserId:      "user_id",
	TreeName:    "tree_name",
	Level:       "level",
	Tree:        "tree",
	Version:     "version",
	CreatedDept: "created_dept",
	CreatedAt:   "created_at",
	CreatedBy:   "created_by",
	UpdatedAt:   "updated_at",
	UpdatedBy:   "updated_by",
	DeletedBy:   "deleted_by",
	DeletedAt:   "deleted_at",
}

// NewTestTreeDao creates and returns a new DAO object for table data access.
func NewTestTreeDao(handlers ...gdb.ModelHandler) *TestTreeDao {
	return &TestTreeDao{
		group:    "default",
		table:    "test_tree",
		columns:  testTreeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TestTreeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TestTreeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TestTreeDao) Columns() TestTreeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TestTreeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TestTreeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TestTreeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
