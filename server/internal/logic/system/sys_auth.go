// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"errors"
	"slices"
	"strings"
	"time"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/library/bcache"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/event"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/request"
	"xiuadmin/internal/service"
	"xiuadmin/utility"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v4"
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

func (s *sSysAuth) Login(ctx context.Context, param *model.LoginParams) (res *model.LoginUserOut, token string, err error) {
	ip := g.RequestFromCtx(ctx).GetClientIp()
	logininfor := &model.SysLogininforAddModel{
		TenantId:      param.TenantId,
		UserName:      param.Username,
		ClientKey:     "web",
		DeviceType:    "web",
		Ipaddr:        ip,
		LoginLocation: utility.GetCityByIp(ip),
		Browser:       utility.GetBrowser(ctx),
		Os:            utility.GetOs(ctx),
		Status:        "0",
		Msg:           "登录成功",
		LoginTime:     gtime.Now(),
	}
	// 验证验证码
	err = service.SysCaptcha().VerifyCaptcha(ctx, param.CaptchaID, param.CaptchaValue)
	if err != nil {
		logininfor.Status = "1"
		logininfor.Msg = err.Error()
		return nil, "", err
	}
	// 获取用户信息
	user, err := service.SysUser().GetUserByUsernameAndPassword(ctx, param.TenantId, param.Username, param.Password)
	if err != nil {
		logininfor.Status = "1"
		logininfor.Msg = err.Error()
		return nil, "", err
	}
	userOut := &model.LoginUserOut{
		ID:       user.UserId,
		Username: user.UserName,
		NickName: user.NickName,
		Avatar:   user.Avatar,
		TenantId: user.TenantId,
		DeptId:   user.DeptId,
	}
	// 构建临时的上下文
	identity := &model.Identity{
		BaseClaims: model.BaseClaims{
			ID:       user.UserId,
			Username: user.UserName,
			NickName: user.NickName,
			TenantId: user.TenantId,
			DeptId:   user.DeptId,
			Roles:    []model.Role{},
			LoginAt:  time.Now().Unix(),
		},
	}
	contextModel := new(model.Context)
	contextModel.User = identity

	ctx = contexts.Set(ctx, contextModel)

	// 生成token
	claims, token, err := s.GenerateToken(ctx, userOut)
	if err != nil {
		logininfor.Status = "1"
		logininfor.Msg = err.Error()
		return nil, "", err
	}

	// 保存登录日志
	go func() {
		id, err := service.SysLogininfor().AddLogininfor(ctx, logininfor)
		if err != nil {
			return
		}

		// 保存在线列表
		service.SysUserOnline().Add(ctx, &model.SysUserOnlineAddModel{
			OnlineId:      id,
			TenantId:      user.TenantId,
			Uuid:          claims.UUID,
			UserName:      user.UserName,
			ClientKey:     "web",
			DeviceType:    "web",
			Ipaddr:        ip,
			LoginLocation: utility.GetCityByIp(ip),
			Browser:       utility.GetBrowser(ctx),
			Os:            utility.GetOs(ctx),
			Token:         token,
			LoginTime:     gtime.Now(),
			ExpireTime:    gtime.NewFromTime(claims.ExpiresAt.Time),
		})
	}()
	event.EventsInstance().Emit(ctx, consts.EventKeyUserLogin, user.UserId)
	return userOut, token, nil
}

// 生成Token
func (s *sSysAuth) GenerateToken(ctx context.Context, user *model.LoginUserOut) (claims *model.CustomClaims, token string, err error) {
	ets := g.Cfg().MustGet(ctx, "jwt.expiresTime", "7d").String()
	bts := g.Cfg().MustGet(ctx, "jwt.bufferTime", "1d").String()
	iss := g.Cfg().MustGet(ctx, "jwt.issuer", "XiuAdmin").String()
	sk := g.Cfg().MustGet(ctx, "jwt.signingKey", "39c54195e73304e74a8429b178965865").String()
	et, _ := utility.ParseDuration(ets)
	bt, _ := utility.ParseDuration(bts)
	g.Log().Infof(ctx, "sSysAuth.GenerateToken 生成Token: user:%+v, et:%s, bt:%s, iss:%s, sk:%s", user, et, bt, iss, sk)
	authorityIds, err := service.SysUser().GetUserRoleIds(ctx, user.ID)
	if err != nil {
		return nil, "", err
	}
	rl, _, err := service.SysRole().List(ctx, &model.SysRoleListParam{
		PageInfo: request.PageInfo{
			Page:     1,
			PageSize: 5000,
		},
		RoleIds: authorityIds,
	})
	if err != nil {
		return nil, "", err
	}
	roles := make([]model.Role, 0)
	for _, r := range rl {
		roles = append(roles, model.Role{
			RoleId:    r.RoleId,
			DataScope: r.DataScope,
		})
	}
	g.Log().Infof(ctx, "sSysAuth.GenerateToken roles: %v", roles)
	claims = &model.CustomClaims{
		BaseClaims: model.BaseClaims{
			UUID:     strings.ReplaceAll(uuid.New().String(), "-", ""),
			ID:       user.ID,
			Username: user.Username,
			NickName: user.NickName,
			TenantId: user.TenantId,
			DeptId:   user.DeptId,
			Roles:    roles,
			LoginAt:  time.Now().Unix(),
		},
		BufferTime: int64(bt / time.Second),
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"XiuAdmin"},              // 受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(et)),    // 过期时间 7天  配置文件
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			Issuer:    iss,                                       // 发行人
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString([]byte(sk))
	if err != nil {
		return nil, "", err
	}
	g.Log().Infof(ctx, "sSysAuth.GenerateToken token: %s, claims: %+v", token, claims)
	bcache.SetSysAuthToken(ctx, claims.BaseClaims.ID, claims.BaseClaims.UUID, token, et)
	return claims, token, nil
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
	_, err = bcache.GetSysAuthToken(ctx, t.Claims.(*model.CustomClaims).BaseClaims.ID, t.Claims.(*model.CustomClaims).BaseClaims.UUID)
	if err != nil {
		return nil, err
	}
	return t.Claims.(*model.CustomClaims), nil
}

