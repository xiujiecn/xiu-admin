// package model
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package model

import (
	"xiuadmin/internal/model/request"

	"github.com/gogf/gf/v2/os/gtime"
)

type SysRoleListModel struct {
	RoleId            int64       `json:"roleId"            orm:"role_id"             description:"角色ID"`
	TenantId          string      `json:"tenantId"          orm:"tenant_id"           description:"租户编号"`
	DeptId            int64       `json:"deptId"            orm:"dept_id"             description:"部门组织id"`
	RoleName          string      `json:"roleName"          orm:"role_name"           description:"角色名称"`
	RoleKey           string      `json:"roleKey"           orm:"role_key"            description:"角色权限字符串"`
	RoleSort          int         `json:"roleSort"          orm:"role_sort"           description:"显示顺序"`
	DataScope         string      `json:"dataScope"         orm:"data_scope"          description:"数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）"`
	MenuCheckStrictly int         `json:"menuCheckStrictly" orm:"menu_check_strictly" description:"菜单树选择项是否关联显示"`
	DeptCheckStrictly int         `json:"deptCheckStrictly" orm:"dept_check_strictly" description:"部门树选择项是否关联显示"`
	Status            string      `json:"status"            orm:"status"              description:"角色状态（0正常 1停用）"`
	CreatedDept       int64       `json:"createdDept"       orm:"created_dept"        description:"创建部门"`
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"          description:"创建时间"`
	Remark            string      `json:"remark"            orm:"remark"              description:"备注"`
	DeptName          string      `json:"deptName"          orm:"dept_name"           description:"部门名称"`
}

type SysRoleListParam struct {
	request.PageInfo
	RoleName  string   `json:"roleName" orm:"role_name"`
	Status    string   `json:"status" orm:"status"`
	RoleKey   string   `json:"roleKey"                  description:"角色权限字符串"`
	CreatedAt []string `json:"createdAt"                description:"创建时间"`
	RoleIds   []int64  `json:"roleIds"                  description:"角色ID"`
	DeptId    *int64   `json:"deptId"                   description:"部门组织id"`
}

type SysRoleMiniModel struct {
	RoleId    int64  `json:"roleId"   orm:"role_id"   description:"角色ID"`
	RoleName  string `json:"roleName" orm:"role_name" description:"角色名称"`
	DataScope string `json:"dataScope" orm:"data_scope" description:"数据范围"`
}
type SysRoleViewParam struct {
	RoleId int64 `json:"roleId" orm:"role_id"`
}

type SysRoleViewModel struct {
	RoleId             int64            `json:"roleId"            orm:"role_id"             description:"角色ID"`
	TenantId           string           `json:"tenantId"          orm:"tenant_id"           description:"租户编号"`
	DeptId             int64            `json:"deptId"            orm:"dept_id"             description:"部门组织id"`
	RoleName           string           `json:"roleName"          orm:"role_name"           description:"角色名称"`
	RoleKey            string           `json:"roleKey"           orm:"role_key"            description:"角色权限字符串"`
	RoleSort           int              `json:"roleSort"          orm:"role_sort"           description:"显示顺序"`
	DataScope          string           `json:"dataScope"         orm:"data_scope"          description:"数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）"`
	MenuCheckStrictly  int              `json:"menuCheckStrictly" orm:"menu_check_strictly" description:"菜单树选择项是否关联显示"`
	DeptCheckStrictly  int              `json:"deptCheckStrictly" orm:"dept_check_strictly" description:"部门树选择项是否关联显示"`
	Status             string           `json:"status"            orm:"status"              description:"角色状态（0正常 1停用）"`
	CreatedDept        int64            `json:"createdDept"       orm:"created_dept"        description:"创建部门"`
	CreatedAt          *gtime.Time      `json:"createdAt"         orm:"created_at"          description:"创建时间"`
	Remark             string           `json:"remark"            orm:"remark"              description:"备注"`
	MenuIds            []int64          `json:"menuIds"           orm:"menu_ids"            description:"角色菜单ID列表"`
	RoleMenuDataScopes map[int64]string `json:"roleMenuDataScopes" description:"角色菜单数据范围"`
	DeptIds            []int64          `json:"deptIds"           orm:"dept_ids"            description:"角色部门ID列表"`
}

