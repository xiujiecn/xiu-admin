package system

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"
)

type sSysTenant struct {
}

func NewSysTenant() *sSysTenant {
	return &sSysTenant{}
}

func init() {
	service.RegisterSysTenant(NewSysTenant())
}

// 获取租户信息
func (l *sSysTenant) GetTenantInfo(ctx context.Context, tenantId string) (data *entity.SysTenant, err error) {
	err = dao.SysTenant.Ctx(ctx).Where(dao.SysTenant.Columns().TenantId, tenantId).Scan(&data)
	return
}

// 获取租户套餐
func (l *sSysTenant) GetTenantPackage(ctx context.Context, packageId int64) (data *entity.SysTenantPackage, err error) {
	err = dao.SysTenantPackage.Ctx(ctx).Where(dao.SysTenantPackage.Columns().PackageId, packageId).Scan(&data)
	return
}
