// Package genin
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
// @AutoGenerate Version
package genin

import (
	"context"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/library/xgorm/hook"
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/model/request"
	"xiuadmin/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// TestCategoryUpdateFields 修改测试分类字段过滤
type TestCategoryUpdateFields struct {
	Name        string `json:"name"        dc:"分类名称"`
	ShortName   string `json:"shortName"   dc:"简称"`
	Description string `json:"description" dc:"描述"`
	Sort        int    `json:"sort"        dc:"排序"`
	Remark      string `json:"remark"      dc:"备注"`
	Status      int    `json:"status"      dc:"状态"`
	UpdatedBy   int64  `json:"updatedBy"   dc:"修改者"`
}

// TestCategoryInsertFields 新增测试分类字段过滤
type TestCategoryInsertFields struct {
	Name        string `json:"name"        dc:"分类名称"`
	ShortName   string `json:"shortName"   dc:"简称"`
	Description string `json:"description" dc:"描述"`
	Sort        int    `json:"sort"        dc:"排序"`
	Remark      string `json:"remark"      dc:"备注"`
	Status      int    `json:"status"      dc:"状态"`
	CreatedBy   int64  `json:"createdBy"   dc:"创建者"`
}

// TestCategoryEditParam 修改/新增测试分类
type TestCategoryEditParam struct {
	entity.TestCategory
}

func (in *TestCategoryEditParam) Filter(ctx context.Context) (err error) {

	return
}

type TestCategoryEditModel struct{}

// TestCategoryDeleteParam 删除测试分类
type TestCategoryDeleteParam struct {
	Id interface{} `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *TestCategoryDeleteParam) Filter(ctx context.Context) (err error) {
	return
}

type TestCategoryDeleteModel struct{}

// TestCategoryViewParam 获取指定测试分类信息
type TestCategoryViewParam struct {
	Id int64 `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *TestCategoryViewParam) Filter(ctx context.Context) (err error) {
	return
}

type TestCategoryViewModel struct {
	entity.TestCategory
	CreatedBySumma *hook.MemberSumma `json:"createdBySumma" dc:"创建者摘要信息"`
	UpdatedBySumma *hook.MemberSumma `json:"updatedBySumma" dc:"修改者摘要信息"`
	DeletedBySumma *hook.MemberSumma `json:"deletedBySumma" dc:"删除者摘要信息"`
}

// TestCategoryListParam 获取测试分类列表
type TestCategoryListParam struct {
	request.PageInfo
	Id        int64         `json:"id"        dc:"分类ID"`
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *TestCategoryListParam) Filter(ctx context.Context) (err error) {
	return
}

type TestCategoryListModel struct {
	Id             int64             `json:"id"             dc:"分类ID"`
	Name           string            `json:"name"           dc:"分类名称"`
	ShortName      string            `json:"shortName"      dc:"简称"`
	Description    string            `json:"description"    dc:"描述"`
	Sort           int               `json:"sort"           dc:"排序"`
	Remark         string            `json:"remark"         dc:"备注"`
	Status         int               `json:"status"         dc:"状态"`
	CreatedDept    int64             `json:"createdDept"    dc:"创建部门"`
	CreatedAt      *gtime.Time       `json:"createdAt"      dc:"创建时间"`
	CreatedBy      int64             `json:"createdBy"      dc:"创建者"`
	CreatedBySumma *hook.MemberSumma `json:"createdBySumma" dc:"创建者摘要信息"`
	UpdatedAt      *gtime.Time       `json:"updatedAt"      dc:"修改时间"`
	UpdatedBy      int64             `json:"updatedBy"      dc:"修改者"`
	UpdatedBySumma *hook.MemberSumma `json:"updatedBySumma" dc:"修改者摘要信息"`
	DeletedBy      int64             `json:"deletedBy"      dc:"删除者"`
	DeletedBySumma *hook.MemberSumma `json:"deletedBySumma" dc:"删除者摘要信息"`
}

// TestCategoryExportModel 导出测试分类
type TestCategoryExportModel struct {
	Id          int64       `json:"id"          dc:"分类ID"`
	Name        string      `json:"name"        dc:"分类名称"`
	ShortName   string      `json:"shortName"   dc:"简称"`
	Description string      `json:"description" dc:"描述"`
	Sort        int         `json:"sort"        dc:"排序"`
	Remark      string      `json:"remark"      dc:"备注"`
	Status      int         `json:"status"      dc:"状态"`
	CreatedDept int64       `json:"createdDept" dc:"创建部门"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	CreatedBy   int64       `json:"createdBy"   dc:"创建者"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"修改时间"`
	UpdatedBy   int64       `json:"updatedBy"   dc:"修改者"`
	DeletedBy   int64       `json:"deletedBy"   dc:"删除者"`
}

// TestCategoryMaxSortParam 获取测试分类最大排序
type TestCategoryMaxSortParam struct{}

func (in *TestCategoryMaxSortParam) Filter(ctx context.Context) (err error) {
	return
}

type TestCategoryMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// TestCategoryStatusParam 更新测试分类状态
type TestCategoryStatusParam struct {
	Id     int64 `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *TestCategoryStatusParam) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("分类ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type TestCategoryStatusModel struct{}