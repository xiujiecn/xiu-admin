package genview

import (
	genmodel "xiujieadmin/internal/library/xgen/gen_model"

	"github.com/gogf/gf/v2/text/gstr"
)

// 表单组件
const (
	FMInput          = "Input"          // 文本输入
	FMInputNumber    = "InputNumber"    // 数字输入
	FMInputTextarea  = "InputTextarea"  // 文本域
	FMInputEditor    = "InputEditor"    // 富文本
	FMInputDynamic   = "InputDynamic"   // 动态键值对
	FMDate           = "Date"           // 日期选择(Y-M-D)
	FMDateRange      = "DateRange"      // 日期范围选择
	FMTime           = "Time"           // 时间选择(Y-M-D H:i:s)
	FMTimeRange      = "TimeRange"      // 时间范围选择
	FMRadio          = "Radio"          // 单选按钮
	FMCheckbox       = "Checkbox"       // 复选按钮
	FMSelect         = "Select"         // 单选下拉框
	FMSelectMultiple = "SelectMultiple" // 多选下拉框
	FMTreeSelect     = "TreeSelect"     // 树型选择
	FMCascader       = "Cascader"       // 级联选择
	FMUploadImage    = "UploadImage"    // 单图上传
	FMUploadImages   = "UploadImages"   // 多图上传
	FMUploadFile     = "UploadFile"     // 单文件上传
	FMUploadFiles    = "UploadFiles"    // 多文件上传
	FMSwitch         = "Switch"         // 开关
	FMRate           = "Rate"           // 评分
	FMCitySelector   = "CitySelector"   // 省市区选择
	FMPidTreeSelect  = "PidTreeSelect"  // 树型上级选择，树表生成专用
)

var FMs = []string{
	FMInput, FMInputNumber, FMInputTextarea, FMInputEditor, FMInputDynamic,
	FMDate, FMDateRange, FMTime, FMTimeRange,
	FMRadio, FMCheckbox, FMSelect, FMSelectMultiple, FMTreeSelect, FMCascader,
	FMUploadImage, FMUploadImages, FMUploadFile, FMUploadFiles,
	FMSwitch,
	FMRate,
	FMCitySelector,
}

var FMMap = map[string]string{
	FMInput:          "文本输入",
	FMInputNumber:    "数字输入",
	FMInputTextarea:  "文本域",
	FMInputEditor:    "富文本",
	FMInputDynamic:   "动态键值对",
	FMDate:           "日期选择(Y-M-D)",
	FMDateRange:      "日期范围选择",
	FMTime:           "时间选择(Y-M-D H:i:s)",
	FMTimeRange:      "时间范围选择",
	FMRadio:          "单选按钮",
	FMCheckbox:       "复选按钮",
	FMSelect:         "单选下拉框",
	FMSelectMultiple: "多选下拉框",
	FMTreeSelect:     "树型选择",
	FMCascader:       "级联选择",
	FMUploadImage:    "单图上传",
	FMUploadImages:   "多图上传",
	FMUploadFile:     "单文件上传",
	FMUploadFiles:    "多文件上传",
	FMSwitch:         "开关",
	FMRate:           "评分",
	FMCitySelector:   "省市区选择",
}

// 表单验证
const (
	FRNone       = "none"
	FRIp         = "ip"
	FRPercentage = "percentage"
	FRTel        = "tel"
	FRPhone      = "phone"
	FRQq         = "qq"
	FREmail      = "email"
	FRIdCard     = "idCard"
	FRNum        = "num"
	FRBankCard   = "bankCard"
	FRWeibo      = "weibo"
	FRUserName   = "userName"
	FRAccount    = "account"
	FRPassword   = "password"
	FRAmount     = "amount"
)

