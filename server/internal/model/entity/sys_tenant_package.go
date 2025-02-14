// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTenantPackage is the golang structure for table sys_tenant_package.
type SysTenantPackage struct {
	PackageId         int64       `json:"packageId"         orm:"package_id"          description:"租户套餐id"`
	PackageName       string      `json:"packageName"       orm:"package_name"        description:"套餐名称"`
	MenuIds           string      `json:"menuIds"           orm:"menu_ids"            description:"关联菜单id"`
	Remark            string      `json:"remark"            orm:"remark"              description:"备注"`
	MenuCheckStrictly int         `json:"menuCheckStrictly" orm:"menu_check_strictly" description:"菜单树选择项是否关联显示"`
	Status            string      `json:"status"            orm:"status"              description:"状态（0正常 1停用）"`
	CreatedDept       int64       `json:"createdDept"       orm:"created_dept"        description:"创建部门"`
	CreatedBy         int64       `json:"createdBy"         orm:"created_by"          description:"创建者"`
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"          description:"创建时间"`
	UpdatedBy         int64       `json:"updatedBy"         orm:"updated_by"          description:"更新者"`
	UpdatedAt         *gtime.Time `json:"updatedAt"         orm:"updated_at"          description:"更新时间"`
	DeletedBy         int64       `json:"deletedBy"         orm:"deleted_by"          description:"删除人"`
	DeletedAt         *gtime.Time `json:"deletedAt"         orm:"deleted_at"          description:"删除时间"`
}
