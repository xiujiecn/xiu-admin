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
	localGenTestDemo IGenTestDemo
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
