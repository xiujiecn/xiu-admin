// Package v1
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package v1

import (
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/request"
	"xiuadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type RoleListReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"系统-角色管理" summary:"获取角色列表" x-check-permission:"cpm:system:role:list"`
	request.PageInfo
	model.SysRoleListParam
}

type RoleListRes struct {
	response.PageResult
	Data []*model.SysRoleListModel `json:"items" dc:"角色列表"`
}

type RoleAddReq struct {
	g.Meta `path:"/role/add" method:"post" tags:"系统-角色管理" summary:"新增角色" x-check-permission:"cpm:system:role:add"`
	model.SysRoleAddParam
}

type RoleAddRes struct {
	model.SysRoleAddModel
}

type RoleEditReq struct {
	g.Meta `path:"/role/edit" method:"post" tags:"系统-角色管理" summary:"编辑角色" x-check-permission:"cpm:system:role:edit"`
	model.SysRoleEditParam
}

type RoleEditRes struct {
	model.SysRoleEditModel
}

type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" method:"post" tags:"系统-角色管理" summary:"删除角色" x-check-permission:"cpm:system:role:remove"`
	model.SysRoleDeleteParam
}

type RoleDeleteRes struct {
	model.SysRoleDeleteModel
}

type RoleViewReq struct {
	g.Meta `path:"/role/view" method:"get" tags:"系统-角色管理" summary:"获取角色详情" x-check-permission:"cpm:system:role:query"`
	model.SysRoleViewParam
}

type RoleViewRes struct {
	model.SysRoleViewModel
}

type RoleDataScopeEditReq struct {
	g.Meta `path:"/role/dataScope" method:"post" tags:"系统-角色管理" summary:"编辑角色数据权限" x-check-permission:"cpm:system:role:edit"`
	model.SysRoleDataScopeEditParam
}

type RoleDataScopeEditRes struct {
}
