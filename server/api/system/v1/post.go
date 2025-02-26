package v1

import (
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type PostListReq struct {
	g.Meta `path:"/post/list" method:"get" tags:"系统" summary:"获取岗位列表"`
	request.PageInfo
	model.SysPostListParam
}

type PostListRes struct {
	response.PageResult
	Data []*model.SysPostListModel `json:"items" dc:"岗位列表"`
}
