// Package sysnoticeuser
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
// @AutoGenerate Version
package v1 //sysnoticeuser

import (
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询用户通知公告表列表
type SysNoticeUserListReq struct {
	g.Meta `path:"/sysNoticeUser/list" method:"get" tags:"系统-用户通知公告表" summary:"获取用户通知公告表列表" x-check-permission:"cpm:system:sysNoticeUser:list,cpc:current:user"`
	model.SysNoticeUserListParam
}

type SysNoticeUserListRes struct {
	response.PageResult
	Items []*model.SysNoticeUserListModel `json:"items"   dc:"数据列表"`
}

// ExportReq 导出用户通知公告表列表
type SysNoticeUserExportReq struct {
	g.Meta `path:"/sysNoticeUser/export" method:"post" tags:"系统-用户通知公告表" summary:"导出用户通知公告表列表" x-check-permission:"cpm:system:sysNoticeUser:export,cpc:current:user"`
	model.SysNoticeUserListParam
}

type SysNoticeUserExportRes struct{}

// ViewReq 获取用户通知公告表指定信息
type SysNoticeUserViewReq struct {
	g.Meta `path:"/sysNoticeUser/view" method:"get" tags:"系统-用户通知公告表" summary:"获取用户通知公告表指定信息" x-check-permission:"cpm:system:sysNoticeUser:view,cpc:current:user"`
	model.SysNoticeUserViewParam
}

type SysNoticeUserViewRes struct {
	*model.SysNoticeUserViewModel
}

// EditReq 修改/新增用户通知公告表
type SysNoticeUserEditReq struct {
	g.Meta `path:"/sysNoticeUser/edit" method:"post" tags:"系统-用户通知公告表" summary:"修改/新增用户通知公告表" x-check-permission:"cpm:system:sysNoticeUser:edit,cpc:current:user"`
	model.SysNoticeUserEditParam
}

type SysNoticeUserEditRes struct{}

// DeleteReq 删除用户通知公告表
type SysNoticeUserDeleteReq struct {
	g.Meta `path:"/sysNoticeUser/delete" method:"post" tags:"系统-用户通知公告表" summary:"删除用户通知公告表" x-check-permission:"cpm:system:sysNoticeUser:delete,cpc:current:user"`
	model.SysNoticeUserDeleteParam
}

type SysNoticeUserDeleteRes struct{}

// StatusReq 更新用户通知公告表状态
type SysNoticeUserStatusReq struct {
	g.Meta `path:"/sysNoticeUser/status" method:"post" tags:"系统-用户通知公告表" summary:"更新用户通知公告表状态" x-check-permission:"cpm:system:sysNoticeUser:status,cpc:current:user"`
	model.SysNoticeUserStatusParam
}

type SysNoticeUserStatusRes struct{}

// ReadReq 已读
type SysNoticeUserReadReq struct {
	g.Meta `path:"/sysNoticeUser/read" method:"post" tags:"系统-用户通知公告表" summary:"已读" x-check-permission:"cpm:system:sysNoticeUser:read,cpc:current:user"`
	model.SysNoticeUserReadParam
}

type SysNoticeUserReadRes struct{}
