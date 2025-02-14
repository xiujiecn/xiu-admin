// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTenant is the golang structure of table sys_tenant for DAO operations like Where/Data.
type SysTenant struct {
	g.Meta          `orm:"table:sys_tenant, do:true"`
	Id              interface{} // id
	TenantId        interface{} // 租户编号
	ContactUserName interface{} // 联系人
	ContactPhone    interface{} // 联系电话
	CompanyName     interface{} // 企业名称
	LicenseNumber   interface{} // 统一社会信用代码
	Address         interface{} // 地址
	Intro           interface{} // 企业简介
	Domain          interface{} // 域名
	Remark          interface{} // 备注
	PackageId       interface{} // 租户套餐编号
	ExpireTime      *gtime.Time // 过期时间
	AccountCount    interface{} // 用户数量（-1不限制）
	Status          interface{} // 租户状态（0正常 1停用）
	CreatedDept     interface{} // 创建部门
	CreatedBy       interface{} // 创建者
	CreatedAt       *gtime.Time // 创建时间
	UpdatedBy       interface{} // 更新者
	UpdatedAt       *gtime.Time // 更新时间
	DeletedBy       interface{} // 删除人
	DeletedAt       *gtime.Time // 删除时间
}
