package gen_codes

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

type sSysGenTable struct {
}

func SysGenTableNew() *sSysGenTable {
	return &sSysGenTable{}
}

func init() {
	service.RegisterSysGenTable(SysGenTableNew())
}
func (l *sSysGenTable) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) > 0 {
		option = append(option, &handler.Option{
			FilterTenant: false,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.SysGenTable.Ctx(ctx), option...)
}

// 列表
func (s *sSysGenTable) List(ctx context.Context, param *model.SysGenTableListParam) (output []*model.SysGenTableListModel, total int, err error) {
	m := s.Model(ctx)
	if param.GenType != 0 {
		m = m.Where(dao.SysGenTable.Columns().GenType, param.GenType)
	}
	if param.VarName != "" {
		m = m.Where(dao.SysGenTable.Columns().VarName, param.VarName)
	}
	if param.Status != "" {
		m = m.Where(dao.SysGenTable.Columns().Status, param.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	output = make([]*model.SysGenTableListModel, 0)
	err = m.Page(param.Page, param.PageSize).OrderDesc(dao.SysGenTable.Columns().TableId).Scan(&output)
	if err != nil {
		return
	}
	return
}

// 详情
func (s *sSysGenTable) View(ctx context.Context, param *model.SysGenTableViewParam) (output *model.SysGenTableViewModel, err error) {
	if param.TableId == 0 {
		err = gerror.New("参数错误")
		return
	}
	m := s.Model(ctx)
	err = m.Where(dao.SysGenTable.Columns().TableId, param.TableId).Scan(&output)
	if err != nil {
		return
	}
	return
}

// 新增
func (s *sSysGenTable) Add(ctx context.Context, param *model.SysGenTableAddParam) (output *model.SysGenTableAddModel, err error) {
	data := do.SysGenTable{}
	gconv.Struct(param, &data)
	data.Status = "1"
	data.CreatedDept = contexts.GetDeptId(ctx)
	data.CreatedBy = contexts.GetUserId(ctx)
	data.CreatedAt = gtime.Now()
	m := s.Model(ctx)
	id, err := m.Data(data).InsertAndGetId()
	if err != nil {
		return
	}
	output = &model.SysGenTableAddModel{
		TableId: id,
	}
	return
}

func (s *sSysGenTable) Delete(ctx context.Context, param *model.SysGenTableDeleteParam) (output *model.SysGenTableDeleteModel, err error) {
	if len(param.TableIds) == 0 {
		err = gerror.New("参数错误")
		return
	}
	m := s.Model(ctx)
	_, err = m.WhereIn(dao.SysGenTable.Columns().TableId, param.TableIds).Delete()
	if err != nil {
		return
	}
	output = &model.SysGenTableDeleteModel{
		TableIds: param.TableIds,
	}
	return
}
