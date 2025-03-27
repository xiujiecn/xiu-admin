// Package v1
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package v1

import (
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type SysNoticeListReq struct {
	g.Meta `path:"/notice/list" method:"get" tags:"系统-公告管理" summary:"公告列表" x-check-permission:"cpm:system:notice:list"`
	*model.SysNoticeListParam
}

type SysNoticeListRes struct {
	response.PageResult
	Items []*model.SysNoticeListModel `json:"items"`
}

type SysNoticeAddReq struct {
	g.Meta `path:"/notice/add" method:"post" tags:"系统-公告管理" summary:"新增公告" x-check-permission:"cpm:system:notice:add"`
	*model.SysNoticeAddParam
}

type SysNoticeAddRes struct {
	*model.SysNoticeAddModel
}

type SysNoticeEditReq struct {
	g.Meta `path:"/notice/edit" method:"post" tags:"系统-公告管理" summary:"编辑公告" x-check-permission:"cpm:system:notice:edit"`
	*model.SysNoticeEditParam
}

type SysNoticeEditRes struct {
	*model.SysNoticeEditModel
}

type SysNoticeDeleteReq struct {
	g.Meta `path:"/notice/delete" method:"post" tags:"系统-公告管理" summary:"删除公告" x-check-permission:"cpm:system:notice:remove"`
	*model.SysNoticeDeleteParam
}

type SysNoticeDeleteRes struct {
	*model.SysNoticeDeleteModel
}

type SysNoticeViewReq struct {
	g.Meta `path:"/notice/view" method:"get" tags:"系统-公告管理" summary:"查看公告" x-check-permission:"cpm:system:notice:query"`
	*model.SysNoticeViewParam
}

type SysNoticeViewRes struct {
	*model.SysNoticeViewModel
}
