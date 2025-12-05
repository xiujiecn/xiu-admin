package model

import "github.com/gogf/gf/v2/os/gtime"

type QrcodeLoginModel struct {
	TempUserId string `json:"tempUserId" dc:"临时用户ID"`
	QrcodeCacheModel
	Bound        bool                  `json:"bound" dc:"是否已绑定"`
	Expired      bool                  `json:"expired" dc:"是否过期"`
	LoginUserOut *LoginUserOut         `json:"loginUserOut" dc:"登录用户信息"`
	Token        string                `json:"token" dc:"Token"`
	UserList     []*SysUserSocialModel `json:"userList" dc:"用户列表"`
}

type ApiCallQrcodeParam struct {
	ThirdId int    `json:"thirdId" dc:"第三方ID"`
	SceneId int    `json:"sceneId" dc:"场景ID"`
	Token   string `json:"token" dc:"Token"`
}

type ApiCallQrcodeModel struct {
	Ticket     string `json:"ticket" dc:"Ticket"`
	ExpireTime int64  `json:"expireTime" dc:"过期时间"`
	QrUrl      string `json:"qrUrl" dc:"二维码地址"`
}

type QrcodeCacheModel struct {
	Ticket     string      `json:"ticket" dc:"Ticket"`
	ExpireTime *gtime.Time `json:"expireTime" dc:"过期时间"`
	QrUrl      string      `json:"qrUrl" dc:"二维码地址"`
	Scanned    bool        `json:"scanned" dc:"是否已扫码"`
	OpenId     string      `json:"openId" dc:"微信OpenId"`
}

type QrcodeRegisterAndBindOpenIdParam struct {
	OpenId string `json:"openId" dc:"微信OpenId" v:"required#微信OpenId不能为空"`
	SysUserRegisterModel
}

type QrcodeLoginAndBindOpenIdModel struct {
	LoginUserOut *LoginUserOut `json:"loginUserOut" dc:"登录用户信息"`
	Token        string        `json:"token" dc:"Token"`
}

type QrcodeLoginAndBindOpenIdParam struct {
	OpenId       string `json:"openId" dc:"微信OpenId" v:"required#微信OpenId不能为空"`
	TenantId     string `json:"tenantId" v:"required#租户ID不能为空" dc:"租户ID"`
	Username     string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password     string `json:"password" v:"required#密码不能为空" dc:"密码"`
	CaptchaID    string `json:"captchaID" dc:"验证码ID"`
	CaptchaValue string `json:"captchaValue" v:"required#验证码不能为空" dc:"验证码"`
}

type QrcodeBindStatusModel struct {
	QrcodeCacheModel
	Bound   bool `json:"bound" dc:"是否已绑定"`
	Expired bool `json:"expired" dc:"是否过期"`
}

type QrcodeScanCallbackParam struct {
	SceneId   int    `json:"sceneId" dc:"场景ID"`
	ThirdId   int    `json:"thirdId" dc:"第三方ID"`
	Ticket    string `json:"ticket" dc:"Ticket"`
	OpenId    string `json:"openId" dc:"微信OpenId"`
	Event     string `json:"event" dc:"事件类型"`
	TimeStamp int64  `json:"timeStamp" dc:"时间戳"`
	Sign      string `json:"sign" dc:"签名"`
}

type ApiCallQrcodeResult struct {
	Code    int    `json:"code" dc:"状态码"`
	Message string `json:"message" dc:"消息"`
	Data    struct {
		Ticket     string `json:"ticket" dc:"Ticket"`
		ExpireTime int64  `json:"expireTime" dc:"过期时间"`
		QrUrl      string `json:"qrUrl" dc:"二维码地址"`
	} `json:"data" dc:"数据"`
	Error string `json:"error" dc:"错误信息"`
}

type QrcodeLoginSelectUserIdParam struct {
	OpenId   string `json:"openId" dc:"微信OpenId" v:"required#微信OpenId不能为空"`
	UserName string `json:"userName" dc:"用户名" v:"required#用户名不能为空"`
	TenantId string `json:"tenantId" dc:"租户ID" v:"required#租户ID不能为空"`
}
