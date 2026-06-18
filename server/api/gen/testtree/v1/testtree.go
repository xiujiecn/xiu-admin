// Package testtree
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2026 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
// @AutoGenerate Version
package v1 //testtree

import (
	"xiuadmin/internal/model/genin"
	"xiuadmin/internal/model/response"
	"xiuadmin/utility/tree"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询测试树表列表
type TestTreeListReq struct {
	g.Meta `path:"/testTree/list" method:"get" tags:"测试树表" summary:"获取测试树表列表" x-check-permission:"cpm:gen:testTree:list"`
	genin.TestTreeListParam
}

type TestTreeListRes struct {
	response.PageResult
	Items []*genin.TestTreeListModel `json:"items"   dc:"数据列表"`
}

// ExportReq 导出测试树表列表
type TestTreeExportReq struct {
	g.Meta `path:"/testTree/export" method:"post" tags:"测试树表" summary:"导出测试树表列表" x-check-permission:"cpm:gen:testTree:export"`
	genin.TestTreeListParam
}

type TestTreeExportRes struct{}

// ViewReq 获取测试树表指定信息
type TestTreeViewReq struct {
	g.Meta `path:"/testTree/view" method:"get" tags:"测试树表" summary:"获取测试树表指定信息" x-check-permission:"cpm:gen:testTree:view"`
	genin.TestTreeViewParam
}

type TestTreeViewRes struct {
	*genin.TestTreeViewModel
}

// EditReq 修改/新增测试树表
type TestTreeEditReq struct {
	g.Meta `path:"/testTree/edit" method:"post" tags:"测试树表" summary:"修改/新增测试树表" x-check-permission:"cpm:gen:testTree:edit"`
	genin.TestTreeEditParam
}

type TestTreeEditRes struct{}

// DeleteReq 删除测试树表
type TestTreeDeleteReq struct {
	g.Meta `path:"/testTree/delete" method:"post" tags:"测试树表" summary:"删除测试树表" x-check-permission:"cpm:gen:testTree:delete"`
	genin.TestTreeDeleteParam
}

type TestTreeDeleteRes struct{}

// TreeOptionReq 获取测试树表关系树选项
type TestTreeTreeOptionReq struct {
	g.Meta `path:"/testTree/treeOption" method:"get" tags:"测试树表" summary:"获取测试树表关系树选项" x-check-permission:"cpm:gen:testTree:treeOption"`
}

type TestTreeTreeOptionRes []tree.Node