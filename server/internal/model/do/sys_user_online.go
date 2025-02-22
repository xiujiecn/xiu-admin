// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserOnline is the golang structure of table sys_user_online for DAO operations like Where/Data.
type SysUserOnline struct {
	g.Meta        `orm:"table:sys_user_online, do:true"`
	OnlineId      interface{} // 访问ID
	TenantId      interface{} // 租户编号
	Uuid          interface{} // UUID
	UserName      interface{} // 用户账号
	ClientKey     interface{} // 客户端
	DeviceType    interface{} // 设备类型
	Ipaddr        interface{} // 登录IP地址
	LoginLocation interface{} // 登录地点
	Browser       interface{} // 浏览器类型
	Os            interface{} // 操作系统
	Token         interface{} // Token
	LoginTime     *gtime.Time // 访问时间
	ExpireTime    *gtime.Time // 过期时间
	DeletedAt     *gtime.Time // 删除时间
}
