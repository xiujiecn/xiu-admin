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

type SysSocialListReq struct {
	g.Meta `path:"/social/list" method:"get" tags:"系统-第三方授权管理" summary:"社交列表" x-check-permission:"cpc:current:user"`
	model.SysSocialListParam
	request.PageInfo
}

type SysSocialListRes struct {
	Items []*model.SysSocialListModel `json:"items"`
	response.PageResult
}

// 删除绑定关系
type SysSocialDeleteReq struct {
	g.Meta `path:"/social/delete" method:"post" tags:"系统-第三方授权管理" summary:"删除社交" x-check-permission:"cpc:current:user"`
	Id     int64 `json:"id" dc:"主键"`
}

type SysSocialDeleteRes struct {
}
