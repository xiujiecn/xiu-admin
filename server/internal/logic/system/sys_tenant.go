package system

import (
	"context"
	"errors"
	"slices"
	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/library/contexts"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/do"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
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
func (l *sSysTenant) View(ctx context.Context, param *model.SysTenantViewParam) (data *model.SysTenantViewModel, err error) {
	db := dao.SysTenant.Ctx(ctx)
	if param.Id != 0 {
		db = db.Where(dao.SysTenant.Columns().Id, param.Id)
	} else if param.TenantId != "" {
		db = db.Where(dao.SysTenant.Columns().TenantId, param.TenantId)
	} else {
		return nil, errors.New("参数错误")
	}
	err = db.Scan(&data)
	return
}

// 获取租户列表
func (l *sSysTenant) List(ctx context.Context, param *model.SysTenantListParam) (data []*model.SysTenantListModel, total int, err error) {
	db := dao.SysTenant.Ctx(ctx)
	if param.TenantId != "" {
		db = db.Where(dao.SysTenant.Columns().TenantId, param.TenantId)
	}
	if param.ContactUserName != "" {
		db = db.WhereLike(dao.SysTenant.Columns().ContactUserName, "%"+param.ContactUserName+"%")
	}
	if param.ContactPhone != "" {
		db = db.WhereLike(dao.SysTenant.Columns().ContactPhone, "%"+param.ContactPhone+"%")
	}
	if param.CompanyName != "" {
		db = db.WhereLike(dao.SysTenant.Columns().CompanyName, "%"+param.CompanyName+"%")
	}
	if param.LicenseNumber != "" {
		db = db.WhereLike(dao.SysTenant.Columns().LicenseNumber, "%"+param.LicenseNumber+"%")
	}

	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}

	err = db.Page(param.Page, param.PageSize).OrderAsc(dao.SysTenant.Columns().PackageId).Scan(&data)
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func (l *sSysTenant) Add(ctx context.Context, param *model.SysTenantAddParam) (output *model.SysTenantAddModel, err error) {
	data := &do.SysTenant{}
	gconv.Struct(param, &data)
	data.CreatedAt = gtime.Now()
	data.CreatedBy = contexts.GetUserId(ctx)
	data.CreatedDept = contexts.GetDeptId(ctx)

	db := dao.SysTenant.Ctx(ctx)
	id, err := db.Data(data).OmitNil().InsertAndGetId()
	if err != nil {
		return nil, err
	}
	output = &model.SysTenantAddModel{
		Id: id,
	}
	return
}

func (l *sSysTenant) Edit(ctx context.Context, param *model.SysTenantEditParam) (output *model.SysTenantEditModel, err error) {
	if param.Id == 0 {
		return nil, errors.New("参数错误")
	}
	data := &do.SysTenant{}
	gconv.Struct(param, &data)
	data.UpdatedAt = gtime.Now()
	data.UpdatedBy = contexts.GetUserId(ctx)
	db := dao.SysTenant.Ctx(ctx)
	db = db.Where(dao.SysTenant.Columns().Id, param.Id)
	_, err = db.Data(data).OmitNil().Update()
	if err != nil {
		return nil, err
	}
	output = &model.SysTenantEditModel{
		Id: param.Id,
	}
	return
}

func (l *sSysTenant) Delete(ctx context.Context, param *model.SysTenantDeleteParam) (output *model.SysTenantDeleteModel, err error) {
	if len(param.Ids) == 0 {
		return nil, errors.New("参数错误")
	}
	if slices.Contains(param.Ids, 1) {
		return nil, errors.New("不能删除默认租户")
	}
	data := &do.SysTenant{}
	gconv.Struct(param, &data)
	data.DeletedAt = gtime.Now()
	data.DeletedBy = contexts.GetUserId(ctx)

	db := dao.SysTenant.Ctx(ctx)
	db = db.WhereIn(dao.SysTenant.Columns().Id, param.Ids)
	_, err = db.Data(data).OmitNil().Update()
	if err != nil {
		return nil, err
	}
	output = &model.SysTenantDeleteModel{
		Ids: param.Ids,
	}
	return
}

func (l *sSysTenant) Status(ctx context.Context, param *model.SysTenantStatusParam) (output *model.SysTenantStatusModel, err error) {
	if param.Id == 0 {
		return nil, errors.New("参数错误")
	}
	data := &do.SysTenant{}
	gconv.Struct(param, &data)
	data.UpdatedAt = gtime.Now()
	data.UpdatedBy = contexts.GetUserId(ctx)
	db := dao.SysTenant.Ctx(ctx)
	db = db.Where(dao.SysTenant.Columns().Id, param.Id)
	_, err = db.Data(data).OmitNil().Update()
	if err != nil {
		return nil, err
	}
	output = &model.SysTenantStatusModel{
		Id: param.Id,
	}
	return
}
