// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysOssConfigDao is the data access object for the table sys_oss_config.
type SysOssConfigDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  SysOssConfigColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// SysOssConfigColumns defines and stores column names for the table sys_oss_config.
type SysOssConfigColumns struct {
	OssConfigId  string // 主键
	TenantId     string // 租户编号
	ConfigKey    string // 配置key
	AccessKey    string // accessKey
	SecretKey    string // 秘钥
	BucketName   string // 桶名称
	Prefix       string // 前缀
	Endpoint     string // 访问站点
	Domain       string // 自定义域名
	IsHttps      string // 是否https（Y=是,N=否）
	Region       string // 域
	AccessPolicy string // 桶权限类型(0=private 1=public 2=custom)
	Status       string // 是否默认（0=是,1=否）
	Ext1         string // 扩展字段
	CreatedDept  string // 创建部门
	CreatedBy    string // 创建者
	CreatedAt    string // 创建时间
	UpdatedBy    string // 更新者
	UpdatedAt    string // 更新时间
	Remark       string // 备注
}

// sysOssConfigColumns holds the columns for the table sys_oss_config.
var sysOssConfigColumns = SysOssConfigColumns{
	OssConfigId:  "oss_config_id",
	TenantId:     "tenant_id",
	ConfigKey:    "config_key",
	AccessKey:    "access_key",
	SecretKey:    "secret_key",
	BucketName:   "bucket_name",
	Prefix:       "prefix",
	Endpoint:     "endpoint",
	Domain:       "domain",
	IsHttps:      "is_https",
	Region:       "region",
	AccessPolicy: "access_policy",
	Status:       "status",
	Ext1:         "ext1",
	CreatedDept:  "created_dept",
	CreatedBy:    "created_by",
	CreatedAt:    "created_at",
	UpdatedBy:    "updated_by",
	UpdatedAt:    "updated_at",
	Remark:       "remark",
}

// NewSysOssConfigDao creates and returns a new DAO object for table data access.
func NewSysOssConfigDao(handlers ...gdb.ModelHandler) *SysOssConfigDao {
	return &SysOssConfigDao{
		group:    "default",
		table:    "sys_oss_config",
		columns:  sysOssConfigColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysOssConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysOssConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysOssConfigDao) Columns() SysOssConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysOssConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysOssConfigDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysOssConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
