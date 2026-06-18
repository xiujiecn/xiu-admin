// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TestTree is the golang structure for table test_tree.
type TestTree struct {
	Id          int64       `json:"id"          orm:"id"           description:"主键"`
	TenantId    string      `json:"tenantId"    orm:"tenant_id"    description:"租户编号"`
	ParentId    int64       `json:"parentId"    orm:"parent_id"    description:"父id"`
	DeptId      int64       `json:"deptId"      orm:"dept_id"      description:"部门id"`
	UserId      int64       `json:"userId"      orm:"user_id"      description:"用户id"`
	TreeName    string      `json:"treeName"    orm:"tree_name"    description:"值"`
	Level       int         `json:"level"       orm:"level"        description:"关系树等级"`
	Tree        string      `json:"tree"        orm:"tree"         description:"关系树"`
	Version     int         `json:"version"     orm:"version"      description:"版本"`
	CreatedDept int64       `json:"createdDept" orm:"created_dept" description:"创建部门"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:"创建者"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"   description:"更新者"`
	DeletedBy   int64       `json:"deletedBy"   orm:"deleted_by"   description:"删除人"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:"删除时间"`
}
