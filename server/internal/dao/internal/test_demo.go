// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TestDemoDao is the data access object for the table test_demo.
type TestDemoDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TestDemoColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TestDemoColumns defines and stores column names for the table test_demo.
type TestDemoColumns struct {
	Id          string // 主键
	TenantId    string // 租户编号
	DeptId      string // 部门id
	UserId      string // 用户id
	OrderNum    string // 排序号
	TestKey     string // key键
	Value       string // 值
	Version     string // 版本
	CreatedDept string // 创建部门
	CreatedAt   string // 创建时间
	CreatedBy   string // 创建者
	UpdatedAt   string // 更新时间
	UpdatedBy   string // 更新者
	DeletedBy   string // 删除人
	DeletedAt   string // 删除时间
}

// testDemoColumns holds the columns for the table test_demo.
var testDemoColumns = TestDemoColumns{
	Id:          "id",
	TenantId:    "tenant_id",
	DeptId:      "dept_id",
	UserId:      "user_id",
	OrderNum:    "order_num",
	TestKey:     "test_key",
	Value:       "value",
	Version:     "version",
	CreatedDept: "created_dept",
	CreatedAt:   "created_at",
	CreatedBy:   "created_by",
	UpdatedAt:   "updated_at",
	UpdatedBy:   "updated_by",
	DeletedBy:   "deleted_by",
	DeletedAt:   "deleted_at",
}

// NewTestDemoDao creates and returns a new DAO object for table data access.
func NewTestDemoDao(handlers ...gdb.ModelHandler) *TestDemoDao {
	return &TestDemoDao{
		group:    "default",
		table:    "test_demo",
		columns:  testDemoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TestDemoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TestDemoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TestDemoDao) Columns() TestDemoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TestDemoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TestDemoDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TestDemoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
