// package model
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package model

import (
	genmodel "xiuadmin/internal/library/xgen/gen_model"
	"xiuadmin/internal/model/request"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

type SysGenTableListParam struct {
	request.PageInfo
	GenType uint   `json:"genType"      description:"生成类型"`
	VarName string `json:"varName"      description:"实体命名"`
	Status  string `json:"status"       description:"生成状态（0成功 1未开始）"`
}

type SysGenTableListModel struct {
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
}

type SysGenTableViewParam struct {
	TableId int64 `json:"tableId" description:"生成ID"`
}

type SysGenTableViewModel struct {
	TableId       int64       `json:"tableId"       orm:"table_id"       description:"生成ID"`
	GenType       uint        `json:"genType"       orm:"gen_type"       description:"生成类型"`
	GenTemplate   int         `json:"genTemplate"   orm:"gen_template"   description:"生成模板"`
	VarName       string      `json:"varName"       orm:"var_name"       description:"实体命名"`
	Options       *gjson.Json `json:"options"       orm:"options"        description:"配置选项"`
	DbName        string      `json:"dbName"        orm:"db_name"        description:"数据库名称"`
	TableName     string      `json:"tableName"     orm:"table_name"     description:"主表名称"`
	TableComment  string      `json:"tableComment"  orm:"table_comment"  description:"主表注释"`
	DaoName       string      `json:"daoName"       orm:"dao_name"       description:"主表dao模型"`
	MasterColumns *gjson.Json `json:"masterColumns" orm:"master_columns" description:"主表字段"`
	AddonName     string      `json:"addonName"     orm:"addon_name"     description:"插件名称"`
	Status        string      `json:"status"        orm:"status"         description:"生成状态（0成功 1未开始）"`
	CreatedDept   int64       `json:"createdDept"   orm:"created_dept"   description:"创建部门"`
	CreatedBy     int64       `json:"createdBy"     orm:"created_by"     description:"创建者"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
	UpdatedBy     int64       `json:"updatedBy"     orm:"updated_by"     description:"更新者"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`
}

type SysGenTableAddParam struct {
	GenType       uint   `json:"genType"       description:"生成类型"`
	GenTemplate   int    `json:"genTemplate"   description:"生成模板"`
	VarName       string `json:"varName"       description:"实体命名"`
	Options       string `json:"options"       description:"配置选项"`
	DbName        string `json:"dbName"        description:"数据库名称"`
	TableName     string `json:"tableName"     description:"主表名称"`
	TableComment  string `json:"tableComment"  description:"主表注释"`
	DaoName       string `json:"daoName"       description:"主表dao模型"`
	MasterColumns string `json:"masterColumns" description:"主表字段"`
	AddonName     string `json:"addonName"     description:"插件名称"`
	// Status        string `json:"status"        description:"生成状态（0成功 1未开始）"`
}

type SysGenTableAddModel struct {
	TableId int64 `json:"tableId"       orm:"table_id"       description:"生成ID"`
}
type SysGenTableEditParam struct {
	TableId       int64   `json:"tableId"       description:"生成ID"`
	GenType       *uint   `json:"genType"       description:"生成类型"`
	GenTemplate   *int    `json:"genTemplate"   description:"生成模板"`
	VarName       *string `json:"varName"       description:"实体命名"`
	Options       *string `json:"options"       description:"配置选项"`
	DbName        *string `json:"dbName"        description:"数据库名称"`
	TableName     *string `json:"tableName"     description:"主表名称"`
	TableComment  *string `json:"tableComment"  description:"主表注释"`
	DaoName       *string `json:"daoName"       description:"主表dao模型"`
	MasterColumns *string `json:"masterColumns" description:"主表字段"`
	AddonName     *string `json:"addonName"     description:"插件名称"`
	Status        *string `json:"status"        description:"生成状态（0成功 1未开始）"`
}

type SysGenTableEditModel struct {
	TableId int64 `json:"tableId"       orm:"table_id"       description:"生成ID"`
}
type SysGenTableDeleteParam struct {
	TableIds []int64 `json:"tableIds" description:"生成ID"`
}
type SysGenTableDeleteModel struct {
	TableIds []int64 `json:"tableIds" description:"生成ID"`
}

type SelectsModel = genmodel.SelectsModel

type GenCodesTableSelectParam = genmodel.GenCodesTableSelectParam

type GenCodesTableSelectModel = genmodel.GenCodesTableSelectModel

type GenCodesColumnSelectParam = genmodel.GenCodesColumnSelectParam

type GenCodesColumnSelectModel = genmodel.GenCodesColumnSelectModel

type GenCodesColumnListParam = genmodel.GenCodesColumnListParam

type GenCodesColumnListModel = genmodel.GenCodesColumnListModel

type GenCodesPreviewParam = genmodel.GenCodesPreviewParam

type GenCodesPreviewModel = genmodel.GenCodesPreviewModel

type GenCodesBuildParam = genmodel.GenCodesBuildParam

type GenCodesBuildModel = genmodel.GenCodesBuildModel
