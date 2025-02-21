package v1

import (
	"server/internal/model"
	"server/internal/model/request"
	"server/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type GetConfigListReq struct {
	g.Meta `path:"/config/list" method:"get" tags:"系统配置" summary:"获取系统配置列表"`
	request.PageInfo
	model.SysConfigListQuery
}

type GetConfigListRes struct {
	response.PageResult
	Data []*model.SysConfig `json:"items" dc:"系统配置列表"`
}
