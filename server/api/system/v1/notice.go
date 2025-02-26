package v1

import (
	"server/internal/model"
	"server/internal/model/request"
	"server/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type SysNoticeListReq struct {
	g.Meta `path:"/notice/list" method:"get" tags:"公告" summary:"公告列表"`
	request.PageInfo
	model.SysNoticeListParam
}

type SysNoticeListRes struct {
	response.PageResult
	Items []*model.SysNotice `json:"items"`
}
