// Package v1
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package v1

import (
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type SysOssConfigListReq struct {
	g.Meta `path:"/oss-config/list" method:"get" tags:"系统-存储管理" summary:"获取存储配置列表" x-check-permission:"cpm:system:ossConfig:list"`
	*model.SysOssConfigListParam
}
type SysOssConfigListRes struct {
	response.PageResult
	Items []*model.SysOssConfigListModel `json:"items"`
}

type SysOssConfigViewReq struct {
	g.Meta `path:"/oss-config/view" method:"get" tags:"系统-存储管理" summary:"获取存储配置详情" x-check-permission:"cpm:system:ossConfig:list"`
	*model.SysOssConfigViewParam
}
type SysOssConfigViewRes struct {
	*model.SysOssConfigViewModel
}

type SysOssConfigAddReq struct {
	g.Meta `path:"/oss-config/add" method:"post" tags:"系统-存储管理" summary:"新增存储配置" x-check-permission:"cpm:system:ossConfig:add"`
	*model.SysOssConfigAddParam
}
type SysOssConfigAddRes struct {
	*model.SysOssConfigAddModel
}

type SysOssConfigEditReq struct {
	g.Meta `path:"/oss-config/edit" method:"post" tags:"系统-存储管理" summary:"编辑存储配置" x-check-permission:"cpm:system:ossConfig:edit"`
	*model.SysOssConfigEditParam
}
type SysOssConfigEditRes struct {
	*model.SysOssConfigEditModel
}

type SysOssConfigDeleteReq struct {
	g.Meta `path:"/oss-config/delete" method:"post" tags:"系统-存储管理" summary:"删除存储配置" x-check-permission:"cpm:system:ossConfig:remove"`
	*model.SysOssConfigDeleteParam
}
type SysOssConfigDeleteRes struct {
	*model.SysOssConfigDeleteModel
}
