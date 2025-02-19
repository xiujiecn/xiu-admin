package model

import "github.com/gogf/gf/v2/os/gtime"

type SysPost struct {
	PostId       int64       `json:"postId"       orm:"post_id"       description:"岗位ID"`
	TenantId     string      `json:"tenantId"     orm:"tenant_id"     description:"租户编号"`
	DeptId       int64       `json:"deptId"       orm:"dept_id"       description:"部门id"`
	PostCode     string      `json:"postCode"     orm:"post_code"     description:"岗位编码"`
	PostCategory string      `json:"postCategory" orm:"post_category" description:"岗位类别编码"`
	PostName     string      `json:"postName"     orm:"post_name"     description:"岗位名称"`
	PostSort     int         `json:"postSort"     orm:"post_sort"     description:"显示顺序"`
	Status       string      `json:"status"       orm:"status"        description:"状态（0正常 1停用）"`
	Remark       string      `json:"remark"       orm:"remark"        description:"备注"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
}

type SysPostListQuery struct {
	DeptId       int64  `json:"deptId"       orm:"dept_id"       description:"部门id"`
	PostCode     string `json:"postCode"     orm:"post_code"     description:"岗位编码"`
	PostCategory string `json:"postCategory" orm:"post_category" description:"岗位类别编码"`
	PostName     string `json:"postName"     orm:"post_name"     description:"岗位名称"`
}
