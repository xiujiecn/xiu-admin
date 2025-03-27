// package model
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package model

import "github.com/gogf/gf/v2/os/gtime"

type SysSocialListModel struct {
	Id               int64       `json:"id"               orm:"id"                 description:"主键"`
	UserId           int64       `json:"userId"           orm:"user_id"            description:"用户ID"`
	TenantId         string      `json:"tenantId"         orm:"tenant_id"          description:"租户id"`
	AuthId           string      `json:"authId"           orm:"auth_id"            description:"平台+平台唯一id"`
	Source           string      `json:"source"           orm:"source"             description:"用户来源"`
	OpenId           string      `json:"openId"           orm:"open_id"            description:"平台编号唯一id"`
	UserName         string      `json:"userName"         orm:"user_name"          description:"登录账号"`
	NickName         string      `json:"nickName"         orm:"nick_name"          description:"用户昵称"`
	Email            string      `json:"email"            orm:"email"              description:"用户邮箱"`
	Avatar           string      `json:"avatar"           orm:"avatar"             description:"头像地址"`
	AccessToken      string      `json:"accessToken"      orm:"access_token"       description:"用户的授权令牌"`
	ExpireIn         int         `json:"expireIn"         orm:"expire_in"          description:"用户的授权令牌的有效期，部分平台可能没有"`
	RefreshToken     string      `json:"refreshToken"     orm:"refresh_token"      description:"刷新令牌，部分平台可能没有"`
	AccessCode       string      `json:"accessCode"       orm:"access_code"        description:"平台的授权信息，部分平台可能没有"`
	UnionId          string      `json:"unionId"          orm:"union_id"           description:"用户的 unionid"`
	Scope            string      `json:"scope"            orm:"scope"              description:"授予的权限，部分平台可能没有"`
	TokenType        string      `json:"tokenType"        orm:"token_type"         description:"个别平台的授权信息，部分平台可能没有"`
	IdToken          string      `json:"idToken"          orm:"id_token"           description:"id token，部分平台可能没有"`
	MacAlgorithm     string      `json:"macAlgorithm"     orm:"mac_algorithm"      description:"小米平台用户的附带属性，部分平台可能没有"`
	MacKey           string      `json:"macKey"           orm:"mac_key"            description:"小米平台用户的附带属性，部分平台可能没有"`
	Code             string      `json:"code"             orm:"code"               description:"用户的授权code，部分平台可能没有"`
	OauthToken       string      `json:"oauthToken"       orm:"oauth_token"        description:"Twitter平台用户的附带属性，部分平台可能没有"`
	OauthTokenSecret string      `json:"oauthTokenSecret" orm:"oauth_token_secret" description:"Twitter平台用户的附带属性，部分平台可能没有"`
	CreatedDept      int64       `json:"createdDept"      orm:"created_dept"       description:"创建部门"`
	CreatedBy        int64       `json:"createdBy"        orm:"created_by"         description:"创建者"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:"创建时间"`
}

type SysSocialListParam struct {
	UserId int64 `json:"userId" description:"用户ID"`
}
