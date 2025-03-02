package system

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "xiujieadmin/api/system/v1"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	userOut, token, err := service.SysAuth().Login(ctx, &model.LoginParams{
		CaptchaID:    req.CaptchaID,
		CaptchaValue: req.CaptchaValue,
		Username:     req.Username,
		Password:     req.Password,
		TenantId:     req.TenantId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.LoginRes{
		ID:          userOut.ID,
		Username:    userOut.Username,
		NickName:    userOut.NickName,
		Avatar:      userOut.Avatar,
		AccessToken: token,
	}, nil
}
func (c *ControllerV1) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerV1) GetAccessCodes(ctx context.Context, req *v1.GetAccessCodesReq) (res *v1.GetAccessCodesRes, err error) {
	accessCodes := make([]string, 0)
	accessCodes = append(accessCodes, "123456")
	res = &v1.GetAccessCodesRes{
		Data: accessCodes,
	}
	return res, nil
}
func (c *ControllerV1) TenantList(ctx context.Context, req *v1.TenantListReq) (res *v1.TenantListRes, err error) {
	tenantList, _, err := service.SysTenant().List(ctx, &model.SysTenantListParam{}, &request.PageInfo{})
	if err != nil {
		return nil, err
	}
	tenantEnabled := false
	if len(tenantList) > 1 {
		tenantEnabled = true
	}
	voList := make([]v1.TenantListData, 0)
	for _, tenant := range tenantList {
		voList = append(voList, v1.TenantListData{
			CompanyName: tenant.CompanyName,
			Domain:      tenant.Domain,
			TenantId:    tenant.TenantId,
		})
	}
	return &v1.TenantListRes{
		TenantEnabled: tenantEnabled,
		VoList:        voList,
	}, nil
}
