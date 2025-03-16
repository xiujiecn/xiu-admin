package model

import (
	"xiujieadmin/internal/model/entity"
	"xiujieadmin/internal/model/request"

	"github.com/gogf/gf/v2/os/gtime"
)

type SysTenantListModel struct {
	Id              int64       `json:"id"              orm:"id"                description:"id"`
	TenantId        string      `json:"tenantId"        orm:"tenant_id"         description:"租户编号"`
	ContactUserName string      `json:"contactUserName" orm:"contact_user_name" description:"联系人"`
	ContactPhone    string      `json:"contactPhone"    orm:"contact_phone"     description:"联系电话"`
	CompanyName     string      `json:"companyName"     orm:"company_name"      description:"企业名称"`
	LicenseNumber   string      `json:"licenseNumber"   orm:"license_number"    description:"统一社会信用代码"`
	Address         string      `json:"address"         orm:"address"           description:"地址"`
	Intro           string      `json:"intro"           orm:"intro"             description:"企业简介"`
	Domain          string      `json:"domain"          orm:"domain"            description:"域名"`
	Remark          string      `json:"remark"          orm:"remark"            description:"备注"`
	PackageId       int64       `json:"packageId"       orm:"package_id"        description:"租户套餐编号"`
	ExpireTime      *gtime.Time `json:"expireTime"      orm:"expire_time"       description:"过期时间"`
	AccountCount    int         `json:"accountCount"    orm:"account_count"     description:"用户数量（-1不限制）"`
	Status          string      `json:"status"          orm:"status"            description:"租户状态（0正常 1停用）"`
	CreatedDept     int64       `json:"createdDept"     orm:"created_dept"      description:"创建部门"`
	CreatedBy       int64       `json:"createdBy"       orm:"created_by"        description:"创建者"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"        description:"创建时间"`
	UpdatedBy       int64       `json:"updatedBy"       orm:"updated_by"        description:"更新者"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"        description:"更新时间"`
}

type SysTenantListParam struct {
	request.PageInfo
	TenantId        string `json:"tenantId"        orm:"tenant_id"         description:"租户编号"`
	ContactUserName string `json:"contactUserName" orm:"contact_user_name" description:"联系人"`
	ContactPhone    string `json:"contactPhone"    orm:"contact_phone"     description:"联系电话"`
	CompanyName     string `json:"companyName"     orm:"company_name"      description:"企业名称"`
	LicenseNumber   string `json:"licenseNumber"   orm:"license_number"    description:"统一社会信用代码"`
}

type SysTenantStatusParam struct {
	Id     int64  `json:"id"     orm:"id"     description:"租户id"`
	Status string `json:"status" orm:"status" description:"租户状态"`
}

type SysTenantStatusModel struct {
	Id int64 `json:"id"     orm:"id"     description:"租户id"`
}

type SysTenantDeleteParam struct {
	Ids []int64 `json:"ids"  description:"租户id"`
}

type SysTenantDeleteModel struct {
	Ids []int64 `json:"ids"  description:"租户id"`
}
type SysTenantViewParam struct {
	Id       int64  `json:"id"  description:"租户id"`
	TenantId string `json:"tenantId"         description:"租户编号"`
}

type SysTenantViewModel struct {
	entity.SysTenant
}

type SysTenantAddParam struct {
	// TenantId        string      `json:"tenantId"         description:"租户编号"`
	ContactUserName string      `json:"contactUserName"  description:"联系人"`
	ContactPhone    string      `json:"contactPhone"     description:"联系电话"`
	CompanyName     string      `json:"companyName"      description:"企业名称"`
	LicenseNumber   string      `json:"licenseNumber"    description:"统一社会信用代码"`
	Address         string      `json:"address"          description:"地址"`
	Intro           string      `json:"intro"            description:"企业简介"`
	Domain          string      `json:"domain"           description:"域名"`
	Remark          string      `json:"remark"           description:"备注"`
	PackageId       int64       `json:"packageId"        description:"租户套餐编号"`
	ExpireTime      *gtime.Time `json:"expireTime"       description:"过期时间"`
	AccountCount    int         `json:"accountCount"     description:"用户数量（-1不限制）"`
	Status          *string     `json:"status,omitempty" description:"租户状态（0正常 1停用）"`
	Username        string      `json:"username"         description:"用户账号"`
	Password        string      `json:"password"         description:"用户密码"`
}

type SysTenantAddModel struct {
	Id int64 `json:"id"                   description:"id"`
}

type SysTenantEditParam struct {
	Id              int64       `json:"id"              description:"id"`
	ContactUserName *string     `json:"contactUserName,omitempty" description:"联系人"`
	ContactPhone    *string     `json:"contactPhone,omitempty"    description:"联系电话"`
	CompanyName     *string     `json:"companyName,omitempty"     description:"企业名称"`
	LicenseNumber   *string     `json:"licenseNumber,omitempty"   description:"统一社会信用代码"`
	Address         *string     `json:"address,omitempty"         description:"地址"`
	Intro           *string     `json:"intro,omitempty"           description:"企业简介"`
	Domain          *string     `json:"domain,omitempty"          description:"域名"`
	Remark          *string     `json:"remark,omitempty"          description:"备注"`
	PackageId       *int64      `json:"packageId,omitempty"       description:"租户套餐编号"`
	ExpireTime      *gtime.Time `json:"expireTime,omitempty"      description:"过期时间"`
	AccountCount    *int        `json:"accountCount,omitempty"    description:"用户数量（-1不限制）"`
	Status          *string     `json:"status,omitempty"          description:"租户状态（0正常 1停用）"`
}

type SysTenantEditModel struct {
	Id int64 `json:"id"              orm:"id"                description:"id"`
}
