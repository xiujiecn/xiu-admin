package v1

import (
	"server/internal/model"
	"server/internal/model/request"
	"server/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type RoleListReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"系统" summary:"获取角色列表"`
	request.PageInfo
	model.SysRoleListQuery
}

type RoleListRes struct {
	response.PageResult
	Data []*model.SysRole `json:"items" dc:"角色列表"`
}
