package system

import (
	"context"

	v1 "xiuadmin/api/system/v1"
	"xiuadmin/internal/service"
)

// 获取登录二维码
func (c *ControllerV1) QrcodeLogin(ctx context.Context, req *v1.QrcodeLoginReq) (res *v1.QrcodeLoginRes, err error) {
	qrcodeLoginModel, err := service.SysQrcode().GetQrcodeLogin(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.QrcodeLoginRes{
		QrcodeLoginModel: *qrcodeLoginModel,
	}, nil
}

// 获取登录二维码扫码结果
func (c *ControllerV1) QrcodeLoginStatus(ctx context.Context, req *v1.QrcodeLoginStatusReq) (res *v1.QrcodeLoginStatusRes, err error) {
	qrcodeLoginModel, err := service.SysQrcode().GetQrcodeLoginStatus(ctx, req.TempUserId)
	if err != nil {
		return nil, err
	}
	return &v1.QrcodeLoginStatusRes{
		QrcodeLoginModel: *qrcodeLoginModel,
	}, nil
}

// 获取绑定二维码
func (c *ControllerV1) QrcodeBind(ctx context.Context, req *v1.QrcodeBindReq) (res *v1.QrcodeBindRes, err error) {
	qrcodeBindModel, err := service.SysQrcode().GetQrcodeBind(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.QrcodeBindRes{
		QrcodeCacheModel: *qrcodeBindModel,
	}, nil
}

// 获取绑定二维码扫码结果
func (c *ControllerV1) QrcodeBindStatus(ctx context.Context, req *v1.QrcodeBindStatusReq) (res *v1.QrcodeBindStatusRes, err error) {
	qrcodeBindStatusModel, err := service.SysQrcode().GetQrcodeBindStatus(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.QrcodeBindStatusRes{
		QrcodeBindStatusModel: *qrcodeBindStatusModel,
	}, nil
}

// 扫码成功回调
func (c *ControllerV1) QrcodeScanCallback(ctx context.Context, req *v1.QrcodeScanCallbackReq) (res *v1.QrcodeScanCallbackRes, err error) {
	err = service.SysQrcode().QrcodeScanCallback(ctx, &req.QrcodeScanCallbackParam)
	if err != nil {
		return nil, err
	}
	return &v1.QrcodeScanCallbackRes{}, nil
}

// 注册并关联OpenId
func (c *ControllerV1) QrcodeRegisterAndBindOpenId(ctx context.Context, req *v1.QrcodeRegisterAndBindOpenIdReq) (res *v1.QrcodeRegisterAndBindOpenIdRes, err error) {
	qrcodeLoginAndBindOpenIdModel, err := service.SysQrcode().QrcodeRegisterAndBindOpenId(ctx, &req.QrcodeRegisterAndBindOpenIdParam)
	if err != nil {
		return nil, err
	}
	return &v1.QrcodeRegisterAndBindOpenIdRes{
		QrcodeLoginAndBindOpenIdModel: *qrcodeLoginAndBindOpenIdModel,
	}, nil
}

// 登录并关联OpenId
func (c *ControllerV1) QrcodeLoginAndBindOpenId(ctx context.Context, req *v1.QrcodeLoginAndBindOpenIdReq) (res *v1.QrcodeLoginAndBindOpenIdRes, err error) {
	qrcodeLoginAndBindOpenIdModel, err := service.SysQrcode().QrcodeLoginAndBindOpenId(ctx, &req.QrcodeLoginAndBindOpenIdParam)
	if err != nil {
		return nil, err
	}
	return &v1.QrcodeLoginAndBindOpenIdRes{
		QrcodeLoginAndBindOpenIdModel: *qrcodeLoginAndBindOpenIdModel,
	}, nil
}

// 扫码后选择用户登录
func (c *ControllerV1) QrcodeLoginSelectUser(ctx context.Context, req *v1.QrcodeLoginSelectUserReq) (res *v1.QrcodeLoginSelectUserRes, err error) {
	qrcodeLoginAndBindOpenIdModel, err := service.SysQrcode().QrcodeLoginSelectUserId(ctx, &req.QrcodeLoginSelectUserIdParam)
	if err != nil {
		return nil, err
	}
	return &v1.QrcodeLoginSelectUserRes{
		QrcodeLoginAndBindOpenIdModel: *qrcodeLoginAndBindOpenIdModel,
	}, nil
}
