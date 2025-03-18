// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGenTable is the golang structure of table sys_gen_table for DAO operations like Where/Data.
type SysGenTable struct {
	g.Meta        `orm:"table:sys_gen_table, do:true"`
	TableId       interface{} // 生成ID
	GenType       interface{} // 生成类型
	GenTemplate   interface{} // 生成模板
	VarName       interface{} // 实体命名
	Options       interface{} // 配置选项
	DbName        interface{} // 数据库名称
	TableName     interface{} // 主表名称
	TableComment  interface{} // 主表注释
	DaoName       interface{} // 主表dao模型
	MasterColumns interface{} // 主表字段
	AddonName     interface{} // 插件名称
	Status        interface{} // 生成状态（0成功 1未开始）
	CreatedDept   interface{} // 创建部门
	CreatedBy     interface{} // 创建者
	CreatedAt     *gtime.Time // 创建时间
	UpdatedBy     interface{} // 更新者
	UpdatedAt     *gtime.Time // 更新时间
}
