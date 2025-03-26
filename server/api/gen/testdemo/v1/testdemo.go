// Package testdemo
// @Link  https://github.com/xiujie/xiujie-admin
// @Copyright  Copyright (c) 2025 XiuJieZhiLian CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujie/xiujie-admin/blob/master/LICENSE
// @AutoGenerate Version
package v1 //testdemo

import (
	"xiujieadmin/internal/model/genin"
	"xiujieadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询测试单表列表
type TestDemoListReq struct {
	g.Meta `path:"/testDemo/list" method:"get" tags:"测试单表" summary:"获取测试单表列表" x-check-permission:"cpm:gen:testDemo:list"`
	genin.TestDemoListParam
}

type TestDemoListRes struct {
	response.PageResult
	Items []*genin.TestDemoListModel `json:"items"   dc:"数据列表"`
}

// ExportReq 导出测试单表列表
type TestDemoExportReq struct {
	g.Meta `path:"/testDemo/export" method:"post" tags:"测试单表" summary:"导出测试单表列表" x-check-permission:"cpm:gen:testDemo:export"`
	genin.TestDemoListParam
}

type TestDemoExportRes struct{}

// ViewReq 获取测试单表指定信息
type TestDemoViewReq struct {
	g.Meta `path:"/testDemo/view" method:"get" tags:"测试单表" summary:"获取测试单表指定信息" x-check-permission:"cpm:gen:testDemo:view"`
	genin.TestDemoViewParam
}

type TestDemoViewRes struct {
	*genin.TestDemoViewModel
}

// EditReq 修改/新增测试单表
type TestDemoEditReq struct {
	g.Meta `path:"/testDemo/edit" method:"post" tags:"测试单表" summary:"修改/新增测试单表" x-check-permission:"cpm:gen:testDemo:edit"`
	genin.TestDemoEditParam
}

type TestDemoEditRes struct{}

// DeleteReq 删除测试单表
type TestDemoDeleteReq struct {
	g.Meta `path:"/testDemo/delete" method:"post" tags:"测试单表" summary:"删除测试单表" x-check-permission:"cpm:gen:testDemo:delete"`
	genin.TestDemoDeleteParam
}

type TestDemoDeleteRes struct{}