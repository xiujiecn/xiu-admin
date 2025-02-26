package v1

import (
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type SysClientListReq struct {
	g.Meta `path:"/client/list" method:"get" tags:"系统" summary:"获取客户端列表"`
	model.SysClientListParam
	request.PageInfo
}

type SysClientListRes struct {
	Items []*model.SysClientListModel `json:"items"`
	response.PageResult
}
