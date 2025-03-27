// Package v1
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package v1

import (
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type ListLogininforReq struct {
	g.Meta `path:"/logininfor/list" method:"get" tags:"系统-日志管理" summary:"登录信息列表" x-check-permission:"cpm:monitor:logininfor:list"`
	model.SysLogininforListParam
}

type ListLogininforRes struct {
	Items []*model.SysLogininforListModel `json:"items"`
	response.PageResult
}

type DeleteLogininforReq struct {
	g.Meta `path:"/logininfor/delete" method:"post" tags:"系统-日志管理" summary:"删除登录信息" x-check-permission:"cpm:monitor:logininfor:remove"`
	model.SysLogininforDeleteParam
}

type DeleteLogininforRes struct {
	model.SysLogininforDeleteModel
}
