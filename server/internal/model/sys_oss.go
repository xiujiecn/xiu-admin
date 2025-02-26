package model

import (
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

type SysOssViewModel struct {
	entity.SysOss
}

type SysOssListModel struct {
	OssId        int64       `json:"ossId"        orm:"oss_id"        description:"对象存储主键"`
	TenantId     string      `json:"tenantId"     orm:"tenant_id"     description:"租户编号"`
	FileName     string      `json:"fileName"     orm:"file_name"     description:"文件名"`
	OriginalName string      `json:"originalName" orm:"original_name" description:"原名"`
	FileSuffix   string      `json:"fileSuffix"   orm:"file_suffix"   description:"文件后缀名"`
	Url          string      `json:"url"          orm:"url"           description:"URL地址"`
	CreatedDept  int64       `json:"createdDept"  orm:"created_dept"  description:"创建部门"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	CreatedBy    int64       `json:"createdBy"    orm:"created_by"    description:"创建者"`
	Service      string      `json:"service"      orm:"service"       description:"服务商"`
}

type SysOssListParam struct {
	FileName     string   `json:"fileName"     orm:"file_name"     description:"文件名"`
	OriginalName string   `json:"originalName" orm:"original_name" description:"原名"`
	FileSuffix   string   `json:"fileSuffix"   orm:"file_suffix"   description:"文件后缀名"`
	Service      string   `json:"service"      orm:"service"       description:"服务商"`
	CreatedAt    []string `json:"createdAt"    orm:"created_at"    description:"创建时间"`
}
