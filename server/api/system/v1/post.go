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

type PostListReq struct {
	g.Meta `path:"/post/list" method:"get" tags:"系统-岗位管理" summary:"获取岗位列表" x-check-permission:"cpm:system:post:list"`
	*model.SysPostListParam
}

type PostListRes struct {
	response.PageResult
	Data []*model.SysPostListModel `json:"items" dc:"岗位列表"`
}

type PostAddReq struct {
	g.Meta `path:"/post/add" method:"post" tags:"系统-岗位管理" summary:"新增岗位" x-check-permission:"cpm:system:post:add"`
	*model.SysPostAddParam
}

type PostAddRes struct {
	*model.SysPostAddModel
}

type PostEditReq struct {
	g.Meta `path:"/post/edit" method:"post" tags:"系统-岗位管理" summary:"编辑岗位" x-check-permission:"cpm:system:post:edit"`
	*model.SysPostEditParam
}

type PostEditRes struct {
	*model.SysPostEditModel
}

type PostDeleteReq struct {
	g.Meta `path:"/post/delete" method:"post" tags:"系统-岗位管理" summary:"删除岗位"`
	*model.SysPostDeleteParam
}

type PostDeleteRes struct {
	*model.SysPostDeleteModel
}

type PostViewReq struct {
	g.Meta `path:"/post/view" method:"get" tags:"系统-岗位管理" summary:"获取岗位详情" x-check-permission:"cpm:system:post:list"`
	*model.SysPostViewParam
}

type PostViewRes struct {
	*model.SysPostViewModel
}

type PostExportReq struct {
	g.Meta `path:"/post/export" method:"post" tags:"系统-岗位管理" summary:"导出岗位" x-check-permission:"cpm:system:post:export"`
	*model.SysPostExportParam
}

type PostExportRes struct {
	*model.SysPostExportModel
}
