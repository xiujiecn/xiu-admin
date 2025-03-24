package consts

var (
	//  默认系统租客编号,禁止修改编号
	DefaultSystemTenantCode = "000000"
	SuperAdminRoleId        = int64(1)
)

var (
	// 菜单状态: 0正常 1停用
	SysMenuStatusNormal  = "0"
	SysMenuStatusDisable = "1"
)

var (
	// 菜单类型: M目录 C菜单 F按钮
	SysMenuTypeDir  = "M"
	SysMenuTypeMenu = "C"
	SysMenuTypeBtn  = "F"
)

var (
	// 菜单可见: 0显示 1隐藏
	SysMenuVisibleShow = "0"
	SysMenuVisibleHide = "1"
)

var (
	// 菜单是否为外链: 0是 1否
	SysMenuIsFrameNo  = 1
	SysMenuIsFrameYes = 0
)

var (
	// 用户状态: 0正常 1停用
	SysUserStatusNormal  = "0"
	SysUserStatusDisable = "1"
	SysUserStatusLocked  = "2"
	SysUserStatusExpired = "3"
	SysUserStatusDeleted = "4"
)

var (
	// 角色数据范围: 1全部数据权限 2自定数据权限 3本部门数据权限 4本部门及以下数据权限 5仅本人数据权限 6本部门及以下或本人数据权限
	SysRoleDataScopeAll                    = "1"
	SysRoleDataScopeCustom                 = "2"
	SysRoleDataScopeDept                   = "3"
	SysRoleDataScopeDeptAndBelow           = "4"
	SysRoleDataScopePersonal               = "5"
	SysRoleDataScopeDeptAndBelowOrPersonal = "6"
)

var (
	// 上传文件类型: 1文件 2图片
	SysUploadFileTypeFile = "file"
	SysUploadFileTypeImg  = "img"
)

var (
	// 用户性别: 0男 1女 2未知
	SysUserSexMale    = "0"
	SysUserSexFemale  = "1"
	SysUserSexUnknown = "2"
)

var (
	// 用户类型: pc端  app端
	SysUserTypeSys = "sys_user"
	SysUserTypeAPP = "app_user"
)

var (
	// 部门状态: 0正常 1停用
	SysDeptStatusNormal  = "0"
	SysDeptStatusDisable = "1"
)

var (
	// 任务状态: 0正常 1停用
	SysJobStatusNormal  = "0"
	SysJobStatusDisable = "1"
)
