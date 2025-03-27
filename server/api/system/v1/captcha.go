// Package v1
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package v1

import "github.com/gogf/gf/v2/frame/g"

// 获取验证码
type GetCaptchaReq struct {
	g.Meta `path:"/captcha" method:"get" tags:"系统-授权" summary:"获取验证码"`
}

type GetCaptchaRes struct {
	CaptchaID    string `json:"captchaID" dc:"验证码ID"`
	CaptchaImage string `json:"captchaImage" dc:"验证码图片"`
}
