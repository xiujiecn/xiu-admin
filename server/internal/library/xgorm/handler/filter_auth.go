// package handler
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package handler

import (
	"context"
	"fmt"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/model/entity"
	"xiuadmin/utility/convert"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// FilterAuth 过滤数据权限
// 通过上下文中的用户角色权限和表中是否含有需要过滤的字段附加查询条件
func FilterAuth(m *gdb.Model) *gdb.Model {
	var (
		ctx    = m.GetCtx()
		fields = convert.EscapeFieldsToSlice(m.GetFieldsStr())
	)

	// 超管拥有全部权限
	if contexts.IsSuperAdmin(ctx) {
		return m
	}

	// 获取用户信息
	user := contexts.GetUser(ctx)
	if user == nil {
		return m
	}
	if len(user.Roles) == 0 {
		return m
	}

	deptIds := make([]int64, 0)
	userIds := make([]int64, 0)
	currentRoleId := ""
	currentRoleDataScope := ""
	currentRoleIdIf := contexts.GetDataValue(ctx, "currentRoleId")
	if currentRoleIdIf != nil {
		currentRoleId = currentRoleIdIf.(string)
	}
	currentRoleDataScopeIf := contexts.GetDataValue(ctx, "currentRoleDataScope")
	if currentRoleDataScopeIf != nil {
		currentRoleDataScope = currentRoleDataScopeIf.(string)
	}
	if len(currentRoleDataScope) > 0 && len(currentRoleId) > 0 {
		dataScope := currentRoleDataScope
		roleId := gconv.Int64(currentRoleId)
		g.Log().Infof(ctx, "uid:%d, currentRoleId: %d, currentRoleDataScope: %s", user.ID, roleId, dataScope)
		if dataScope == consts.SysRoleDataScopeAll { // 角色数据范围: 1全部数据权限
			return m
		} else if dataScope == consts.SysRoleDataScopeCustom { // 角色数据范围: 2自定数据权限
			var roleDepts []*entity.SysRoleDept
			err := g.Model("sys_role_dept").Where("role_id", roleId).Scan(&roleDepts)
			if err != nil {
				g.Log().Panicf(ctx, "failed to get role dept data err:%+v", err)
			}
			for _, roleDept := range roleDepts {
				deptIds = append(deptIds, roleDept.DeptId)
			}
		} else if dataScope == consts.SysRoleDataScopeDept { // 角色数据范围: 3本部门数据权限
			deptIds = append(deptIds, user.DeptId)
		} else if dataScope == consts.SysRoleDataScopeDeptAndBelow { // 角色数据范围: 4本部门及以下数据权限
			deptIds = append(deptIds, GetDeptAndSub(ctx, user.DeptId)...)
		} else if dataScope == consts.SysRoleDataScopePersonal { // 角色数据范围: 5仅本人数据权限
			userIds = append(userIds, user.ID)
		} else if dataScope == consts.SysRoleDataScopeDeptAndBelowOrPersonal { // 角色数据范围: 6本部门及以下或本人数据权限
			deptIds = append(deptIds, GetDeptAndSub(ctx, user.DeptId)...)
			userIds = append(userIds, user.ID)
		} else if dataScope == consts.SysRoleDataScopeOrgAndNextLevel { // 角色数据范围: 7本组织及本组织下一级数据权限
			deptIds = append(deptIds, GetOrgAndNextLevelDeptIds(ctx, user.DeptId)...)
		}
	} else {
		for _, role := range user.Roles {
			dataScope := role.DataScope
			roleId := role.RoleId
			if dataScope == consts.SysRoleDataScopeAll { // 角色数据范围: 1全部数据权限
				return m
			} else if dataScope == consts.SysRoleDataScopeCustom { // 角色数据范围: 2自定数据权限
				var roleDepts []*entity.SysRoleDept
				err := g.Model("sys_role_dept").Where("role_id", roleId).Scan(&roleDepts)
				if err != nil {
					g.Log().Panicf(ctx, "failed to get role dept data err:%+v", err)
				}
				for _, roleDept := range roleDepts {
					deptIds = append(deptIds, roleDept.DeptId)
				}
			} else if dataScope == consts.SysRoleDataScopeDept { // 角色数据范围: 3本部门数据权限
				deptIds = append(deptIds, user.DeptId)
			} else if dataScope == consts.SysRoleDataScopeDeptAndBelow { // 角色数据范围: 4本部门及以下数据权限
				deptIds = append(deptIds, GetDeptAndSub(ctx, user.DeptId)...)
			} else if dataScope == consts.SysRoleDataScopePersonal { // 角色数据范围: 5仅本人数据权限
				userIds = append(userIds, user.ID)
			} else if dataScope == consts.SysRoleDataScopeDeptAndBelowOrPersonal { // 角色数据范围: 6本部门及以下或本人数据权限
				deptIds = append(deptIds, GetDeptAndSub(ctx, user.DeptId)...)
				userIds = append(userIds, user.ID)
			} else if dataScope == consts.SysRoleDataScopeOrgAndNextLevel { // 角色数据范围: 7本组织及本组织下一级数据权限
				deptIds = append(deptIds, GetOrgAndNextLevelDeptIds(ctx, user.DeptId)...)
			}
		}
	}

	deptFilterField := ""
	if gstr.InArray(fields, "dept_id") {
		deptFilterField = "dept_id"
	} else if gstr.InArray(fields, "created_dept") {
		deptFilterField = "created_dept"
	}
	userFilterField := ""
	if gstr.InArray(fields, "user_id") {
		userFilterField = "user_id"
	} else if gstr.InArray(fields, "created_by") {
		userFilterField = "created_by"
	}

	if len(deptIds) > 0 && len(userIds) > 0 {
		if deptFilterField != "" && userFilterField != "" {
			m = m.Where(deptFilterField+" IN (?) OR "+userFilterField+" IN (?)", deptIds, userIds)
		} else if deptFilterField != "" {
			m = m.WhereIn(deptFilterField, deptIds)
		} else if userFilterField != "" {
			m = m.WhereIn(userFilterField, userIds)
		}
	} else if len(deptIds) > 0 && deptFilterField != "" {
		m = m.WhereIn(deptFilterField, deptIds)
	} else if len(userIds) > 0 && userFilterField != "" {
		m = m.WhereIn(userFilterField, userIds)
	}
	return m
}

// GetParentCompanyDeptId 获取用户所属上级公司组织ID
func GetParentCompanyDeptId(ctx context.Context, deptId int64) int64 {
	if deptId <= 0 {
		return deptId
	}
	var dept entity.SysDept
	err := g.Model("sys_dept").Where("dept_id", deptId).Scan(&dept)
	if err != nil || dept.DeptId == 0 {
		return deptId
	}
	if dept.ParentId == 0 || dept.DeptType == 1 {
		return dept.DeptId
	}
	return GetParentCompanyDeptId(ctx, dept.ParentId)
}

// GetDeptDirectSub 获取指定部门的直接下级部门ID
func GetDeptDirectSub(ctx context.Context, deptId int64) (ids []int64) {
	array, err := g.Model("sys_dept").
		Where("parent_id", deptId).
		Fields("dept_id").
		Array()
	if err != nil {
		g.Log().Panicf(ctx, "GetDeptDirectSub err:%+v", err)
		return
	}
	for _, v := range array {
		ids = append(ids, v.Int64())
	}
	return
}

// GetOrgAndNextLevelDeptIds 获取本组织及本组织下一级数据权限范围：所属公司组织及直接下级
func GetOrgAndNextLevelDeptIds(ctx context.Context, deptId int64) (ids []int64) {
	orgDeptId := GetParentCompanyDeptId(ctx, deptId)
	ids = append(ids, orgDeptId)
	ids = append(ids, GetDeptDirectSub(ctx, orgDeptId)...)
	return
}

// GetDeptAndSub 获取指定部门的所有下级，含本部门
func GetDeptAndSub(ctx context.Context, deptId int64) (ids []int64) {
	array, err := g.Model("sys_dept").
		WhereLike("ancestors", "%,"+fmt.Sprintf("%d", deptId)+",%").
		Fields("dept_id").
		Array()
	if err != nil {
		g.Log().Panicf(ctx, "GetDeptAndSub err:%+v", err)
		return
	}

	for _, v := range array {
		ids = append(ids, v.Int64())
	}

	ids = append(ids, deptId)
	return
}
