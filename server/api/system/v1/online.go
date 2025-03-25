package v1

import (
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type SysUserOnlineListReq struct {
	g.Meta `path:"/user-online/list" method:"get" tags:"监控-在线用户管理" summary:"系统用户在线列表" x-check-permission:"cpm:monitor:online:list"`
	request.PageInfo
	model.SysUserOnlineListParam
}

type SysUserOnlineListRes struct {
	Items []*model.SysUserOnlineListModel `json:"items"`
	response.PageResult
}

type SysUserOnlineDeleteReq struct {
	g.Meta `path:"/user-online/delete" method:"delete" tags:"监控-在线用户管理" summary:"系统用户在线删除" x-check-permission:"cpm:monitor:online:remove"`
	ID     int64 `json:"id"`
}

type SysUserOnlineDeleteRes struct {
}
