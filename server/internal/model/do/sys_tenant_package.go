// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTenantPackage is the golang structure of table sys_tenant_package for DAO operations like Where/Data.
type SysTenantPackage struct {
	g.Meta            `orm:"table:sys_tenant_package, do:true"`
	PackageId         interface{} // 租户套餐id
	PackageName       interface{} // 套餐名称
	MenuIds           interface{} // 关联菜单id
	Remark            interface{} // 备注
	MenuCheckStrictly interface{} // 菜单树选择项是否关联显示
	Status            interface{} // 状态（0正常 1停用）
	CreatedDept       interface{} // 创建部门
	CreatedBy         interface{} // 创建者
	CreatedAt         *gtime.Time // 创建时间
	UpdatedBy         interface{} // 更新者
	UpdatedAt         *gtime.Time // 更新时间
	DeletedBy         interface{} // 删除人
	DeletedAt         *gtime.Time // 删除时间
}
