package middleware

import (
	"context"
	"slices"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"xiujieadmin/internal/consts"
	"xiujieadmin/internal/library/contexts"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/packed/response"
	"xiujieadmin/internal/service"
)

type sMiddleware struct {
}

func New() *sMiddleware {
	return &sMiddleware{}
}

func init() {
	service.RegisterMiddleware(New())
}

// 响应处理中间件
func (m *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if r.Response.BufferLength() > 0 {
		return
	}
	err := r.GetError()
	res := r.GetHandlerResponse()
	var code gcode.Code = gcode.CodeOK

	if err != nil {
		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		response.JsonExit(r, code.Code(), err.Error())
	} else {
		response.Json(r, code.Code(), "", res)
	}
}

// Ctx 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	ctx := r.GetCtx()
	r.SetCtx(r.GetNeverDoneCtx())

	// 初始化登录用户信息
	data, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		r.Middleware.Next()
		return
	}
	if data != nil {
		contextModel := new(model.Context)
		contextModel.User = &model.Identity{
			BaseClaims: data.BaseClaims,
		}
		contexts.Init(r, contextModel)
	}
	r.Middleware.Next()
}

// CORS 跨域处理
func (s *sMiddleware) CORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsConfig := g.Cfg().MustGet(context.Background(), "server.allowedDomains").Strings()
	if len(corsConfig) == 0 {
		r.Response.CORSDefault()
	} else {
		corsOptions.AllowDomain = corsConfig
		r.Response.CORS(corsOptions)
	}
	r.Middleware.Next()
}

// Auth 认证处理
func (s *sMiddleware) Auth(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		path = r.URL.Path
	)
	// 不需要验证登录的路由地址
	if s.IsExceptLogin(ctx, path) {
		r.Middleware.Next()
		return
	}

	userId := contexts.GetUserId(ctx)
	if userId == 0 {
		g.Log().Error(ctx, "sMiddleware.Auth userId is 0", "path", path)
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), consts.CodeLoginExpired.Message())
		return
	}
	if !contexts.IsSuperAdmin(r.GetCtx()) {
		r.Middleware.Next()
		return
	}
	// TODO: 检查API权限
	r.Middleware.Next()
}

// IsExceptLogin 判断是否需要验证登录
func (s *sMiddleware) IsExceptLogin(ctx context.Context, path string) bool {
	pathList := g.Cfg().MustGet(ctx, "router.exceptLogin").Strings()

	for i := 0; i < len(pathList); i++ {
		if slices.Contains(pathList, path) {
			return true
		}
	}
	return false
}
