package system

import (
	"context"
	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/library/contexts"
	"xiujieadmin/internal/library/xgorm/handler"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/do"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysDictData struct{}

func NewSysDictData() *sSysDictData {
	return &sSysDictData{}
}

func init() {
	service.RegisterSysDictData(NewSysDictData())
}

func (l *sSysDictData) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) > 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.SysDictData.Ctx(ctx), option...)
}

func (s *sSysDictData) List(ctx context.Context, param *model.SysDictDataListParam) (items []model.SysDictDataListModel, total int, err error) {
	// 获取当前用户租户编码
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, 0, err
	}
	tenantId := claims.TenantId
	typeData := &model.SysDictTypeViewModel{}
	// 获取字典类型
	if param.DictType == "" && param.DictId != 0 {
		typeData, err = service.SysDictType().View(ctx, &model.SysDictTypeViewParam{
			DictId: param.DictId,
		})
		if err != nil {
			return nil, 0, err
		}
		if typeData == nil {
			return nil, 0, gerror.New("字典类型不存在")
		}
		param.DictType = typeData.DictType
	}
	if param.DictType == "" {
		return nil, 0, gerror.New("字典类型不能为空")
	}
	// 获取字典数据
	items = make([]model.SysDictDataListModel, 0)
	m := s.Model(ctx).Where(dao.SysDictData.Columns().TenantId, tenantId)
	m = m.Where(dao.SysDictData.Columns().DictType, param.DictType)
	total, err = m.Count()
	if err != nil {
		return nil, 0, err
	}
	err = m.Order(dao.SysDictData.Columns().DictSort, "ASC").Page(param.Page, param.PageSize).Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return
}

func (s *sSysDictData) View(ctx context.Context, param *model.SysDictDataViewParam) (data *model.SysDictDataViewModel, err error) {
	if param.DictCode == 0 {
		return nil, gerror.New("字典编码不能为空")
	}

	m := s.Model(ctx)
	if param.DictCode != 0 {
		m = m.Where(dao.SysDictData.Columns().DictCode, param.DictCode)
	}
	err = m.Scan(&data)
	if err != nil {
		return nil, err
	}
	return
}

func (s *sSysDictData) Add(ctx context.Context, param *model.SysDictDataAddParam) (output *model.SysDictDataAddModel, err error) {
	m := s.Model(ctx)
	data := do.SysDictData{}
	gconv.Struct(param, &data)
	data.CreatedAt = gtime.Now()
	data.CreatedBy = contexts.GetUserId(ctx)
	data.CreatedDept = contexts.GetDeptId(ctx)
	m = m.Data(data).OmitNil()
	output = &model.SysDictDataAddModel{}
	output.DictCode, err = m.InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return
}

func (s *sSysDictData) Edit(ctx context.Context, param *model.SysDictDataEditParam) (output *model.SysDictDataEditModel, err error) {
	if param.DictCode == 0 {
		return nil, gerror.New("字典编码不能为空")
	}
	m := s.Model(ctx)
	data := do.SysDictData{}
	gconv.Struct(param, &data)
	data.UpdatedAt = gtime.Now()
	data.UpdatedBy = contexts.GetUserId(ctx)
	m = m.Data(data).OmitNil().Where(dao.SysDictData.Columns().DictCode, param.DictCode)
	output = &model.SysDictDataEditModel{
		DictCode: param.DictCode,
	}
	_, err = m.Update()
	if err != nil {
		return nil, err
	}
	return
}

func (s *sSysDictData) Delete(ctx context.Context, param *model.SysDictDataDeleteParam) (output *model.SysDictDataDeleteModel, err error) {
	if len(param.DictCodes) == 0 {
		return nil, gerror.New("字典编码不能为空")
	}
	m := s.Model(ctx)
	m = m.WhereIn(dao.SysDictData.Columns().DictCode, param.DictCodes)
	_, err = m.Delete()
	if err != nil {
		return nil, err
	}
	output = &model.SysDictDataDeleteModel{
		DictCodes: param.DictCodes,
	}
	return
}
