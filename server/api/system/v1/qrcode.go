package v1

import (
	"xiuadmin/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取登录二维码
type QrcodeLoginReq struct {
	g.Meta `path:"/qrcode/login" method:"get" tags:"系统-二维码" summary:"获取登录二维码"`
}

type QrcodeLoginRes struct {
	model.QrcodeLoginModel
}

// 获取登录二维码扫码结果
type QrcodeLoginStatusReq struct {
	g.Meta     `path:"/qrcode/login/status" method:"get" tags:"系统-二维码" summary:"获取登录二维码扫码状态"`
	TempUserId string `json:"tempUserId" dc:"临时用户ID" v:"required#临时用户ID不能为空"`
}

type QrcodeLoginStatusRes struct {
	model.QrcodeLoginModel
}

// 获取绑定二维码
type QrcodeBindReq struct {
	g.Meta `path:"/qrcode/bind" method:"get" tags:"系统-二维码" summary:"获取绑定二维码"`
}

type QrcodeBindRes struct {
	model.QrcodeCacheModel
}

// 获取绑定二维码扫码结果
type QrcodeBindStatusReq struct {
	g.Meta `path:"/qrcode/bind/status" method:"get" tags:"系统-二维码" summary:"获取绑定二维码扫码状态"`
}

type QrcodeBindStatusRes struct {
	model.QrcodeBindStatusModel
}

// 扫码成功回调
type QrcodeScanCallbackReq struct {
	g.Meta `path:"/qrcode/scan/callback" method:"post" tags:"系统-二维码" summary:"扫码成功回调"`
	model.QrcodeScanCallbackParam
}

type QrcodeScanCallbackRes struct {
}

// 注册并关联OpenId
type QrcodeRegisterAndBindOpenIdReq struct {
	g.Meta `path:"/qrcode/registerAndBindOpenId" method:"post" tags:"系统-二维码" summary:"注册并关联OpenId"`
	model.QrcodeRegisterAndBindOpenIdParam
}

type QrcodeRegisterAndBindOpenIdRes struct {
	model.QrcodeLoginAndBindOpenIdModel
}

// 登录并关联OpenId
type QrcodeLoginAndBindOpenIdReq struct {
	g.Meta `path:"/qrcode/loginAndBindOpenId" method:"post" tags:"系统-二维码" summary:"登录并关联OpenId"`
	model.QrcodeLoginAndBindOpenIdParam
}

type QrcodeLoginAndBindOpenIdRes struct {
	model.QrcodeLoginAndBindOpenIdModel
}

// 扫码后选择用户登录
type QrcodeLoginSelectUserReq struct {
	g.Meta `path:"/qrcode/loginSelectUser" method:"post" tags:"系统-二维码" summary:"扫码后选择用户登录"`
	model.QrcodeLoginSelectUserIdParam
}

type QrcodeLoginSelectUserRes struct {
	model.QrcodeLoginAndBindOpenIdModel
}
