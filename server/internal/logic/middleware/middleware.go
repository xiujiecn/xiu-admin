// package middleware
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package middleware

import (
	"context"
	"net/http"
	"slices"
	"strings"

	"xiuadmin/utility"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gmeta"

	"xiuadmin/internal/consts"
	"xiuadmin/internal/library/bcache"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/mcache"
	"xiuadmin/internal/model"
	"xiuadmin/internal/packed/response"
	"xiuadmin/internal/service"
	"xiuadmin/utility/queue"
)

type sMiddleware struct {
}

func New() *sMiddleware {
	return &sMiddleware{}
}

func init() {
	service.RegisterMiddleware(New())
}

func getContentType(r *ghttp.Request) (contentType string) {
	contentType = r.Response.Header().Get("Content-Type")
	if contentType != "" {
		return
	}

	mime := gmeta.Get(r.GetHandlerResponse(), "mime").String()
	if mime == "" {
		contentType = consts.HTTPContentTypeJson
	} else {
		contentType = mime
	}
	return
}

// 响应处理中间件
func (m *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if r.Response.BufferLength() > 0 {
		return
	}
	if contexts.IsWebSocket(r.GetCtx()) {
		return
	}
	contentType := getContentType(r)
	if contentType == consts.HTTPContentTypeExcel {
		r.Response.Status = http.StatusOK
		return
	}
	// g.Log().Infof(r.GetCtx(), "sMiddleware.ResponseHandler,contentType %v", contentType)
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
	r.SetCtx(r.GetNeverDoneCtx())
	if contexts.IsWebSocket(r.GetCtx()) {
		r.Middleware.Next()
		return
	}
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
		ctx    = r.Context()
		method = r.Method
		path   = ""
	)
	if r.URL != nil {
		path = r.URL.Path
	}
	// 不需要验证登录的路由地址
	if s.IsExceptLogin(ctx, path) {
		r.Middleware.Next()
		return
	}

	userId := contexts.GetUserId(ctx)
	tenantId := contexts.GetTenantId(ctx)
	g.Log().Debugf(ctx, "sMiddleware.Auth userId: %d tenantId: %s method: %s path: %s, ================>>", userId, tenantId, method, path)
	if userId == 0 {
		g.Log().Error(ctx, "sMiddleware.Auth userId is 0", "path", path, "clientIp", r.GetClientIp(), utility.GetUserAgent(ctx))
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), consts.CodeLoginExpired.Message())
		// r.Response.WriteStatus(http.StatusUnauthorized)
		return
	}
	// 如果是websocket，则不进行权限验证
	// g.Log().Infof(ctx, "sMiddleware.Auth is websocket: %v", contexts.IsWebSocket(ctx))
	if contexts.IsWebSocket(ctx) {
		r.Middleware.Next()
		return
	}
	// 是否校验强行退出用户
	forceLogout, err := mcache.GetSystemConfig(ctx, consts.ConfigOnlineForceLogout, consts.ConfigOnlineForceLogoutFalse)
	if err != nil {
		g.Log().Errorf(ctx, "sMiddleware.Auth GetSystemConfig error: %v", err)
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), consts.CodeLoginExpired.Message())
		return
	}
	if forceLogout == consts.ConfigOnlineForceLogoutTrue {
		uuid := contexts.GetUserUuid(ctx)
		token, err := bcache.GetSysAuthTokenReject(ctx, userId, uuid)
		if err != nil {
			g.Log().Errorf(ctx, "sMiddleware.Auth GetSysAuthTokenReject error: %v", err)
			response.JsonExit(r, gcode.CodeNotAuthorized.Code(), consts.CodeLoginExpired.Message())
			return
		}
		if token != "" {
			g.Log().Infof(ctx, "sMiddleware.Auth token is reject, userId: %d, uuid: %s, token: %s", userId, uuid, token)
			response.JsonExit(r, gcode.CodeNotAuthorized.Code(), consts.CodeLoginExpired.Message())
			return
		}
	}
	// 超级管理员
	if contexts.IsSuperAdmin(r.GetCtx()) {
		r.Middleware.Next()
		return
	}
	// 权限验证
	serveHandler := r.GetServeHandler()
	userAccessCodeList, menuRoleDataAccessCodeList, err := mcache.GetUserAccessCodeList(ctx, userId)
	if err != nil {
		g.Log().Errorf(ctx, "sMiddleware.Auth GetUserAccessCodeList error: %v", err)
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), consts.CodeLoginExpired.Message())
		return
	}
	if serveHandler != nil {
		g.Log().Infof(ctx, "sMiddleware.Auth serveHandler GetMetaTag(x-check-permission): %+v", serveHandler.GetMetaTag("x-check-permission"))
		accessCode := serveHandler.GetMetaTag("x-check-permission")
		if accessCode != "" {
			accessCodeList := strings.Split(accessCode, ",")
			hasPermission := false
			for _, code := range accessCodeList {
				if strings.HasPrefix(code, consts.SysCheckPermissionPrefix) {
					accessCodeItemList := strings.Split(code, "&")
					hasItem := true
					for _, item := range accessCodeItemList {
						if !slices.Contains(userAccessCodeList, item) {
							g.Log().Errorf(ctx, "sMiddleware.Auth userAccessCodeList not contains item: %v, userAccessCodeList: %v", item, userAccessCodeList)
							hasItem = false
							break
						}
						for _, roleDataAccessCode := range menuRoleDataAccessCodeList {
							if strings.HasPrefix(roleDataAccessCode, item+"|") {
								roleDataAccessCodeItemList := strings.Split(roleDataAccessCode, "|")
								g.Log().Debugf(ctx, "sMiddleware.Auth roleDataAccessCodeItemList: %v", roleDataAccessCodeItemList)
								if len(roleDataAccessCodeItemList) == 3 {
									roleId := roleDataAccessCodeItemList[1]
									roleDataScope := roleDataAccessCodeItemList[2]
									contexts.SetDataValue(ctx, "currentRoleId", roleId)
									contexts.SetDataValue(ctx, "currentRoleDataScope", roleDataScope)
									g.Log().Debugf(ctx, "sMiddleware.Auth currentRoleId: %s, currentRoleDataScope: %s", roleId, roleDataScope)
								}
								break
							}
						}
					}
					if hasItem {
						hasPermission = true
						break
					}
				}
			}
			if !hasPermission {
				response.JsonExit(r, consts.CodeNoPermission.Code(), consts.CodeNoPermission.Message())
				return
			}
		}
	}
	r.Middleware.Next()
}

