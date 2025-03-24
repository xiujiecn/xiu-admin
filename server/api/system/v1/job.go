package v1

import (
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取任务列表
type SysJobListReq struct {
	g.Meta `path:"/job/list" method:"get" tags:"系统" summary:"获取任务列表"`
	model.SysJobListParam
	request.PageInfo
}

type SysJobListRes struct {
	Data []*model.SysJobListModel `json:"items" dc:"任务列表"`
	response.PageResult
}

// 获取任务详情
type SysJobViewReq struct {
	g.Meta `path:"/job/view" method:"get" tags:"系统" summary:"获取任务详情"`
	JobId  int64 `v:"required" json:"jobId" dc:"任务ID"`
}

type SysJobViewRes struct {
	model.SysJobViewModel
}

// 添加任务
type SysJobAddReq struct {
	g.Meta `path:"/job/add" method:"post" tags:"系统" summary:"添加任务"`
	model.SysJobAddModel
}

type SysJobAddRes struct {
	model.SysJobViewModel
}

// 更新任务
type SysJobUpdateReq struct {
	g.Meta `path:"/job/update" method:"post" tags:"系统" summary:"更新任务"`
	model.SysJobUpdateModel
}

type SysJobUpdateRes struct {
	model.SysJobViewModel
}

// 更新任务
type SysJobUpdateStatusReq struct {
	g.Meta `path:"/job/status" method:"post" tags:"系统" summary:"更新任务状态"`
	model.SysJobUpdateStatusModel
}

// 删除任务
type SysJobDeleteReq struct {
	g.Meta `path:"/job/delete" method:"post" tags:"系统" summary:"删除任务"`
	JobIds []int64 `json:"jobId" dc:"任务ID"`
}

type SysJobDeleteRes struct {
	JobIds []int64 `json:"jobId" dc:"任务ID"`
}
