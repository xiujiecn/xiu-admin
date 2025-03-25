// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TestDemo is the golang structure of table test_demo for DAO operations like Where/Data.
type TestDemo struct {
	g.Meta      `orm:"table:test_demo, do:true"`
	Id          interface{} // 主键
	TenantId    interface{} // 租户编号
	DeptId      interface{} // 部门id
	UserId      interface{} // 用户id
	OrderNum    interface{} // 排序号
	TestKey     interface{} // key键
	Value       interface{} // 值
	Version     interface{} // 版本
	CreatedDept interface{} // 创建部门
	CreatedAt   *gtime.Time // 创建时间
	CreatedBy   interface{} // 创建者
	UpdatedAt   *gtime.Time // 更新时间
	UpdatedBy   interface{} // 更新者
	DeletedBy   interface{} // 删除人
	DeletedAt   *gtime.Time // 删除时间
}
