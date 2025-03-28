// Package v1
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025	 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package v1

import (
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type SysGenTableListReq struct {
	g.Meta `path:"/genTable/list" method:"get" tags:"工具-代码生成" summary:"获取代码生成列表" x-check-permission:"cpm:tool:gen:list"`
	*model.SysGenTableListParam
}

type SysGenTableListRes struct {
	response.PageResult
	Items []*model.SysGenTableListModel `json:"items"`
}

type SysGenTableViewReq struct {
	g.Meta `path:"/genTable/view" method:"get" tags:"工具-代码生成" summary:"获取代码生成详情" x-check-permission:"cpm:tool:gen:query"`
	*model.SysGenTableViewParam
}

type SysGenTableViewRes struct {
	*model.SysGenTableViewModel
}

type SysGenTableAddReq struct {
	g.Meta `path:"/genTable/add" method:"post" tags:"工具-代码生成" summary:"新增代码生成" x-check-permission:"cpm:tool:gen:import"`
	*model.SysGenTableAddParam
}

type SysGenTableAddRes struct {
	*model.SysGenTableAddModel
}

type SysGenTableEditReq struct {
	g.Meta `path:"/genTable/edit" method:"post" tags:"工具-代码生成" summary:"修改代码生成" x-check-permission:"cpm:tool:gen:edit"`
	*model.SysGenTableEditParam
}

type SysGenTableEditRes struct {
	*model.SysGenTableEditModel
}

type SysGenTableDeleteReq struct {
	g.Meta `path:"/genTable/delete" method:"post" tags:"工具-代码生成" summary:"删除代码生成" x-check-permission:"cpm:tool:gen:remove"`
	*model.SysGenTableDeleteParam
}

type SysGenTableDeleteRes struct {
	*model.SysGenTableDeleteModel
}

type SysGenTableSelectsReq struct {
	g.Meta `path:"/genTable/selects" method:"get" tags:"工具-代码生成" summary:"获取代码生成选择项" x-check-permission:"cpm:tool:gen:list"`
}

type SysGenTableSelectsRes struct {
	*model.SelectsModel
}

type SysGenTableTableSelectReq struct {
	g.Meta `path:"/genTable/tableSelect" method:"get" tags:"工具-代码生成" summary:"获取代码生成表选择项" x-check-permission:"cpm:tool:gen:list"`
	*model.GenCodesTableSelectParam
}

type SysGenTableTableSelectRes struct {
	Items []*model.GenCodesTableSelectModel `json:"items"`
}

type SysGenTableColumnSelectReq struct {
	g.Meta `path:"/genTable/columnSelect" method:"get" tags:"工具-代码生成" summary:"获取代码生成字段选择项" x-check-permission:"cpm:tool:gen:list"`
	*model.GenCodesColumnSelectParam
}

type SysGenTableColumnSelectRes struct {
	Items []*model.GenCodesColumnSelectModel `json:"items"`
}

type SysGenTableColumnListReq struct {
	g.Meta `path:"/genTable/columnList" method:"get" tags:"工具-代码生成" summary:"获取代码生成字段列表" x-check-permission:"cpm:tool:gen:list"`
	*model.GenCodesColumnListParam
}

type SysGenTableColumnListRes struct {
	Items []*model.GenCodesColumnListModel `json:"items"`
}

type SysGenTablePreviewReq struct {
	g.Meta `path:"/genTable/preview" method:"post" tags:"工具-代码生成" summary:"生成预览" x-check-permission:"cpm:tool:gen:preview"`
	*model.GenCodesPreviewParam
}

type SysGenTablePreviewRes struct {
	*model.GenCodesPreviewModel
}

type SysGenTableBuildReq struct {
	g.Meta `path:"/genTable/build" method:"post" tags:"工具-代码生成" summary:"构建" x-check-permission:"cpm:tool:gen:code"`
	*model.GenCodesBuildParam
}

type SysGenTableBuildRes struct {
	*model.GenCodesBuildModel
}