type SysRoleAddParam struct {
	RoleName           string           `json:"roleName"          orm:"role_name"           description:"角色名称"`
	TenantId           string           `json:"tenantId"          orm:"tenant_id"           description:"租户编号"`
	RoleKey            string           `json:"roleKey"           orm:"role_key"            description:"角色权限字符串"`
	RoleSort           int              `json:"roleSort"          orm:"role_sort"           description:"显示顺序"`
	DataScope          string           `json:"dataScope"         orm:"data_scope"          description:"数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）"`
	MenuCheckStrictly  int              `json:"menuCheckStrictly" orm:"menu_check_strictly" description:"菜单树选择项是否关联显示"`
	Status             string           `json:"status"            orm:"status"              description:"角色状态（0正常 1停用）"`
	CreatedDept        int64            `json:"createdDept"       orm:"created_dept"        description:"创建部门"`
	CreatedBy          int64            `json:"createdBy"         orm:"created_by"          description:"创建者"`
	CreatedAt          *gtime.Time      `json:"createdAt"         orm:"created_at"          description:"创建时间"`
	UpdatedBy          int64            `json:"updatedBy"         orm:"updated_by"          description:"更新者"`
	UpdatedAt          *gtime.Time      `json:"updatedAt"         orm:"updated_at"          description:"更新时间"`
	Remark             string           `json:"remark"            orm:"remark"              description:"备注"`
	MenuIds            []int64          `json:"menuIds"           orm:"menu_ids"            description:"菜单ID"`
	RoleMenuDataScopes map[int64]string `json:"roleMenuDataScopes" description:"角色菜单数据范围"`
	DeptId             *int64           `json:"deptId"            orm:"dept_id"             description:"部门组织id"`
}

type SysRoleAddModel struct {
	RoleId int64 `json:"roleId" orm:"role_id" description:"角色ID"`
}

type SysRoleEditParam struct {
	RoleId             int64            `json:"roleId"   orm:"role_id"   description:"角色ID"`
	RoleName           *string          `json:"roleName"          orm:"role_name"           description:"角色名称"`
	RoleKey            *string          `json:"roleKey"           orm:"role_key"            description:"角色权限字符串"`
	RoleSort           *int             `json:"roleSort"          orm:"role_sort"           description:"显示顺序"`
	MenuCheckStrictly  *int             `json:"menuCheckStrictly" orm:"menu_check_strictly" description:"菜单树选择项是否关联显示"`
	Status             *string          `json:"status"            orm:"status"              description:"角色状态（0正常 1停用）"`
	UpdatedBy          *int64           `json:"updatedBy"         orm:"updated_by"          description:"更新者"`
	UpdatedAt          *gtime.Time      `json:"updatedAt"         orm:"updated_at"          description:"更新时间"`
	Remark             *string          `json:"remark"            orm:"remark"              description:"备注"`
	MenuIds            []int64          `json:"menuIds"           orm:"menu_ids"            description:"菜单ID"`
	RoleMenuDataScopes map[int64]string `json:"roleMenuDataScopes" description:"角色菜单数据范围"`
	DeptId             *int64           `json:"deptId"            orm:"dept_id"             description:"部门组织id"`
}

type SysRoleEditModel struct {
	RoleId int64 `json:"roleId" orm:"role_id"`
}

type SysRoleDeleteParam struct {
	RoleId  int64   `json:"roleId" orm:"role_id"`
	RoleIds []int64 `json:"roleIds" orm:"role_ids"`
}
type SysRoleDeleteModel struct {
	RoleId  int64   `json:"roleId" orm:"role_id"`
	RoleIds []int64 `json:"roleIds" orm:"role_ids"`
}

type SysRoleDataScopeEditParam struct {
	RoleId            int64   `json:"roleId" description:"角色ID"`
	DataScope         string  `json:"dataScope" description:"数据范围"`
	DeptCheckStrictly *int    `json:"deptCheckStrictly" description:"部门树选择项是否关联显示"`
	DeptIds           []int64 `json:"deptIds" description:"部门ID"`
}
