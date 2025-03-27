// package model
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package model

import (
	"xiuadmin/internal/model/request"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

type SysDictTypeListModel struct {
	gmeta.Meta  `orm:"table:sys_dict_type"`
	DictId      int64       `json:"dictId"      orm:"dict_id"      description:"字典主键"`
	TenantId    string      `json:"tenantId"    orm:"tenant_id"    description:"租户编号"`
	DictName    string      `json:"dictName"    orm:"dict_name"    description:"字典名称"`
	DictType    string      `json:"dictType"    orm:"dict_type"    description:"字典类型"`
	CreatedDept int64       `json:"createdDept" orm:"created_dept" description:"创建部门"`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:"创建者"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
}

type SysDictTypeListParam struct {
	request.PageInfo
	DictName  string   `json:"dictName"    description:"字典名称"`
	DictType  string   `json:"dictType"    description:"字典类型"`
	CreatedAt []string `json:"createdAt"   description:"创建时间"`
}

type SysDictTypeViewParam struct {
	DictId   int64  `json:"dictId"    description:"字典主键"`
	DictType string `json:"dictType"    orm:"dict_type"    description:"字典类型"`
}

type SysDictTypeViewModel struct {
	gmeta.Meta  `orm:"table:sys_dict_type"`
	DictId      int64       `json:"dictId"      orm:"dict_id"      description:"字典主键"`
	TenantId    string      `json:"tenantId"    orm:"tenant_id"    description:"租户编号"`
	DictName    string      `json:"dictName"    orm:"dict_name"    description:"字典名称"`
	DictType    string      `json:"dictType"    orm:"dict_type"    description:"字典类型"`
	CreatedDept int64       `json:"createdDept" orm:"created_dept" description:"创建部门"`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:"创建者"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"   description:"更新者"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
}

type SysDictTypeAddParam struct {
	DictName string `json:"dictName"    description:"字典名称"`
	DictType string `json:"dictType"    description:"字典类型"`
	Remark   string `json:"remark"      description:"备注"`
}

type SysDictTypeAddModel struct {
	DictId int64 `json:"dictId"        description:"字典主键"`
}

type SysDictTypeEditParam struct {
	DictId   int64  `json:"dictId"       description:"字典主键"`
	DictName string `json:"dictName"     description:"字典名称"`
	DictType string `json:"dictType"     description:"字典类型"`
	Remark   string `json:"remark"       description:"备注"`
}

type SysDictTypeEditModel struct {
	DictId int64 `json:"dictId"        description:"字典主键"`
}

type SysDictTypeDeleteParam struct {
	gmeta.Meta `orm:"table:sys_dict_type"`
	DictIds    []int64 `json:"dictIds"      description:"字典主键"`
}

type SysDictTypeDeleteModel struct {
	DictIds []int64 `json:"dictIds"       description:"字典主键"`
}

type SysDictDataListModel struct {
	gmeta.Meta  `orm:"table:sys_dict_data"`
	DictCode    int64       `json:"dictCode"    orm:"dict_code"    description:"字典编码"`
	TenantId    string      `json:"tenantId"    orm:"tenant_id"    description:"租户编号"`
	DictSort    int         `json:"dictSort"    orm:"dict_sort"    description:"字典排序"`
	DictLabel   string      `json:"dictLabel"   orm:"dict_label"   description:"字典标签"`
	DictValue   string      `json:"dictValue"   orm:"dict_value"   description:"字典键值"`
	DictType    string      `json:"dictType"    orm:"dict_type"    description:"字典类型"`
	CssClass    string      `json:"cssClass"    orm:"css_class"    description:"样式属性（其他样式扩展）"`
	ListClass   string      `json:"listClass"   orm:"list_class"   description:"表格回显样式"`
	IsDefault   string      `json:"isDefault"   orm:"is_default"   description:"是否默认（Y是 N否）"`
	CreatedDept int64       `json:"createdDept" orm:"created_dept" description:"创建部门"`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:"创建者"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
}

type SysDictDataList struct {
	SysDictTypeListModel
	Items []SysDictDataListModel `json:"items" orm:"with:dict_type=dict_type"`
}

type SysDictDataListParam struct {
	request.PageInfo
	DictId   int64  `json:"dictId"`
	DictType string `json:"dictType"`
}

type SysDictDataViewParam struct {
	DictCode int64 `json:"dictCode"     description:"字典编码"`
}

type SysDictDataViewModel struct {
	DictCode    int64       `json:"dictCode"    description:"字典编码"`
	TenantId    string      `json:"tenantId"    description:"租户编号"`
	DictSort    int         `json:"dictSort"    description:"字典排序"`
	DictLabel   string      `json:"dictLabel"   description:"字典标签"`
	DictValue   string      `json:"dictValue"   description:"字典键值"`
	DictType    string      `json:"dictType"    description:"字典类型"`
	CssClass    string      `json:"cssClass"    description:"样式属性（其他样式扩展）"`
	ListClass   string      `json:"listClass"   description:"表格回显样式"`
	IsDefault   string      `json:"isDefault"   description:"是否默认（Y是 N否）"`
	CreatedDept int64       `json:"createdDept" description:"创建部门"`
	CreatedBy   int64       `json:"createdBy"   description:"创建者"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedBy   int64       `json:"updatedBy"   description:"更新者"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"更新时间"`
	Remark      string      `json:"remark"      description:"备注"`
}

type SysDictDataAddParam struct {
	DictSort  int    `json:"dictSort"    description:"字典排序"`
	DictLabel string `json:"dictLabel"   description:"字典标签"`
	DictValue string `json:"dictValue"   description:"字典键值"`
	DictType  string `json:"dictType"    description:"字典类型"`
	CssClass  string `json:"cssClass"    description:"样式属性（其他样式扩展）"`
	ListClass string `json:"listClass"   description:"表格回显样式"`
	IsDefault string `json:"isDefault"   description:"是否默认（Y是 N否）"`
	Remark    string `json:"remark"      description:"备注"`
}

type SysDictDataAddModel struct {
	DictCode int64 `json:"dictCode"    description:"字典编码"`
}

type SysDictDataEditParam struct {
	DictCode  int64  `json:"dictCode"     description:"字典编码"`
	DictSort  int    `json:"dictSort"     description:"字典排序"`
	DictLabel string `json:"dictLabel"    description:"字典标签"`
	DictValue string `json:"dictValue"    description:"字典键值"`
	DictType  string `json:"dictType"     description:"字典类型"`
	CssClass  string `json:"cssClass"     description:"样式属性（其他样式扩展）"`
	ListClass string `json:"listClass"    description:"表格回显样式"`
	IsDefault string `json:"isDefault"    description:"是否默认（Y是 N否）"`
	Remark    string `json:"remark"       description:"备注"`
}

type SysDictDataEditModel struct {
	DictCode int64 `json:"dictCode"    description:"字典编码"`
}

type SysDictDataDeleteParam struct {
	DictCodes []int64 `json:"dictCodes"      description:"字典编码"`
}

type SysDictDataDeleteModel struct {
	DictCodes []int64 `json:"dictCodes"       description:"字典编码"`
}
