package system

import (
	"context"
	"errors"
	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/library/contexts"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/do"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
)

type sSysClient struct {
}

func NewSysClient() *sSysClient {
	return &sSysClient{}
}

func init() {
	service.RegisterSysClient(NewSysClient())
}

func (s *sSysClient) List(ctx context.Context, query *model.SysClientListParam) (items []*model.SysClientListModel, total int, err error) {

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

	err = db.Page(query.Page, query.PageSize).OrderAsc(dao.SysClient.Columns().Id).Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
func (s *sSysClient) View(ctx context.Context, param *model.SysClientViewParam) (item *model.SysClientViewModel, err error) {
	db := dao.SysClient.Ctx(ctx)
	if param.Id != 0 {
		db = db.Where(dao.SysClient.Columns().Id, param.Id)
	} else if param.ClientId != "" {
		db = db.Where(dao.SysClient.Columns().ClientId, param.ClientId)
	} else {
		return nil, errors.New("未填写查询条件")
	}
	err = db.Scan(&item)
	return
}

func (s *sSysClient) Add(ctx context.Context, param *model.SysClientAddParam) (output *model.SysClientAddModel, err error) {
	data := &do.SysClient{}
	gconv.Struct(param, data)
	data.ClientId = guid.S()
	data.CreatedAt = gtime.Now()
	data.CreatedBy = contexts.GetUserId(ctx)
	data.CreatedDept = contexts.GetDeptId(ctx)

	db := dao.SysClient.Ctx(ctx)
	id, err := db.Data(data).OmitNil().InsertAndGetId()
	if err != nil {
		return
	}
	output = &model.SysClientAddModel{
		Id: id,
	}
	return
}

func (s *sSysClient) Edit(ctx context.Context, param *model.SysClientEditParam) (output *model.SysClientEditModel, err error) {
	data := &do.SysClient{}
	gconv.Struct(param, data)
	data.ClientId = nil
	data.UpdatedAt = gtime.Now()
	data.UpdatedBy = contexts.GetUserId(ctx)

	db := dao.SysClient.Ctx(ctx)
	_, err = db.Data(data).OmitNil().Where(dao.SysClient.Columns().Id, param.Id).Update()
	if err != nil {
		return
	}
	output = &model.SysClientEditModel{
		Id: param.Id,
	}
	return
}

func (s *sSysClient) Delete(ctx context.Context, param *model.SysClientDeleteParam) (output *model.SysClientDeleteModel, err error) {
	db := dao.SysClient.Ctx(ctx)
	data := &do.SysClient{}
	data.DeletedAt = gtime.Now()
	data.DeletedBy = contexts.GetUserId(ctx)
	_, err = db.Data(data).OmitNil().Where(dao.SysClient.Columns().Id, param.Ids).Update()
	if err != nil {
		return
	}
	output = &model.SysClientDeleteModel{
		Ids: param.Ids,
	}
	return
}

func (s *sSysClient) Status(ctx context.Context, param *model.SysClientStatusParam) (output *model.SysClientStatusModel, err error) {
	db := dao.SysClient.Ctx(ctx)
	_, err = db.Data(param).OmitNil().Where(dao.SysClient.Columns().Id, param.Id).Update()
	if err != nil {
		return
	}
	output = &model.SysClientStatusModel{
		Id: param.Id,
	}
	return
}
