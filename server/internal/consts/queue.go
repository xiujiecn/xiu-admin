// Package consts
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package consts

const (
	QueueSysOptLog = "queue:sys_opt_log"
)

// 数据库变动类型
const (
	DbChgTypeAdd  = "add"  // 数据库变动新增
	DbChgTypeEdit = "edit" // 数据库变动编辑
	DbChgTypeDel  = "del"  // 数据库变动删除
)
const (
	BroadcastDbChg = "broadcast:db:chg" // 数据库配置表变动更新广播消息 obj{table_name,pk,type:add/edit/del,data:{主要字段数据}}
)