// IsExceptLogin 判断是否需要验证登录
func (s *sMiddleware) IsExceptLogin(ctx context.Context, path string) bool {
	pathList := g.Cfg().MustGet(ctx, "router.exceptLogin").Strings()
	return slices.Contains(pathList, path)
}

// OperationLog 操作日志
func (s *sMiddleware) OperationLog(r *ghttp.Request) {
	r.Middleware.Next()
	ctx := r.Context()
	enabledDBOperationLog := g.Cfg().MustGet(ctx, "server.enabledDBOperationLog").Bool()
	if !enabledDBOperationLog {
		return
	}
	data, err := service.SysOperLog().AnalysisLog(ctx)
	if err != nil {
		g.Log().Errorf(context.TODO(), "sMiddleware.OperationLog error:%v", err)
		return
	}
	// 写入队列
	// logData, _ := json.Marshal(data)
	queue.Push(r.GetCtx(), consts.QueueSysOptLog, data)
}

// ResponseHandlerRaw 返回处理中间件
func (s *sMiddleware) ResponseHandlerRaw(r *ghttp.Request) {
	r.Middleware.Next()
	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		g.Log().Infof(r.Context(), "ResponseHandler r.Response.BufferLength()  %d buf:%s", r.Response.BufferLength(), r.Response.BufferString())
		return
	}

	var (
		err = r.GetError()
		res = r.GetHandlerResponse()
	)
	if err != nil {
		r.Response.WriteJson(res)
		r.ExitAll()
	} else {
		r.Response.WriteJson(res)
	}
}
