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
	EventKeyDBSysConfigUpdate = "db.sys_config_update" // 系统配置创建更新删除  args=[configId, tenantId,key?]
	EventKeyDBSysConfigCreate = "db.sys_config_create" // 系统配置创建  args=[configId, tenantId,key?]
	EventKeyDBSysConfigDelete = "db.sys_config_delete" // 系统配置删除  args=[configId, tenantId]
)

const (
	EventKeyDBSysDeptCreate = "db.sys_dept_create" // 系统部门创建  args=[deptId,]
	EventKeyDBSysDeptUpdate = "db.sys_dept_update" // 系统部门更新  args=[deptId,]
	EventKeyDBSysDeptDelete = "db.sys_dept_delete" // 系统部门删除  args=[deptId[],]
)

const (
	EventKeyDBSysTenantCreate = "db.sys_tenant_create" // 租户创建  args=[pk,]
	EventKeyDBSysTenantUpdate = "db.sys_tenant_update" // 租户更新  args=[pk,]
	EventKeyDBSysTenantDelete = "db.sys_tenant_delete" // 租户删除  args=[pk[],]

	EventKeyDBSysRoleCreate = "db.sys_role_create" // 角色创建  args=[pk,]
	EventKeyDBSysRoleUpdate = "db.sys_role_update" // 角色更新  args=[pk,]
	EventKeyDBSysRoleDelete = "db.sys_role_delete" // 角色删除  args=[pk[],]

	EventKeyDBSysMenuCreate = "db.sys_menu_create" // 菜单创建  args=[pk,]
	EventKeyDBSysMenuUpdate = "db.sys_menu_update" // 菜单更新  args=[pk,]
	EventKeyDBSysMenuDelete = "db.sys_menu_delete" // 菜单删除  args=[pk[],]

	EventKeyDBSysPostCreate = "db.sys_post_create" // 岗位创建  args=[pk,]
	EventKeyDBSysPostUpdate = "db.sys_post_update" // 岗位更新  args=[pk,]
	EventKeyDBSysPostDelete = "db.sys_post_delete" // 岗位删除  args=[pk[],]

	EventKeyDBSysDictDataCreate = "db.sys_dict_data_create" // 字典数据创建  args=[pk,]
	EventKeyDBSysDictDataUpdate = "db.sys_dict_data_update" // 字典数据更新  args=[pk,]
	EventKeyDBSysDictDataDelete = "db.sys_dict_data_delete" // 字典数据删除  args=[pk[],]

	EventKeyDBSysDictTypeCreate = "db.sys_dict_type_create" // 字典类型创建  args=[pk,]
	EventKeyDBSysDictTypeUpdate = "db.sys_dict_type_update" // 字典类型更新  args=[pk,]
	EventKeyDBSysDictTypeDelete = "db.sys_dict_type_delete" // 字典类型删除  args=[pk[],]
)
