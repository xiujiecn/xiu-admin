package model

import "github.com/gogf/gf/v2/os/gtime"

type SysRole struct {
	RoleId            int64       `json:"roleId"            orm:"role_id"             description:"角色ID"`
	TenantId          string      `json:"tenantId"          orm:"tenant_id"           description:"租户编号"`
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
}

type SysRoleListQuery struct {
	RoleName  string      `json:"roleName" orm:"role_name"`
	Status    string      `json:"status" orm:"status"`
	RoleKey   string      `json:"roleKey"           orm:"role_key"            description:"角色权限字符串"`
	CreatedAt *gtime.Time `json:"createdAt"         orm:"created_at"          description:"创建时间"`
}

type SysRoleMiniModel struct {
	RoleId   int64  `json:"roleId"   orm:"role_id"   description:"角色ID"`
	RoleName string `json:"roleName" orm:"role_name" description:"角色名称"`
}
