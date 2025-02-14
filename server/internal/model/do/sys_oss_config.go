// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOssConfig is the golang structure of table sys_oss_config for DAO operations like Where/Data.
type SysOssConfig struct {
	g.Meta       `orm:"table:sys_oss_config, do:true"`
	OssConfigId  interface{} // 主键
	TenantId     interface{} // 租户编号
	ConfigKey    interface{} // 配置key
	AccessKey    interface{} // accessKey
	SecretKey    interface{} // 秘钥
	BucketName   interface{} // 桶名称
	Prefix       interface{} // 前缀
	Endpoint     interface{} // 访问站点
	Domain       interface{} // 自定义域名
	IsHttps      interface{} // 是否https（Y=是,N=否）
	Region       interface{} // 域
	AccessPolicy interface{} // 桶权限类型(0=private 1=public 2=custom)
	Status       interface{} // 是否默认（0=是,1=否）
	Ext1         interface{} // 扩展字段
	CreatedDept  interface{} // 创建部门
	CreatedBy    interface{} // 创建者
	CreatedAt    *gtime.Time // 创建时间
	UpdatedBy    interface{} // 更新者
	UpdatedAt    *gtime.Time // 更新时间
	Remark       interface{} // 备注
}
