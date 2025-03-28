// package model
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package model

import (
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/model/request"

	"github.com/gogf/gf/v2/os/gtime"
)

type SysTenantPackageListModel struct {
	PackageId         int64       `json:"packageId"         orm:"package_id"          description:"租户套餐id"`
	PackageName       string      `json:"packageName"       orm:"package_name"        description:"套餐名称"`
	MenuIds           string      `json:"menuIds"           orm:"menu_ids"            description:"关联菜单id"`
	Remark            string      `json:"remark"            orm:"remark"              description:"备注"`
	MenuCheckStrictly int         `json:"menuCheckStrictly" orm:"menu_check_strictly" description:"菜单树选择项是否关联显示"`
	Status            string      `json:"status"            orm:"status"              description:"状态（0正常 1停用）"`
	CreatedDept       int64       `json:"createdDept"       orm:"created_dept"        description:"创建部门"`
	CreatedBy         int64       `json:"createdBy"         orm:"created_by"          description:"创建者"`
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"          description:"创建时间"`
	UpdatedBy         int64       `json:"updatedBy"         orm:"updated_by"          description:"更新者"`
	UpdatedAt         *gtime.Time `json:"updatedAt"         orm:"updated_at"          description:"更新时间"`
}

type SysTenantPackageListParam struct {
	request.PageInfo
	PackageName string `json:"packageName"       orm:"package_name"        description:"套餐名称"`
}
type SysTenantPackageStatusParam struct {
	PackageId int64  `json:"packageId"              description:"租户套餐id"`
	Status    string `json:"status"  description:"状态（0正常 1停用）"`
}
type SysTenantPackageStatusModel struct {
	PackageId int64 `json:"packageId"              description:"租户套餐id"`
}
type SysTenantPackageViewParam struct {
	PackageId int64 `json:"packageId"              description:"租户套餐id"`
}
type SysTenantPackageViewModel struct {
	entity.SysTenantPackage
}
type SysTenantPackageAddParam struct {
	PackageName       string `json:"packageName"       description:"套餐名称"`
	MenuIds           string `json:"menuIds"           description:"关联菜单id"`
	Remark            string `json:"remark"            description:"备注"`
	MenuCheckStrictly int    `json:"menuCheckStrictly" description:"菜单树选择项是否关联显示"`
	Status            string `json:"status"            description:"状态（0正常 1停用）"`
}
type SysTenantPackageAddModel struct {
	PackageId int64 `json:"packageId"              description:"租户套餐id"`
}
type SysTenantPackageEditParam struct {
	PackageId         int64   `json:"packageId"              description:"租户套餐id"`
	PackageName       *string `json:"packageName,omitempty"       description:"套餐名称"`
	MenuIds           *string `json:"menuIds,omitempty"           description:"关联菜单id"`
	Remark            *string `json:"remark,omitempty"            description:"备注"`
	MenuCheckStrictly *int    `json:"menuCheckStrictly,omitempty" description:"菜单树选择项是否关联显示"`
	Status            *string `json:"status,omitempty"            description:"状态（0正常 1停用）"`
}
type SysTenantPackageEditModel struct {
	PackageId int64 `json:"packageId"              description:"租户套餐id"`
}

type SysTenantPackageDeleteParam struct {
	PackageIds []int64 `json:"packageIds"              description:"租户套餐id"`
}
type SysTenantPackageDeleteModel struct {
	PackageIds []int64 `json:"packageIds"              description:"租户套餐id"`
}
