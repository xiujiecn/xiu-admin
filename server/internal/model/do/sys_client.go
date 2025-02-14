// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysClient is the golang structure of table sys_client for DAO operations like Where/Data.
type SysClient struct {
	g.Meta        `orm:"table:sys_client, do:true"`
	Id            interface{} // id
	ClientId      interface{} // 客户端id
	ClientKey     interface{} // 客户端key
	ClientSecret  interface{} // 客户端秘钥
	GrantType     interface{} // 授权类型
	DeviceType    interface{} // 设备类型
	ActiveTimeout interface{} // token活跃超时时间
	Timeout       interface{} // token固定超时
	Status        interface{} // 状态（0正常 1停用）
	CreatedDept   interface{} // 创建部门
	CreatedBy     interface{} // 创建者
	CreatedAt     *gtime.Time // 创建时间
	UpdatedBy     interface{} // 更新者
	UpdatedAt     *gtime.Time // 更新时间
	DeletedBy     interface{} // 删除人
	DeletedAt     *gtime.Time // 删除时间
}
