package v1

import (
	"server/internal/model"
	"server/internal/model/request"
	"server/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type SysSocialListReq struct {
	g.Meta `path:"/social/list" method:"get" tags:"社交" summary:"社交列表"`
	model.SysSocialListParam
	request.PageInfo
}

type SysSocialListRes struct {
	Items []*model.SysSocialListModel `json:"items"`
	response.PageResult
}
