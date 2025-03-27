// Package gen
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
// @AutoGenerate Version
package gen

import (
	"context"
	v1 "xiuadmin/api/gen/testdemo/v1"
	"xiuadmin/internal/model/genin"
	"xiuadmin/internal/service"
)

var (
	TestDemo = cTestDemo{}
)

type cTestDemo struct{}

// List 查看测试单表列表
func (c *cTestDemo) TestDemoList(ctx context.Context, req *v1.TestDemoListReq) (res *v1.TestDemoListRes, err error) {
	list, totalCount, err := service.GenTestDemo().List(ctx, &req.TestDemoListParam)
	if err != nil {
		return
	}

	if list == nil {
		list = []*genin.TestDemoListModel{}
	}

	res = new(v1.TestDemoListRes)
	res.Items = list
	res.PageResult.Page = req.Page
	res.PageResult.PageSize = req.PageSize
	res.PageResult.Total = totalCount
	return
}

// Export 导出测试单表列表
func (c *cTestDemo) TestDemoExport(ctx context.Context, req *v1.TestDemoExportReq) (res *v1.TestDemoExportRes, err error) {
	err = service.GenTestDemo().Export(ctx, &req.TestDemoListParam)
	return
}

// Edit 更新测试单表
func (c *cTestDemo) TestDemoEdit(ctx context.Context, req *v1.TestDemoEditReq) (res *v1.TestDemoEditRes, err error) {
	err = service.GenTestDemo().Edit(ctx, &req.TestDemoEditParam)
	return
}

// View 获取指定测试单表信息
func (c *cTestDemo) TestDemoView(ctx context.Context, req *v1.TestDemoViewReq) (res *v1.TestDemoViewRes, err error) {
	data, err := service.GenTestDemo().View(ctx, &req.TestDemoViewParam)
	if err != nil {
		return
	}

	res = new(v1.TestDemoViewRes)
	res.TestDemoViewModel = data
	return
}

// Delete 删除测试单表
func (c *cTestDemo) TestDemoDelete(ctx context.Context, req *v1.TestDemoDeleteReq) (res *v1.TestDemoDeleteRes, err error) {
	err = service.GenTestDemo().Delete(ctx, &req.TestDemoDeleteParam)
	return
}
