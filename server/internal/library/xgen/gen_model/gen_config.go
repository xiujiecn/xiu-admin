package genmodel

import (
	genconsts "xiujieadmin/internal/library/xgen/gen_consts"

	"github.com/gogf/gf/v2/util/gconv"
)

// 模板配置 config/config.yaml  xgen.application.crud.templates
type GenCodesConfigCrudTemplate struct {
	Group          string `name:"group" json:"group"`
	IsAddon        bool   `name:"isAddon" json:"isAddon"`
	MasterPackage  string `name:"masterPackage" json:"masterPackage"`
	TemplatePath   string `name:"templatePath" json:"templatePath"`
	ApiPath        string `name:"apiPath" json:"apiPath"`
	InputPath      string `name:"inputPath" json:"inputPath"`
	ControllerPath string `name:"controllerPath" json:"controllerPath"`
	LogicPath      string `name:"logicPath" json:"logicPath"`
	RouterPath     string `name:"routerPath" json:"routerPath"`
	SqlPath        string `name:"sqlPath" json:"sqlPath"`
	WebApiPath     string `name:"webApiPath" json:"webApiPath"`
	WebViewsPath   string `name:"webViewsPath" json:"webViewsPath"`
}
type GenCodesConfigTreeTemplate struct {
	Group        string `json:"group"`
	TemplatePath string `json:"templatePath"`
}

type GenCodesConfigBuildAddon struct {
	SrcPath      string `json:"srcPath"`
	WebApiPath   string `json:"webApiPath"`
	WebViewsPath string `json:"webViewsPath"`
}

// 工具生成配置 hack/config.yaml gfcli.gen.dao
type GenDaoConfig struct {
	Group        string `name:"group" json:"group"`
	Link         string `name:"link" json:"link"`
	DbName       string `name:"dbName" json:"dbName"`
	TableName    string `name:"tableName" json:"tableName"`
	JsonCase     string `name:"jsonCase" json:"jsonCase"`
	StdTime      bool   `name:"stdTime" json:"stdTime"`
	GJsonSupport bool   `name:"gJsonSupport" json:"gJsonSupport"`
}

// 工具生成配置 hack/config.yaml gfcli.gen.service
type GenServiceConfig struct {
	SrcFolder       string   `name:"srcFolder"`
	DstFolder       string   `name:"dstFolder"`
	DstFileNameCase string   `name:"dstFileNameCase"`
	WatchFile       string   `name:"watchFile" `
	StPattern       string   `name:"stPattern" `
	Packages        []string `name:"packages" `
	ImportPrefix    string   `name:"importPrefix" `
	Clear           bool     `name:"clear" `
}

// 模板配置 config/config.yaml  xgen
type GenCodesConfig struct {
	AllowedIPs  []string `json:"allowedIPs"`
	Application struct {
		Crud struct {
			Templates []*GenCodesConfigCrudTemplate `json:"templates"`
		} `json:"crud"`
		Tree struct {
			Templates []*GenCodesConfigTreeTemplate `json:"templates"`
		} `json:"tree"`
	} `json:"application"`
	Delimiters    []string                  `json:"delimiters"`
	DevPath       string                    `json:"devPath"`
	DisableTables []string                  `json:"disableTables"`
	SelectDbs     []string                  `json:"selectDbs"`
	Addon         *GenCodesConfigBuildAddon `json:"addon"`
}

// ConvType 类型转换
func ConvType(val interface{}, t string) interface{} {
	switch t {
	case genconsts.ConfigTypeString:
		val = gconv.String(val)
	case genconsts.ConfigTypeInt:
		val = gconv.Int(val)
	case genconsts.ConfigTypeInt8:
		val = gconv.Int8(val)
	case genconsts.ConfigTypeInt16:
		val = gconv.Int16(val)
	case genconsts.ConfigTypeInt32:
		val = gconv.Int32(val)
	case genconsts.ConfigTypeInt64:
		val = gconv.Int64(val)
	case genconsts.ConfigTypeUint:
		val = gconv.Uint(val)
	case genconsts.ConfigTypeUint8:
		val = gconv.Uint8(val)
	case genconsts.ConfigTypeUint16:
		val = gconv.Uint16(val)
	case genconsts.ConfigTypeUint32:
		val = gconv.Uint32(val)
	case genconsts.ConfigTypeUint64:
		val = gconv.Uint64(val)
	case genconsts.ConfigTypeFloat32:
		val = gconv.Float32(val)
	case genconsts.ConfigTypeFloat64:
		val = gconv.Float64(val)
	case genconsts.ConfigTypeBool:
		val = gconv.Bool(val)
	case genconsts.ConfigTypeDate:
		val = gconv.Time(val, "Y-m-d")
	case genconsts.ConfigTypeDateTime:
		val = gconv.Time(val, "Y-m-d H:i:s")
	case genconsts.ConfigTypeSliceInt:
		val = gconv.SliceInt(val)
	case genconsts.ConfigTypeSliceInt64:
		val = gconv.SliceInt64(val)
	case genconsts.ConfigTypeSliceString:
		val = gconv.SliceStr(val)
	default:
		val = gconv.String(val)
	}
	return val
}
