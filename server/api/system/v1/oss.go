package v1

import (
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type SysOssListReq struct {
	g.Meta `path:"/oss/list" method:"get" tags:"系统" summary:"获取OSS列表"`
	model.SysOssListParam
	request.PageInfo
}

type SysOssListRes struct {
	Items []*model.SysOssListModel `json:"items"`
	response.PageResult
}

type SysOssViewReq struct {
	g.Meta `path:"/oss/view" method:"get" tags:"系统" summary:"查看文件"`
	*model.SysOssViewParam
}

type SysOssViewRes struct {
	*model.SysOssViewModel
}

type SysOssDeleteReq struct {
	g.Meta `path:"/oss/delete" method:"post" tags:"系统" summary:"删除文件"`
	*model.SysOssDeleteParam
}

type SysOssDeleteRes struct {
	*model.SysOssDeleteModel
}
