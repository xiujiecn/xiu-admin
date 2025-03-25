package genmodel

import "github.com/gogf/gf/v2/encoding/gjson"

type SelectItemModel struct {
	Value    interface{} `json:"value"`    // 选择项值
	Label    string      `json:"label"`    // 选择项标签
	Name     string      `json:"name"`     // 选择项名称
	Disabled bool        `json:"disabled"` // 是否禁用
}

// 生成类型选择项
type SelectGenTypeModel struct {
	SelectItemModel
	Templates []*SelectGenTemplateModel `json:"templates"` // 选择项模板
}

// 生成模板选择项
type SelectGenTemplateModel struct {
	SelectItemModel
	IsAddon bool `json:"isAddon"` // 是否为插件
}

type SelectsModel struct {
	GenType       []*SelectGenTypeModel `json:"genType"`       // 生成类型    0:增删改查列表 1:关系树列表
	Db            []*SelectItemModel    `json:"db"`            // 数据库选择项
	Status        []*SelectItemModel    `json:"status"`        // 状态选择项  0:生成成功 1:未生产 2:生成失败
	LinkMode      []*SelectItemModel    `json:"linkMode"`      // 链接模式选择项 1左关联 2右关联 3内关联
	BuildMethod   []*SelectItemModel    `json:"buildMethod"`   // 构建标志选择项 1:创建文件 2强制覆盖 3已存在跳过 4不生成
	FormMode      []*SelectItemModel    `json:"formMode"`      // 表单项类型选择项
	FormRole      []*SelectItemModel    `json:"formRole"`      // 表单项校验规则选择项
	DictMode      []*SelectItemModel    `json:"dictMode"`      // 字典模式选择项
	WhereMode     []*SelectItemModel    `json:"whereMode"`     // 查询条件选择项
	Addons        []*SelectItemModel    `json:"addons"`        // 插件选择项
	TableAlign    []*SelectItemModel    `json:"tableAlign"`    // 表格对齐方式选择项
	TreeStyleType []*SelectItemModel    `json:"treeStyleType"` // 树形结构选择项
}

// GenCodesColumn 生成表字段属性
type GenCodesColumn struct {
	// 表属性
	Id           int64       `json:"id" dc:"序号"`
	Name         string      `json:"name" dc:"字段列名"`
	Dc           string      `json:"dc" dc:"字段描述"`
	DataType     string      `json:"dataType" dc:"字段类型"`
	SqlType      string      `json:"sqlType" dc:"物理类型"`
	Length       int64       `json:"length" dc:"字段长度"`
	IsAllowNull  string      `json:"isAllowNull" dc:"是否允许为空"`
	DefaultValue interface{} `json:"defaultValue" dc:"默认值"`
	Index        string      `json:"index" dc:"索引"`
	Extra        string      `json:"extra" dc:"额外选项"`
	// 自定义生成属性
	// Alias      string `json:"alias" dc:"字段别名"`
	GoName       string `json:"goName" dc:"Go属性"`
	GoType       string `json:"goType" dc:"Go类型"`
	TsName       string `json:"tsName" dc:"Ts属性"`
	TsType       string `json:"tsType" dc:"Ts类型"`
	IsList       bool   `json:"isList" dc:"列表"`
	IsExport     bool   `json:"isExport" dc:"导出"`
	IsSort       bool   `json:"isSort" dc:"排序"`
	IsQuery      bool   `json:"isQuery" dc:"查询"`
	QueryWhere   string `json:"queryWhere" dc:"查询条件"`
	IsEdit       bool   `json:"isEdit" dc:"编辑"`
	Required     bool   `json:"required" dc:"必填"`
	Unique       bool   `json:"unique" dc:"唯一性"`
	FormMode     string `json:"formMode" dc:"表单组件"`
	FormRole     string `json:"formRole" dc:"表单验证"`
	FormGridSpan int    `json:"formGridSpan" dc:"栅格占位数量"`
	DictType     string `json:"dictType" dc:"字典类型ID"`
	Align        string `json:"align"    dc:"排列方式"`
	Width        int64  `json:"width"    dc:"列宽"`
}

type GenCodesTableSelectParam struct {
	DbGroup string `json:"dbGroup"`
}

type GenCodesTableSelectModel struct {
	Value           string `json:"value"`
	Label           string `json:"label"`
	Name            string `json:"name"`
	DaoName         string `json:"daoName" dc:"orm模型名称"`
	DefVarName      string `json:"defVarName" dc:"默认实体名称"`
	DefAlias        string `json:"defAlias" dc:"默认关联表别名"`
	DefTableComment string `json:"defTableComment" dc:"默认菜单名称"`
}

type GenCodesColumnSelectParam struct {
	DbGroup   string `json:"dbGroup"`
	TableName string `json:"tableName"`
}

type GenCodesColumnSelectModel struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Name  string `json:"name"`
}

type GenCodesColumnListParam struct {
	DbGroup   string `json:"dbGroup" dc:"数据库配置名称"`
	TableName string `json:"tableName" dc:"表名称"`
	IsLink    int64  `json:"isLink" dc:"是否是关联表"`
	Alias     string `json:"alias" dc:"关联表别名"`
}

type GenCodesColumnListModel struct {
	GenCodesColumn
}

type GenCodesPreviewParam struct {
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
	Status        int         `json:"status"        orm:"status"         description:"生成状态"`
}

// GenFile 生成文件配置
type GenFile struct {
	Content  string `json:"content" dc:"页面内容"`
	Path     string `json:"path" dc:"生成路径"`
	Meth     int    `json:"meth" dc:"生成方式"`
	Required bool   `json:"required" dc:"是否是必要构建文件"`
}

type GenCodesPreviewModel struct {
	Config *GenCodesConfig     `json:"config"`
	Views  map[string]*GenFile `json:"views" dc:"页面"`
}

type GenCodesBuildParam = GenCodesPreviewParam

type GenCodesBuildModel struct {
}