// 删除Token
func (s *sSysAuth) DeleteToken(ctx context.Context, token string) (err error) {
	sk := g.Cfg().MustGet(ctx, "jwt.signingKey", "39c54195e73304e74a8429b178965865").String()
	t, err := jwt.ParseWithClaims(token, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(sk), nil
	})
	if err != nil {
		return err
	}
	timeout := t.Claims.(*model.CustomClaims).ExpiresAt.Unix() - time.Now().Unix()
	if timeout <= 0 {
		return nil
	}
	err = bcache.DelSysAuthToken(ctx, t.Claims.(*model.CustomClaims).BaseClaims.ID, t.Claims.(*model.CustomClaims).BaseClaims.UUID)
	if err != nil {
		return err
	}
	bcache.SetSysAuthTokenReject(ctx, t.Claims.(*model.CustomClaims).BaseClaims.ID, t.Claims.(*model.CustomClaims).BaseClaims.UUID, token, time.Duration(timeout)*time.Second)
	return nil
}

// 获取token
func (s *sSysAuth) GetAccessToken(ctx context.Context) (token string, err error) {
	authorization := g.RequestFromCtx(ctx).Header.Get("Authorization")
	authParam := g.RequestFromCtx(ctx).Get("access_token").String()
	if authorization == "" && authParam == "" {
		return "", gerror.NewCode(gcode.CodeNotImplemented)
	}
	token = strings.TrimPrefix(authorization, "Bearer ")
	if token == "" {
		token = authParam
	}
	return token, nil
}

