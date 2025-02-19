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
