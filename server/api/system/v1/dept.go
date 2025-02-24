package v1

import (
	"server/internal/model"
	"server/internal/model/request"
	"server/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type DeptListReq struct {
	g.Meta `path:"/dept/list" method:"get" tags:"系统" summary:"获取部门列表"`
	request.PageInfo
	model.DeptListQuery
}

type DeptListRes struct {
	response.PageResult
	Data []*model.SysDept `json:"items" dc:"部门列表"`
}

type DeptTreeReq struct {
	g.Meta `path:"/dept/tree" method:"get" tags:"系统" summary:"获取部门树"`
}

type DeptTreeRes struct {
	Data []*model.SysDeptTreeModel `json:"items" dc:"部门树"`
}
