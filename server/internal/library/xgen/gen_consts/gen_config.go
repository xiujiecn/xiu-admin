// package genconsts
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package genconsts

import "github.com/gogf/gf/v2/util/gconv"

// 配置数据类型
const (
	ConfigTypeString      = "string"
	ConfigTypeInt         = "int"
	ConfigTypeInt8        = "int8"
	ConfigTypeInt16       = "int16"
	ConfigTypeInt32       = "int32"
	ConfigTypeInt64       = "int64"
	ConfigTypeUint        = "uint"
	ConfigTypeUint8       = "uint8"
	ConfigTypeUint16      = "uint16"
	ConfigTypeUint32      = "uint32"
	ConfigTypeUint64      = "uint64"
	ConfigTypeFloat32     = "float32"
	ConfigTypeFloat64     = "float64"
	ConfigTypeBool        = "bool"
	ConfigTypeDate        = "date"
	ConfigTypeDateTime    = "datetime"
	ConfigTypeSliceString = "[]string"
	ConfigTypeSliceInt    = "[]int"
	ConfigTypeSliceInt64  = "[]int64"
)

var ConfigTypes = []string{ConfigTypeString,
	ConfigTypeInt, ConfigTypeInt8, ConfigTypeInt16, ConfigTypeInt32, ConfigTypeInt64,
	ConfigTypeUint, ConfigTypeUint8, ConfigTypeUint16, ConfigTypeUint32, ConfigTypeUint64,
	ConfigTypeFloat32, ConfigTypeFloat64,
	ConfigTypeBool,
	ConfigTypeDate, ConfigTypeDateTime,
	ConfigTypeSliceString, ConfigTypeSliceInt, ConfigTypeSliceInt64,
}

// ConvType 类型转换
func ConvType(val interface{}, t string) interface{} {
	switch t {
	case ConfigTypeString:
		val = gconv.String(val)
	case ConfigTypeInt:
		val = gconv.Int(val)
	case ConfigTypeInt8:
		val = gconv.Int8(val)
	case ConfigTypeInt16:
		val = gconv.Int16(val)
	case ConfigTypeInt32:
		val = gconv.Int32(val)
	case ConfigTypeInt64:
		val = gconv.Int64(val)
	case ConfigTypeUint:
		val = gconv.Uint(val)
	case ConfigTypeUint8:
		val = gconv.Uint8(val)
	case ConfigTypeUint16:
		val = gconv.Uint16(val)
	case ConfigTypeUint32:
		val = gconv.Uint32(val)
	case ConfigTypeUint64:
		val = gconv.Uint64(val)
	case ConfigTypeFloat32:
		val = gconv.Float32(val)
	case ConfigTypeFloat64:
		val = gconv.Float64(val)
	case ConfigTypeBool:
		val = gconv.Bool(val)
	case ConfigTypeDate:
		val = gconv.Time(val, "Y-m-d")
	case ConfigTypeDateTime:
		val = gconv.Time(val, "Y-m-d H:i:s")
	case ConfigTypeSliceInt:
		val = gconv.SliceInt(val)
	case ConfigTypeSliceInt64:
		val = gconv.SliceInt64(val)
	case ConfigTypeSliceString:
		val = gconv.SliceStr(val)
	default:
		val = gconv.String(val)
	}
	return val
}
