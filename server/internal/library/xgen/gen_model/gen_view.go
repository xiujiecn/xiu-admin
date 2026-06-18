// package genmodel
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package genmodel

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gview"
)

// ImportModel 导包 - model.ts
type ImportModel struct {
	NaiveUI       []string
	UtilsIs       []string
	UtilsUrl      []string
	UtilsDate     []string
	UtilsValidate []string
	UtilsIndex    []string
}

// CurdStep 生成选项
type CurdStep struct {
	HasMaxSort            bool        // 最大排序
	HasAdd                bool        // 表单添加
	HasBatchDel           bool        // 批量删除
	HasExport             bool        // 表格导出
	HasNotFilterAuth      bool        // 不过滤认证权限
	HasEdit               bool        // 表单编辑
	HasDel                bool        // 删除
	HasView               bool        // 查看详情
	HasStatus             bool        // 修改状态
	HasSwitch             bool        // 数值开关
	HasCheck              bool        // 勾选列
	HasMenu               bool        // 菜单权限
	IsTreeTable           bool        // 树型列表
	IsOptionTreeTable     bool        // 选项式树型列表
	HasRules              bool        // 表单验证规则
	HasRulesValidator     bool        // 表单验证器
	HasSearchForm         bool        // 列表搜索
	HasDict               bool        // 字典
	HasFuncDict           bool        // 注册方法字典
	HasQueryMemberSummary bool        // 查询用户摘要
	HasHookMemberSummary  bool        // hook用户摘要
	ImportModel           ImportModel // 公用导包 - model.ts
	ActionColumnWidth     int64       // 列表操作栏宽度
	IsAddon               bool        // 是否是插件
	IsEditModal           bool        // 新增修改模态框模式
}

type FuncDict struct {
	ValueColumn string // 选项值
	LabelColumn string // 选项名称
	Value       *GenCodesColumnListModel
	Label       *GenCodesColumnListModel
}

// PresetStep 预设生成流程参数
type PresetStep struct {
	FormGridCols int `json:"formGridCols" dc:"表单显示的栅格数量"`
}

// OptionsTree 树形结构选项
type OptionsTree struct {
	TitleColumn string `json:"titleColumn"`
	PidColumn   string `json:"pidColumn"`
	LevelColumn string `json:"levelColumn"`
	TreeColumn  string `json:"treeColumn"`
	StyleType   int    `json:"styleType"`
	TitleField  *GenCodesColumnListModel
	PidField    *GenCodesColumnListModel
	LevelField  *GenCodesColumnListModel
	TreeField   *GenCodesColumnListModel
	HasTreePath bool
}

// CurdOptionsMenu 菜单选项
type CurdOptionsMenu struct {
	Icon string `json:"icon"`
	Pid  int    `json:"pid"`
	Sort int    `json:"sort"`
}

// CurdOptionsJoin 关联表选项
type CurdOptionsJoin struct {
	Uuid        string                     `json:"uuid"`
	LinkTable   string                     `json:"linkTable"`
	Alias       string                     `json:"alias"`
	LinkMode    int                        `json:"linkMode"`
	Field       string                     `json:"field"`
	MasterField string                     `json:"masterField"`
	DaoName     string                     `json:"daoName"`
	Columns     []*GenCodesColumnListModel `json:"columns"`
}

type OptionsSchemasField struct {
	Field string
	Type  string
}

type CurdOptionsDict struct {
	Has     bool
	Types   []string
	Schemas []*OptionsSchemasField
}

// CurdOptions 生成选项
type CurdOptions struct {
	AutoOps       []string           `json:"autoOps"`
	ColumnOps     []string           `json:"columnOps"`
	HeadOps       []string           `json:"headOps"`
	Join          []*CurdOptionsJoin `json:"join"`
	Menu          *CurdOptionsMenu   `json:"menu"`
	Tree          *OptionsTree       `json:"tree"`
	TemplateGroup string             `json:"templateGroup"`
	ApiPrefix     string             `json:"apiPrefix"`
	ImportWebApi  string             `json:"importWebApi"`
	FuncDict      *FuncDict          `json:"funcDict"`
	PresetStep    *PresetStep        `json:"presetStep"`
	Step          *CurdStep          // 转换后的流程控制条件
	DictOps       CurdOptionsDict    // 字典选项
	DictMap       g.Map              // 字典选项 -> 字段映射关系
}

type CurdPreviewParam struct {
	In           *GenCodesPreviewParam      // 提交参数
	DaoConfig    *GenDaoConfig              // 生成dao配置
	Config       *GenCodesConfig            // 生成配置
	View         *gview.View                // 视图模板
	Content      *GenCodesPreviewModel      // 页面代码
	MasterFields []*GenCodesColumnListModel // 主表字段属性
	Pk           *GenCodesColumnListModel   // 主键属性
	Options      *CurdOptions               // 生成选项
}

type CurdBuildEvent map[string]func(ctx context.Context) (err error)

type CurdBuildParam struct {
	PreviewIn   *CurdPreviewParam // 预览参数
	BeforeEvent CurdBuildEvent    // 前置事件
	AfterEvent  CurdBuildEvent    // 后置事件
}
