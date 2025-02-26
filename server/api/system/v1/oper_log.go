package v1

import (
	"server/internal/model"
	"server/internal/model/request"
	"server/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type GetOperLogListReq struct {
	g.Meta `path:"/oper-log/list" method:"get" tags:"操作日志" summary:"操作日志列表"`
	model.SysOperLogListParam
	request.PageInfo
}

type GetOperLogListRes struct {
	Items []*model.SysOperLogListModel `json:"items"`
	response.PageResult
}
