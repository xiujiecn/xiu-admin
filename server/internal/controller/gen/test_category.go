// Package gen
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
// @AutoGenerate Version
package gen

import (
	"context"
	"xiuadmin/api/gen/testcategory/v1"
	"xiuadmin/internal/model/genin"
	"xiuadmin/internal/service"
)

var (
	TestCategory = cTestCategory{}
)

type cTestCategory struct{}

// List 查看测试分类列表
func (c *cTestCategory) TestCategoryList(ctx context.Context, req *v1.TestCategoryListReq) (res *v1.TestCategoryListRes, err error) {
	list, totalCount, err := service.GenTestCategory().List(ctx, &req.TestCategoryListParam)
	if err != nil {
		return
	}

	if list == nil {
		list = []*genin.TestCategoryListModel{}
	}

	res = new(v1.TestCategoryListRes)
	res.Items = list
	res.PageResult.Page = req.Page
	res.PageResult.PageSize = req.PageSize
	res.PageResult.Total = totalCount
	return
}

// Export 导出测试分类列表
func (c *cTestCategory) TestCategoryExport(ctx context.Context, req *v1.TestCategoryExportReq) (res *v1.TestCategoryExportRes, err error) {
	err = service.GenTestCategory().Export(ctx, &req.TestCategoryListParam)
	return
}

// Edit 更新测试分类
func (c *cTestCategory) TestCategoryEdit(ctx context.Context, req *v1.TestCategoryEditReq) (res *v1.TestCategoryEditRes, err error) {
	err = service.GenTestCategory().Edit(ctx, &req.TestCategoryEditParam)
	return
}

// MaxSort 获取测试分类最大排序
func (c *cTestCategory) TestCategoryMaxSort(ctx context.Context, req *v1.TestCategoryMaxSortReq) (res *v1.TestCategoryMaxSortRes, err error) {
	data, err := service.GenTestCategory().MaxSort(ctx, &req.TestCategoryMaxSortParam)
	if err != nil {
		return
	}

	res = new(testcategory.MaxSortRes)
	res.TestCategoryMaxSortModel = data
	return
}

// View 获取指定测试分类信息
func (c *cTestCategory) TestCategoryView(ctx context.Context, req *v1.TestCategoryViewReq) (res *v1.TestCategoryViewRes, err error) {
	data, err := service.GenTestCategory().View(ctx, &req.TestCategoryViewParam)
	if err != nil {
		return
	}

	res = new(v1.TestCategoryViewRes)
	res.TestCategoryViewModel = data
	return
}

// Delete 删除测试分类
func (c *cTestCategory) TestCategoryDelete(ctx context.Context, req *v1.TestCategoryDeleteReq) (res *v1.TestCategoryDeleteRes, err error) {
	err = service.GenTestCategory().Delete(ctx, &req.TestCategoryDeleteParam)
	return
}

// Status 更新测试分类状态
func (c *cTestCategory) TestCategoryStatus(ctx context.Context, req *v1.TestCategoryStatusReq) (res *v1.TestCategoryStatusRes, err error) {
	err = service.GenTestCategory().Status(ctx, &req.TestCategoryStatusParam)
	return
}