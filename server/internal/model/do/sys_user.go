// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure of table sys_user for DAO operations like Where/Data.
type SysUser struct {
	g.Meta      `orm:"table:sys_user, do:true"`
	UserId      interface{} // 用户ID
	TenantId    interface{} // 租户编号
	DeptId      interface{} // 部门ID
	UserName    interface{} // 用户账号
	NickName    interface{} // 用户昵称
	UserType    interface{} // 用户类型（sys_user系统用户）
	Email       interface{} // 用户邮箱
	Phonenumber interface{} // 手机号码
	Sex         interface{} // 用户性别（0男 1女 2未知）
	Avatar      interface{} // 头像地址
	Salt        interface{} // 加密盐
	Password    interface{} // 密码
	Status      interface{} // 帐号状态（0正常 1停用）
	LoginIp     interface{} // 最后登录IP
	LoginDate   *gtime.Time // 最后登录时间
	CreatedDept interface{} // 创建部门
	CreatedBy   interface{} // 创建者
	CreatedAt   *gtime.Time // 创建时间
	UpdatedBy   interface{} // 更新者
	UpdatedAt   *gtime.Time // 更新时间
	DeletedBy   interface{} // 删除人
	DeletedAt   *gtime.Time // 删除时间
	Remark      interface{} // 备注
}
