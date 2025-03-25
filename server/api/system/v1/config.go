package v1

import (
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type GetConfigListReq struct {
	g.Meta `path:"/config/list" method:"get" tags:"系统-配置管理" summary:"获取系统配置列表" x-check-permission:"cpm:system:config:list"`
	model.SysConfigListParam
}

type GetConfigListRes struct {
	response.PageResult
	Data []*model.SysConfigListModel `json:"items" dc:"系统配置列表"`
}

type AddConfigReq struct {
	g.Meta `path:"/config/add" method:"post" tags:"系统-配置管理" summary:"添加系统配置" x-check-permission:"cpm:system:config:add"`
	model.SysConfigAddParam
}

type AddConfigRes struct {
	model.SysConfigAddModel
}

type EditConfigReq struct {
	g.Meta `path:"/config/edit" method:"post" tags:"系统-配置管理" summary:"编辑系统配置" x-check-permission:"cpm:system:config:edit"`
	model.SysConfigEditParam
}

type EditConfigRes struct {
	model.SysConfigEditModel
}

type DeleteConfigReq struct {
	g.Meta `path:"/config/delete" method:"post" tags:"系统-配置管理" summary:"删除系统配置" x-check-permission:"cpm:system:config:remove"`
	model.SysConfigDeleteParam
}

type DeleteConfigRes struct {
	model.SysConfigDeleteModel
}

type GetConfigByIdReq struct {
	g.Meta `path:"/config/view" method:"get" tags:"系统-配置管理" summary:"获取系统配置详情" x-check-permission:"cpm:system:config:query"`
	model.SysConfigViewParam
}

type GetConfigByIdRes struct {
	model.SysConfigViewModel
}
