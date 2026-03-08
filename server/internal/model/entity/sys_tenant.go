// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTenant is the golang structure for table sys_tenant.
type SysTenant struct {
	Id              int64       `json:"id"              orm:"id"                description:"id"`
	TenantId        string      `json:"tenantId"        orm:"tenant_id"         description:"租户编号"`
	ContactUserName string      `json:"contactUserName" orm:"contact_user_name" description:"联系人"`
	ContactPhone    string      `json:"contactPhone"    orm:"contact_phone"     description:"联系电话"`
	CompanyName     string      `json:"companyName"     orm:"company_name"      description:"企业名称"`
	LicenseNumber   string      `json:"licenseNumber"   orm:"license_number"    description:"统一社会信用代码"`
	Address         string      `json:"address"         orm:"address"           description:"地址"`
	Intro           string      `json:"intro"           orm:"intro"             description:"企业简介"`
	Domain          string      `json:"domain"          orm:"domain"            description:"域名"`
	Remark          string      `json:"remark"          orm:"remark"            description:"备注"`
	PackageId       int64       `json:"packageId"       orm:"package_id"        description:"租户套餐编号"`
	AdminRoleId     int64       `json:"adminRoleId"     orm:"admin_role_id"     description:"管理员角色ID"`
	AdminDeptId     int64       `json:"adminDeptId"     orm:"admin_dept_id"     description:"管理员部门ID"`
	AdminUserId     int64       `json:"adminUserId"     orm:"admin_user_id"     description:"管理员用户ID"`
	ExpireTime      *gtime.Time `json:"expireTime"      orm:"expire_time"       description:"过期时间"`
	AccountCount    int         `json:"accountCount"    orm:"account_count"     description:"用户数量（-1不限制）"`
	Status          string      `json:"status"          orm:"status"            description:"租户状态（0正常 1停用）"`
	CreatedDept     int64       `json:"createdDept"     orm:"created_dept"      description:"创建部门"`
	CreatedBy       int64       `json:"createdBy"       orm:"created_by"        description:"创建者"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"        description:"创建时间"`
	UpdatedBy       int64       `json:"updatedBy"       orm:"updated_by"        description:"更新者"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"        description:"更新时间"`
	DeletedBy       int64       `json:"deletedBy"       orm:"deleted_by"        description:"删除人"`
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"        description:"删除时间"`
}
