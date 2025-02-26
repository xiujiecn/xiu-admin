package v1

import (
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type SysUserOnlineListReq struct {
	g.Meta `path:"/user-online/list" method:"get" tags:"系统" summary:"系统用户在线列表"`
	request.PageInfo
	model.SysUserOnlineListParam
}

type SysUserOnlineListRes struct {
	Items []*model.SysUserOnlineListModel `json:"items"`
	response.PageResult
}

type SysUserOnlineDeleteReq struct {
	g.Meta `path:"/user-online/delete" method:"delete" tags:"系统" summary:"系统用户在线删除"`
	ID     int64 `json:"id"`
}

type SysUserOnlineDeleteRes struct {
}
