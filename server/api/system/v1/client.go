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

type SysClientListReq struct {
	g.Meta `path:"/client/list" method:"get" tags:"系统-客户端管理" summary:"获取客户端列表" x-check-permission:"cpm:system:client:list"`
	model.SysClientListParam
}

type SysClientListRes struct {
	Items []*model.SysClientListModel `json:"items"`
	response.PageResult
}

type SysClientViewReq struct {
	g.Meta `path:"/client/view" method:"get" tags:"系统-客户端管理" summary:"获取客户端详情" x-check-permission:"cpm:system:client:query"`
	*model.SysClientViewParam
}

type SysClientViewRes struct {
	*model.SysClientViewModel
}

type SysClientAddReq struct {
	g.Meta `path:"/client/add" method:"post" tags:"系统-客户端管理" summary:"添加客户端" x-check-permission:"cpm:system:client:add"`
	*model.SysClientAddParam
}

type SysClientAddRes struct {
	*model.SysClientAddModel
}

type SysClientEditReq struct {
	g.Meta `path:"/client/edit" method:"post" tags:"系统-客户端管理" summary:"编辑客户端" x-check-permission:"cpm:system:client:edit"`
	*model.SysClientEditParam
}

type SysClientEditRes struct {
	*model.SysClientEditModel
}

type SysClientDeleteReq struct {
	g.Meta `path:"/client/delete" method:"post" tags:"系统-客户端管理" summary:"删除客户端" x-check-permission:"cpm:system:client:remove"`
	*model.SysClientDeleteParam
}

type SysClientDeleteRes struct {
	*model.SysClientDeleteModel
}

type SysClientStatusReq struct {
	g.Meta `path:"/client/status" method:"post" tags:"系统-客户端管理" summary:"修改客户端状态" x-check-permission:"cpm:system:client:edit"`
	*model.SysClientStatusParam
}

type SysClientStatusRes struct {
	*model.SysClientStatusModel
}
