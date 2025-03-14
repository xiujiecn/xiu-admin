package v1

import (
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type SysOssConfigListReq struct {
	g.Meta `path:"/oss-config/list" method:"get" tags:"系统配置" summary:"获取系统配置列表"`
	*model.SysOssConfigListParam
}
type SysOssConfigListRes struct {
	response.PageResult
	Items []*model.SysOssConfigListModel `json:"items"`
}

type SysOssConfigViewReq struct {
	g.Meta `path:"/oss-config/view" method:"get" tags:"系统配置" summary:"获取系统配置详情"`
	*model.SysOssConfigViewParam
}
type SysOssConfigViewRes struct {
	*model.SysOssConfigViewModel
}

type SysOssConfigAddReq struct {
	g.Meta `path:"/oss-config/add" method:"post" tags:"系统配置" summary:"新增系统配置"`
	*model.SysOssConfigAddParam
}
type SysOssConfigAddRes struct {
	*model.SysOssConfigAddModel
}

type SysOssConfigEditReq struct {
	g.Meta `path:"/oss-config/edit" method:"post" tags:"系统配置" summary:"编辑系统配置"`
	*model.SysOssConfigEditParam
}
type SysOssConfigEditRes struct {
	*model.SysOssConfigEditModel
}

type SysOssConfigDeleteReq struct {
	g.Meta `path:"/oss-config/delete" method:"post" tags:"系统配置" summary:"删除系统配置"`
	*model.SysOssConfigDeleteParam
}
type SysOssConfigDeleteRes struct {
	*model.SysOssConfigDeleteModel
}
