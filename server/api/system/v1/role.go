package v1

import (
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type RoleListReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"系统" summary:"获取角色列表"`
	request.PageInfo
	model.SysRoleListParam
}

type RoleListRes struct {
	response.PageResult
	Data []*model.SysRoleListModel `json:"items" dc:"角色列表"`
}

type RoleAddReq struct {
	g.Meta `path:"/role/add" method:"post" tags:"系统" summary:"新增角色"`
	model.SysRoleAddParam
}

type RoleAddRes struct {
	model.SysRoleAddModel
}

type RoleEditReq struct {
	g.Meta `path:"/role/edit" method:"post" tags:"系统" summary:"编辑角色"`
	model.SysRoleEditParam
}

type RoleEditRes struct {
	model.SysRoleEditModel
}

type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" method:"post" tags:"系统" summary:"删除角色"`
	model.SysRoleDeleteParam
}

type RoleDeleteRes struct {
	model.SysRoleDeleteModel
}

type RoleViewReq struct {
	g.Meta `path:"/role/view" method:"get" tags:"系统" summary:"获取角色详情"`
	model.SysRoleViewParam
}

type RoleViewRes struct {
	model.SysRoleViewModel
}

type RoleDataScopeEditReq struct {
	g.Meta `path:"/role/dataScope" method:"post" tags:"系统" summary:"编辑角色数据权限"`
	model.SysRoleDataScopeEditParam
}

type RoleDataScopeEditRes struct {
}
