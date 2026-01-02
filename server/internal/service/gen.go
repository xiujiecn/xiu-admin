// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/model/genin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IGenEmployees interface {
		// Model 员工信息表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取员工信息表列表
		List(ctx context.Context, in *genin.EmployeesListParam) (list []*genin.EmployeesListModel, totalCount int, err error)
		// Export 导出员工信息表
		Export(ctx context.Context, in *genin.EmployeesListParam) (err error)
		// Edit 修改/新增员工信息表
		Edit(ctx context.Context, in *genin.EmployeesEditParam) (err error)
		// Delete 删除员工信息表
		Delete(ctx context.Context, in *genin.EmployeesDeleteParam) (err error)
		// View 获取员工信息表指定信息
		View(ctx context.Context, in *genin.EmployeesViewParam) (res *genin.EmployeesViewModel, err error)
	}
	IGenTestDemo interface {
		// Model 测试单表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取测试单表列表
		List(ctx context.Context, in *genin.TestDemoListParam) (list []*genin.TestDemoListModel, totalCount int, err error)
		// Export 导出测试单表
		Export(ctx context.Context, in *genin.TestDemoListParam) (err error)
		// Edit 修改/新增测试单表
		Edit(ctx context.Context, in *genin.TestDemoEditParam) (err error)
		// Delete 删除测试单表
		Delete(ctx context.Context, in *genin.TestDemoDeleteParam) (err error)
		// View 获取测试单表指定信息
		View(ctx context.Context, in *genin.TestDemoViewParam) (res *genin.TestDemoViewModel, err error)
	}
)

var (
	localGenEmployees IGenEmployees
	localGenTestDemo  IGenTestDemo
)

func GenEmployees() IGenEmployees {
	if localGenEmployees == nil {
		panic("implement not found for interface IGenEmployees, forgot register?")
	}
	return localGenEmployees
}

func RegisterGenEmployees(i IGenEmployees) {
	localGenEmployees = i
}

func GenTestDemo() IGenTestDemo {
	if localGenTestDemo == nil {
		panic("implement not found for interface IGenTestDemo, forgot register?")
	}
	return localGenTestDemo
}

func RegisterGenTestDemo(i IGenTestDemo) {
	localGenTestDemo = i
}
