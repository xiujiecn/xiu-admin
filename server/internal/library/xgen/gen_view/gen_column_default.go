// package genview
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package genview

import (
	genconsts "xiuadmin/internal/library/xgen/gen_consts"
	genmodel "xiuadmin/internal/library/xgen/gen_model"

	"github.com/gogf/gf/v2/text/gstr"
)

// 默认表单组件映射 Ts -> 表单组件
var defaultFormModeMap = map[string]string{
	TsTypeString:  FMInput,
	TsTypeNumber:  FMInputNumber,
	TsTypeBoolean: FMInputNumber,
	TsTypeArray:   FMInputDynamic,
	TsTypeTuple:   FMInputDynamic,
	TsTypeAny:     FMInput,
}

var defaultEditFields = map[string]bool{
	"id":           false,
	"created_dept": false,
	"created_by":   false,
	"created_at":   false,
	"updated_by":   false,
	"updated_at":   false,
	"deleted_by":   false,
	"deleted_at":   false,
}

var defaultEditSwitch = map[string]bool{
	"id":           false,
	"level":        false,
	"tree":         false,
	"created_by":   false,
	"updated_by":   false,
	"created_dept": false,
	"created_at":   false,
	"updated_at":   false,
	"deleted_at":   false,
}

var defaultListSwitch = map[string]bool{
	"level":      false,
	"tree":       false,
	"deleted_at": false,
}

var defaultExportSwitch = map[string]bool{
	"level":      false,
	"tree":       false,
	"deleted_at": false,
}

var defaultQuerySwitch = map[string]bool{
	"level":      false,
	"tree":       false,
	"deleted_at": false,
}

var defaultSort = map[string]bool{
	"id":   true,
	"sort": true,
}

var defaultTreeFields = []string{"pid", "level", "tree"}

// 默认表单验证映射 物理类型命名识别
var defaultFormRoleMap = map[string]string{
	"mobile":    FRPhone,
	"qq":        FRQq,
	"email":     FREmail,
	"id_card":   FRIdCard,
	"bank_card": FRBankCard,
	"password":  FRPassword,
	"pass":      FRPassword,
	"price":     FRAmount,
}

// 默认查询条件映射 go类型识别
var defaultWhereModeMap = map[string]string{
	GoTypeString:      WMLike,
	GoTypeDate:        WMEq,
	GoTypeDatetime:    WMEq,
	GoTypeInt:         WMEq,
	GoTypeUint:        WMEq,
	GoTypeInt64:       WMEq,
	GoTypeUint64:      WMEq,
	GoTypeIntSlice:    WMIn,
	GoTypeInt64Slice:  WMIn,
	GoTypeUint64Slice: WMIn,
	GoTypeFloat32:     WMEq,
	GoTypeFloat64:     WMEq,
	GoTypeBytes:       WMEq,
	GoTypeTime:        WMEq,
	GoTypeGTime:       WMEq,
	GoTypeJson:        WMJsonContains,
}

// IsIndexPK 是否是主键
func IsIndexPK(index string) bool {
	return gstr.ToUpper(index) == gstr.ToUpper(genconsts.GenCodesIndexPK)
}

// IsIndexUNI 是否是唯一索引
func IsIndexUNI(index string) bool {
	return gstr.ToUpper(index) == gstr.ToUpper(genconsts.GenCodesIndexUNI)
}

// setDefault 设置默认属性
func setDefault(field *genmodel.GenCodesColumnListModel) {
	setDefaultEdit(field)

	setDefaultFormMode(field)

	setDefaultFormRole(field)

	setDefaultDictType(field)

	setDefaultList(field)

	setDefaultExport(field)

	setDefaultQuery(field)

	setDefaultQueryWhere(field)

	setDefaultValue(field)

	if field.IsAllowNull == "NO" {
		field.Required = true
	}

	if IsIndexUNI(field.Index) {
		field.Unique = true
	}

	if df, ok := defaultSort[field.Name]; ok {
		field.IsSort = df
	}

	if field.Dc == "" {
		field.Dc = field.Name
	}
}

// setDefaultEdit 设置默认编辑
func setDefaultEdit(field *genmodel.GenCodesColumnListModel) {
	field.IsEdit = true

	if IsIndexPK(field.Index) {
		field.IsEdit = false
		return
	}

	if df, ok := defaultEditFields[field.Name]; ok {
		field.IsEdit = df
	}

	if df, ok := defaultEditSwitch[field.Name]; ok {
		field.IsEdit = df
	}
}

