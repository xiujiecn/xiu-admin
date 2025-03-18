package v1

import (
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type SysGenTableListReq struct {
	g.Meta `path:"/genTable/list" method:"get" tags:"代码生成" summary:"获取代码生成列表"`
	*model.SysGenTableListParam
}

type SysGenTableListRes struct {
	response.PageResult
	Items []*model.SysGenTableListModel `json:"items"`
}

type SysGenTableViewReq struct {
	g.Meta `path:"/genTable/view" method:"get" tags:"代码生成" summary:"获取代码生成详情"`
	*model.SysGenTableViewParam
}

type SysGenTableViewRes struct {
	*model.SysGenTableViewModel
}

type SysGenTableAddReq struct {
	g.Meta `path:"/genTable/add" method:"post" tags:"代码生成" summary:"新增代码生成"`
	*model.SysGenTableAddParam
}

type SysGenTableAddRes struct {
	*model.SysGenTableAddModel
}

type SysGenTableDeleteReq struct {
	g.Meta `path:"/genTable/delete" method:"post" tags:"代码生成" summary:"删除代码生成"`
	*model.SysGenTableDeleteParam
}

type SysGenTableDeleteRes struct {
	*model.SysGenTableDeleteModel
}
