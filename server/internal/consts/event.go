// Package consts
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package consts

const (
	EventKeyServerClose = "server_close" // 服务关闭  args=[]
)
const (
	EventKeyUserLogin  = "user_login"  // 用户登录  args=[uid,]
	EventKeyUserLogout = "user_logout" // 用户退出  args=[uid,]
	EventKeyUserCreate = "user_create" // 用户创建  args=[uid,]
	EventKeyUserDelete = "user_delete" // 用户删除  args=[uidArray,]
	EventKeyUserUpdate = "user_update" // 用户更新  args=[uid,]
)
const (
	EventKeySysConfigUpdate = "sys_config_update" // 系统配置创建更新删除  args=[tenantId,key?]
)
const (
	EventKeySysDeptUpdate = "sys_dept_update" // 系统部门创建更新删除  args=[deptId,]
)
