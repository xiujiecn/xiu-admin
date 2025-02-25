package model

import (
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
	TenantId    string   `json:"tenantId"    description:"租户编号"`
	UserId      int64    `json:"userId"      description:"用户ID"`
	DeptId      int64    `json:"deptId"      description:"部门ID"`
	UserName    string   `json:"userName"    description:"用户账号"`
	NickName    string   `json:"nickName"    description:"用户昵称"`
	Email       string   `json:"email"       description:"用户邮箱"`
	Phonenumber string   `json:"phonenumber" description:"手机号码"`
	Status      string   `json:"status"      description:"帐号状态（0正常 1停用）"`
	CreatedAt   []string `json:"createdAt"   description:"创建时间"`
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
	UserId      int64               `json:"userId"      orm:"user_id"      description:"用户ID"`
	TenantId    string              `json:"tenantId"    orm:"tenant_id"    description:"租户编号"`
	DeptId      int64               `json:"deptId"      orm:"dept_id"      description:"部门ID"`
	UserName    string              `json:"userName"    orm:"user_name"    description:"用户账号"`
	NickName    string              `json:"nickName"    orm:"nick_name"    description:"用户昵称"`
	UserType    string              `json:"userType"    orm:"user_type"    description:"用户类型（sys_user系统用户）"`
	Email       string              `json:"email"       orm:"email"        description:"用户邮箱"`
	Phonenumber string              `json:"phonenumber" orm:"phonenumber"  description:"手机号码"`
	Sex         string              `json:"sex"         orm:"sex"          description:"用户性别（0男 1女 2未知）"`
	Avatar      string              `json:"avatar"      orm:"avatar"       description:"头像地址"`
	Status      string              `json:"status"      orm:"status"       description:"帐号状态（0正常 1停用）"`
	LoginIp     string              `json:"loginIp"     orm:"login_ip"     description:"最后登录IP"`
	LoginDate   *gtime.Time         `json:"loginDate"   orm:"login_date"   description:"最后登录时间"`
	CreatedDept int64               `json:"createdDept" orm:"created_dept" description:"创建部门"`
	CreatedBy   int64               `json:"createdBy"   orm:"created_by"   description:"创建者"`
	CreatedAt   *gtime.Time         `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	Remark      string              `json:"remark"      orm:"remark"       description:"备注"`
	DeptInfo    *DeptMiniModel      `json:"deptInfo"    orm:"with:dept_id=dept_id"    description:"部门信息"`
	Roles       []*SysRoleMiniModel `json:"roles"       description:"角色信息"`
	Posts       []*SysPostMiniModel `json:"posts"       description:"岗位信息"`
}

type UserProfileModel struct {
	SysUserViewModel
}

type UpdateCurrentUserModel struct {
	NickName    string `json:"nickName"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Avatar      string `json:"avatar"`
	Sex         string `json:"sex"`
}

type UpdateCurrentUserPasswordModel struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type SysUserAddModel struct {
	TenantId    string      `json:"tenantId"    orm:"tenant_id"    description:"租户编号"`
	DeptId      int64       `json:"deptId"      orm:"dept_id"      description:"部门ID"`
	UserName    string      `json:"userName"    orm:"user_name"    description:"用户账号"`
	NickName    string      `json:"nickName"    orm:"nick_name"    description:"用户昵称"`
	UserType    string      `json:"userType"    orm:"user_type"    description:"用户类型（sys_user系统用户）"`
	Email       string      `json:"email"       orm:"email"        description:"用户邮箱"`
	Phonenumber string      `json:"phonenumber" orm:"phonenumber"  description:"手机号码"`
	Sex         string      `json:"sex"         orm:"sex"          description:"用户性别（0男 1女 2未知）"`
	Avatar      string      `json:"avatar"      orm:"avatar"       description:"头像地址"`
	Salt        string      `json:"salt"        orm:"salt"         description:"加密盐"`
	Password    string      `json:"password"    orm:"password"     description:"密码"`
	Status      string      `json:"status"      orm:"status"       description:"帐号状态（0正常 1停用）"`
	LoginIp     string      `json:"loginIp"     orm:"login_ip"     description:"最后登录IP"`
	LoginDate   *gtime.Time `json:"loginDate"   orm:"login_date"   description:"最后登录时间"`
	CreatedDept int64       `json:"createdDept" orm:"created_dept" description:"创建部门"`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:"创建者"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"   description:"更新者"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
	PostIds     []int64     `json:"postIds"     orm:"post_ids"     description:"岗位ID列表"`
	RoleIds     []int64     `json:"roleIds"     orm:"role_ids"     description:"角色ID列表"`
}

type SysUserUpdateModel struct {
	UserId      int64       `json:"userId"      orm:"user_id"      description:"用户ID"`
	DeptId      *int64      `json:"deptId"      orm:"dept_id"      description:"部门ID"`
	UserName    *string     `json:"userName"    orm:"user_name"    description:"用户账号"`
	NickName    *string     `json:"nickName"    orm:"nick_name"    description:"用户昵称"`
	UserType    *string     `json:"userType"    orm:"user_type"    description:"用户类型（sys_user系统用户）"`
	Email       *string     `json:"email"       orm:"email"        description:"用户邮箱"`
	Phonenumber *string     `json:"phonenumber" orm:"phonenumber"  description:"手机号码"`
	Sex         *string     `json:"sex"         orm:"sex"          description:"用户性别（0男 1女 2未知）"`
	Avatar      *string     `json:"avatar"      orm:"avatar"       description:"头像地址"`
	Status      *string     `json:"status"      orm:"status"       description:"帐号状态（0正常 1停用）"`
	UpdatedBy   *int64      `json:"updatedBy"   orm:"updated_by"   description:"更新者"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
	Remark      *string     `json:"remark"      orm:"remark"       description:"备注"`
	PostIds     []int64     `json:"postIds"     orm:"post_ids"     description:"岗位ID列表"`
	RoleIds     []int64     `json:"roleIds"     orm:"role_ids"     description:"角色ID列表"`
}
