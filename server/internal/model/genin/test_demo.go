// Package genin
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
// @AutoGenerate Version
package genin

import (
	"context"
	"xiuadmin/internal/library/xgorm/hook"
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/model/request"

	"github.com/gogf/gf/v2/os/gtime"
)

// TestDemoUpdateFields 修改测试单表字段过滤
type TestDemoUpdateFields struct {
	DeptId    int64  `json:"deptId"    dc:"部门id"`
	UserId    int64  `json:"userId"    dc:"用户id"`
	OrderNum  int    `json:"orderNum"  dc:"排序号"`
	TestKey   string `json:"testKey"   dc:"key键"`
	Value     string `json:"value"     dc:"值"`
	Version   int    `json:"version"   dc:"版本"`
	UpdatedBy int64  `json:"updatedBy" dc:"更新者"`
}

// TestDemoInsertFields 新增测试单表字段过滤
type TestDemoInsertFields struct {
	DeptId    int64  `json:"deptId"    dc:"部门id"`
	UserId    int64  `json:"userId"    dc:"用户id"`
	OrderNum  int    `json:"orderNum"  dc:"排序号"`
	TestKey   string `json:"testKey"   dc:"key键"`
	Value     string `json:"value"     dc:"值"`
	Version   int    `json:"version"   dc:"版本"`
	CreatedBy int64  `json:"createdBy" dc:"创建者"`
}

// TestDemoEditParam 修改/新增测试单表
type TestDemoEditParam struct {
	entity.TestDemo
}

func (in *TestDemoEditParam) Filter(ctx context.Context) (err error) {

	return
}

type TestDemoEditModel struct{}

// TestDemoDeleteParam 删除测试单表
type TestDemoDeleteParam struct {
	Id interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *TestDemoDeleteParam) Filter(ctx context.Context) (err error) {
	return
}

type TestDemoDeleteModel struct{}

// TestDemoViewParam 获取指定测试单表信息
type TestDemoViewParam struct {
	Id int64 `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *TestDemoViewParam) Filter(ctx context.Context) (err error) {
	return
}

type TestDemoViewModel struct {
	entity.TestDemo
	CreatedBySumma *hook.MemberSumma `json:"createdBySumma" dc:"创建者摘要信息"`
	UpdatedBySumma *hook.MemberSumma `json:"updatedBySumma" dc:"更新者摘要信息"`
	DeletedBySumma *hook.MemberSumma `json:"deletedBySumma" dc:"删除人摘要信息"`
}

// TestDemoListParam 获取测试单表列表
type TestDemoListParam struct {
	request.PageInfo
	Id           int64         `json:"id"           dc:"主键"`
	CreatedAt    []*gtime.Time `json:"createdAt"    dc:"创建时间"`
	DeptDeptName string        `json:"deptDeptName" dc:"部门名称"`
}

func (in *TestDemoListParam) Filter(ctx context.Context) (err error) {
	return
}

type TestDemoListModel struct {
	Id             int64             `json:"id"             dc:"主键"`
	TenantId       string            `json:"tenantId"       dc:"租户编号"`
	DeptId         int64             `json:"deptId"         dc:"部门id"`
	UserId         int64             `json:"userId"         dc:"用户id"`
	OrderNum       int               `json:"orderNum"       dc:"排序号"`
	TestKey        string            `json:"testKey"        dc:"key键"`
	Value          string            `json:"value"          dc:"值"`
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
	DeptDeptName   string            `json:"deptDeptName"   dc:"部门名称"`
}

// TestDemoExportModel 导出测试单表
type TestDemoExportModel struct {
	Id           int64       `json:"id"           dc:"主键"`
	TenantId     string      `json:"tenantId"     dc:"租户编号"`
	DeptId       int64       `json:"deptId"       dc:"部门id"`
	UserId       int64       `json:"userId"       dc:"用户id"`
	OrderNum     int         `json:"orderNum"     dc:"排序号"`
	TestKey      string      `json:"testKey"      dc:"key键"`
	Value        string      `json:"value"        dc:"值"`
	Version      int         `json:"version"      dc:"版本"`
	CreatedDept  int64       `json:"createdDept"  dc:"创建部门"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"创建时间"`
	CreatedBy    int64       `json:"createdBy"    dc:"创建者"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"更新时间"`
	UpdatedBy    int64       `json:"updatedBy"    dc:"更新者"`
	DeletedBy    int64       `json:"deletedBy"    dc:"删除人"`
	DeptDeptName string      `json:"deptDeptName" dc:"部门名称"`
}
