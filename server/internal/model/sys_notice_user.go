// Package genin
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
// @AutoGenerate Version
package model

import (
	"context"
	"slices"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/library/xgorm/hook"
	"xiuadmin/internal/model/request"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysNoticeUserUpdateFields 修改用户通知公告表字段过滤
type SysNoticeUserUpdateFields struct {
	NoticeId      int64       `json:"noticeId"      dc:"公告ID"`
	TenantId      string      `json:"tenantId"      dc:"租户编号"`
	NoticeTitle   string      `json:"noticeTitle"   dc:"公告标题"`
	NoticeType    string      `json:"noticeType"    dc:"公告类型"`
	NoticeContent []byte      `json:"noticeContent" dc:"公告内容"`
	Status        string      `json:"status"        dc:"公告状态"`
	UserId        int64       `json:"userId"        dc:"用户ID"`
	IsRead        int         `json:"isRead"        dc:"是否已读"`
	UpdatedBy     int64       `json:"updatedBy"     dc:"更新者"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
	Remark        string      `json:"remark"        dc:"备注"`
}

// SysNoticeUserInsertFields 新增用户通知公告表字段过滤
type SysNoticeUserInsertFields struct {
	NoticeId      int64       `json:"noticeId"      dc:"公告ID"`
	TenantId      string      `json:"tenantId"      dc:"租户编号"`
	NoticeTitle   string      `json:"noticeTitle"   dc:"公告标题"`
	NoticeType    string      `json:"noticeType"    dc:"公告类型"`
	NoticeContent []byte      `json:"noticeContent" dc:"公告内容"`
	Status        string      `json:"status"        dc:"公告状态"`
	UserId        int64       `json:"userId"        dc:"用户ID"`
	IsRead        int         `json:"isRead"        dc:"是否已读"`
	CreatedDept   int64       `json:"createdDept"   dc:"创建部门"`
	CreatedBy     int64       `json:"createdBy"     dc:"创建者"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	Remark        string      `json:"remark"        dc:"备注"`
}

// SysNoticeUserEditParam 修改/新增用户通知公告表
type SysNoticeUserEditParam struct {
	Id            int64       `json:"id"            orm:"id"             description:"ID"`
	NoticeId      int64       `json:"noticeId"      orm:"notice_id"      description:"公告ID"`
	TenantId      string      `json:"tenantId"      orm:"tenant_id"      description:"租户编号"`
	NoticeTitle   string      `json:"noticeTitle"   orm:"notice_title"   description:"公告标题"`
	NoticeType    string      `json:"noticeType"    orm:"notice_type"    description:"公告类型（1通知 2公告）"`
	NoticeContent string      `json:"noticeContent" orm:"notice_content" description:"公告内容"`
	Status        string      `json:"status"        orm:"status"         description:"公告状态（0正常 1关闭）"`
	UserId        int64       `json:"userId"        orm:"user_id"        description:"用户ID"`
	IsRead        int         `json:"isRead"        orm:"is_read"        description:"是否已读（0未读 1已读）"`
	CreatedDept   int64       `json:"createdDept"   orm:"created_dept"   description:"创建部门"`
	CreatedBy     int64       `json:"createdBy"     orm:"created_by"     description:"创建者"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
	UpdatedBy     int64       `json:"updatedBy"     orm:"updated_by"     description:"更新者"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`
	Remark        string      `json:"remark"        orm:"remark"         description:"备注"`
}

func (in *SysNoticeUserEditParam) Filter(ctx context.Context) (err error) {

	return
}

type SysNoticeUserEditModel struct{}

// SysNoticeUserDeleteParam 删除用户通知公告表
type SysNoticeUserDeleteParam struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *SysNoticeUserDeleteParam) Filter(ctx context.Context) (err error) {
	return
}

type SysNoticeUserDeleteModel struct{}

// SysNoticeUserViewParam 获取指定用户通知公告表信息
type SysNoticeUserViewParam struct {
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *SysNoticeUserViewParam) Filter(ctx context.Context) (err error) {
	return
}

type SysNoticeUserViewModel struct {
	Id             int64             `json:"id"            orm:"id"             description:"ID"`
	NoticeId       int64             `json:"noticeId"      orm:"notice_id"      description:"公告ID"`
	TenantId       string            `json:"tenantId"      orm:"tenant_id"      description:"租户编号"`
	NoticeTitle    string            `json:"noticeTitle"   orm:"notice_title"   description:"公告标题"`
	NoticeType     string            `json:"noticeType"    orm:"notice_type"    description:"公告类型（1通知 2公告）"`
	NoticeContent  string            `json:"noticeContent" orm:"notice_content" description:"公告内容"`
	Status         string            `json:"status"        orm:"status"         description:"公告状态（0正常 1关闭）"`
	UserId         int64             `json:"userId"        orm:"user_id"        description:"用户ID"`
	IsRead         int               `json:"isRead"        orm:"is_read"        description:"是否已读（0未读 1已读）"`
	CreatedDept    int64             `json:"createdDept"   orm:"created_dept"   description:"创建部门"`
	CreatedBy      int64             `json:"createdBy"     orm:"created_by"     description:"创建者"`
	CreatedAt      *gtime.Time       `json:"createdAt"     orm:"created_at"     description:"创建时间"`
	UpdatedBy      int64             `json:"updatedBy"     orm:"updated_by"     description:"更新者"`
	UpdatedAt      *gtime.Time       `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`
	Remark         string            `json:"remark"        orm:"remark"         description:"备注"`
	CreatedBySumma *hook.MemberSumma `json:"createdBySumma" dc:"创建者摘要信息"`
	UpdatedBySumma *hook.MemberSumma `json:"updatedBySumma" dc:"更新者摘要信息"`
}

// SysNoticeUserListParam 获取用户通知公告表列表
type SysNoticeUserListParam struct {
	request.PageInfo
	IsRead      string `json:"isRead" dc:"是否已读"`
	NoticeId    int64  `json:"noticeId" dc:"公告ID"`
	NoticeTitle string `json:"noticeTitle"    dc:"公告标题"`
	NoticeType  string `json:"noticeType"     dc:"公告类型"`
}

func (in *SysNoticeUserListParam) Filter(ctx context.Context) (err error) {
	return
}

type SysNoticeUserListModel struct {
	Id             int64             `json:"id"             dc:"ID"`
	NoticeId       int64             `json:"noticeId"       dc:"公告ID"`
	TenantId       string            `json:"tenantId"       dc:"租户编号"`
	NoticeTitle    string            `json:"noticeTitle"    dc:"公告标题"`
	NoticeType     string            `json:"noticeType"     dc:"公告类型"`
	NoticeContent  string            `json:"noticeContent"  dc:"公告内容"`
	Status         string            `json:"status"         dc:"公告状态"`
	UserId         int64             `json:"userId"         dc:"用户ID"`
	IsRead         int               `json:"isRead"         dc:"是否已读"`
	CreatedDept    int64             `json:"createdDept"    dc:"创建部门"`
	CreatedBy      int64             `json:"createdBy"      dc:"创建者"`
	CreatedBySumma *hook.MemberSumma `json:"createdBySumma" dc:"创建者摘要信息"`
	CreatedAt      *gtime.Time       `json:"createdAt"      dc:"创建时间"`
	UpdatedBy      int64             `json:"updatedBy"      dc:"更新者"`
	UpdatedBySumma *hook.MemberSumma `json:"updatedBySumma" dc:"更新者摘要信息"`
	UpdatedAt      *gtime.Time       `json:"updatedAt"      dc:"更新时间"`
	Remark         string            `json:"remark"         dc:"备注"`
}

// SysNoticeUserExportModel 导出用户通知公告表
type SysNoticeUserExportModel struct {
	Id            int64       `json:"id"            dc:"ID"`
	NoticeId      int64       `json:"noticeId"      dc:"公告ID"`
	TenantId      string      `json:"tenantId"      dc:"租户编号"`
	NoticeTitle   string      `json:"noticeTitle"   dc:"公告标题"`
	NoticeType    string      `json:"noticeType"    dc:"公告类型"`
	NoticeContent string      `json:"noticeContent" dc:"公告内容"`
	Status        string      `json:"status"        dc:"公告状态"`
	UserId        int64       `json:"userId"        dc:"用户ID"`
	IsRead        int         `json:"isRead"        dc:"是否已读"`
	CreatedDept   int64       `json:"createdDept"   dc:"创建部门"`
	CreatedBy     int64       `json:"createdBy"     dc:"创建者"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedBy     int64       `json:"updatedBy"     dc:"更新者"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
	Remark        string      `json:"remark"        dc:"备注"`
}

// SysNoticeUserStatusParam 更新用户通知公告表状态
type SysNoticeUserStatusParam struct {
	Id     int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *SysNoticeUserStatusParam) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !slices.Contains(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type SysNoticeUserStatusModel struct{}

type SysNoticeUserReadParam struct {
	Ids []int64 `json:"ids" v:"required#ID不能为空" dc:"ID"`
}

func (in *SysNoticeUserReadParam) Filter(ctx context.Context) (err error) {
	return
}
