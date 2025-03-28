// package model
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package model

import (
	"xiuadmin/internal/model/request"

	"github.com/gogf/gf/v2/os/gtime"
)

type SysPostListModel struct {
	PostId       int64            `json:"postId"       orm:"post_id"       description:"岗位ID"`
	TenantId     string           `json:"tenantId"     orm:"tenant_id"     description:"租户编号"`
	DeptId       int64            `json:"deptId"       orm:"dept_id"       description:"部门id"`
	PostCode     string           `json:"postCode"     orm:"post_code"     description:"岗位编码"`
	PostCategory string           `json:"postCategory" orm:"post_category" description:"岗位类别编码"`
	PostName     string           `json:"postName"     orm:"post_name"     description:"岗位名称"`
	PostSort     int              `json:"postSort"     orm:"post_sort"     description:"显示顺序"`
	Status       string           `json:"status"       orm:"status"        description:"状态（0正常 1停用）"`
	Remark       string           `json:"remark"       orm:"remark"        description:"备注"`
	CreatedAt    *gtime.Time      `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	DeptInfo     SysDeptMiniModel `json:"deptInfo"    orm:"with:dept_id=dept_id"    description:"部门信息"`
}

type SysPostListParam struct {
	request.PageInfo
	DeptId       int64    `json:"deptId"        description:"部门id"`
	PostCode     string   `json:"postCode"      description:"岗位编码"`
	PostCategory string   `json:"postCategory"  description:"岗位类别编码"`
	PostName     string   `json:"postName"      description:"岗位名称"`
	BelongDeptId int64    `json:"belongDeptId"  description:"归属部门id"`
	Status       string   `json:"status"        description:"状态"`
	CreatedAt    []string `json:"createdAt"    description:"创建时间"`
	PostIds      []int64  `json:"postIds"      description:"岗位ID"`
}

type SysPostMiniModel struct {
	PostId   int64  `json:"postId"   orm:"post_id"   description:"岗位ID"`
	PostCode string `json:"postCode" orm:"post_code" description:"岗位编码"`
	PostName string `json:"postName" orm:"post_name" description:"岗位名称"`
	DeptId   int64  `json:"deptId"   orm:"dept_id"   description:"部门ID"`
}

type SysPostAddParam struct {
	DeptId       int64  `json:"deptId"       orm:"dept_id"       description:"部门id"`
	PostCode     string `json:"postCode"     orm:"post_code"     description:"岗位编码"`
	PostCategory string `json:"postCategory" orm:"post_category" description:"岗位类别编码"`
	PostName     string `json:"postName"     orm:"post_name"     description:"岗位名称"`
	PostSort     int    `json:"postSort"     orm:"post_sort"     description:"显示顺序"`
	Status       string `json:"status"       orm:"status"        description:"状态（0正常 1停用）"`
	Remark       string `json:"remark"       orm:"remark"        description:"备注"`
}

type SysPostAddModel struct {
	PostId int64 `json:"postId"       orm:"post_id"       description:"岗位ID"`
}

type SysPostEditParam struct {
	PostId       int64   `json:"postId"       orm:"post_id"       description:"岗位ID"`
	DeptId       *int64  `json:"deptId"       orm:"dept_id"       description:"部门id"`
	PostCode     *string `json:"postCode"     orm:"post_code"     description:"岗位编码"`
	PostCategory *string `json:"postCategory" orm:"post_category" description:"岗位类别编码"`
	PostName     *string `json:"postName"     orm:"post_name"     description:"岗位名称"`
	PostSort     *int    `json:"postSort"     orm:"post_sort"     description:"显示顺序"`
	Status       *string `json:"status"       orm:"status"        description:"状态（0正常 1停用）"`
	Remark       *string `json:"remark"       orm:"remark"        description:"备注"`
}

type SysPostEditModel struct {
	PostId int64 `json:"postId"       orm:"post_id"       description:"岗位ID"`
}

type SysPostDeleteParam struct {
	PostId  int64   `json:"postId" orm:"post_id" description:"岗位ID"`
	PostIds []int64 `json:"postIds" orm:"post_ids" description:"岗位ID"`
}

type SysPostDeleteModel struct {
	PostId int64 `json:"postId" orm:"post_id" description:"岗位ID"`
}

type SysPostViewParam struct {
	PostId int64 `json:"postId" description:"岗位ID"`
}

type SysPostViewModel struct {
	PostId       int64            `json:"postId"       orm:"post_id"       description:"岗位ID"`
	TenantId     string           `json:"tenantId"     orm:"tenant_id"     description:"租户编号"`
	DeptId       int64            `json:"deptId"       orm:"dept_id"       description:"部门id"`
	PostCode     string           `json:"postCode"     orm:"post_code"     description:"岗位编码"`
	PostCategory string           `json:"postCategory" orm:"post_category" description:"岗位类别编码"`
	PostName     string           `json:"postName"     orm:"post_name"     description:"岗位名称"`
	PostSort     int              `json:"postSort"     orm:"post_sort"     description:"显示顺序"`
	Status       string           `json:"status"       orm:"status"        description:"状态（0正常 1停用）"`
	CreatedDept  int64            `json:"createdDept"  orm:"created_dept"  description:"创建部门"`
	CreatedBy    int64            `json:"createdBy"    orm:"created_by"    description:"创建者"`
	CreatedAt    *gtime.Time      `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	Remark       string           `json:"remark"       orm:"remark"        description:"备注"`
	DeptInfo     SysDeptMiniModel `json:"deptInfo"     orm:"with:dept_id=dept_id"    description:"部门信息"`
}

type SysPostExportParam struct {
	*SysPostListParam
}

type SysPostExportModel struct {
	*SysPostListModel
}
