// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOss is the golang structure of table sys_oss for DAO operations like Where/Data.
type SysOss struct {
	g.Meta       `orm:"table:sys_oss, do:true"`
	OssId        interface{} // 对象存储主键
	TenantId     interface{} // 租户编号
	FileName     interface{} // 文件名
	OriginalName interface{} // 原名
	FileSuffix   interface{} // 文件后缀名
	Path         interface{} // 存储路径
	Url          interface{} // URL地址
	CreatedDept  interface{} // 创建部门
	CreatedAt    *gtime.Time // 创建时间
	CreatedBy    interface{} // 创建者
	UpdatedAt    *gtime.Time // 更新时间
	UpdatedBy    interface{} // 更新者
	Service      interface{} // 服务商
	Md5          interface{} // 文件MD5
	FileSize     interface{} // 文件大小
	FileCrc16    interface{} // 文件Crc16
	FileSum      interface{} // 文件校验和
}
