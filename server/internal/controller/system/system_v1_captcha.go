package system

import (
	"context"

	v1 "xiuadmin/api/system/v1"
	"xiuadmin/internal/service"
)

func (c *ControllerV1) GetCaptcha(ctx context.Context, req *v1.GetCaptchaReq) (res *v1.GetCaptchaRes, err error) {
	key, image, err := service.SysCaptcha().GenerateCaptcha(ctx)
	res = &v1.GetCaptchaRes{
		CaptchaID:    key,
		CaptchaImage: image,
	}
	return
}
