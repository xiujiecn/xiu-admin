package system

import (
	"context"
	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/service"
)

type sSysConfig struct {
}

func NewSysConfig() *sSysConfig {
	return &sSysConfig{}
}

func init() {
	service.RegisterSysConfig(NewSysConfig())
}

func (s *sSysConfig) GetConfigList(ctx context.Context, query *model.SysConfigListParam, page *request.PageInfo) (items []*model.SysConfig, total int, err error) {
	// 获取当前用户租户ID
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, 0, err
	}
	tenantId := claims.TenantId

	db := dao.SysConfig.Ctx(ctx)

	db = db.Where("tenant_id = ?", tenantId)

	if query.ConfigName != "" {
		db = db.Where("config_name like ?", "%"+query.ConfigName+"%")
	}

	if query.ConfigKey != "" {
		db = db.Where("config_key like ?", "%"+query.ConfigKey+"%")
	}

	if query.ConfigType != "" {
		db = db.Where("config_type = ?", query.ConfigType)
	}

	if query.CreatedAt != nil {
		db = db.Where("created_at >= ?", query.CreatedAt)
	}

	db = db.Order("created_at desc")

	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}

	err = db.Page(page.Page, page.PageSize).Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return
}
