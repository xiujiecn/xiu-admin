// Package testcategory
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
// @AutoGenerate Version
package v1 //testcategory

import (
	"xiuadmin/internal/model/genin"
	"xiuadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询测试分类列表
type TestCategoryListReq struct {
	g.Meta `path:"/testCategory/list" method:"get" tags:"测试分类" summary:"获取测试分类列表" x-check-permission:"cpm:gen:testCategory:list"`
	genin.TestCategoryListParam
}

type TestCategoryListRes struct {
	response.PageResult
	Items []*genin.TestCategoryListModel `json:"items"   dc:"数据列表"`
}

// ExportReq 导出测试分类列表
type TestCategoryExportReq struct {
	g.Meta `path:"/testCategory/export" method:"post" tags:"测试分类" summary:"导出测试分类列表" x-check-permission:"cpm:gen:testCategory:export"`
	genin.TestCategoryListParam
}

type TestCategoryExportRes struct{}

// ViewReq 获取测试分类指定信息
type TestCategoryViewReq struct {
	g.Meta `path:"/testCategory/view" method:"get" tags:"测试分类" summary:"获取测试分类指定信息" x-check-permission:"cpm:gen:testCategory:view"`
	genin.TestCategoryViewParam
}

type TestCategoryViewRes struct {
	*genin.TestCategoryViewModel
}

// EditReq 修改/新增测试分类
type TestCategoryEditReq struct {
	g.Meta `path:"/testCategory/edit" method:"post" tags:"测试分类" summary:"修改/新增测试分类" x-check-permission:"cpm:gen:testCategory:edit"`
	genin.TestCategoryEditParam
}

type TestCategoryEditRes struct{}

// DeleteReq 删除测试分类
type TestCategoryDeleteReq struct {
	g.Meta `path:"/testCategory/delete" method:"post" tags:"测试分类" summary:"删除测试分类" x-check-permission:"cpm:gen:testCategory:delete"`
	genin.TestCategoryDeleteParam
}

type TestCategoryDeleteRes struct{}

// MaxSortReq 获取测试分类最大排序
type TestCategoryMaxSortReq struct {
	g.Meta `path:"/testCategory/maxSort" method:"get" tags:"测试分类" summary:"获取测试分类最大排序" x-check-permission:"cpm:gen:testCategory:maxSort"`
	genin.TestCategoryMaxSortParam
}

type TestCategoryMaxSortRes struct {
	*genin.TestCategoryMaxSortModel
}

// StatusReq 更新测试分类状态
type TestCategoryStatusReq struct {
	g.Meta `path:"/testCategory/status" method:"post" tags:"测试分类" summary:"更新测试分类状态" x-check-permission:"cpm:gen:testCategory:status"`
	genin.TestCategoryStatusParam
}

type TestCategoryStatusRes struct{}