// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysNoticeUser is the golang structure of table sys_notice_user for DAO operations like Where/Data.
type SysNoticeUser struct {
	g.Meta        `orm:"table:sys_notice_user, do:true"`
	Id            interface{} // ID
	NoticeId      interface{} // 公告ID
	TenantId      interface{} // 租户编号
	NoticeTitle   interface{} // 公告标题
	NoticeType    interface{} // 公告类型（1通知 2公告）
	NoticeContent []byte      // 公告内容
	Status        interface{} // 公告状态（0正常 1关闭）
	UserId        interface{} // 用户ID
	IsRead        interface{} // 是否已读（0未读 1已读）
	CreatedDept   interface{} // 创建部门
	CreatedBy     interface{} // 创建者
	CreatedAt     *gtime.Time // 创建时间
	UpdatedBy     interface{} // 更新者
	UpdatedAt     *gtime.Time // 更新时间
	Remark        interface{} // 备注
}
