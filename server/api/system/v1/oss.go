package v1

import (
	"server/internal/model"
	"server/internal/model/request"
	"server/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type SysOssListReq struct {
	g.Meta `path:"/oss/list" method:"get" tags:"系统" summary:"获取OSS列表"`
	model.SysOssListQuery
	request.PageInfo
}

type SysOssListRes struct {
	Items []*model.SysOssListModel `json:"items"`
	response.PageResult
}
