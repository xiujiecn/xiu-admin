// package model
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package model

import (
	"xiuadmin/internal/model/request"

	"github.com/gogf/gf/v2/os/gtime"
)

type SysNoticeListModel struct {
	NoticeId      int64            `json:"noticeId"      orm:"notice_id"      description:"公告ID"`
	TenantId      string           `json:"tenantId"      orm:"tenant_id"      description:"租户编号"`
	NoticeTitle   string           `json:"noticeTitle"   orm:"notice_title"   description:"公告标题"`
	NoticeType    string           `json:"noticeType"    orm:"notice_type"    description:"公告类型（1通知 2公告）"`
	NoticeContent []byte           `json:"noticeContent" orm:"notice_content" description:"公告内容"`
	Status        string           `json:"status"        orm:"status"         description:"公告状态（0正常 1关闭）"`
	NoticeRange   int              `json:"noticeRange"   orm:"notice_range"   description:"公告范围（1全部 2指定机构 3指定用户）"`
	CreatedDept   int64            `json:"createdDept"   orm:"created_dept"   description:"创建部门"`
	CreatedBy     int64            `json:"createdBy"     orm:"created_by"     description:"创建者"`
	CreatedAt     *gtime.Time      `json:"createdAt"     orm:"created_at"     description:"创建时间"`
	Remark        string           `json:"remark"        orm:"remark"         description:"备注"`
	CreatedByUser SysUserMiniModel `json:"createdByUser" orm:"with:user_id=created_by" description:"创建者"`
}

type SysNoticeListParam struct {
	request.PageInfo
	NoticeTitle string   `json:"noticeTitle"    description:"公告标题"`
	NoticeType  string   `json:"noticeType"     description:"公告类型（1通知 2公告）"`
	CreatedBy   string   `json:"createdBy"      description:"创建者"`
	CreatedAt   []string `json:"createdAt"      description:"创建时间断"`
}

type SysNoticeAddParam struct {
	NoticeTitle   string  `json:"noticeTitle"    description:"公告标题"`
	NoticeType    string  `json:"noticeType"     description:"公告类型（1通知 2公告）"`
	NoticeContent string  `json:"noticeContent"  description:"公告内容"`
	Status        string  `json:"status"         description:"公告状态（0正常 1关闭）"`
	Remark        string  `json:"remark"         description:"备注"`
	UserIdList    []int64 `json:"userIdList"  description:"通知用户ID列表"`
	DeptIdList    []int64 `json:"deptIdList"  description:"通知机构ID列表"`
	NoticeRange   int     `json:"noticeRange"    description:"公告范围（1全部 2指定机构 3指定用户）"`
}

type SysNoticeAddModel struct {
	NoticeId int64 `json:"noticeId"       description:"公告ID"`
}

type SysNoticeEditParam struct {
	NoticeId      int64   `json:"noticeId"       description:"公告ID"`
	NoticeTitle   *string `json:"noticeTitle"    description:"公告标题"`
	NoticeType    *string `json:"noticeType"     description:"公告类型（1通知 2公告）"`
	NoticeContent *string `json:"noticeContent"  description:"公告内容"`
	Status        *string `json:"status"         description:"公告状态（0正常 1关闭）"`
	Remark        *string `json:"remark"         description:"备注"`
	UserIdList    []int64 `json:"userIdList"  description:"通知用户ID列表"`
	DeptIdList    []int64 `json:"deptIdList"  description:"通知机构ID列表"`
	NoticeRange   int     `json:"noticeRange"    description:"公告范围（1全部 2指定机构 3指定用户）"`
}

type SysNoticeEditModel struct {
	NoticeId int64 `json:"noticeId"       description:"公告ID"`
}

type SysNoticeDeleteParam struct {
	NoticeIds []int64 `json:"noticeIds" orm:"notice_ids" description:"公告ID"`
}

type SysNoticeDeleteModel struct {
	NoticeIds []int64 `json:"noticeIds" orm:"notice_ids" description:"公告ID"`
}
type SysNoticeViewParam struct {
	NoticeId int64 `json:"noticeId" orm:"notice_id" description:"公告ID"`
}

type SysNoticeViewModel struct {
	NoticeId      int64       `json:"noticeId"      orm:"notice_id"      description:"公告ID"`
	TenantId      string      `json:"tenantId"      orm:"tenant_id"      description:"租户编号"`
	NoticeTitle   string      `json:"noticeTitle"   orm:"notice_title"   description:"公告标题"`
	NoticeType    string      `json:"noticeType"    orm:"notice_type"    description:"公告类型（1通知 2公告）"`
	NoticeContent string      `json:"noticeContent" orm:"notice_content" description:"公告内容"`
	UserIds       string      `json:"userIds"       orm:"user_ids"       description:"通知用户ID列表"`
	DeptIds       string      `json:"deptIds"       orm:"dept_ids"       description:"通知机构ID列表"`
	UserIdList    []int64     `json:"userIdList"  description:"通知用户ID列表"`
	DeptIdList    []int64     `json:"deptIdList"  description:"通知机构ID列表"`
	NoticeRange   int         `json:"noticeRange"    description:"公告范围（1全部 2指定机构 3指定用户）"`
	Status        string      `json:"status"        orm:"status"         description:"公告状态（0正常 1关闭）"`
	CreatedDept   int64       `json:"createdDept"   orm:"created_dept"   description:"创建部门"`
	CreatedBy     int64       `json:"createdBy"     orm:"created_by"     description:"创建者"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
	UpdatedBy     int64       `json:"updatedBy"     orm:"updated_by"     description:"更新者"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`
	Remark        string      `json:"remark"        orm:"remark"         description:"备注"`
}
