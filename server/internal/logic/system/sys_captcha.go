package system

import (
	"context"
	"server/internal/consts"
	"server/internal/service"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.NewMemoryStore(100, 120*time.Second)

type sSysCaptcha struct {
}

func NewSysCaptcha() *sSysCaptcha {
	return &sSysCaptcha{}
}

func init() {
	service.RegisterSysCaptcha(NewSysCaptcha())
}

type CaptchaConfig struct {
	KeyLong            int
	ImgWidth           int
	ImgHeight          int
	OpenCaptcha        int
	OpenCaptchaTimeout int
}

// 生成验证码
func (l *sSysCaptcha) GenerateCaptcha(ctx context.Context) (key string, image string, err error) {
	var captchaConfig CaptchaConfig
	g.Cfg().MustGet(ctx, "captcha").Struct(&captchaConfig)
	driver := base64Captcha.NewDriverDigit(captchaConfig.ImgHeight, captchaConfig.ImgWidth, captchaConfig.KeyLong, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, store)
	key, image, _, err = captcha.Generate()
	return
}

// 验证验证码
func (l *sSysCaptcha) VerifyCaptcha(ctx context.Context, key string, value string) (err error) {
	ok := store.Verify(key, value, true)
	if !ok {
		err = gerror.NewCode(consts.CodeCaptchaError, "验证码错误")
	}
	return
}
