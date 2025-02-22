package v1

import (
	"server/internal/model"
	"server/internal/model/request"
	"server/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type ListLogininforReq struct {
	g.Meta `path:"/logininfor/list" method:"get" tags:"系统" summary:"登录信息列表"`
	model.SysLogininforListQuery
	request.PageInfo
}

type ListLogininforRes struct {
	Items []*model.SysLogininforListModel `json:"items"`
	response.PageResult
}
