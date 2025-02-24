package model

import "github.com/gogf/gf/v2/frame/g"

type DeptMiniModel struct {
	g.Meta   `orm:"table:sys_dept" description:"部门"`
	DeptId   int64  `json:"deptId"  orm:"dept_id" description:"部门id"`
	DeptName string `json:"deptName" orm:"dept_name" description:"部门名称"`
}

type DeptListQuery struct {
	DeptName string `json:"deptName"`
	Status   string `json:"status"`
}

type SysDept struct {
	DeptId       int64  `json:"deptId"       orm:"dept_id"       description:"部门id"`
	TenantId     string `json:"tenantId"     orm:"tenant_id"     description:"租户编号"`
	ParentId     int64  `json:"parentId"     orm:"parent_id"     description:"父部门id"`
	Ancestors    string `json:"ancestors"    orm:"ancestors"     description:"祖级列表"`
	DeptName     string `json:"deptName"     orm:"dept_name"     description:"部门名称"`
	DeptCategory string `json:"deptCategory" orm:"dept_category" description:"部门类别编码"`
	OrderNum     int    `json:"orderNum"     orm:"order_num"     description:"显示顺序"`
	Leader       int64  `json:"leader"       orm:"leader"        description:"负责人"`
	Phone        string `json:"phone"        orm:"phone"         description:"联系电话"`
	Email        string `json:"email"        orm:"email"         description:"邮箱"`
	Status       string `json:"status"       orm:"status"        description:"部门状态（0正常 1停用）"`
}

type SysDeptTreeModel struct {
	DeptId    int64               `json:"deptId"      description:"部门id"`
	ParentId  int64               `json:"parentId"    description:"父部门id"`
	Key       string              `json:"key"         description:"key"`
	DeptName  string              `json:"deptName"    description:"部门名称"`
	Children  []*SysDeptTreeModel `json:"children"   description:"子部门"`
	Ancestors string              `json:"ancestors"   description:"祖级列表"`
}
