// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserOnline is the golang structure for table sys_user_online.
type SysUserOnline struct {
	OnlineId      int64       `json:"onlineId"      orm:"online_id"      description:"访问ID"`
	TenantId      string      `json:"tenantId"      orm:"tenant_id"      description:"租户编号"`
	Uuid          string      `json:"uuid"          orm:"uuid"           description:"UUID"`
	UserName      string      `json:"userName"      orm:"user_name"      description:"用户账号"`
	ClientKey     string      `json:"clientKey"     orm:"client_key"     description:"客户端"`
	DeviceType    string      `json:"deviceType"    orm:"device_type"    description:"设备类型"`
	Ipaddr        string      `json:"ipaddr"        orm:"ipaddr"         description:"登录IP地址"`
	LoginLocation string      `json:"loginLocation" orm:"login_location" description:"登录地点"`
	Browser       string      `json:"browser"       orm:"browser"        description:"浏览器类型"`
	Os            string      `json:"os"            orm:"os"             description:"操作系统"`
	Token         string      `json:"token"         orm:"token"          description:"Token"`
	LoginTime     *gtime.Time `json:"loginTime"     orm:"login_time"     description:"访问时间"`
	ExpireTime    *gtime.Time `json:"expireTime"    orm:"expire_time"    description:"过期时间"`
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"     description:"删除时间"`
}
