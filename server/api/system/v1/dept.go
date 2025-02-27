package v1

import (
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type DeptListReq struct {
	g.Meta `path:"/dept/list" method:"get" tags:"系统" summary:"获取部门列表"`
	request.PageInfo
	model.SysDeptListParam
}

type DeptListRes struct {
	response.PageResult
	Data []*model.SysDeptListModel `json:"items" dc:"部门列表"`
}

type DeptTreeReq struct {
	g.Meta `path:"/dept/tree" method:"get" tags:"系统" summary:"获取部门树"`
}

type DeptTreeRes struct {
	Data []*model.SysDeptTreeModel `json:"items" dc:"部门树"`
}

type DeptAddReq struct {
	g.Meta `path:"/dept/add" method:"post" tags:"系统" summary:"新增部门"`
	model.SysDeptAddModel
}

type DeptAddRes struct {
	DeptId int64 `json:"deptId" dc:"部门id"`
}

type DeptEditReq struct {
	g.Meta `path:"/dept/edit" method:"post" tags:"系统" summary:"编辑部门"`
	model.SysDeptEditModel
}

type DeptEditRes struct {
	DeptId int64 `json:"deptId" dc:"部门id"`
}

type DeptDeleteReq struct {
	g.Meta `path:"/dept/delete" method:"post" tags:"系统" summary:"删除部门"`
	model.SysDeptDeleteModel
}

type DeptDeleteRes struct {
	DeptId int64 `json:"deptId" dc:"部门id"`
}

type DeptViewReq struct {
	g.Meta `path:"/dept/view" method:"get" tags:"系统" summary:"获取部门详情"`
	model.SysDeptViewModel
}

type DeptViewRes struct {
	model.SysDeptViewModel
}
