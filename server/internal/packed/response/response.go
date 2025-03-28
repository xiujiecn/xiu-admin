// package response
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package response

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = g.Map{}
	}
	// g.Log().Infof(r.GetCtx(), "response.Json: %v", responseData)
	r.Response.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

func JsonExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	Json(r, code, message, data...)
	r.ExitAll()
}

// 向前端返回Excel文件 参数 content 为上面生成的io.ReadSeeker， fileTag 为返回前端的文件名
func Excel(r *ghttp.Request, content io.ReadSeeker, fileTag string) {
	fileName := fmt.Sprintf("%s%s%s.xlsx", gtime.Now().Format("20060102150405"), `_`, fileTag)
	r.Response.Header().Set("Content-Type", "application/vnd.ms-excel")
	r.Response.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	http.ServeContent(r.Response.Writer, r.Request, fileTag, time.Now(), content)
}

// 向前端返回Json文件 参数 content 为上面生成的io.ReadSeeker， fileTag 为返回前端的文件名
func JsonFile(r *ghttp.Request, content io.ReadSeeker, fileTag string) {
	fileName := fmt.Sprintf("%s%s%s.json", gtime.Now().Format("20060102150405"), `_`, fileTag)
	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	http.ServeContent(r.Response.Writer, r.Request, fileTag, time.Now(), content)
}
