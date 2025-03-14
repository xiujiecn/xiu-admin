package model

import (
	"xiujieadmin/internal/model/entity"
	"xiujieadmin/internal/model/request"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

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
	request.PageInfo
	FileName     string   `json:"fileName"     orm:"file_name"     description:"文件名"`
	OriginalName string   `json:"originalName" orm:"original_name" description:"原名"`
	FileSuffix   string   `json:"fileSuffix"   orm:"file_suffix"   description:"文件后缀名"`
	Service      string   `json:"service"      orm:"service"       description:"服务商"`
	CreatedAt    []string `json:"createdAt"    orm:"created_at"    description:"创建时间"`
}

type SysOssViewParam struct {
	OssId int64 `json:"ossId" orm:"oss_id" description:"对象存储主键"`
}

type SysOssViewModel struct {
	entity.SysOss
}

type SysOssDownloadParam struct {
	OssId int64 `json:"ossId" orm:"oss_id" description:"对象存储主键"`
}

type SysOssDownloadModel struct {
	*entity.SysOss
}

type SysOssDeleteParam struct {
	OssIds []int64 `json:"ossIds" description:"对象存储主键"`
}

type SysOssDeleteModel struct {
	entity.SysOss
}

type SysOssUploadParam struct {
	File     *ghttp.UploadFile `json:"file" description:"文件"`
	FileType string            `json:"fileType" description:"文件类型"`
}
type SysOssUploadModel struct {
	entity.SysOss
}

type UploadResponse struct {
	Size         int64  `json:"size"`          // 文件大小
	Path         string `json:"path"`          // 文件路径
	FullPath     string `json:"full_path"`     // 文件全路径
	Name         string `json:"name"`          // 文件名
	OriginalName string `json:"original_name"` // 文件原名
	Type         string `json:"type"`          // 文件类型 http Content-type
	Ext          string `json:"ext"`           // 文件扩展名 不包含.
}