var FRMap = map[string]string{
	FRNone:       "不验证",
	FRIp:         "Ipv4或Ipv6",
	FRPercentage: "0-100百分比",
	FRTel:        "固话格式",
	FRPhone:      "手机号",
	FRQq:         "QQ号码",
	FREmail:      "邮箱",
	FRIdCard:     "身份证",
	FRNum:        "非零正整数",
	FRBankCard:   "银行卡",
	FRWeibo:      "微博号",
	FRUserName:   "用户名",
	FRAccount:    "账号",
	FRPassword:   "密码",
	FRAmount:     "金额",
}

// 查询条件
const (
	WMEq           = "="                       // =
	WMNeq          = "!="                      // !=
	WMGt           = ">"                       // >
	WMGte          = ">="                      // >=
	WMLt           = "<"                       // <
	WMLte          = "<="                      // <=
	WMIn           = "IN"                      // IN (...)
	WMNotIn        = "NOT IN"                  // NOT IN (...)
	WMBetween      = "BETWEEN"                 // BETWEEN
	WMNotBetween   = "NOT BETWEEN"             // NOT BETWEEN
	WMLike         = "LIKE"                    // LIKE
	WMLikeAll      = "LIKE %...%"              // LIKE %...%
	WMNotLike      = "NOT LIKE"                // NOT LIKE
	WMJsonContains = "JSON_CONTAINS(doc, val)" // JSON_CONTAINS(json_doc, val[, path]) // 判断是否包含某个json值
)

var WMs = []string{WMEq,
	WMNeq, WMGt, WMGte, WMLt, WMLte,
	WMIn, WMNotIn,
	WMBetween, WMNotBetween,
	WMLike, WMLikeAll, WMNotLike,
	WMJsonContains,
}

// 表格列的排序方式
const (
	TableAlignLeft   = "left"
	TableAlignRight  = "right"
	TableAlignCenter = "center"
)

var TableAligns = []string{TableAlignLeft, TableAlignRight, TableAlignCenter}

var TableAlignMap = map[string]string{
	TableAlignLeft:   "居左",
	TableAlignRight:  "居右",
	TableAlignCenter: "居中",
}

// 是否是数字类型
func IsNumberType(goType string) bool {
	switch goType {
	case GoTypeInt, GoTypeUint, GoTypeInt64, GoTypeUint64:
		return true
	case GoTypeFloat32, GoTypeFloat64:
		return true
	}
	return false
}

// IsSelectFormModel 是否是选择器组件
func IsSelectFormModel(FM string) bool {
	switch FM {
	case FMRadio, FMCheckbox, FMSelect, FMSelectMultiple, FMCitySelector, FMTreeSelect, FMCascader:
		return true
	}
	return false
}

// 字段映射关系

// go类型
const (
	GoTypeString      = "string"
	GoTypeDate        = "date"
	GoTypeDatetime    = "datetime"
	GoTypeInt         = "int"
	GoTypeUint        = "uint"
	GoTypeInt64       = "int64"
	GoTypeUint64      = "uint64"
	GoTypeIntSlice    = "[]int"
	GoTypeInt64Slice  = "[]int64"
	GoTypeUint64Slice = "[]uint64"
	GoTypeFloat32     = "float32"
	GoTypeFloat64     = "float64"
	GoTypeBytes       = "[]byte"
	GoTypeBool        = "bool"
	GoTypeTime        = "time.Time"
	GoTypeGTime       = "*gtime.Time"
	GoTypeJson        = "*gjson.Json"
)

var GoTypeNameMap = map[string]string{
	GoTypeString:      GoTypeString,
	GoTypeDate:        GoTypeDate,
	GoTypeDatetime:    GoTypeDatetime,
	GoTypeInt:         GoTypeInt,
	GoTypeUint:        GoTypeUint,
	GoTypeInt64:       GoTypeInt64,
	GoTypeUint64:      GoTypeUint64,
	GoTypeIntSlice:    GoTypeIntSlice,
	GoTypeInt64Slice:  GoTypeInt64Slice,
	GoTypeUint64Slice: GoTypeUint64Slice,
	GoTypeFloat32:     GoTypeFloat32,
	GoTypeFloat64:     GoTypeFloat64,
	GoTypeBytes:       GoTypeBytes,
	GoTypeBool:        GoTypeBool,
	GoTypeTime:        GoTypeTime,
	GoTypeGTime:       GoTypeGTime,
	GoTypeJson:        GoTypeJson,
}

