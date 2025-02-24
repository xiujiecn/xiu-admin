package v1

import (
	"server/internal/model"
	"server/internal/model/request"
	"server/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type SysDictTypeListReq struct {
	g.Meta `path:"/dict/list" method:"get" tags:"字典类型" summary:"字典类型列表"`
	request.PageInfo
	model.SysDictTypeListQuery
}

type SysDictTypeListRes struct {
	response.PageResult
	Items []model.SysDictType `json:"items"`
}

type SysDictDataListReq struct {
	g.Meta `path:"/dict-data/list/{id}" method:"get" tags:"字典数据" summary:"字典数据列表"`
	request.PageInfo
	model.SysDictDataListQuery
}

type SysDictDataListRes struct {
	response.PageResult
	Items []model.SysDictData `json:"items"`
}
