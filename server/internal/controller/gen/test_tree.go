// Package gen
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2026 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
// @AutoGenerate Version
package gen

import (
	"context"
	v1 "xiuadmin/api/gen/testtree/v1"
	"xiuadmin/internal/model/genin"
	"xiuadmin/internal/service"
)

var (
	TestTree = cTestTree{}
)

type cTestTree struct{}

// List 查看测试树表列表
func (c *cTestTree) TestTreeList(ctx context.Context, req *v1.TestTreeListReq) (res *v1.TestTreeListRes, err error) {
	list, totalCount, err := service.GenTestTree().List(ctx, &req.TestTreeListParam)
	if err != nil {
		return
	}

	if list == nil {
		list = []*genin.TestTreeListModel{}
	}

	res = new(v1.TestTreeListRes)
	res.Items = list
	res.PageResult.Page = req.Page
	res.PageResult.PageSize = req.PageSize
	res.PageResult.Total = totalCount
	return
}

// Export 导出测试树表列表
func (c *cTestTree) TestTreeExport(ctx context.Context, req *v1.TestTreeExportReq) (res *v1.TestTreeExportRes, err error) {
	err = service.GenTestTree().Export(ctx, &req.TestTreeListParam)
	return
}

// Edit 更新测试树表
func (c *cTestTree) TestTreeEdit(ctx context.Context, req *v1.TestTreeEditReq) (res *v1.TestTreeEditRes, err error) {
	err = service.GenTestTree().Edit(ctx, &req.TestTreeEditParam)
	return
}

// View 获取指定测试树表信息
func (c *cTestTree) TestTreeView(ctx context.Context, req *v1.TestTreeViewReq) (res *v1.TestTreeViewRes, err error) {
	data, err := service.GenTestTree().View(ctx, &req.TestTreeViewParam)
	if err != nil {
		return
	}

	res = new(v1.TestTreeViewRes)
	res.TestTreeViewModel = data
	return
}

// Delete 删除测试树表
func (c *cTestTree) TestTreeDelete(ctx context.Context, req *v1.TestTreeDeleteReq) (res *v1.TestTreeDeleteRes, err error) {
	err = service.GenTestTree().Delete(ctx, &req.TestTreeDeleteParam)
	return
}

// TreeOption 获取测试树表关系树选项
func (c *cTestTree) TestTreeTreeOption(ctx context.Context, req *v1.TestTreeTreeOptionReq) (res *v1.TestTreeTreeOptionRes, err error) {
	data, err := service.GenTestTree().TreeOption(ctx)
	if err != nil {
		return nil, err
	}

	if len(data) > 0 {
		res = (*v1.TestTreeTreeOptionRes)(&data)
	} else {
		temp := make(v1.TestTreeTreeOptionRes, 0)
		res = &temp
	}
	return
}
