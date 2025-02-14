// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPost is the golang structure of table sys_post for DAO operations like Where/Data.
type SysPost struct {
	g.Meta       `orm:"table:sys_post, do:true"`
	PostId       interface{} // 岗位ID
	TenantId     interface{} // 租户编号
	DeptId       interface{} // 部门id
	PostCode     interface{} // 岗位编码
	PostCategory interface{} // 岗位类别编码
	PostName     interface{} // 岗位名称
	PostSort     interface{} // 显示顺序
	Status       interface{} // 状态（0正常 1停用）
	CreatedDept  interface{} // 创建部门
	CreatedBy    interface{} // 创建者
	CreatedAt    *gtime.Time // 创建时间
	UpdatedBy    interface{} // 更新者
	UpdatedAt    *gtime.Time // 更新时间
	DeletedBy    interface{} // 删除人
	DeletedAt    *gtime.Time // 删除时间
	Remark       interface{} // 备注
}
