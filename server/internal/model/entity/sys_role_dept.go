// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysRoleDept is the golang structure for table sys_role_dept.
type SysRoleDept struct {
	TenantId string `json:"tenantId" orm:"tenant_id" description:"租户编号"`
	RoleId   int64  `json:"roleId" orm:"role_id" description:"角色ID"`
	DeptId   int64  `json:"deptId" orm:"dept_id" description:"部门ID"`
}
