// Package gen
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
// @AutoGenerate Version
package system

import (
	"context"
	v1 "xiuadmin/api/system/v1"
	"xiuadmin/internal/model"
	"xiuadmin/internal/service"
)

// List 查看用户通知公告表列表
func (c *ControllerV1) SysNoticeUserList(ctx context.Context, req *v1.SysNoticeUserListReq) (res *v1.SysNoticeUserListRes, err error) {
	list, totalCount, err := service.SysNoticeUser().List(ctx, &req.SysNoticeUserListParam)
	if err != nil {
		return
	}

	if list == nil {
		list = []*model.SysNoticeUserListModel{}
	}

	res = new(v1.SysNoticeUserListRes)
	res.Items = list
	res.PageResult.Page = req.Page
	res.PageResult.PageSize = req.PageSize
	res.PageResult.Total = totalCount
	return
}

// Export 导出用户通知公告表列表
func (c *ControllerV1) SysNoticeUserExport(ctx context.Context, req *v1.SysNoticeUserExportReq) (res *v1.SysNoticeUserExportRes, err error) {
	err = service.SysNoticeUser().Export(ctx, &req.SysNoticeUserListParam)
	return
}

// Edit 更新用户通知公告表
func (c *ControllerV1) SysNoticeUserEdit(ctx context.Context, req *v1.SysNoticeUserEditReq) (res *v1.SysNoticeUserEditRes, err error) {
	err = service.SysNoticeUser().Edit(ctx, &req.SysNoticeUserEditParam)
	return
}

// View 获取指定用户通知公告表信息
func (c *ControllerV1) SysNoticeUserView(ctx context.Context, req *v1.SysNoticeUserViewReq) (res *v1.SysNoticeUserViewRes, err error) {
	data, err := service.SysNoticeUser().View(ctx, &req.SysNoticeUserViewParam)
	if err != nil {
		return
	}

	res = new(v1.SysNoticeUserViewRes)
	res.SysNoticeUserViewModel = data
	return
}

// Delete 删除用户通知公告表
func (c *ControllerV1) SysNoticeUserDelete(ctx context.Context, req *v1.SysNoticeUserDeleteReq) (res *v1.SysNoticeUserDeleteRes, err error) {
	err = service.SysNoticeUser().Delete(ctx, &req.SysNoticeUserDeleteParam)
	return
}

// Status 更新用户通知公告表状态
func (c *ControllerV1) SysNoticeUserStatus(ctx context.Context, req *v1.SysNoticeUserStatusReq) (res *v1.SysNoticeUserStatusRes, err error) {
	err = service.SysNoticeUser().Status(ctx, &req.SysNoticeUserStatusParam)
	return
}

// Read 已读
func (c *ControllerV1) SysNoticeUserRead(ctx context.Context, req *v1.SysNoticeUserReadReq) (res *v1.SysNoticeUserReadRes, err error) {
	err = service.SysNoticeUser().Read(ctx, &req.SysNoticeUserReadParam)
	return
}
