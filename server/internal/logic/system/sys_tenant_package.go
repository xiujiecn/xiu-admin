// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"errors"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/do"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysTenantPackage struct {
}

func NewSysTenantPackage() *sSysTenantPackage {
	return &sSysTenantPackage{}
}

func init() {
	service.RegisterSysTenantPackage(NewSysTenantPackage())
}

// 获取租户套餐
func (l *sSysTenantPackage) View(ctx context.Context, param *model.SysTenantPackageViewParam) (data *model.SysTenantPackageViewModel, err error) {
	if param.PackageId == 0 {
		return nil, errors.New("参数错误")
	}
	db := dao.SysTenantPackage.Ctx(ctx)
	db = db.Where(dao.SysTenantPackage.Columns().PackageId, param.PackageId)
	err = db.Scan(&data)
	return
}

func (l *sSysTenantPackage) List(ctx context.Context, param *model.SysTenantPackageListParam) (data []*model.SysTenantPackageListModel, total int, err error) {
	db := dao.SysTenantPackage.Ctx(ctx)
	if param.PackageName != "" {
		db = db.WhereLike(dao.SysTenantPackage.Columns().PackageName, "%"+param.PackageName+"%")
	}

	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}

	err = db.Page(param.Page, param.PageSize).OrderAsc(dao.SysTenantPackage.Columns().PackageId).Scan(&data)
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func (l *sSysTenantPackage) Add(ctx context.Context, param *model.SysTenantPackageAddParam) (output *model.SysTenantPackageAddModel, err error) {
	data := &do.SysTenantPackage{}
	gconv.Struct(param, &data)
	data.CreatedAt = gtime.Now()
	data.CreatedBy = contexts.GetUserId(ctx)
	data.CreatedDept = contexts.GetDeptId(ctx)

	db := dao.SysTenantPackage.Ctx(ctx)
	id, err := db.Data(data).OmitNil().InsertAndGetId()
	if err != nil {
		return nil, err
	}
	output = &model.SysTenantPackageAddModel{
		PackageId: id,
	}
	return
}

func (l *sSysTenantPackage) Edit(ctx context.Context, param *model.SysTenantPackageEditParam) (output *model.SysTenantPackageEditModel, err error) {
	if param.PackageId == 0 {
		return nil, errors.New("参数错误")
	}
	data := &do.SysTenantPackage{}
	gconv.Struct(param, &data)
	data.UpdatedAt = gtime.Now()
	data.UpdatedBy = contexts.GetUserId(ctx)

	db := dao.SysTenantPackage.Ctx(ctx)
	db = db.Where(dao.SysTenantPackage.Columns().PackageId, param.PackageId)
	_, err = db.Data(data).OmitNil().Update()
	if err != nil {
		return nil, err
	}
	output = &model.SysTenantPackageEditModel{
		PackageId: param.PackageId,
	}
	return
}

func (l *sSysTenantPackage) Delete(ctx context.Context, param *model.SysTenantPackageDeleteParam) (output *model.SysTenantPackageDeleteModel, err error) {
	if len(param.PackageIds) == 0 {
		return nil, errors.New("参数错误")
	}
	data := &do.SysTenantPackage{}
	gconv.Struct(param, &data)
	data.DeletedAt = gtime.Now()
	data.DeletedBy = contexts.GetUserId(ctx)

	db := dao.SysTenantPackage.Ctx(ctx)
	db = db.WhereIn(dao.SysTenantPackage.Columns().PackageId, param.PackageIds)
	_, err = db.Data(data).OmitNil().Update()
	if err != nil {
		return nil, err
	}
	output = &model.SysTenantPackageDeleteModel{
		PackageIds: param.PackageIds,
	}
	return
}

func (l *sSysTenantPackage) Status(ctx context.Context, param *model.SysTenantPackageStatusParam) (output *model.SysTenantPackageStatusModel, err error) {
	if param.PackageId == 0 {
		return nil, errors.New("参数错误")
	}
	data := &do.SysTenantPackage{}
	gconv.Struct(param, &data)
	data.UpdatedAt = gtime.Now()
	data.UpdatedBy = contexts.GetUserId(ctx)

	db := dao.SysTenantPackage.Ctx(ctx)
	db = db.Where(dao.SysTenantPackage.Columns().PackageId, param.PackageId)
	_, err = db.Data(data).OmitNil().Update()
	if err != nil {
		return nil, err
	}
	output = &model.SysTenantPackageStatusModel{
		PackageId: param.PackageId,
	}
	return
}
