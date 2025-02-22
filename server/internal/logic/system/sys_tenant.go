package system

import (
	"context"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/entity"
	"server/internal/model/request"
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

// 获取租户列表
func (l *sSysTenant) List(ctx context.Context, query *model.SysTenantListQuery, page *request.PageInfo) (data []*model.SysTenantListModel, total int, err error) {
	db := dao.SysTenant.Ctx(ctx)
	if query.TenantId != "" {
		db = db.Where(dao.SysTenant.Columns().TenantId, query.TenantId)
	}
	if query.ContactUserName != "" {
		db = db.WhereLike(dao.SysTenant.Columns().ContactUserName, "%"+query.ContactUserName+"%")
	}
	if query.ContactPhone != "" {
		db = db.WhereLike(dao.SysTenant.Columns().ContactPhone, "%"+query.ContactPhone+"%")
	}
	if query.CompanyName != "" {
		db = db.WhereLike(dao.SysTenant.Columns().CompanyName, "%"+query.CompanyName+"%")
	}
	if query.LicenseNumber != "" {
		db = db.WhereLike(dao.SysTenant.Columns().LicenseNumber, "%"+query.LicenseNumber+"%")
	}

	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}

	err = db.Page(page.Page, page.PageSize).Scan(&data)
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func (l *sSysTenant) TenantPackageList(ctx context.Context, query *model.SysTenantPackageListQuery, page *request.PageInfo) (data []*model.SysTenantPackageListModel, total int, err error) {
	db := dao.SysTenantPackage.Ctx(ctx)
	if query.PackageName != "" {
		db = db.WhereLike(dao.SysTenantPackage.Columns().PackageName, "%"+query.PackageName+"%")
	}

	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}

	err = db.Page(page.Page, page.PageSize).Scan(&data)
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}
