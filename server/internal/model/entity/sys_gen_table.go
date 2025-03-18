// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGenTable is the golang structure for table sys_gen_table.
type SysGenTable struct {
	TableId       int64       `json:"tableId"       orm:"table_id"       description:"生成ID"`
	GenType       uint        `json:"genType"       orm:"gen_type"       description:"生成类型"`
	GenTemplate   int         `json:"genTemplate"   orm:"gen_template"   description:"生成模板"`
	VarName       string      `json:"varName"       orm:"var_name"       description:"实体命名"`
	Options       string      `json:"options"       orm:"options"        description:"配置选项"`
	DbName        string      `json:"dbName"        orm:"db_name"        description:"数据库名称"`
	TableName     string      `json:"tableName"     orm:"table_name"     description:"主表名称"`
	TableComment  string      `json:"tableComment"  orm:"table_comment"  description:"主表注释"`
	DaoName       string      `json:"daoName"       orm:"dao_name"       description:"主表dao模型"`
	MasterColumns string      `json:"masterColumns" orm:"master_columns" description:"主表字段"`
	AddonName     string      `json:"addonName"     orm:"addon_name"     description:"插件名称"`
	Status        string      `json:"status"        orm:"status"         description:"生成状态（0成功 1未开始）"`
	CreatedDept   int64       `json:"createdDept"   orm:"created_dept"   description:"创建部门"`
	CreatedBy     int64       `json:"createdBy"     orm:"created_by"     description:"创建者"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
	UpdatedBy     int64       `json:"updatedBy"     orm:"updated_by"     description:"更新者"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`
}
