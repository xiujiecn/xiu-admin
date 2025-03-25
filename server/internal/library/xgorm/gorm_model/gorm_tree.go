package gormmodel

// DefaultTree 默认树表字段
type DefaultTree struct {
	Id    int64  `json:"id"             description:"ID"`
	Pid   int64  `json:"pid"            description:"父ID"`
	Level int    `json:"level"          description:"关系树等级"`
	Tree  string `json:"tree"           description:"关系树"`
}
