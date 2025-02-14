package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"

	"server/internal/packed/response"
	"server/internal/service"
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
