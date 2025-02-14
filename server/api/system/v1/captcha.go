package v1

import "github.com/gogf/gf/v2/frame/g"

// 获取验证码
type GetCaptchaReq struct {
	g.Meta `path:"/captcha" method:"get" tags:"系统" summary:"获取验证码"`
}

type GetCaptchaRes struct {
	CaptchaID    string `json:"captchaID" dc:"验证码ID"`
	CaptchaImage string `json:"captchaImage" dc:"验证码图片"`
}
