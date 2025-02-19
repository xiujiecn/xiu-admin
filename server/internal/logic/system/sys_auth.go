package system

import (
	"context"
	"errors"
	"server/internal/model"
	"server/internal/service"
	"server/utility"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type sSysAuth struct {
}

func NewSysAuth() *sSysAuth {
	return &sSysAuth{}
}

func init() {
	service.RegisterSysAuth(NewSysAuth())
}

func (s *sSysAuth) Login(ctx context.Context, captchaID string, captchaValue string, username string, password string) (res *model.LoginUserOut, token string, err error) {
	// 验证验证码
	err = service.SysCaptcha().VerifyCaptcha(ctx, captchaID, captchaValue)
	if err != nil {
		return nil, "", err
	}
	// 获取用户信息
	user, err := service.SysUser().GetUserByUsernameAndPassword(ctx, username, password)
	if err != nil {
		return nil, "", err
	}
	userOut := &model.LoginUserOut{
		ID:       user.UserId,
		Username: user.UserName,
		NickName: user.NickName,
		Avatar:   user.Avatar,
	}
	// TODO: 保存在线列表
	// TODO: 保存登录日志
	// 生成token
	token, err = s.GenerateToken(ctx, userOut)
	if err != nil {
		return nil, "", err
	}
	return userOut, token, nil
}

// 生成Token
func (s *sSysAuth) GenerateToken(ctx context.Context, user *model.LoginUserOut) (token string, err error) {
	ets := g.Cfg().MustGet(ctx, "jwt.expiresTime", "7d").String()
	bts := g.Cfg().MustGet(ctx, "jwt.bufferTime", "1d").String()
	iss := g.Cfg().MustGet(ctx, "jwt.issuer", "XiujieAdmin").String()
	sk := g.Cfg().MustGet(ctx, "jwt.signingKey", "39c54195e73304e74a8429b178965865").String()
	et, _ := utility.ParseDuration(ets)
	bt, _ := utility.ParseDuration(bts)
	g.Log().Infof(ctx, "生成Token: user:%+v, et:%s, bt:%s, iss:%s, sk:%s", user, et, bt, iss, sk)
	claims := model.CustomClaims{
		BaseClaims: model.BaseClaims{
			UUID:     uuid.New(),
			ID:       user.ID,
			Username: user.Username,
			NickName: user.NickName,
		},
		BufferTime: int64(bt / time.Second),
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"XiujieAdmin"},           // 受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(et)),    // 过期时间 7天  配置文件
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			Issuer:    iss,                                       // 发行人
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString([]byte(sk))
	if err != nil {
		return "", err
	}
	return token, nil
}

// 解析Token
func (s *sSysAuth) ParseToken(ctx context.Context, token string) (claims *model.CustomClaims, err error) {
	sk := g.Cfg().MustGet(ctx, "jwt.signingKey", "39c54195e73304e74a8429b178965865").String()
	t, err := jwt.ParseWithClaims(token, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(sk), nil
	})
	if err != nil {

		return nil, err
	}
	if t.Claims.(*model.CustomClaims).ExpiresAt.Unix() < time.Now().Unix() {
		return nil, errors.New("token已过期")
	}
	return t.Claims.(*model.CustomClaims), nil
}

// 根据Token获取当前登录用户信息
func (s *sSysAuth) GetCurrentUser(ctx context.Context) (claims *model.CustomClaims, err error) {
	// 获取token
	authorization := g.RequestFromCtx(ctx).Header.Get("Authorization")
	if authorization == "" {
		return nil, gerror.NewCode(gcode.CodeNotImplemented)
	}
	token := strings.TrimPrefix(authorization, "Bearer ")
	// 解析token
	claims, err = s.ParseToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
