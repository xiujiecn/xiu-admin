// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysSocial is the golang structure of table sys_social for DAO operations like Where/Data.
type SysSocial struct {
	g.Meta           `orm:"table:sys_social, do:true"`
	Id               interface{} // 主键
	UserId           interface{} // 用户ID
	TenantId         interface{} // 租户id
	AuthId           interface{} // 平台+平台唯一id
	Source           interface{} // 用户来源
	OpenId           interface{} // 平台编号唯一id
	UserName         interface{} // 登录账号
	NickName         interface{} // 用户昵称
	Email            interface{} // 用户邮箱
	Avatar           interface{} // 头像地址
	AccessToken      interface{} // 用户的授权令牌
	ExpireIn         interface{} // 用户的授权令牌的有效期，部分平台可能没有
	RefreshToken     interface{} // 刷新令牌，部分平台可能没有
	AccessCode       interface{} // 平台的授权信息，部分平台可能没有
	UnionId          interface{} // 用户的 unionid
	Scope            interface{} // 授予的权限，部分平台可能没有
	TokenType        interface{} // 个别平台的授权信息，部分平台可能没有
	IdToken          interface{} // id token，部分平台可能没有
	MacAlgorithm     interface{} // 小米平台用户的附带属性，部分平台可能没有
	MacKey           interface{} // 小米平台用户的附带属性，部分平台可能没有
	Code             interface{} // 用户的授权code，部分平台可能没有
	OauthToken       interface{} // Twitter平台用户的附带属性，部分平台可能没有
	OauthTokenSecret interface{} // Twitter平台用户的附带属性，部分平台可能没有
	CreatedDept      interface{} // 创建部门
	CreatedBy        interface{} // 创建者
	CreatedAt        *gtime.Time // 创建时间
	UpdatedBy        interface{} // 更新者
	UpdatedAt        *gtime.Time // 更新时间
	DeletedBy        interface{} // 删除人
	DeletedAt        *gtime.Time // 删除时间
}
