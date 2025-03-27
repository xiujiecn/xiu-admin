// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/do"
	"xiuadmin/internal/service"
	"xiuadmin/utility"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
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
	err = g.DB().Transaction(context.TODO(), func(ctx context.Context, tx gdb.TX) error {
		data.TenantId = "100000"
		id, err := tx.Ctx(ctx).Model(dao.SysTenant.Table()).Data(data).OmitNil().InsertAndGetId()
		if err != nil {
			return err
		}
		tenantId := fmt.Sprintf("%06d", 100000+id)
		_, err = tx.Ctx(ctx).Model(dao.SysTenant.Table()).Data(g.Map{
			dao.SysTenant.Columns().TenantId: tenantId,
		}).Where(dao.SysTenant.Columns().Id, id).Update()
		if err != nil {
			return err
		}
		// 创建部门
		dataDeptInsert := do.SysDept{}
		dataDeptInsert.TenantId = tenantId
		dataDeptInsert.DeptName = param.CompanyName
		dataDeptInsert.ParentId = 0
		dataDeptInsert.OrderNum = 0
		dataDeptInsert.Status = consts.SysDeptStatusNormal
		dataDeptInsert.CreatedBy = contexts.GetUserId(ctx)
		dataDeptInsert.CreatedAt = gtime.Now()
		dataDeptInsert.UpdatedBy = contexts.GetUserId(ctx)
		dataDeptInsert.UpdatedAt = gtime.Now()
		deptId, err := tx.Ctx(ctx).Model(dao.SysDept.Table()).Data(dataDeptInsert).OmitNil().InsertAndGetId()
		if err != nil {
			return err
		}

		// 创建系统管理员
		salt := utility.RandomString(5)
		password := utility.PasswordEncrypt(param.Password, salt)

		dataUserInsert := do.SysUser{}
		dataUserInsert.TenantId = tenantId
		dataUserInsert.DeptId = deptId
		dataUserInsert.NickName = param.ContactUserName
		dataUserInsert.UserName = param.Username
		dataUserInsert.Password = password
		dataUserInsert.UserType = consts.SysUserTypeSys
		dataUserInsert.Salt = salt
		dataUserInsert.Phonenumber = param.ContactPhone
		dataUserInsert.Email = ""
		dataUserInsert.Sex = consts.SysUserSexUnknown
		dataUserInsert.Status = consts.SysUserStatusNormal
		dataUserInsert.CreatedDept = contexts.GetDeptId(ctx)
		dataUserInsert.CreatedBy = contexts.GetUserId(ctx)
		dataUserInsert.CreatedAt = gtime.Now()
		dataUserInsert.UpdatedBy = contexts.GetUserId(ctx)
		dataUserInsert.UpdatedAt = gtime.Now()

		userId, err := tx.Ctx(ctx).Model(dao.SysUser.Table()).Data(dataUserInsert).OmitNil().InsertAndGetId()
		if err != nil {
			return err
		}
		// 修改部门负责人
		dataDeptUpdate := do.SysDept{}
		dataDeptUpdate.DeptId = deptId
		dataDeptUpdate.Leader = userId
		dataDeptUpdate.UpdatedBy = contexts.GetUserId(ctx)
		dataDeptUpdate.UpdatedAt = gtime.Now()
		_, err = tx.Ctx(ctx).Model(dao.SysDept.Table()).Data(dataDeptUpdate).OmitNil().Where(dao.SysDept.Columns().DeptId, deptId).Update()
		if err != nil {
			return err
		}
		output = &model.SysTenantAddModel{
			Id: id,
		}
		return nil
	})
	if err != nil {
		g.Log().Errorf(ctx, "sSysTenant.Add err: %v", err)
		return nil, err
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
