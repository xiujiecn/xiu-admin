package model

// IotChannelStatusCache 资源通道表状态缓存
type IotChannelStatusCache struct {
	ChannelId   int64 `json:"channelId"`
	ConnectTime int64 `json:"connectTime"` // 连接时间戳(秒)
}

// 数据库变动更新广播消息 obj{table_name,pk,type:add/edit/del,data:{}}
type BroadcastDbChgCache struct {
	TableName string                 `json:"table_name"` // 表名
	PK        []int64                `json:"pk"`         // 主键
	Type      string                 `json:"type"`       // 类型:add/edit/del
	Data      map[string]interface{} `json:"data"`       // 变动数据,只需要包含主要字段数据
}
