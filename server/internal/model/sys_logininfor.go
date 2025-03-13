package model

import (
	"xiujieadmin/internal/model/entity"
	"xiujieadmin/internal/model/request"

	"github.com/gogf/gf/v2/os/gtime"
)

type SysLogininforListModel struct {
	InfoId        int64       `json:"infoId"        orm:"info_id"        description:"访问ID"`
	TenantId      string      `json:"tenantId"      orm:"tenant_id"      description:"租户编号"`
	UserName      string      `json:"userName"      orm:"user_name"      description:"用户账号"`
	ClientKey     string      `json:"clientKey"     orm:"client_key"     description:"客户端"`
	DeviceType    string      `json:"deviceType"    orm:"device_type"    description:"设备类型"`
	Ipaddr        string      `json:"ipaddr"        orm:"ipaddr"         description:"登录IP地址"`
	LoginLocation string      `json:"loginLocation" orm:"login_location" description:"登录地点"`
	Browser       string      `json:"browser"       orm:"browser"        description:"浏览器类型"`
	Os            string      `json:"os"            orm:"os"             description:"操作系统"`
	Status        string      `json:"status"        orm:"status"         description:"登录状态（0成功 1失败）"`
	Msg           string      `json:"msg"           orm:"msg"            description:"提示消息"`
	LoginTime     *gtime.Time `json:"loginTime"     orm:"login_time"     description:"访问时间"`
}

type SysLogininforAddModel struct {
	TenantId      string      `json:"tenantId"      orm:"tenant_id"      description:"租户编号"`
	UserName      string      `json:"userName"      orm:"user_name"      description:"用户账号"`
	ClientKey     string      `json:"clientKey"     orm:"client_key"     description:"客户端"`
	DeviceType    string      `json:"deviceType"    orm:"device_type"    description:"设备类型"`
	Ipaddr        string      `json:"ipaddr"        orm:"ipaddr"         description:"登录IP地址"`
	LoginLocation string      `json:"loginLocation" orm:"login_location" description:"登录地点"`
	Browser       string      `json:"browser"       orm:"browser"        description:"浏览器类型"`
	Os            string      `json:"os"            orm:"os"             description:"操作系统"`
	Status        string      `json:"status"        orm:"status"         description:"登录状态（0成功 1失败）"`
	Msg           string      `json:"msg"           orm:"msg"            description:"提示消息"`
	LoginTime     *gtime.Time `json:"loginTime"     orm:"login_time"     description:"访问时间"`
}

type SysLogininforViewModel struct {
	entity.SysLogininfor
}

type SysLogininforListParam struct {
	request.PageInfo
	Ipaddr    string   `json:"ipaddr"        description:"登录IP地址"`
	UserName  string   `json:"userName"      description:"用户账号"`
	Status    string   `json:"status"        description:"登录状态（0成功 1失败）"`
	LoginTime []string `json:"loginTime"     description:"访问时间"`
}

type SysLogininforDeleteParam struct {
	InfoIds []int64 `json:"infoIds"         description:"访问ID"`
}
type SysLogininforDeleteModel struct {
	InfoIds []int64 `json:"infoIds"         description:"访问ID"`
}