// 根据Token获取当前登录用户信息
func (s *sSysAuth) GetCurrentUser(ctx context.Context) (claims *model.CustomClaims, err error) {
	// 获取token
	token, err := s.GetAccessToken(ctx)
	if err != nil {
		return nil, err
	}
	// 解析token
	claims, err = s.ParseToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

// 获取用户权限码
func (s *sSysAuth) GetUserAccessCodeList(ctx context.Context, userId int64) (accessCodeList []string, menuRoleAccessCodeList []string, err error) {
	accessCodeList = make([]string, 0)
	menuRoleAccessCodeList = make([]string, 0)
	roleDataScopeMap := make(map[int64]string)

	// 角色权限码
	roleIds, err := service.SysUser().GetUserRoleIds(ctx, userId)
	if err != nil {
		return nil, nil, err
	}
	if len(roleIds) > 0 {
		rl, _, err := service.SysRole().List(ctx, &model.SysRoleListParam{
			PageInfo: request.PageInfo{
				Page:     1,
				PageSize: 5000,
			},
			RoleIds: roleIds,
		})
		if err != nil {
			return nil, nil, err
		}
		for _, r := range rl {
			accessCodeList = append(accessCodeList, consts.SysCheckPermissionRolePrefix+r.RoleKey)
			roleDataScopeMap[r.RoleId] = r.DataScope
		}
	}
	g.Log().Infof(ctx, "sSysAuth.GetUserAccessCodeList roleIds: %v", roleIds)
	// 菜单权限码
	// 获取角色菜单列表
	rmList, err := service.SysRole().GetRoleListMenu(ctx, roleIds)
	if err != nil {
		return nil, nil, err
	}
	menuIds := make([]int64, 0)
	for _, menu := range rmList {
		menuIds = append(menuIds, menu.MenuId)
	}
	g.Log().Infof(ctx, "sSysAuth.GetUserAccessCodeList menuIds: %v", menuIds)
	if len(menuIds) > 0 {
		// 获取用户角色菜单
		menuList, _, err := service.SysMenu().List(ctx, &model.SysMenuListParam{
			MenuIds: menuIds,
		})
		if err != nil {
			return nil, nil, err
		}
		g.Log().Infof(ctx, "sSysAuth.GetUserAccessCodeList menuList: %v", menuList)
		for _, m := range menuList {
			accessCodeList = append(accessCodeList, consts.SysCheckPermissionMenuPrefix+m.Perms)
			for _, r := range rmList {
				if r.MenuId == m.MenuId {
					roleDS := roleDataScopeMap[r.RoleId]
					if len(roleDS) < 1 {
						roleDS = consts.SysRoleDataScopeAll
					}
					menuRoleAccessCodeList = append(menuRoleAccessCodeList, consts.SysCheckPermissionMenuPrefix+m.Perms+"|"+gconv.String(r.RoleId)+"|"+roleDS)
				}
			}
		}
	}
	g.Log().Infof(ctx, "sSysAuth.GetUserAccessCodeList accessCodeList: %v", accessCodeList)
	// 获取用户岗位信息
	postIds, err := service.SysUser().GetUserPostIds(ctx, userId)
	if err != nil {
		return nil, nil, err
	}
	if len(postIds) > 0 {
		pl, _, err := service.SysPost().List(ctx, &model.SysPostListParam{
			PageInfo: request.PageInfo{
				Page:     1,
				PageSize: 5000,
			},
			PostIds: postIds,
		})
		if err != nil {
			return nil, nil, err
		}
		for _, p := range pl {
			accessCodeList = append(accessCodeList, consts.SysCheckPermissionPostPrefix+p.PostCode)
		}
	}
	// 获取用户信息
	user, err := service.SysUser().GetUserById(ctx, userId)
	if err != nil {
		return nil, nil, err
	}
	accessCodeList = append(accessCodeList, consts.SysCheckPermissionUserPrefix+user.UserName)
	accessCodeList = append(accessCodeList, consts.SysCheckPermissionCurrentUser)
	accessCodeList = slices.Compact(accessCodeList)
	return accessCodeList, menuRoleAccessCodeList, nil
}

// 根据openId登录
func (s *sSysAuth) LoginByOpenId(ctx context.Context, social *model.SysSocialListModel) (res *model.LoginUserOut, token string, err error) {
	ip := g.RequestFromCtx(ctx).GetClientIp()
	logininfor := &model.SysLogininforAddModel{
		TenantId:      social.User.TenantId,
		UserName:      social.User.UserName,
		ClientKey:     "web",
		DeviceType:    "web",
		Ipaddr:        ip,
		LoginLocation: utility.GetCityByIp(ip),
		Browser:       utility.GetBrowser(ctx),
		Os:            utility.GetOs(ctx),
		Status:        "0",
		Msg:           "登录成功",
		LoginTime:     gtime.Now(),
	}
	user, err := service.SysUser().GetUserByUsername(ctx, social.User.UserName, social.User.TenantId)
	if err != nil {
		return nil, "用户不存在", err
	}
	userOut := &model.LoginUserOut{
		ID:       user.UserId,
		Username: user.UserName,
		NickName: user.NickName,
		Avatar:   user.Avatar,
		TenantId: user.TenantId,
		DeptId:   user.DeptId,
	}
	// 构建临时的上下文
	identity := &model.Identity{
		BaseClaims: model.BaseClaims{
			ID:       user.UserId,
			Username: user.UserName,
			NickName: user.NickName,
			TenantId: user.TenantId,
			DeptId:   user.DeptId,
			Roles:    []model.Role{},
			LoginAt:  time.Now().Unix(),
		},
	}
	contextModel := new(model.Context)
	contextModel.User = identity

	ctx = contexts.Set(ctx, contextModel)

	// 生成token
	claims, token, err := s.GenerateToken(ctx, userOut)
	if err != nil {
		logininfor.Status = "1"
		logininfor.Msg = err.Error()
		return nil, "", err
	}

	// 保存登录日志
	go func() {
		id, err := service.SysLogininfor().AddLogininfor(ctx, logininfor)
		if err != nil {
			return
		}

		// 保存在线列表
		service.SysUserOnline().Add(ctx, &model.SysUserOnlineAddModel{
			OnlineId:      id,
			TenantId:      user.TenantId,
			Uuid:          claims.UUID,
			UserName:      user.UserName,
			ClientKey:     "web",
			DeviceType:    "web",
			Ipaddr:        ip,
			LoginLocation: utility.GetCityByIp(ip),
			Browser:       utility.GetBrowser(ctx),
			Os:            utility.GetOs(ctx),
			Token:         token,
			LoginTime:     gtime.Now(),
			ExpireTime:    gtime.NewFromTime(claims.ExpiresAt.Time),
		})
	}()
	event.EventsInstance().Emit(ctx, consts.EventKeyUserLogin, user.UserId)
	return userOut, token, nil
}
