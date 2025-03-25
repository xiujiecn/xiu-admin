package genconsts

// 生成类型值定义
const (
	GenCodesTypeCurd = 0
	GenCodesTypeTree = 1
)

// 生成类型标签映射
var GenCodesTypeNameMap = map[int]string{
	GenCodesTypeCurd: "增删改查列表",
	GenCodesTypeTree: "关系树列表",
}

// 生成类型配置映射
var GenCodesTypeConfMap = map[int]string{
	GenCodesTypeCurd: "crud",
	GenCodesTypeTree: "tree",
}

// 生成代码状态
const (
	GenCodesStatusOk   = "0" // 生成成功
	GenCodesStatusWait = "1" // 未生成
	GenCodesStatusFail = "2" // 生成失败
)

var GenCodesStatusNameMap = map[string]string{
	GenCodesStatusOk:   "生成成功",
	GenCodesStatusWait: "未生成",
	GenCodesStatusFail: "生成失败",
}

// 生成代码关联表方式
const (
	GenCodesJoinLeft  = 1 // 左关联
	GenCodesJoinRight = 2 // 右关联
	GenCodesJoinInner = 3 // 内关联
)

var GenCodesJoinNameMap = map[int]string{
	GenCodesJoinLeft:  "左关联",
	GenCodesJoinRight: "右关联",
	GenCodesJoinInner: "内关联",
}

var GenCodesJoinLinkMap = map[int]string{
	GenCodesJoinLeft:  "LeftJoin",
	GenCodesJoinRight: "RightJoin",
	GenCodesJoinInner: "InnerJoin",
}

// 生成代码的生成方式
const (
	GenCodesBuildMethCreate = 1 // 创建
	GenCodesBuildMethCover  = 2 // 覆盖
	GenCodesBuildMethSkip   = 3 // 跳过
	GenCodesBuildIgnore     = 4 // 不生成
)

var GenCodesBuildMethNameMap = map[int]string{
	GenCodesBuildMethCreate: "创建文件",
	GenCodesBuildMethCover:  "强制覆盖",
	GenCodesBuildMethSkip:   "已存在跳过",
	GenCodesBuildIgnore:     "不生成",
}

const (
	GenCodesIndexPK  = "PRI" // 主键索引
	GenCodesIndexUNI = "UNI" // 唯一索引
)

const (
	GenCodesTreeStyleTypeNormal = 1 // 普通树表格
	GenCodesTreeStyleTypeOption = 2 // 选项式树表
)

var GenCodesTreeStyleTypeNameMap = map[int]string{
	GenCodesTreeStyleTypeNormal: "普通树表格",
	GenCodesTreeStyleTypeOption: "选项式树表",
}
