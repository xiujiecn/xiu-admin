// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TestTree is the golang structure of table test_tree for DAO operations like Where/Data.
type TestTree struct {
	g.Meta      `orm:"table:test_tree, do:true"`
	Id          any         // 主键
	TenantId    any         // 租户编号
	ParentId    any         // 父id
	DeptId      any         // 部门id
	UserId      any         // 用户id
	TreeName    any         // 值
	Level       any         // 关系树等级
	Tree        any         // 关系树
	Version     any         // 版本
	CreatedDept any         // 创建部门
	CreatedAt   *gtime.Time // 创建时间
	CreatedBy   any         // 创建者
	UpdatedAt   *gtime.Time // 更新时间
	UpdatedBy   any         // 更新者
	DeletedBy   any         // 删除人
	DeletedAt   *gtime.Time // 删除时间
}
