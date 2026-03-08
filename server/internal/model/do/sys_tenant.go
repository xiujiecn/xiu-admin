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
	Id              any         // id
	TenantId        any         // 租户编号
	ContactUserName any         // 联系人
	ContactPhone    any         // 联系电话
	CompanyName     any         // 企业名称
	LicenseNumber   any         // 统一社会信用代码
	Address         any         // 地址
	Intro           any         // 企业简介
	Domain          any         // 域名
	Remark          any         // 备注
	PackageId       any         // 租户套餐编号
	AdminRoleId     any         // 管理员角色ID
	AdminDeptId     any         // 管理员部门ID
	AdminUserId     any         // 管理员用户ID
	ExpireTime      *gtime.Time // 过期时间
	AccountCount    any         // 用户数量（-1不限制）
	Status          any         // 租户状态（0正常 1停用）
	CreatedDept     any         // 创建部门
	CreatedBy       any         // 创建者
	CreatedAt       *gtime.Time // 创建时间
	UpdatedBy       any         // 更新者
	UpdatedAt       *gtime.Time // 更新时间
	DeletedBy       any         // 删除人
	DeletedAt       *gtime.Time // 删除时间
}
