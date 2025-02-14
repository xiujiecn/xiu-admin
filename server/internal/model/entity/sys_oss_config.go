// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOssConfig is the golang structure for table sys_oss_config.
type SysOssConfig struct {
	OssConfigId  int64       `json:"ossConfigId"  orm:"oss_config_id" description:"主键"`
	TenantId     string      `json:"tenantId"     orm:"tenant_id"     description:"租户编号"`
	ConfigKey    string      `json:"configKey"    orm:"config_key"    description:"配置key"`
	AccessKey    string      `json:"accessKey"    orm:"access_key"    description:"accessKey"`
	SecretKey    string      `json:"secretKey"    orm:"secret_key"    description:"秘钥"`
	BucketName   string      `json:"bucketName"   orm:"bucket_name"   description:"桶名称"`
	Prefix       string      `json:"prefix"       orm:"prefix"        description:"前缀"`
	Endpoint     string      `json:"endpoint"     orm:"endpoint"      description:"访问站点"`
	Domain       string      `json:"domain"       orm:"domain"        description:"自定义域名"`
	IsHttps      string      `json:"isHttps"      orm:"is_https"      description:"是否https（Y=是,N=否）"`
	Region       string      `json:"region"       orm:"region"        description:"域"`
	AccessPolicy string      `json:"accessPolicy" orm:"access_policy" description:"桶权限类型(0=private 1=public 2=custom)"`
	Status       string      `json:"status"       orm:"status"        description:"是否默认（0=是,1=否）"`
	Ext1         string      `json:"ext1"         orm:"ext1"          description:"扩展字段"`
	CreatedDept  int64       `json:"createdDept"  orm:"created_dept"  description:"创建部门"`
	CreatedBy    int64       `json:"createdBy"    orm:"created_by"    description:"创建者"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	UpdatedBy    int64       `json:"updatedBy"    orm:"updated_by"    description:"更新者"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`
	Remark       string      `json:"remark"       orm:"remark"        description:"备注"`
}
