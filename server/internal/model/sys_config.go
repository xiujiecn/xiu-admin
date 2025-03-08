package model

import (
	"xiujieadmin/internal/model/request"

	"github.com/gogf/gf/v2/os/gtime"
)

type SysConfigListModel struct {
	ConfigId    int64       `json:"configId"    orm:"config_id"    description:"参数主键"`
	TenantId    string      `json:"tenantId"    orm:"tenant_id"    description:"租户编号"`
	ConfigName  string      `json:"configName"  orm:"config_name"  description:"参数名称"`
	ConfigKey   string      `json:"configKey"   orm:"config_key"   description:"参数键名"`
	ConfigValue string      `json:"configValue" orm:"config_value" description:"参数键值"`
	ConfigType  string      `json:"configType"  orm:"config_type"  description:"系统内置（Y是 N否）"`
	CreatedDept int64       `json:"createdDept" orm:"created_dept" description:"创建部门"`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:"创建者"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
}

type SysConfigListParam struct {
	request.PageInfo
	ConfigName  string   `json:"configName" form:"configName" comment:"参数名称"`
	ConfigKey   string   `json:"configKey"  form:"configKey"  comment:"参数键名"`
	ConfigType  string   `json:"configType" form:"configType" comment:"系统内置"`
	ConfigValue string   `json:"configValue" orm:"config_value" description:"参数键值"`
	CreatedAt   []string `json:"createdAt" form:"createdAt" comment:"创建时间"`
}

type SysConfigAddParam struct {
	ConfigName  string `json:"configName"  description:"参数名称"`
	ConfigKey   string `json:"configKey"   description:"参数键名"`
	ConfigValue string `json:"configValue" description:"参数键值"`
	ConfigType  string `json:"configType"  description:"系统内置（Y是 N否）"`
	Remark      string `json:"remark"      description:"备注"`
}

type SysConfigAddModel struct {
	ConfigId int64 `json:"configId" form:"configId" comment:"参数主键"`
}

type SysConfigEditParam struct {
	ConfigId    int64   `json:"configId"    description:"参数主键"`
	ConfigName  *string `json:"configName"  description:"参数名称"`
	ConfigKey   *string `json:"configKey"   description:"参数键名"`
	ConfigValue *string `json:"configValue" description:"参数键值"`
	ConfigType  *string `json:"configType"  description:"系统内置（Y是 N否）"`
	Remark      *string `json:"remark"      description:"备注"`
}
type SysConfigEditModel struct {
	ConfigId int64 `json:"configId" form:"configId" comment:"参数主键"`
}

type SysConfigDeleteParam struct {
	ConfigIds []int64 `json:"configIds" form:"configIds" comment:"参数主键"`
}

type SysConfigDeleteModel struct {
	ConfigIds []int64 `json:"configIds" form:"configIds" comment:"参数主键"`
}

type SysConfigViewParam struct {
	ConfigId int64 `json:"configId" form:"configId" comment:"参数主键"`
}

type SysConfigViewModel struct {
	ConfigId    int64       `json:"configId"    orm:"config_id"    description:"参数主键"`
	TenantId    string      `json:"tenantId"    orm:"tenant_id"    description:"租户编号"`
	ConfigName  string      `json:"configName"  orm:"config_name"  description:"参数名称"`
	ConfigKey   string      `json:"configKey"   orm:"config_key"   description:"参数键名"`
	ConfigValue string      `json:"configValue" orm:"config_value" description:"参数键值"`
	ConfigType  string      `json:"configType"  orm:"config_type"  description:"系统内置（Y是 N否）"`
	CreatedDept int64       `json:"createdDept" orm:"created_dept" description:"创建部门"`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:"创建者"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
}
