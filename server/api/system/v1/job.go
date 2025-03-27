// Package v1
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package v1

import (
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/request"
	"xiuadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取任务列表
type SysJobListReq struct {
	g.Meta `path:"/job/list" method:"get" tags:"系统-任务管理" summary:"获取任务列表" x-check-permission:"cpm:system:job:list"`
	model.SysJobListParam
	request.PageInfo
}

type SysJobListRes struct {
	Data []*model.SysJobListModel `json:"items" dc:"任务列表"`
	response.PageResult
}

// 获取任务详情
type SysJobViewReq struct {
	g.Meta `path:"/job/view" method:"get" tags:"系统-任务管理" summary:"获取任务详情" x-check-permission:"cpm:system:job:view"`
	JobId  int64 `v:"required" json:"jobId" dc:"任务ID"`
}

type SysJobViewRes struct {
	model.SysJobViewModel
}

// 添加任务
type SysJobAddReq struct {
	g.Meta `path:"/job/add" method:"post" tags:"系统-任务管理" summary:"添加任务" x-check-permission:"cpm:system:job:add"`
	model.SysJobAddModel
}

type SysJobAddRes struct {
	model.SysJobViewModel
}

// 更新任务
type SysJobUpdateReq struct {
	g.Meta `path:"/job/update" method:"post" tags:"系统-任务管理" summary:"更新任务" x-check-permission:"cpm:system:job:update"`
	model.SysJobUpdateModel
}

type SysJobUpdateRes struct {
	model.SysJobViewModel
}

// 更新任务
type SysJobUpdateStatusReq struct {
	g.Meta `path:"/job/status" method:"post" tags:"系统" summary:"更新任务状态"  x-check-permission:"cpm:system:job:status"`
	model.SysJobUpdateStatusModel
}

type SysJobUpdateStatusRes = SysJobUpdateRes

// 删除任务
type SysJobDeleteReq struct {
	g.Meta `path:"/job/delete" method:"post" tags:"系统-任务管理" summary:"删除任务"  x-check-permission:"cpm:system:job:delete"`
	JobIds []int64 `json:"jobId" dc:"任务ID"`
}

type SysJobExecReq struct {
	g.Meta `path:"/job/exec" method:"post" tags:"系统" summary:"执行任务"  x-check-permission:"cpm:system:job:exec"`
	JobId  int64 `v:"required" json:"jobId" dc:"任务ID"`
}

type SysJobDeleteRes struct {
	JobIds []int64 `json:"jobId" dc:"任务ID"`
}

type SysJobExecRes struct {
	JobId int64 `json:"jobId" dc:"任务ID"`
}
