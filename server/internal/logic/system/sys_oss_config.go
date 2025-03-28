// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/do"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysOssConfig struct {
}

func NewSysOssConfig() *sSysOssConfig {
	return &sSysOssConfig{}
}

func init() {
	service.RegisterSysOssConfig(NewSysOssConfig())
}

func (s *sSysOssConfig) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.SysOssConfig.Ctx(ctx), option...)
}

func (s *sSysOssConfig) List(ctx context.Context, param *model.SysOssConfigListParam) (items []*model.SysOssConfigListModel, total int, err error) {

	db := s.Model(ctx)

	if param.ConfigKey != "" {
		db = db.WhereLike(dao.SysOssConfig.Columns().ConfigKey, "%"+param.ConfigKey+"%")
	}

	if param.BucketName != "" {
		db = db.WhereLike(dao.SysOssConfig.Columns().BucketName, "%"+param.BucketName+"%")
	}
	if param.Status != "" {
		db = db.Where(dao.SysOssConfig.Columns().Status, param.Status)
	}

	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}

	err = db.Page(param.Page, param.PageSize).OrderAsc(dao.SysOssConfig.Columns().OssConfigId).Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return
}

func (s *sSysOssConfig) View(ctx context.Context, param *model.SysOssConfigViewParam) (item *model.SysOssConfigViewModel, err error) {
	db := s.Model(ctx)
	db = db.Where(dao.SysOssConfig.Columns().OssConfigId, param.OssConfigId)
	err = db.Scan(&item)
	return
}

func (s *sSysOssConfig) Add(ctx context.Context, param *model.SysOssConfigAddParam) (item *model.SysOssConfigAddModel, err error) {
	db := s.Model(ctx)
	data := &do.SysOssConfig{}
	gconv.Struct(param, data)
	data.CreatedAt = gtime.Now()
	data.CreatedBy = contexts.GetUserId(ctx)
	data.CreatedDept = contexts.GetDeptId(ctx)
	data.TenantId = contexts.GetTenantId(ctx)
	if len(param.Status) == 0 {
		data.Status = "1"
	}

	id, err := db.Data(data).OmitNil().InsertAndGetId()
	if err != nil {
		return nil, err
	}
	item = &model.SysOssConfigAddModel{
		OssConfigId: id,
	}
	return
}

func (s *sSysOssConfig) Edit(ctx context.Context, param *model.SysOssConfigEditParam) (item *model.SysOssConfigEditModel, err error) {
	item = &model.SysOssConfigEditModel{
		OssConfigId: param.OssConfigId,
	}
	db := s.Model(ctx)
	data := &do.SysOssConfig{}
	gconv.Struct(param, data)
	data.OssConfigId = nil
	data.UpdatedAt = gtime.Now()
	data.UpdatedBy = contexts.GetUserId(ctx)
	if param.Status == nil || *param.Status == "" {
		data.Status = nil
	}

	_, err = db.Data(data).OmitNil().Where(dao.SysOssConfig.Columns().OssConfigId, param.OssConfigId).Update()
	if err != nil {
		return item, err
	}
	return
}

func (s *sSysOssConfig) Delete(ctx context.Context, param *model.SysOssConfigDeleteParam) (item *model.SysOssConfigDeleteModel, err error) {
	item = &model.SysOssConfigDeleteModel{
		OssConfigIds: param.OssConfigIds,
	}
	db := s.Model(ctx)
	_, err = db.WhereIn(dao.SysOssConfig.Columns().OssConfigId, param.OssConfigIds).Delete()
	if err != nil {
		return item, err
	}
	return
}
