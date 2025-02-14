// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDept is the golang structure of table sys_dept for DAO operations like Where/Data.
type SysDept struct {
	g.Meta       `orm:"table:sys_dept, do:true"`
	DeptId       interface{} // 部门id
	TenantId     interface{} // 租户编号
	ParentId     interface{} // 父部门id
	Ancestors    interface{} // 祖级列表
	DeptName     interface{} // 部门名称
	DeptCategory interface{} // 部门类别编码
	OrderNum     interface{} // 显示顺序
	Leader       interface{} // 负责人
	Phone        interface{} // 联系电话
	Email        interface{} // 邮箱
	Status       interface{} // 部门状态（0正常 1停用）
	CreatedDept  interface{} // 创建部门
	CreatedBy    interface{} // 创建者
	CreatedAt    *gtime.Time // 创建时间
	UpdatedBy    interface{} // 更新者
	UpdatedAt    *gtime.Time // 更新时间
	DeletedBy    interface{} // 删除人
	DeletedAt    *gtime.Time // 删除时间
}
