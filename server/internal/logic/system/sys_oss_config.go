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
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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
	if param.OssConfigId > 0 {
		db = db.Where(dao.SysOssConfig.Columns().OssConfigId, param.OssConfigId)
	} else if param.ConfigKey != "" {
		db = db.Where(dao.SysOssConfig.Columns().ConfigKey, param.ConfigKey)
	} else {
		g.Log().Errorf(ctx, "sSysOssConfig View 参数错误: %+v", param)
		return nil, gerror.New("参数错误")
	}
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

func (s *sSysOssConfig) GetAllUrlByService(ctx context.Context, path string, tenantId string, ossService string) (allUrl string, err error) {
	ossConfig := &entity.SysOssConfig{}
	db := s.Model(ctx)
	db = db.Where(dao.SysOssConfig.Columns().ConfigKey, ossService)
	db = db.Where(dao.SysOssConfig.Columns().TenantId, tenantId)
	err = db.Scan(&ossConfig)
	if err != nil {
		g.Log().Errorf(ctx, "sSysOss GetAllUrlByService error: %+v, url: %s ,service: %s", err, path, ossService)
		return "", err
	}
	if ossConfig == nil {
		g.Log().Errorf(ctx, "sSysOss GetAllUrlByService error: ossConfig is empty, url: %s ,service: %s", path, ossService)
		err = gerror.New("ossConfig is empty")
		return
	}
	domain := ossConfig.Endpoint
	if ossConfig.Domain != "" {
		domain = ossConfig.Domain
	}
	if ossConfig.IsHttps == "Y" {
		allUrl = "https://" + domain + path
	} else {
		allUrl = "http://" + domain + path
	}
	g.Log().Infof(ctx, "sSysOss GetAllUrlByService success: %s, url: %s, ossConfig: %+v", ossService, allUrl, ossConfig)
	return
}
