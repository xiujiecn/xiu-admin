package consts

import "github.com/gogf/gf/v2/errors/gcode"

var (
	// 账号不存在
	CodeUserNotFound = gcode.New(10001, "账号不存在", nil)
	// 账号已删除
	CodeUserDeleted = gcode.New(10002, "账号已删除", nil)
	// 密码错误
	CodeUserPasswordError = gcode.New(10003, "密码错误", nil)
	// 账号已禁用
	CodeUserDisabled = gcode.New(10004, "账号已禁用", nil)
	// 账号已锁定
	CodeUserLocked = gcode.New(10005, "账号已锁定", nil)
	// 账号已过期
	CodeUserExpired = gcode.New(10006, "账号已过期", nil)
	// 验证码错误
	CodeCaptchaError = gcode.New(10007, "验证码错误", nil)
)
