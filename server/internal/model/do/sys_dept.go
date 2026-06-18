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
	DeptId       any         // 部门id
	TenantId     any         // 租户编号
	ParentId     any         // 父部门id
	Ancestors    any         // 祖级列表
	DeptName     any         // 部门名称
	DeptType     any         // 部门类型(0:部门 1:公司)
	DeptCategory any         // 部门类别编码
	OrderNum     any         // 显示顺序
	Leader       any         // 负责人
	Phone        any         // 联系电话
	Email        any         // 邮箱
	Status       any         // 部门状态（0正常 1停用）
	CreatedDept  any         // 创建部门
	CreatedBy    any         // 创建者
	CreatedAt    *gtime.Time // 创建时间
	UpdatedBy    any         // 更新者
	UpdatedAt    *gtime.Time // 更新时间
	DeletedBy    any         // 删除人
	DeletedAt    *gtime.Time // 删除时间
}
