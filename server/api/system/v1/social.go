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
	g.Meta `path:"/social/list" method:"get" tags:"系统-第三方授权管理" summary:"社交列表" x-check-permission:"cpm:system:social:list"`
	model.SysSocialListParam
	request.PageInfo
}

type SysSocialListRes struct {
	Items []*model.SysSocialListModel `json:"items"`
	response.PageResult
}
