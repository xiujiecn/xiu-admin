// Package v1
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package v1

import (
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/request"
	"xiuadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type GetOperLogListReq struct {
	g.Meta `path:"/oper-log/list" method:"get" tags:"系统-日志管理" summary:"操作日志列表" x-check-permission:"cpm:monitor:operlog:list"`
	model.SysOperLogListParam
	request.PageInfo
}

type GetOperLogListRes struct {
	Items []*model.SysOperLogListModel `json:"items"`
	response.PageResult
}