// setDefaultFormMode 设置默认表单组件
func setDefaultFormMode(field *genmodel.GenCodesColumnListModel) {
	field.FormMode = FMInput
	if df, ok := defaultFormModeMap[field.TsType]; ok {
		field.FormMode = df
	}

	if gstr.HasSuffix(field.GoName, "Status") && IsNumberType(field.GoType) {
		field.FormMode = FMSelect
		return
	}

	if field.GoName == "CreatedAt" {
		field.FormMode = FMDateRange
		return
	}

	if (field.GoName == "ProvinceId" || field.GoName == "CityId") && IsNumberType(field.GoType) {
		field.FormMode = FMCitySelector
		return
	}

	if field.DataType == "datetime" || field.DataType == "timestamp" || field.DataType == "timestamptz" {
		field.FormMode = FMTime
		return
	}

	if field.DataType == "date" {
		field.FormMode = FMDate
		return
	}

	if field.GoType == GoTypeString && field.Length >= 256 && field.Length <= 512 {
		field.FormMode = FMInputTextarea
		return
	}

	if field.GoType == GoTypeString && field.Length > 512 {
		field.FormMode = FMInputEditor
		return
	}
}

// setDefaultFormRole 设置默认表单验证
func setDefaultFormRole(field *genmodel.GenCodesColumnListModel) {
	field.FormRole = FRNone

	switch field.GoType {
	case GoTypeUint, GoTypeUint64:
		field.FormRole = FRNum
		return
	}

	if df, ok := defaultFormRoleMap[field.Name]; ok {
		field.FormRole = df
	}
}

// setDefaultDictType 设置默认字典类型
func setDefaultDictType(field *genmodel.GenCodesColumnListModel) {
	if gstr.HasSuffix(field.GoName, "Status") && IsNumberType(field.GoType) {
		field.DictType = "sys_common_status" // 默认系统状态ID
		return
	}
}

// setDefaultList 设置默认列表
func setDefaultList(field *genmodel.GenCodesColumnListModel) {
	field.IsList = true
	switch field.GoType {
	case GoTypeIntSlice, GoTypeInt64Slice, GoTypeUint64Slice, GoTypeBytes, GoTypeJson:
		field.IsList = false
		return
	}

	if field.Length >= 500 {
		field.IsList = false
		return
	}

	if df, ok := defaultListSwitch[field.Name]; ok {
		field.IsList = df
	}
}

// setDefaultExport 设置默认导出
func setDefaultExport(field *genmodel.GenCodesColumnListModel) {
	field.IsExport = true
	switch field.GoType {
	case GoTypeIntSlice, GoTypeInt64Slice, GoTypeUint64Slice, GoTypeBytes, GoTypeJson:
		field.IsExport = false
		return
	}

	if field.Length >= 500 {
		field.IsExport = false
		return
	}

	if df, ok := defaultExportSwitch[field.Name]; ok {
		field.IsExport = df
	}
}

// setDefaultQuery 设置默认查询
func setDefaultQuery(field *genmodel.GenCodesColumnListModel) {
	field.IsQuery = false

	if IsIndexPK(field.Index) {
		field.IsQuery = true
		return
	}

	if gstr.HasSuffix(field.GoName, "Status") && IsNumberType(field.GoType) {
		field.IsQuery = true
		return
	}

	if field.GoName == "CreatedAt" {
		field.IsQuery = true
		return
	}

	if df, ok := defaultQuerySwitch[field.Name]; ok {
		field.IsQuery = df
	}
}

// setDefaultQueryWhere 设置默认查询条件
func setDefaultQueryWhere(field *genmodel.GenCodesColumnListModel) {
	field.QueryWhere = WMEq

	if field.GoName == "CreatedAt" {
		field.QueryWhere = WMBetween
		return
	}

	if field.Length >= 500 {
		field.QueryWhere = WMLikeAll
		return
	}

	if df, ok := defaultWhereModeMap[field.GoType]; ok {
		field.QueryWhere = df
	}
}

// setDefaultValue 设置默认value
func setDefaultValue(field *genmodel.GenCodesColumnListModel) {
	var value interface{}
	if field.DefaultValue == nil {
		switch field.GoType {
		case GoTypeString, GoTypeBytes, GoTypeDate, GoTypeDatetime, GoTypeTime, GoTypeGTime:
			value = ""
		case GoTypeIntSlice, GoTypeInt64Slice, GoTypeUint64Slice, GoTypeJson:
			value = nil
		case GoTypeInt, GoTypeUint, GoTypeInt64, GoTypeUint64:
			value = 0
		case GoTypeBool:
			value = false
		}
	} else {
		value = genconsts.ConvType(field.DefaultValue, field.GoType)
	}

	// 时间类型不做默认值处理
	if field.GoType == GoTypeGTime {
		value = ""
	}

	field.DefaultValue = value
}
