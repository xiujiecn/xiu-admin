// Package genin
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2026 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
// @AutoGenerate Version
package genin

import (
	"context"
	"xiuadmin/internal/library/xgorm/hook"
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/model/request"
	"xiuadmin/utility/tree"

	"github.com/gogf/gf/v2/os/gtime"
)

// TestTreeUpdateFields 修改测试树表字段过滤
type TestTreeUpdateFields struct {
	TenantId  string      `json:"tenantId"  dc:"租户编号"`
	ParentId  int64       `json:"parentId"  dc:"父id"`
	DeptId    int64       `json:"deptId"    dc:"部门id"`
	UserId    int64       `json:"userId"    dc:"用户id"`
	TreeName  string      `json:"treeName"  dc:"值"`
	Level     int         `json:"level"     dc:"关系树等级"`
	Tree      string      `json:"tree"      dc:"关系树"`
	Version   int         `json:"version"   dc:"版本"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
	UpdatedBy int64       `json:"updatedBy" dc:"更新者"`
}

// TestTreeInsertFields 新增测试树表字段过滤
type TestTreeInsertFields struct {
	TenantId    string      `json:"tenantId"    dc:"租户编号"`
	ParentId    int64       `json:"parentId"    dc:"父id"`
	DeptId      int64       `json:"deptId"      dc:"部门id"`
	UserId      int64       `json:"userId"      dc:"用户id"`
	TreeName    string      `json:"treeName"    dc:"值"`
	Level       int         `json:"level"       dc:"关系树等级"`
	Tree        string      `json:"tree"        dc:"关系树"`
	Version     int         `json:"version"     dc:"版本"`
	CreatedDept int64       `json:"createdDept" dc:"创建部门"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	CreatedBy   int64       `json:"createdBy"   dc:"创建者"`
}

// TestTreeEditParam 修改/新增测试树表
type TestTreeEditParam struct {
	//entity.TestTree
	Id          int64       `json:"id"          dc:"主键"`
	TenantId    *string     `json:"tenantId"    dc:"租户编号"`
	ParentId    int64       `json:"parentId"    dc:"父id"`
	DeptId      *int64      `json:"deptId"      dc:"部门id"`
	UserId      *int64      `json:"userId"      dc:"用户id"`
	TreeName    *string     `json:"treeName"    dc:"值"`
	Level       int         `json:"level"       dc:"关系树等级"`
	Tree        string      `json:"tree"        dc:"关系树"`
	Version     *int        `json:"version"     dc:"版本"`
	CreatedDept *int64      `json:"createdDept" dc:"创建部门"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	CreatedBy   *int64      `json:"createdBy"   dc:"创建者"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"更新时间"`
	UpdatedBy   *int64      `json:"updatedBy"   dc:"更新者"`
	DeletedBy   *int64      `json:"deletedBy"   dc:"删除人"`
	DeletedAt   *gtime.Time `json:"deletedAt"   dc:"删除时间"`
}

func (in *TestTreeEditParam) Filter(ctx context.Context) (err error) {

	return
}

type TestTreeEditModel struct{}

// TestTreeDeleteParam 删除测试树表
type TestTreeDeleteParam struct {
	Id interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *TestTreeDeleteParam) Filter(ctx context.Context) (err error) {
	return
}

type TestTreeDeleteModel struct{}

// TestTreeViewParam 获取指定测试树表信息
type TestTreeViewParam struct {
	Id int64 `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *TestTreeViewParam) Filter(ctx context.Context) (err error) {
	return
}

type TestTreeViewModel struct {
	entity.TestTree
	CreatedBySumma *hook.MemberSumma `json:"createdBySumma" dc:"创建者摘要信息"`
	UpdatedBySumma *hook.MemberSumma `json:"updatedBySumma" dc:"更新者摘要信息"`
	DeletedBySumma *hook.MemberSumma `json:"deletedBySumma" dc:"删除人摘要信息"`
}

// TestTreeListParam 获取测试树表列表
type TestTreeListParam struct {
	request.PageInfo
	Id        int64         `json:"id"        dc:"主键"`
	ParentId  int64         `json:"parentId"  dc:"父id"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *TestTreeListParam) Filter(ctx context.Context) (err error) {
	return
}

type TestTreeListModel struct {
	Id             int64             `json:"id"             dc:"主键"`
	TenantId       string            `json:"tenantId"       dc:"租户编号"`
	ParentId       int64             `json:"parentId"       dc:"父id"`
	DeptId         int64             `json:"deptId"         dc:"部门id"`
	UserId         int64             `json:"userId"         dc:"用户id"`
	TreeName       string            `json:"treeName"       dc:"值"`
	Version        int               `json:"version"        dc:"版本"`
	CreatedDept    int64             `json:"createdDept"    dc:"创建部门"`
	CreatedAt      *gtime.Time       `json:"createdAt"      dc:"创建时间"`
	CreatedBy      int64             `json:"createdBy"      dc:"创建者"`
	CreatedBySumma *hook.MemberSumma `json:"createdBySumma" dc:"创建者摘要信息"`
	UpdatedAt      *gtime.Time       `json:"updatedAt"      dc:"更新时间"`
	UpdatedBy      int64             `json:"updatedBy"      dc:"更新者"`
	UpdatedBySumma *hook.MemberSumma `json:"updatedBySumma" dc:"更新者摘要信息"`
	DeletedBy      int64             `json:"deletedBy"      dc:"删除人"`
	DeletedBySumma *hook.MemberSumma `json:"deletedBySumma" dc:"删除人摘要信息"`
}

// TestTreeExportModel 导出测试树表
type TestTreeExportModel struct {
	Id          int64       `json:"id"          dc:"主键"`
	TenantId    string      `json:"tenantId"    dc:"租户编号"`
	ParentId    int64       `json:"parentId"    dc:"父id"`
	DeptId      int64       `json:"deptId"      dc:"部门id"`
	UserId      int64       `json:"userId"      dc:"用户id"`
	TreeName    string      `json:"treeName"    dc:"值"`
	Version     int         `json:"version"     dc:"版本"`
	CreatedDept int64       `json:"createdDept" dc:"创建部门"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	CreatedBy   int64       `json:"createdBy"   dc:"创建者"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"更新时间"`
	UpdatedBy   int64       `json:"updatedBy"   dc:"更新者"`
	DeletedBy   int64       `json:"deletedBy"   dc:"删除人"`
}

// TestTreeTreeOption 关系树选项
type TestTreeTreeOption struct {
	Id       int64       `json:"id"       dc:"主键"`
	ParentId int64       `json:"parentId" dc:"父id"`
	TreeName string      `json:"treeName" dc:"值"`
	Children []tree.Node `json:"children"  dc:"子节点"`
}

// ID 获取节点ID
func (t *TestTreeTreeOption) ID() int64 {
	return t.Id
}

// PID 获取父级节点ID
func (t *TestTreeTreeOption) PID() int64 {
	return t.ParentId
}

// SetChildren 设置子节点数据
func (t *TestTreeTreeOption) SetChildren(children []tree.Node) {
	t.Children = children
}