package system

import (
	"context"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/request"
	"server/internal/service"
)

type sSysClient struct {
}

func NewSysClient() *sSysClient {
	return &sSysClient{}
}

func init() {
	service.RegisterSysClient(NewSysClient())
}

func (s *sSysClient) List(ctx context.Context, query *model.SysClientListQuery, pageInfo *request.PageInfo) (items []*model.SysClientListModel, total int, err error) {

	db := dao.SysClient.Ctx(ctx)
	if query.ClientId != "" {
		db = db.Where(dao.SysClient.Columns().ClientId, query.ClientId)
	}

	if query.ClientKey != "" {
		db = db.Where(dao.SysClient.Columns().ClientKey, query.ClientKey)
	}

	if query.ClientSecret != "" {
		db = db.Where(dao.SysClient.Columns().ClientSecret, query.ClientSecret)
	}

	if query.Status != "" {
		db = db.Where(dao.SysClient.Columns().Status, query.Status)
	}

	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}

	err = db.Page(pageInfo.Page, pageInfo.PageSize).Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
