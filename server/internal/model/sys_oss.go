// package model
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package model

import (
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/model/request"

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
	File             *ghttp.UploadFile `json:"file" description:"文件"`
	FileType         string            `json:"fileType" description:"文件类型"`
	IsDevice         bool              `json:"isDevice" description:"是否是设备文件"`
	NewFileType      int               `json:"newFileType" description:"新文件类型 1:用户文件 2:用户图片 3:设备文件 4:iot配置图片 5:iot固件文件 6:项目图片 7:项目文件"`
	SubDirName       string            `json:"subDirName" description:"子目录名称或路径"`
	SaveOriginalName int               `json:"saveOriginalName" description:"是否保存原始文件名 0:否 1:是 默认否"`
	NotAddDate       int               `json:"notAddDate" description:"存储路径不添加日期 0:添加 1:不添加 默认添加"`
}

type SysOssUploadModel struct {
	entity.SysOss
}

type SysOssMoveFileParam struct {
	FilePath         string `json:"filePath" description:"文件路径"`
	NewFileType      int    `json:"newFileType" description:"新文件类型 1:用户文件 2:用户图片 3:设备文件 4:iot配置图片 5:iot固件文件 6:项目图片 7:项目文件"`
	SubDirName       string `json:"subDirName" description:"子目录名称或路径"`
	SaveOriginalName int    `json:"saveOriginalName" description:"是否保存原始文件名 0:否 1:是 默认否"`
	NotAddDate       int    `json:"notAddDate" description:"存储路径不添加日期 0:添加 1:不添加 默认添加"`
}

type SysOssMoveFileModel struct {
	*entity.SysOss
}

type SysOssSaveContentParam struct {
	Content          []byte `json:"content" description:"内容"`
	FileName         string `json:"fileName" description:"文件名"`
	NewFileType      int    `json:"newFileType" description:"新文件类型 1:用户文件 2:用户图片 3:设备文件 4:iot配置图片 5:iot固件文件 6:项目图片 7:项目文件"`
	SubDirName       string `json:"subDirName" description:"子目录名称或路径"`
	SaveOriginalName int    `json:"saveOriginalName" description:"是否保存原始文件名 0:否 1:是 默认否"`
	NotAddDate       int    `json:"notAddDate" description:"存储路径不添加日期 0:添加 1:不添加 默认添加"`
}

type SysOssSaveContentModel struct {
	*entity.SysOss
}

type UploadResponse struct {
	Size         int64  `json:"size"`          // 文件大小
	Path         string `json:"path"`          // 文件路径
	FullPath     string `json:"full_path"`     // 文件全路径
	Name         string `json:"name"`          // 文件名
	OriginalName string `json:"original_name"` // 文件原名
	Type         string `json:"type"`          // 文件类型 http Content-type
	Ext          string `json:"ext"`           // 文件扩展名 不包含.
	Md5          string `json:"md5"`           // 文件MD5
	Crc16        uint16 `json:"crc16"`         // 文件CRC-16/MODBUS
	Sum16        uint16 `json:"sum16"`         // 文件SUM16
}
