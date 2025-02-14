// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictData is the golang structure for table sys_dict_data.
type SysDictData struct {
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
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"   description:"更新者"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
}
