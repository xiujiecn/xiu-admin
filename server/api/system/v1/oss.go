// Package v1
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package v1

import (
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/request"
	"xiuadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type SysOssListReq struct {
	g.Meta `path:"/oss/list" method:"get" tags:"系统-存储管理" summary:"获取OSS列表" x-check-permission:"cpm:system:oss:list"`
	model.SysOssListParam
	request.PageInfo
}

type SysOssListRes struct {
	Items []*model.SysOssListModel `json:"items"`
	response.PageResult
}

type SysOssViewReq struct {
	g.Meta `path:"/oss/view" method:"get" tags:"系统-存储管理" summary:"查看文件" x-check-permission:"cpm:system:oss:query"`
	*model.SysOssViewParam
}

type SysOssViewRes struct {
	*model.SysOssViewModel
}

type SysOssDeleteReq struct {
	g.Meta `path:"/oss/delete" method:"post" tags:"系统-存储管理" summary:"删除文件" x-check-permission:"cpm:system:oss:remove"`
	*model.SysOssDeleteParam
}

type SysOssDeleteRes struct {
	*model.SysOssDeleteModel
}
