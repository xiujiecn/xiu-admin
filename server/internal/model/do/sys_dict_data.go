// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictData is the golang structure of table sys_dict_data for DAO operations like Where/Data.
type SysDictData struct {
	g.Meta      `orm:"table:sys_dict_data, do:true"`
	DictCode    interface{} // 字典编码
	TenantId    interface{} // 租户编号
	DictSort    interface{} // 字典排序
	DictLabel   interface{} // 字典标签
	DictValue   interface{} // 字典键值
	DictType    interface{} // 字典类型
	CssClass    interface{} // 样式属性（其他样式扩展）
	ListClass   interface{} // 表格回显样式
	IsDefault   interface{} // 是否默认（Y是 N否）
	CreatedDept interface{} // 创建部门
	CreatedBy   interface{} // 创建者
	CreatedAt   *gtime.Time // 创建时间
	UpdatedBy   interface{} // 更新者
	UpdatedAt   *gtime.Time // 更新时间
	Remark      interface{} // 备注
}
