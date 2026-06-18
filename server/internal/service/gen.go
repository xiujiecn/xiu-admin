// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/model/genin"
	"xiuadmin/utility/tree"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
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
	IGenTestTree interface {
		// Model 测试树表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取测试树表列表
		List(ctx context.Context, in *genin.TestTreeListParam) (list []*genin.TestTreeListModel, totalCount int, err error)
		// Export 导出测试树表
		Export(ctx context.Context, in *genin.TestTreeListParam) (err error)
		// Edit 修改/新增测试树表
		Edit(ctx context.Context, in *genin.TestTreeEditParam) (err error)
		// Delete 删除测试树表
		Delete(ctx context.Context, in *genin.TestTreeDeleteParam) (err error)
		// View 获取测试树表指定信息
		View(ctx context.Context, in *genin.TestTreeViewParam) (res *genin.TestTreeViewModel, err error)
		// TreeOption 获取测试树表关系树选项
		TreeOption(ctx context.Context) (nodes []tree.Node, err error)
	}
)

var (
	localGenTestDemo IGenTestDemo
	localGenTestTree IGenTestTree
)

func GenTestDemo() IGenTestDemo {
	if localGenTestDemo == nil {
		panic("implement not found for interface IGenTestDemo, forgot register?")
	}
	return localGenTestDemo
}

func RegisterGenTestDemo(i IGenTestDemo) {
	localGenTestDemo = i
}

func GenTestTree() IGenTestTree {
	if localGenTestTree == nil {
		panic("implement not found for interface IGenTestTree, forgot register?")
	}
	return localGenTestTree
}

func RegisterGenTestTree(i IGenTestTree) {
	localGenTestTree = i
}
