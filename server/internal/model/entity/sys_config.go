// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure for table sys_config.
type SysConfig struct {
	ConfigId    int64       `json:"configId"    orm:"config_id"    description:"参数主键"`
	TenantId    string      `json:"tenantId"    orm:"tenant_id"    description:"租户编号"`
	ConfigName  string      `json:"configName"  orm:"config_name"  description:"参数名称"`
	ConfigKey   string      `json:"configKey"   orm:"config_key"   description:"参数键名"`
	ConfigValue string      `json:"configValue" orm:"config_value" description:"参数键值"`
	ConfigType  string      `json:"configType"  orm:"config_type"  description:"系统内置（Y是 N否）"`
	CreatedDept int64       `json:"createdDept" orm:"created_dept" description:"创建部门"`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:"创建者"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"   description:"更新者"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
}
