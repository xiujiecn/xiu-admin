package v1

import (
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type ListLogininforReq struct {
	g.Meta `path:"/logininfor/list" method:"get" tags:"系统" summary:"登录信息列表"`
	model.SysLogininforListParam
	request.PageInfo
}

type ListLogininforRes struct {
	Items []*model.SysLogininforListModel `json:"items"`
	response.PageResult
}
