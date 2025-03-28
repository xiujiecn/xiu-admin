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
	IGenTestCategory interface {
		// Model 测试分类ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取测试分类列表
		List(ctx context.Context, in *genin.TestCategoryListParam) (list []*genin.TestCategoryListModel, totalCount int, err error)
		// Export 导出测试分类
		Export(ctx context.Context, in *genin.TestCategoryListParam) (err error)
		// Edit 修改/新增测试分类
		Edit(ctx context.Context, in *genin.TestCategoryEditParam) (err error)
		// Delete 删除测试分类
		Delete(ctx context.Context, in *genin.TestCategoryDeleteParam) (err error)
		// MaxSort 获取测试分类最大排序
		MaxSort(ctx context.Context, in *genin.TestCategoryMaxSortParam) (res *genin.TestCategoryMaxSortModel, err error)
		// View 获取测试分类指定信息
		View(ctx context.Context, in *genin.TestCategoryViewParam) (res *genin.TestCategoryViewModel, err error)
		// Status 更新测试分类状态
		Status(ctx context.Context, in *genin.TestCategoryStatusParam) (err error)
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
	localGenTestCategory IGenTestCategory
	localGenTestDemo     IGenTestDemo
)

func GenTestCategory() IGenTestCategory {
	if localGenTestCategory == nil {
		panic("implement not found for interface IGenTestCategory, forgot register?")
	}
	return localGenTestCategory
}

func RegisterGenTestCategory(i IGenTestCategory) {
	localGenTestCategory = i
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