// ts类型
const (
	TsTypeString  = "string"
	TsTypeNumber  = "number"
	TsTypeBoolean = "boolean"
	TsTypeArray   = "array"
	TsTypeTuple   = "tuple"
	TsTypeAny     = "any"
)

var TsTypeNameMap = map[string]string{
	TsTypeString:  TsTypeString,
	TsTypeNumber:  TsTypeNumber,
	TsTypeBoolean: TsTypeBoolean,
	TsTypeArray:   TsTypeArray,
	TsTypeTuple:   TsTypeTuple,
	TsTypeAny:     TsTypeAny,
}

// ShiftMap Go -> Ts 类型转换
var ShiftMap = map[string]string{
	GoTypeString:      TsTypeString,
	GoTypeDate:        TsTypeString,
	GoTypeDatetime:    TsTypeString,
	GoTypeInt:         TsTypeNumber,
	GoTypeUint:        TsTypeNumber,
	GoTypeInt64:       TsTypeNumber,
	GoTypeUint64:      TsTypeNumber,
	GoTypeIntSlice:    TsTypeArray,
	GoTypeInt64Slice:  TsTypeArray,
	GoTypeUint64Slice: TsTypeArray,
	GoTypeFloat32:     TsTypeNumber,
	GoTypeFloat64:     TsTypeNumber,
	GoTypeBytes:       TsTypeString,
	GoTypeBool:        TsTypeBoolean,
	GoTypeTime:        TsTypeString,
	GoTypeGTime:       TsTypeString,
	GoTypeJson:        TsTypeAny,
}

func HasColumn(fields []*genmodel.GenCodesColumnListModel, column string) bool {
	for _, field := range fields {
		if field.GoName == column {
			return true
		}
	}
	return false
}

func HasColumnWithFormMode(fields []*genmodel.GenCodesColumnListModel, formMode string) bool {
	for _, field := range fields {
		if field.FormMode == formMode {
			return true
		}
	}
	return false
}

func HasMaxSort(fields []*genmodel.GenCodesColumnListModel) bool {
	return HasColumn(fields, "Sort")
}

func HasStatus(headOps []string, fields []*genmodel.GenCodesColumnListModel) bool {
	if !gstr.InArray(headOps, "status") {
		return false
	}
	return HasColumn(fields, "Status")
}

func HasSwitch(fields []*genmodel.GenCodesColumnListModel) bool {
	return HasColumnWithFormMode(fields, FMSwitch)
}

func HasHookMemberSummary(fields []*genmodel.GenCodesColumnListModel) bool {
	for _, field := range fields {
		if IsMemberSummaryField(field.Name) {
			if field.IsList {
				return true
			}
		}
	}
	return false
}

func HasQueryMemberSummary(fields []*genmodel.GenCodesColumnListModel) bool {
	for _, field := range fields {
		if IsMemberSummaryField(field.Name) {
			if field.IsQuery {
				return true
			}
		}
	}
	return false
}

func IsMemberSummaryField(name string) bool {
	switch name {
	case "created_by", "updated_by", "deleted_by":
		return true
	}
	return false
}

// ReviseFields 校正字段值，兼容版本升级前的老数据格式
func ReviseFields(fields []*genmodel.GenCodesColumnListModel) []*genmodel.GenCodesColumnListModel {
	for _, field := range fields {
		if !gstr.InArray(TableAligns, field.Align) {
			field.Align = TableAlignLeft
		}

		if field.Width < 1 {
			field.Width = -1
		}
		if field.Width > 2000 {
			field.Width = 2000
		}

		if field.FormGridSpan < 1 {
			field.FormGridSpan = 1
		}
	}
	return fields
}
