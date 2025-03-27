// package model
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package model

import (
	"xiuadmin/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

type SysMenuTree struct {
	entity.SysMenu
	Children []*SysMenuTree `json:"children"`
}

type SysMenuListModel struct {
	MenuId      int64       `json:"menuId"      orm:"menu_id"      description:"菜单ID"`
	MenuName    string      `json:"menuName"    orm:"menu_name"    description:"菜单名称"`
	ParentId    int64       `json:"parentId"    orm:"parent_id"    description:"父菜单ID"`
	OrderNum    int         `json:"orderNum"    orm:"order_num"    description:"显示顺序"`
	Path        string      `json:"path"        orm:"path"         description:"路由地址"`
	Component   string      `json:"component"   orm:"component"    description:"组件路径"`
	QueryParam  string      `json:"queryParam"  orm:"query_param"  description:"路由参数"`
	IsFrame     int         `json:"isFrame"     orm:"is_frame"     description:"是否为外链（0是 1否）"`
	IsCache     int         `json:"isCache"     orm:"is_cache"     description:"是否缓存（0缓存 1不缓存）"`
	MenuType    string      `json:"menuType"    orm:"menu_type"    description:"菜单类型（M目录 C菜单 F按钮）"`
	Visible     string      `json:"visible"     orm:"visible"      description:"显示状态（0显示 1隐藏）"`
	Status      string      `json:"status"      orm:"status"       description:"菜单状态（0正常 1停用）"`
	Perms       string      `json:"perms"       orm:"perms"        description:"权限标识"`
	Icon        string      `json:"icon"        orm:"icon"         description:"菜单图标"`
	CreatedDept int64       `json:"createdDept" orm:"created_dept" description:"创建部门"`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:"创建者"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
}

type SysMenuListParam struct {
	MenuName string  `json:"menuName"`
	Status   string  `json:"status"`
	MenuIds  []int64 `json:"menuIds"`
}

type SysMenuViewModel struct {
	entity.SysMenu
}

type SysMenuAddModel struct {
	MenuId      int64       `json:"menuId"      orm:"menu_id"      description:"菜单ID"`
	MenuName    string      `json:"menuName"    orm:"menu_name"    description:"菜单名称"`
	ParentId    int64       `json:"parentId"    orm:"parent_id"    description:"父菜单ID"`
	OrderNum    int         `json:"orderNum"    orm:"order_num"    description:"显示顺序"`
	Path        string      `json:"path"        orm:"path"         description:"路由地址"`
	Component   string      `json:"component"   orm:"component"    description:"组件路径"`
	QueryParam  string      `json:"queryParam"  orm:"query_param"  description:"路由参数"`
	IsFrame     int         `json:"isFrame"     orm:"is_frame"     description:"是否为外链（0是 1否）"`
	IsCache     int         `json:"isCache"     orm:"is_cache"     description:"是否缓存（0缓存 1不缓存）"`
	MenuType    string      `json:"menuType"    orm:"menu_type"    description:"菜单类型（M目录 C菜单 F按钮）"`
	Visible     string      `json:"visible"     orm:"visible"      description:"显示状态（0显示 1隐藏）"`
	Status      string      `json:"status"      orm:"status"       description:"菜单状态（0正常 1停用）"`
	Perms       string      `json:"perms"       orm:"perms"        description:"权限标识"`
	Icon        string      `json:"icon"        orm:"icon"         description:"菜单图标"`
	CreatedDept int64       `json:"createdDept" orm:"created_dept" description:"创建部门"`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:"创建者"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"   description:"更新者"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
}

type SysMenuUpdateModel struct {
	MenuId     int64       `json:"menuId"`
	MenuName   *string     `json:"menuName"    orm:"menu_name"    description:"菜单名称"`
	ParentId   *int64      `json:"parentId"    orm:"parent_id"    description:"父菜单ID"`
	OrderNum   *int        `json:"orderNum"    orm:"order_num"    description:"显示顺序"`
	Path       *string     `json:"path"        orm:"path"         description:"路由地址"`
	Component  *string     `json:"component"   orm:"component"    description:"组件路径"`
	QueryParam *string     `json:"queryParam"  orm:"query_param"  description:"路由参数"`
	IsFrame    *int        `json:"isFrame"     orm:"is_frame"     description:"是否为外链（0是 1否）"`
	IsCache    *int        `json:"isCache"     orm:"is_cache"     description:"是否缓存（0缓存 1不缓存）"`
	MenuType   *string     `json:"menuType"    orm:"menu_type"    description:"菜单类型（M目录 C菜单 F按钮）"`
	Visible    *string     `json:"visible"     orm:"visible"      description:"显示状态（0显示 1隐藏）"`
	Status     *string     `json:"status"      orm:"status"       description:"菜单状态（0正常 1停用）"`
	Perms      *string     `json:"perms"       orm:"perms"        description:"权限标识"`
	Icon       *string     `json:"icon"        orm:"icon"         description:"菜单图标"`
	UpdatedBy  int64       `json:"updatedBy"   orm:"updated_by"   description:"更新者"`
	UpdatedAt  *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
	Remark     *string     `json:"remark"      orm:"remark"       description:"备注"`
}

type SysMenuDeleteModel struct {
	MenuId int64 `json:"menuId"`
}

type SysMenuMini struct {
	MenuId   int64   `json:"menuId"`
	MenuName *string `json:"menuName"    orm:"menu_name"    description:"菜单名称"`
	ParentId *int64  `json:"parentId"    orm:"parent_id"    description:"父菜单ID"`
}

type SysMenuTreeModel struct {
	Title    string              `json:"title"`
	Key      string              `json:"key"`
	ParentId int64               `json:"parentId"`
	MenuId   int64               `json:"menuId"`
	MenuName string              `json:"menuName"`
	MenuType string              `json:"menuType"    orm:"menu_type"    description:"菜单类型（M目录 C菜单 F按钮）"`
	Children []*SysMenuTreeModel `json:"children"`
}
