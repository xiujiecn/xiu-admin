// Package v1
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025	 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package v1

import (
	"xiuadmin/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type SysOssUploadReq struct {
	g.Meta `path:"/oss/upload" method:"post" tags:"系统-存储管理" summary:"上传文件" x-check-permission:"cpm:system.oss.upload"`
	*model.SysOssUploadParam
}

type SysOssUploadRes struct {
	*model.SysOssUploadModel
}
type SysOssDownloadReq struct {
	g.Meta `path:"/oss/download" method:"get" tags:"系统-存储管理" summary:"下载文件" x-check-permission:"cpm:system.oss.download"`
	*model.SysOssDownloadParam
}

type SysOssDownloadRes struct {
	*model.SysOssDownloadModel
}
