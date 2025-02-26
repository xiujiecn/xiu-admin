package model

import (
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

type SysClientViewModel struct {
	entity.SysClient
}

type SysClientListModel struct {
	Id            int64       `json:"id"            orm:"id"             description:"id"`
	ClientId      string      `json:"clientId"      orm:"client_id"      description:"客户端id"`
	ClientKey     string      `json:"clientKey"     orm:"client_key"     description:"客户端key"`
	ClientSecret  string      `json:"clientSecret"  orm:"client_secret"  description:"客户端秘钥"`
	GrantType     string      `json:"grantType"     orm:"grant_type"     description:"授权类型"`
	DeviceType    string      `json:"deviceType"    orm:"device_type"    description:"设备类型"`
	ActiveTimeout int         `json:"activeTimeout" orm:"active_timeout" description:"token活跃超时时间"`
	Timeout       int         `json:"timeout"       orm:"timeout"        description:"token固定超时"`
	Status        string      `json:"status"        orm:"status"         description:"状态（0正常 1停用）"`
	CreatedDept   int64       `json:"createdDept"   orm:"created_dept"   description:"创建部门"`
	CreatedBy     int64       `json:"createdBy"     orm:"created_by"     description:"创建者"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
}

type SysClientListParam struct {
	ClientId     string `json:"clientId"      orm:"client_id"      description:"客户端id"`
	ClientKey    string `json:"clientKey"     orm:"client_key"     description:"客户端key"`
	ClientSecret string `json:"clientSecret"  orm:"client_secret"  description:"客户端秘钥"`
	Status       string `json:"status"        orm:"status"         description:"状态（0正常 1停用）"`
}
