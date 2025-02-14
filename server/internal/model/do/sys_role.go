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
	RoleId            interface{} // 角色ID
	TenantId          interface{} // 租户编号
	RoleName          interface{} // 角色名称
	RoleKey           interface{} // 角色权限字符串
	RoleSort          interface{} // 显示顺序
	DataScope         interface{} // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	MenuCheckStrictly interface{} // 菜单树选择项是否关联显示
	DeptCheckStrictly interface{} // 部门树选择项是否关联显示
	Status            interface{} // 角色状态（0正常 1停用）
	CreatedDept       interface{} // 创建部门
	CreatedBy         interface{} // 创建者
	CreatedAt         *gtime.Time // 创建时间
	UpdatedBy         interface{} // 更新者
	UpdatedAt         *gtime.Time // 更新时间
	DeletedBy         interface{} // 删除人
	DeletedAt         *gtime.Time // 删除时间
	Remark            interface{} // 备注
}
