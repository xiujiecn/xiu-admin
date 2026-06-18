// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure of table sys_role for DAO operations like Where/Data.
type SysRole struct {
	g.Meta            `orm:"table:sys_role, do:true"`
	RoleId            any         // 角色ID
	TenantId          any         // 租户编号
	DeptId            any         // 部门机构id
	RoleName          any         // 角色名称
	RoleKey           any         // 角色权限字符串
	RoleSort          any         // 显示顺序
	DataScope         any         // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	MenuCheckStrictly any         // 菜单树选择项是否关联显示
	DeptCheckStrictly any         // 部门树选择项是否关联显示
	Status            any         // 角色状态（0正常 1停用）
	CreatedDept       any         // 创建部门
	CreatedBy         any         // 创建者
	CreatedAt         *gtime.Time // 创建时间
	UpdatedBy         any         // 更新者
	UpdatedAt         *gtime.Time // 更新时间
	DeletedBy         any         // 删除人
	DeletedAt         *gtime.Time // 删除时间
	Remark            any         // 备注
}
