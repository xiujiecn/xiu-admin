// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDept is the golang structure for table sys_dept.
type SysDept struct {
	DeptId       int64       `json:"deptId"       orm:"dept_id"       description:"部门id"`
	TenantId     string      `json:"tenantId"     orm:"tenant_id"     description:"租户编号"`
	ParentId     int64       `json:"parentId"     orm:"parent_id"     description:"父部门id"`
	Ancestors    string      `json:"ancestors"    orm:"ancestors"     description:"祖级列表"`
	DeptName     string      `json:"deptName"     orm:"dept_name"     description:"部门名称"`
	DeptType     int         `json:"deptType"     orm:"dept_type"     description:"部门类型(0:部门 1:公司)"`
	DeptCategory string      `json:"deptCategory" orm:"dept_category" description:"部门类别编码"`
	OrderNum     int         `json:"orderNum"     orm:"order_num"     description:"显示顺序"`
	Leader       int64       `json:"leader"       orm:"leader"        description:"负责人"`
	Phone        string      `json:"phone"        orm:"phone"         description:"联系电话"`
	Email        string      `json:"email"        orm:"email"         description:"邮箱"`
	Status       string      `json:"status"       orm:"status"        description:"部门状态（0正常 1停用）"`
	CreatedDept  int64       `json:"createdDept"  orm:"created_dept"  description:"创建部门"`
	CreatedBy    int64       `json:"createdBy"    orm:"created_by"    description:"创建者"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	UpdatedBy    int64       `json:"updatedBy"    orm:"updated_by"    description:"更新者"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`
	DeletedBy    int64       `json:"deletedBy"    orm:"deleted_by"    description:"删除人"`
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:"删除时间"`
}
