// Package consts
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package consts

// 内存缓存KEY常量 注明 key组成和 过期时间 清理方式
const (
	MemCacheDeptName                   = "dept_name_%d"                       // 部门名称缓存 部门id:name 24小时过期
	MemCacheUserInfo                   = "user_info_%d"                       // 用户信息缓存 用户id:miniinfo 24小时过期 用户修改是清理
	MemCacheUserAccessCodeList         = "user_access_code_list_%d"           // 用户访问码列表缓存 用户id:access_code_list 24小时过期 用户退出清理
	MemCacheUserRoleDataAccessCodeList = "user_role_data_access_code_list_%d" // 用户角色数据访问码列表缓存 用户id:role_id:data_access_code_list 24小时过期 用户退出清理
	MemCacheSystemConfig               = "system_config_%s"                   // 系统配置缓存 租户ID:map[key]value  24小时过期
)
