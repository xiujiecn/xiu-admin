package model

import "github.com/gogf/gf/v2/os/gtime"

type SysNotice struct {
	NoticeId      int64       `json:"noticeId"      orm:"notice_id"      description:"公告ID"`
	TenantId      string      `json:"tenantId"      orm:"tenant_id"      description:"租户编号"`
	NoticeTitle   string      `json:"noticeTitle"   orm:"notice_title"   description:"公告标题"`
	NoticeType    string      `json:"noticeType"    orm:"notice_type"    description:"公告类型（1通知 2公告）"`
	NoticeContent []byte      `json:"noticeContent" orm:"notice_content" description:"公告内容"`
	Status        string      `json:"status"        orm:"status"         description:"公告状态（0正常 1关闭）"`
	CreatedDept   int64       `json:"createdDept"   orm:"created_dept"   description:"创建部门"`
	CreatedBy     int64       `json:"createdBy"     orm:"created_by"     description:"创建者"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`
	Remark        string      `json:"remark"        orm:"remark"         description:"备注"`
}

type SysNoticeListQuery struct {
	NoticeTitle string   `json:"noticeTitle"   orm:"notice_title"   description:"公告标题"`
	NoticeType  string   `json:"noticeType"    orm:"notice_type"    description:"公告类型（1通知 2公告）"`
	CreatedBy   string   `json:"createdBy"    orm:"created_by"     description:"创建者"`
	CreatedAt   []string `json:"createdAt"     orm:"created_at"     description:"创建时间断"`
}
