package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

type SysDictType struct {
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

type SysDictTypeListQuery struct {
	DictName  string      `json:"dictName"    orm:"dict_name"    description:"字典名称"`
	DictType  string      `json:"dictType"    orm:"dict_type"    description:"字典类型"`
	CreatedAt *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
}

type SysDictData struct {
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
	SysDictType
	Items []SysDictData `json:"items" orm:"with:dict_type=dict_type"`
}
