package v1

import (
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type SysTenantListReq struct {
	g.Meta `path:"/tenant/list" method:"get" tags:"系统-租户管理" summary:"租户列表" x-check-permission:"cpm:system:tenant:list"`
	request.PageInfo
	*model.SysTenantListParam
}

type SysTenantListRes struct {
	response.PageResult
	Items []*model.SysTenantListModel `json:"items"`
}

type SysTenantAddReq struct {
	g.Meta `path:"/tenant/add" method:"post" tags:"系统-租户管理" summary:"租户添加" x-check-permission:"cpm:system:tenant:add"`
	*model.SysTenantAddParam
}

type SysTenantAddRes struct {
	*model.SysTenantAddModel
}

type SysTenantEditReq struct {
	g.Meta `path:"/tenant/edit" method:"post" tags:"系统-租户管理" summary:"租户编辑" x-check-permission:"cpm:system:tenant:edit"`
	*model.SysTenantEditParam
}

type SysTenantEditRes struct {
	*model.SysTenantEditModel
}

type SysTenantDeleteReq struct {
	g.Meta `path:"/tenant/delete" method:"post" tags:"系统-租户管理" summary:"租户删除" x-check-permission:"cpm:system:tenant:remove"`
	*model.SysTenantDeleteParam
}

type SysTenantDeleteRes struct {
	*model.SysTenantDeleteModel
}

type SysTenantStatusReq struct {
	g.Meta `path:"/tenant/status" method:"post" tags:"系统-租户管理" summary:"租户状态" x-check-permission:"cpm:system:tenant:edit"`
	*model.SysTenantStatusParam
}

type SysTenantStatusRes struct {
	*model.SysTenantStatusModel
}

type SysTenantViewReq struct {
	g.Meta `path:"/tenant/view" method:"get" tags:"系统-租户管理" summary:"租户详情" x-check-permission:"cpm:system:tenant:query"`
	*model.SysTenantViewParam
}

type SysTenantViewRes struct {
	*model.SysTenantViewModel
}

type SysTenantPackageListReq struct {
	g.Meta `path:"/tenant/package/list" method:"get" tags:"系统-租户管理" summary:"租户套餐列表" x-check-permission:"cpm:system:tenant:package:list"`
	request.PageInfo
	model.SysTenantPackageListParam
}

type SysTenantPackageListRes struct {
	response.PageResult
	Items []*model.SysTenantPackageListModel `json:"items"`
}

type SysTenantPackageViewReq struct {
	g.Meta `path:"/tenant/package/view" method:"get" tags:"系统-租户管理" summary:"租户套餐详情" x-check-permission:"cpm:system:tenantPackage:query"`
	*model.SysTenantPackageViewParam
}

type SysTenantPackageViewRes struct {
	*model.SysTenantPackageViewModel
}

type SysTenantPackageStatusReq struct {
	g.Meta `path:"/tenant/package/status" method:"post" tags:"系统-租户管理" summary:"租户套餐状态" x-check-permission:"cpm:system:tenantPackage:edit"`
	*model.SysTenantPackageStatusParam
}

type SysTenantPackageStatusRes struct {
	*model.SysTenantPackageStatusModel
}

type SysTenantPackageAddReq struct {
	g.Meta `path:"/tenant/package/add" method:"post" tags:"系统-租户管理" summary:"租户套餐添加" x-check-permission:"cpm:system:tenantPackage:add"`
	*model.SysTenantPackageAddParam
}

type SysTenantPackageAddRes struct {
	*model.SysTenantPackageAddModel
}

type SysTenantPackageEditReq struct {
	g.Meta `path:"/tenant/package/edit" method:"post" tags:"系统-租户管理" summary:"租户套餐编辑" x-check-permission:"cpm:system:tenantPackage:edit"`
	*model.SysTenantPackageEditParam
}

type SysTenantPackageEditRes struct {
	*model.SysTenantPackageEditModel
}

type SysTenantPackageDeleteReq struct {
	g.Meta `path:"/tenant/package/delete" method:"post" tags:"系统-租户管理" summary:"租户套餐删除" x-check-permission:"cpm:system:tenantPackage:remove"`
	*model.SysTenantPackageDeleteParam
}

type SysTenantPackageDeleteRes struct {
	*model.SysTenantPackageDeleteModel
}
