package consts

// 内存缓存KEY常量
const (
	DeptName           = "dept_name_%d"             // 部门名称缓存 部门id:name 24小时过期
	UserInfo           = "user_info_%d"             // 用户信息缓存 用户id:miniinfo 24小时过期
	UserAccessCodeList = "user_access_code_list_%d" // 用户访问码列表缓存 用户id:access_code_list 24小时过期 用户退出清理
)
