package v1

import (
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type ListLogininforReq struct {
	g.Meta `path:"/logininfor/list" method:"get" tags:"系统-日志管理" summary:"登录信息列表" x-check-permission:"cpm:system:logininfor:list"`
	model.SysLogininforListParam
}

type ListLogininforRes struct {
	Items []*model.SysLogininforListModel `json:"items"`
	response.PageResult
}

type DeleteLogininforReq struct {
	g.Meta `path:"/logininfor/delete" method:"post" tags:"系统-日志管理" summary:"删除登录信息" x-check-permission:"cpm:system:logininfor:remove"`
	model.SysLogininforDeleteParam
}

type DeleteLogininforRes struct {
	model.SysLogininforDeleteModel
}
