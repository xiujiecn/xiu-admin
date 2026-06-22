package v1

import (
	"xiuadmin/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type InfoReq struct {
	g.Meta `path:"/info" method:"get" tags:"服务信息" summary:"服务信息"`
}

type InfoRes struct {
	*model.InfoModel
}
