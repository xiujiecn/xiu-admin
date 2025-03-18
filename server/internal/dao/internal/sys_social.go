// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysSocialDao is the data access object for the table sys_social.
type SysSocialDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysSocialColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysSocialColumns defines and stores column names for the table sys_social.
type SysSocialColumns struct {
	Id               string // 主键
	UserId           string // 用户ID
	TenantId         string // 租户id
	AuthId           string // 平台+平台唯一id
	Source           string // 用户来源
	OpenId           string // 平台编号唯一id
	UserName         string // 登录账号
	NickName         string // 用户昵称
	Email            string // 用户邮箱
	Avatar           string // 头像地址
	AccessToken      string // 用户的授权令牌
	ExpireIn         string // 用户的授权令牌的有效期，部分平台可能没有
	RefreshToken     string // 刷新令牌，部分平台可能没有
	AccessCode       string // 平台的授权信息，部分平台可能没有
	UnionId          string // 用户的 unionid
	Scope            string // 授予的权限，部分平台可能没有
	TokenType        string // 个别平台的授权信息，部分平台可能没有
	IdToken          string // id token，部分平台可能没有
	MacAlgorithm     string // 小米平台用户的附带属性，部分平台可能没有
	MacKey           string // 小米平台用户的附带属性，部分平台可能没有
	Code             string // 用户的授权code，部分平台可能没有
	OauthToken       string // Twitter平台用户的附带属性，部分平台可能没有
	OauthTokenSecret string // Twitter平台用户的附带属性，部分平台可能没有
	CreatedDept      string // 创建部门
	CreatedBy        string // 创建者
	CreatedAt        string // 创建时间
	UpdatedBy        string // 更新者
	UpdatedAt        string // 更新时间
	DeletedBy        string // 删除人
	DeletedAt        string // 删除时间
}

// sysSocialColumns holds the columns for the table sys_social.
var sysSocialColumns = SysSocialColumns{
	Id:               "id",
	UserId:           "user_id",
	TenantId:         "tenant_id",
	AuthId:           "auth_id",
	Source:           "source",
	OpenId:           "open_id",
	UserName:         "user_name",
	NickName:         "nick_name",
	Email:            "email",
	Avatar:           "avatar",
	AccessToken:      "access_token",
	ExpireIn:         "expire_in",
	RefreshToken:     "refresh_token",
	AccessCode:       "access_code",
	UnionId:          "union_id",
	Scope:            "scope",
	TokenType:        "token_type",
	IdToken:          "id_token",
	MacAlgorithm:     "mac_algorithm",
	MacKey:           "mac_key",
	Code:             "code",
	OauthToken:       "oauth_token",
	OauthTokenSecret: "oauth_token_secret",
	CreatedDept:      "created_dept",
	CreatedBy:        "created_by",
	CreatedAt:        "created_at",
	UpdatedBy:        "updated_by",
	UpdatedAt:        "updated_at",
	DeletedBy:        "deleted_by",
	DeletedAt:        "deleted_at",
}

// NewSysSocialDao creates and returns a new DAO object for table data access.
func NewSysSocialDao(handlers ...gdb.ModelHandler) *SysSocialDao {
	return &SysSocialDao{
		group:    "default",
		table:    "sys_social",
		columns:  sysSocialColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysSocialDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysSocialDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysSocialDao) Columns() SysSocialColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysSocialDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysSocialDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysSocialDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
