package system

import (
	"context"
	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
)

type sSysNotice struct {
}

func NewSysNotice() *sSysNotice {
	return &sSysNotice{}
}

func init() {
	service.RegisterSysNotice(NewSysNotice())
}

func (s *sSysNotice) GetNoticeList(ctx context.Context, query *model.SysNoticeListParam, page *request.PageInfo) (items []*model.SysNotice, total int, err error) {
	// 获取当前用户租户编码
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, 0, err
	}
	tenantId := claims.TenantId

	db := dao.SysNotice.Ctx(ctx).Where(dao.SysNotice.Columns().TenantId, tenantId)

	if query.NoticeTitle != "" {
		db = db.WhereLike(dao.SysNotice.Columns().NoticeTitle, "%"+query.NoticeTitle+"%")
	}

	if query.NoticeType != "" {
		db = db.Where(dao.SysNotice.Columns().NoticeType, query.NoticeType)
	}

	if query.CreatedBy != "" {
		db = db.Where(dao.SysNotice.Columns().CreatedBy, query.CreatedBy)
	}

	if query.CreatedAt != nil && len(query.CreatedAt) == 2 {
		createdAt1 := gtime.NewFromStr(query.CreatedAt[0])
		createdAt2 := gtime.NewFromStr(query.CreatedAt[1])

		db = db.WhereBetween(dao.SysNotice.Columns().CreatedAt, createdAt1, createdAt2.EndOfDay())
	}

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
