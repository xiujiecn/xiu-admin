// package model
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package model

import (
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/model/request"

	"github.com/gogf/gf/v2/os/gtime"
)

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
	request.PageInfo
	ClientId     string `json:"clientId"      orm:"client_id"      description:"客户端id"`
	ClientKey    string `json:"clientKey"     orm:"client_key"     description:"客户端key"`
	ClientSecret string `json:"clientSecret"  orm:"client_secret"  description:"客户端秘钥"`
	Status       string `json:"status"        orm:"status"         description:"状态（0正常 1停用）"`
}

type SysClientViewParam struct {
	Id       int64  `json:"id"            orm:"id"             description:"id"`
	ClientId string `json:"clientId"      orm:"client_id"      description:"客户端id"`
}

type SysClientViewModel struct {
	entity.SysClient
}

type SysClientAddParam struct {
	ClientId      string `json:"clientId"      orm:"client_id"      description:"客户端id"`
	ClientKey     string `json:"clientKey"     orm:"client_key"     description:"客户端key"`
	ClientSecret  string `json:"clientSecret"  orm:"client_secret"  description:"客户端秘钥"`
	GrantType     string `json:"grantType"     orm:"grant_type"     description:"授权类型"`
	DeviceType    string `json:"deviceType"    orm:"device_type"    description:"设备类型"`
	ActiveTimeout int    `json:"activeTimeout" orm:"active_timeout" description:"token活跃超时时间"`
	Timeout       int    `json:"timeout"       orm:"timeout"        description:"token固定超时"`
	Status        string `json:"status"        orm:"status"         description:"状态（0正常 1停用）"`
}

type SysClientAddModel struct {
	Id int64 `json:"id"            orm:"id"             description:"id"`
}

type SysClientEditParam struct {
	Id            int64   `json:"id"            description:"id"`
	ClientId      *string `json:"clientId,omitempty"      description:"客户端id"`
	ClientKey     *string `json:"clientKey,omitempty"     description:"客户端key"`
	ClientSecret  *string `json:"clientSecret,omitempty"  description:"客户端秘钥"`
	GrantType     *string `json:"grantType,omitempty"     description:"授权类型"`
	DeviceType    *string `json:"deviceType,omitempty"    description:"设备类型"`
	ActiveTimeout *int    `json:"activeTimeout,omitempty" description:"token活跃超时时间"`
	Timeout       *int    `json:"timeout,omitempty"       description:"token固定超时"`
	Status        *string `json:"status,omitempty"        description:"状态（0正常 1停用）"`
}

type SysClientEditModel struct {
	Id int64 `json:"id"            orm:"id"             description:"id"`
}

type SysClientDeleteParam struct {
	Ids []int64 `json:"ids"            description:"id"`
}

type SysClientDeleteModel struct {
	Ids []int64 `json:"ids"           description:"id"`
}

type SysClientStatusParam struct {
	Id     int64  `json:"id"            description:"id"`
	Status string `json:"status"        description:"状态（0正常 1停用）"`
}

type SysClientStatusModel struct {
	Id int64 `json:"id"            orm:"id"             description:"id"`
}
