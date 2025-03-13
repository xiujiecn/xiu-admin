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

type sSysNotice struct {
}

func NewSysNotice() *sSysNotice {
	return &sSysNotice{}
}

func init() {
	service.RegisterSysNotice(NewSysNotice())
}

func (l *sSysNotice) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) > 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.SysNotice.Ctx(ctx), option...)
}

func (s *sSysNotice) List(ctx context.Context, param *model.SysNoticeListParam) (items []*model.SysNoticeListModel, total int, err error) {
	db := s.Model(ctx)

	if param.NoticeTitle != "" {
		db = db.WhereLike(dao.SysNotice.Columns().NoticeTitle, "%"+param.NoticeTitle+"%")
	}

	if param.NoticeType != "" {
		db = db.Where(dao.SysNotice.Columns().NoticeType, param.NoticeType)
	}

	if param.CreatedBy != "" {
		db = db.Where(dao.SysNotice.Columns().CreatedBy, param.CreatedBy)
	}

	if len(param.CreatedAt) == 2 {
		createdAt1 := gtime.NewFromStr(param.CreatedAt[0])
		createdAt2 := gtime.NewFromStr(param.CreatedAt[1])

		db = db.WhereBetween(dao.SysNotice.Columns().CreatedAt, createdAt1, createdAt2.EndOfDay())
	}

	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}

	err = db.Page(param.Page, param.PageSize).WithAll().Scan(&items)
	if err != nil {
		return nil, 0, err
	}

	return
}

func (s *sSysNotice) Add(ctx context.Context, param *model.SysNoticeAddParam) (err error) {
	if param == nil || param.NoticeTitle == "" {
		return errors.New("公告标题不能为空")
	}
	if param.NoticeType == "" {
		return errors.New("公告类型不能为空")
	}
	if param.NoticeContent == "" {
		return errors.New("公告内容不能为空")
	}

	data := do.SysNotice{}
	gconv.Struct(param, &data)
	data.TenantId = contexts.GetTenantId(ctx)
	data.CreatedDept = contexts.GetDeptId(ctx)
	data.CreatedBy = contexts.GetUserId(ctx)
	data.CreatedAt = gtime.Now()
	data.UpdatedBy = contexts.GetUserId(ctx)
	data.UpdatedAt = gtime.Now()

	_, err = s.Model(ctx).Data(data).OmitNil().Insert()
	if err != nil {
		return err
	}

	return
}

func (s *sSysNotice) Edit(ctx context.Context, param *model.SysNoticeEditParam) (err error) {
	if param == nil || param.NoticeId == 0 {
		return errors.New("公告ID不能为空")
	}
	if param.NoticeTitle == nil || *param.NoticeTitle == "" {
		return errors.New("公告标题不能为空")
	}
	if param.NoticeType == nil || *param.NoticeType == "" {
		return errors.New("公告类型不能为空")
	}
	if param.NoticeContent == nil || len(*param.NoticeContent) == 0 {
		return errors.New("公告内容不能为空")
	}

	data := do.SysNotice{}
	gconv.Struct(param, &data)
	data.NoticeContent = []byte(*param.NoticeContent)
	data.UpdatedBy = contexts.GetUserId(ctx)
	data.UpdatedAt = gtime.Now()

	m := s.Model(ctx)

	m = m.Where(dao.SysNotice.Columns().NoticeId, param.NoticeId)
	g.Log().Debugf(ctx, "param: %+v", string(*param.NoticeContent))
	g.Log().Debugf(ctx, "data: %+v", string(data.NoticeContent))
	_, err = m.Data(data).OmitNil().Update()
	if err != nil {
		return err
	}

	return
}

func (s *sSysNotice) Delete(ctx context.Context, param *model.SysNoticeDeleteParam) (err error) {
	_, err = s.Model(ctx).WhereIn(dao.SysNotice.Columns().NoticeId, param.NoticeIds).Delete()
	if err != nil {
		return err
	}

	return
}

func (s *sSysNotice) View(ctx context.Context, param *model.SysNoticeViewParam) (data *model.SysNoticeViewModel, err error) {
	if param == nil || param.NoticeId == 0 {
		return nil, errors.New("公告ID不能为空")
	}
	data = &model.SysNoticeViewModel{}
	err = s.Model(ctx).Where(dao.SysNotice.Columns().NoticeId, param.NoticeId).Scan(&data)
	g.Log().Debugf(ctx, "data: %+v", string(data.NoticeContent))
	if err != nil {
		return nil, err
	}
	return
}
