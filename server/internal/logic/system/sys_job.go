package system

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"

	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/library/xgorm/handler"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/service"
)

func init() {
	service.RegisterSysJob(NewSysJob())
}

type sSysJob struct{}

func NewSysJob() *sSysJob {
	return &sSysJob{}
}

func (s *sSysJob) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) > 0 {
		option = append(option, &handler.Option{
			FilterTenant: false,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.SysJob.Ctx(ctx), option...)
}

func (c *sSysJob) List(ctx context.Context, query *model.SysJobListParam, pageInfo *request.PageInfo) (Data []*model.SysJobListModel, total int, err error) {
	m := c.Model(ctx)
	if query.JobName != "" {
		m = m.WhereLike(dao.SysJob.Columns().JobName, "%"+strings.Trim(query.JobName, " ")+"%")
	}
	if query.JobGroup != "" {
		m = m.Where(dao.SysJob.Columns().JobGroup, query.JobGroup)
	}
	if query.Status != "" {
		m = m.Where(dao.SysJob.Columns().Status, query.Status)
	}
	t, err := m.Count()
	if err != nil {
		return nil, 0, err
	}
	Data = make([]*model.SysJobListModel, 0)
	if err := m.Page(pageInfo.Page, pageInfo.PageSize).Order(dao.SysJob.Columns().JobId, "DESC").Scan(&Data); err != nil {
		return nil, 0, err
	}
	return Data, t, nil
}

func (c *sSysJob) View(ctx context.Context, jobId int64) (Data *model.SysJobViewModel, err error) {
	m := c.Model(ctx)
	Data = &model.SysJobViewModel{}
	if err := m.Where(dao.SysJob.Columns().JobId, jobId).Scan(&Data); err != nil {
		return nil, err
	}

	return Data, nil
}

func (c *sSysJob) Add(ctx context.Context, jobAdd *model.SysJobAddModel) (LastInsertId int64, err error) {
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return 0, err
	}
	jobAdd.CreatedDept = claims.BaseClaims.DeptId
	jobAdd.CreatedAt = gtime.Now()
	jobAdd.CreatedBy = claims.BaseClaims.ID
	jobAdd.UpdatedAt = gtime.Now()
	jobAdd.UpdatedBy = claims.BaseClaims.ID

	result, err := c.Model(ctx).Data(jobAdd).Insert()
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (c *sSysJob) Update(ctx context.Context, jobUpdate *model.SysJobUpdateModel) (RowsAffected int64, err error) {
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return 0, err
	}

	jobUpdate.UpdatedAt = gtime.Now()
	jobUpdate.UpdatedBy = claims.BaseClaims.ID
	result, err := c.Model(ctx).Where(dao.SysJob.Columns().JobId, jobUpdate.JobId).Update(jobUpdate)

	if err != nil {
		return 0, err
	}
	row, err := result.RowsAffected()
	return row, err
}
func (c *sSysJob) Delete(ctx context.Context, jobDelete *model.SysJobDeleteModel) (RowsAffected int64, err error) {
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return 0, err
	}

	jobDelete.DeletedAt = gtime.Now()
	jobDelete.DeletedBy = claims.BaseClaims.ID
	result, err := c.Model(ctx).WhereIn(dao.SysJob.Columns().JobId, jobDelete.JobIds).Update(map[string]interface{}{
		dao.SysJob.Columns().DeletedAt: jobDelete.DeletedAt,
		dao.SysJob.Columns().DeletedBy: jobDelete.DeletedBy})

	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
