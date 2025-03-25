package v1

import (
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type SysDictTypeListReq struct {
	g.Meta `path:"/dict/list" method:"get" tags:"系统-字典管理" summary:"字典类型列表" x-check-permission:"cpm:system:dict:list"`
	model.SysDictTypeListParam
}

type SysDictTypeListRes struct {
	response.PageResult
	Items []model.SysDictTypeListModel `json:"items"`
}

type SysDictTypeAddReq struct {
	g.Meta `path:"/dict-type/add" method:"post" tags:"系统-字典管理" summary:"新增字典类型" x-check-permission:"cpm:system:dict:add"`
	model.SysDictTypeAddParam
}

type SysDictTypeAddRes struct {
	model.SysDictTypeAddModel
}

type SysDictTypeEditReq struct {
	g.Meta `path:"/dict-type/edit" method:"post" tags:"系统-字典管理" summary:"编辑字典类型" x-check-permission:"cpm:system:dict:edit"`
	model.SysDictTypeEditParam
}

type SysDictTypeEditRes struct {
	model.SysDictTypeEditModel
}

type SysDictTypeDeleteReq struct {
	g.Meta `path:"/dict-type/delete" method:"post" tags:"系统-字典管理" summary:"删除字典类型" x-check-permission:"cpm:system:dict:remove"`
	model.SysDictTypeDeleteParam
}

type SysDictTypeDeleteRes struct {
	model.SysDictTypeDeleteModel
}

type SysDictTypeViewReq struct {
	g.Meta `path:"/dict-type/view" method:"get" tags:"系统-字典管理" summary:"字典类型详情" x-check-permission:"cpm:system:dict:query"`
	model.SysDictTypeViewParam
}

type SysDictTypeViewRes struct {
	model.SysDictTypeViewModel
}

type SysDictDataListReq struct {
	g.Meta `path:"/dict-data/list/{id}" method:"get" tags:"系统-字典管理" summary:"字典数据列表" `
	request.PageInfo
	model.SysDictDataListParam
}

type SysDictDataListRes struct {
	response.PageResult
	Items []model.SysDictDataListModel `json:"items"`
	Type  *model.SysDictTypeViewModel  `json:"type"`
}

type SysDictDataAddReq struct {
	g.Meta `path:"/dict-data/add" method:"post" tags:"系统-字典管理" summary:"新增字典数据" x-check-permission:"cpm:system:dict:add"`
	model.SysDictDataAddParam
}

type SysDictDataAddRes struct {
	model.SysDictDataAddModel
}

type SysDictDataEditReq struct {
	g.Meta `path:"/dict-data/edit" method:"post" tags:"系统-字典管理" summary:"编辑字典数据" x-check-permission:"cpm:system:dict:edit"`
	model.SysDictDataEditParam
}

type SysDictDataEditRes struct {
	model.SysDictDataEditModel
}

type SysDictDataDeleteReq struct {
	g.Meta `path:"/dict-data/delete" method:"post" tags:"系统-字典管理" summary:"删除字典数据" x-check-permission:"cpm:system:dict:remove"`
	model.SysDictDataDeleteParam
}

type SysDictDataDeleteRes struct {
	model.SysDictDataDeleteModel
}

type SysDictDataViewReq struct {
	g.Meta `path:"/dict-data/view" method:"get" tags:"系统-字典管理" summary:"字典数据详情" x-check-permission:"cpm:system:dict:query"`
	model.SysDictDataViewParam
}

type SysDictDataViewRes struct {
	model.SysDictDataViewModel
}
