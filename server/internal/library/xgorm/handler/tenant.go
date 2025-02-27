package handler

import (
	"xiujieadmin/internal/library/contexts"
	"xiujieadmin/utility/convert"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
)

// FilterTenant 过滤多租户数据权限
func FilterTenant(m *gdb.Model) *gdb.Model {
	var (
		ctx    = m.GetCtx()
		fields = convert.EscapeFieldsToSlice(m.GetFieldsStr())
	)
	// 如果包含tenant_id字段，则过滤多租户数据权限
	if gstr.InArray(fields, "tenant_id") {
		m = m.Where("tenant_id", contexts.GetTenantId(ctx))
	}
	return m
}
