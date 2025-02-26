package v1

import (
	"server/internal/model"
	"server/internal/model/request"
	"server/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type SysTenantListReq struct {
	g.Meta `path:"/tenant/list" method:"get" tags:"租户" summary:"租户列表"`
	request.PageInfo
	model.SysTenantListParam
}

type SysTenantListRes struct {
	response.PageResult
	Items []*model.SysTenantListModel `json:"items"`
}

type SysTenantPackageListReq struct {
	g.Meta `path:"/tenant/package/list" method:"get" tags:"租户" summary:"租户套餐列表"`
	request.PageInfo
	model.SysTenantPackageListParam
}

type SysTenantPackageListRes struct {
	response.PageResult
	Items []*model.SysTenantPackageListModel `json:"items"`
}
