package system

import (
	"context"
	"errors"
	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/library/contexts"
	"xiujieadmin/internal/library/xgorm/handler"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/do"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysDictType struct{}

func NewSysDictType() *sSysDictType {
	return &sSysDictType{}
}

func init() {
	service.RegisterSysDictType(NewSysDictType())
}

func (l *sSysDictType) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) > 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.SysDictType.Ctx(ctx), option...)
}

func (s *sSysDictType) List(ctx context.Context, param *model.SysDictTypeListParam) (items []model.SysDictTypeListModel, total int, err error) {
	m := s.Model(ctx)
	if param.DictName != "" {
		m = m.WhereLike(dao.SysDictType.Columns().DictName, "%"+param.DictName+"%")
	}
	if param.DictType != "" {
		m = m.WhereLike(dao.SysDictType.Columns().DictType, "%"+param.DictType+"%")
	}
	if len(param.CreatedAt) == 2 {
		createdAt := gtime.NewFromStr(param.CreatedAt[0])
		createdAtEnd := gtime.NewFromStr(param.CreatedAt[1])
		m = m.WhereBetween(dao.SysDictType.Columns().CreatedAt, createdAt, createdAtEnd.EndOfDay())
	}
	total, err = m.Count()
	if err != nil {
		return nil, 0, err
	}
	err = m.Order(dao.SysDictType.Columns().DictId, "DESC").Page(param.Page, param.PageSize).Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return
}

func (s *sSysDictType) View(ctx context.Context, param *model.SysDictTypeViewParam) (data *model.SysDictTypeViewModel, err error) {
	m := s.Model(ctx)

	if param.DictId != 0 {
		m = m.Where(dao.SysDictType.Columns().DictId, param.DictId)
	} else if param.DictType != "" {
		m = m.Where(dao.SysDictType.Columns().DictType, param.DictType)
	} else {
		return nil, errors.New("字典类型ID不能为空")
	}
	err = m.Scan(&data)
	if err != nil {
		return nil, err
	}
	return
}

func (s *sSysDictType) Add(ctx context.Context, param *model.SysDictTypeAddParam) (output *model.SysDictTypeAddModel, err error) {
	m := s.Model(ctx)
	data := do.SysDictType{}
	gconv.Struct(param, &data)
	data.CreatedAt = gtime.Now()
	data.CreatedBy = contexts.GetUserId(ctx)
	data.CreatedDept = contexts.GetDeptId(ctx)
	m = m.Data(data).OmitNil()
	output = &model.SysDictTypeAddModel{}
	output.DictId, err = m.InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return
}

func (s *sSysDictType) Edit(ctx context.Context, param *model.SysDictTypeEditParam) (output *model.SysDictTypeEditModel, err error) {
	if param.DictId == 0 {
		g.Log().Errorf(ctx, "字典类型ID不能为空,param:%+v", param)
		return nil, errors.New("字典类型ID不能为空")
	}
	m := s.Model(ctx)
	data := do.SysDictType{}
	gconv.Struct(param, &data)
	data.UpdatedAt = gtime.Now()
	data.UpdatedBy = contexts.GetUserId(ctx)
	m = m.Data(data).OmitNil().Where(dao.SysDictType.Columns().DictId, param.DictId)
	output = &model.SysDictTypeEditModel{
		DictId: param.DictId,
	}
	_, err = m.Update()
	if err != nil {
		return nil, err
	}
	return
}

func (s *sSysDictType) Delete(ctx context.Context, param *model.SysDictTypeDeleteParam) (output *model.SysDictTypeDeleteModel, err error) {
	if len(param.DictIds) == 0 {
		return nil, errors.New("字典类型ID不能为空")
	}
	m := s.Model(ctx)
	m = m.WhereIn(dao.SysDictType.Columns().DictId, param.DictIds)
	_, err = m.Delete()
	if err != nil {
		return nil, err
	}
	output = &model.SysDictTypeDeleteModel{
		DictIds: param.DictIds,
	}
	return
}
