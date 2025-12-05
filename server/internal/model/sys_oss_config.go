// package model
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package model

import (
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/model/request"

	"github.com/gogf/gf/v2/os/gtime"
)

type SysOssConfigViewParam struct {
	OssConfigId int64  `json:"ossConfigId" description:"主键"`
	ConfigKey   string `json:"configKey"    orm:"config_key"    description:"配置key"`
}

type SysOssConfigViewModel struct {
	entity.SysOssConfig
}

type SysOssConfigListModel struct {
	OssConfigId  int64       `json:"ossConfigId"  orm:"oss_config_id" description:"主键"`
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
	Remark       string      `json:"remark"       orm:"remark"        description:"备注"`
}

type SysOssConfigListParam struct {
	request.PageInfo
	ConfigKey  string `json:"configKey"     description:"配置key"`
	BucketName string `json:"bucketName"    description:"桶名称"`
	Status     string `json:"status"        description:"是否默认（0=是,1=否）"`
}

type SysOssConfigAddParam struct {
	ConfigKey    string `json:"configKey"    orm:"config_key"    description:"配置key"`
	AccessKey    string `json:"accessKey"    orm:"access_key"    description:"accessKey"`
	SecretKey    string `json:"secretKey"    orm:"secret_key"    description:"秘钥"`
	BucketName   string `json:"bucketName"   orm:"bucket_name"   description:"桶名称"`
	Prefix       string `json:"prefix"       orm:"prefix"        description:"前缀"`
	Endpoint     string `json:"endpoint"     orm:"endpoint"      description:"访问站点"`
	Domain       string `json:"domain"       orm:"domain"        description:"自定义域名"`
	IsHttps      string `json:"isHttps"      orm:"is_https"      description:"是否https（Y=是,N=否）"`
	Region       string `json:"region"       orm:"region"        description:"域"`
	AccessPolicy string `json:"accessPolicy" orm:"access_policy" description:"桶权限类型(0=private 1=public 2=custom)"`
	Status       string `json:"status"       orm:"status"        description:"是否默认（0=是,1=否）"`
	Ext1         string `json:"ext1"         orm:"ext1"          description:"扩展字段"`
	Remark       string `json:"remark"       orm:"remark"        description:"备注"`
}

type SysOssConfigAddModel struct {
	OssConfigId int64 `json:"ossConfigId" description:"主键"`
}

type SysOssConfigEditParam struct {
	OssConfigId  int64   `json:"ossConfigId"   description:"主键"`
	ConfigKey    *string `json:"configKey"     description:"配置key"`
	AccessKey    *string `json:"accessKey"     description:"accessKey"`
	SecretKey    *string `json:"secretKey"     description:"秘钥"`
	BucketName   *string `json:"bucketName"    description:"桶名称"`
	Prefix       *string `json:"prefix"        description:"前缀"`
	Endpoint     *string `json:"endpoint"      description:"访问站点"`
	Domain       *string `json:"domain"        description:"自定义域名"`
	IsHttps      *string `json:"isHttps,omitempty"       description:"是否https（Y=是,N=否）"`
	Region       *string `json:"region"        description:"域"`
	AccessPolicy *string `json:"accessPolicy,omitempty"  description:"桶权限类型(0=private 1=public 2=custom)"`
	Status       *string `json:"status,omitempty"        description:"是否默认（0=是,1=否）"`
	Ext1         *string `json:"ext1"          description:"扩展字段"`
	Remark       *string `json:"remark"        description:"备注"`
}

type SysOssConfigEditModel struct {
	OssConfigId int64 `json:"ossConfigId" description:"主键"`
}

type SysOssConfigDeleteParam struct {
	OssConfigIds []int64 `json:"ossConfigIds" description:"主键"`
}

type SysOssConfigDeleteModel struct {
	OssConfigIds []int64 `json:"ossConfigIds" description:"主键"`
}
