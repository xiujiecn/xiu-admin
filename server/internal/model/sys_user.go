package model

import (
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type UserMiniModel struct {
	g.Meta   `orm:"table:sys_user" description:"用户"`
	UserId   int64  `json:"userId"   orm:"user_id"   description:"用户ID"`
	UserName string `json:"userName" orm:"user_name" description:"用户账号"`
	NickName string `json:"nickName" orm:"nick_name" description:"用户昵称"`
	Avatar   string `json:"avatar"   orm:"avatar"    description:"头像地址"`
	TenantId string `json:"tenantId"    orm:"tenant_id"    description:"租户编号"`
	DeptId   int64  `json:"deptId"      orm:"dept_id"      description:"部门ID"`
}

type LoginUserOut struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	TenantId string `json:"tenantId"    orm:"tenant_id"    description:"租户编号"`
	DeptId   int64  `json:"deptId"      orm:"dept_id"      description:"部门ID"`
}

type UserListQuery struct {
	TenantId    string `json:"tenantId"    orm:"tenant_id"    description:"租户编号"`
	UserId      int64  `json:"userId"      orm:"user_id"      description:"用户ID"`
	DeptId      int64  `json:"deptId"      orm:"dept_id"      description:"部门ID"`
	UserName    string `json:"userName"    orm:"user_name"    description:"用户账号"`
	NickName    string `json:"nickName"    orm:"nick_name"    description:"用户昵称"`
	Email       string `json:"email"       orm:"email"        description:"用户邮箱"`
	Phonenumber string `json:"phonenumber" orm:"phonenumber"  description:"手机号码"`
	Status      string `json:"status"      orm:"status"       description:"帐号状态（0正常 1停用）"`
}

type SysUserListModel struct {
	UserId      int64          `json:"userId"      orm:"user_id"      description:"用户ID"`
	TenantId    string         `json:"tenantId"    orm:"tenant_id"    description:"租户编号"`
	DeptId      int64          `json:"deptId"      orm:"dept_id"      description:"部门ID"`
	UserName    string         `json:"userName"    orm:"user_name"    description:"用户账号"`
	NickName    string         `json:"nickName"    orm:"nick_name"    description:"用户昵称"`
	UserType    string         `json:"userType"    orm:"user_type"    description:"用户类型（sys_user系统用户）"`
	Email       string         `json:"email"       orm:"email"        description:"用户邮箱"`
	Phonenumber string         `json:"phonenumber" orm:"phonenumber"  description:"手机号码"`
	Sex         string         `json:"sex"         orm:"sex"          description:"用户性别（0男 1女 2未知）"`
	Avatar      string         `json:"avatar"      orm:"avatar"       description:"头像地址"`
	Status      string         `json:"status"      orm:"status"       description:"帐号状态（0正常 1停用）"`
	LoginIp     string         `json:"loginIp"     orm:"login_ip"     description:"最后登录IP"`
	LoginDate   *gtime.Time    `json:"loginDate"   orm:"login_date"   description:"最后登录时间"`
	CreatedDept int64          `json:"createdDept" orm:"created_dept" description:"创建部门"`
	CreatedBy   int64          `json:"createdBy"   orm:"created_by"   description:"创建者"`
	CreatedAt   *gtime.Time    `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	Remark      string         `json:"remark"      orm:"remark"       description:"备注"`
	DeptInfo    *DeptMiniModel `json:"deptInfo"    orm:"with:dept_id=dept_id"    description:"部门信息"`
}

type SysUserViewModel struct {
	entity.SysUser
	DeptInfo *DeptMiniModel      `json:"deptInfo"    orm:"with:dept_id=dept_id"    description:"部门信息"`
	Roles    []*SysRoleMiniModel `json:"roles"       description:"角色信息"`
	Posts    []*SysPostMiniModel `json:"posts"       description:"岗位信息"`
}

type AddUser struct {
	UserId   int64   `json:"userId"`
	TenantId string  `json:"tenantId"`
	UserName string  `json:"userName"`
	NickName string  `json:"nickName"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Password string  `json:"password"`
	Status   string  `json:"status"`
	DeptId   int64   `json:"deptId"`
	RoleIds  []int64 `json:"roleIds"`
	PostIds  []int64 `json:"postIds"`
	Remark   string  `json:"remark"`
}

type UserProfileModel struct {
	SysUserViewModel
}
