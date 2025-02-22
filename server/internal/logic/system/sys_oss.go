package system

import (
	"context"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/request"
	"server/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
)

type sSysOss struct {
}

func NewSysOss() *sSysOss {
	return &sSysOss{}
}

func init() {
	service.RegisterSysOss(NewSysOss())
}

func (s *sSysOss) List(ctx context.Context, query *model.SysOssListQuery, pageInfo *request.PageInfo) (items []*model.SysOssListModel, total int, err error) {
	// 获取当前用户租户编码
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, 0, err
	}
	tenantId := claims.TenantId
	db := dao.SysOss.Ctx(ctx).Where(dao.SysOss.Columns().TenantId, tenantId)

	if query.FileName != "" {
		db = db.WhereLike(dao.SysOss.Columns().FileName, "%"+query.FileName+"%")
	}

	if query.OriginalName != "" {
		db = db.WhereLike(dao.SysOss.Columns().OriginalName, "%"+query.OriginalName+"%")
	}

	if query.FileSuffix != "" {
		db = db.Where(dao.SysOss.Columns().FileSuffix, query.FileSuffix)
	}

	if query.Service != "" {
		db = db.Where(dao.SysOss.Columns().Service, query.Service)
	}

	if len(query.CreatedAt) == 2 {
		startTime, err := gtime.StrToTime(query.CreatedAt[0])
		if err != nil {
			return nil, 0, err
		}
		endTime, err := gtime.StrToTime(query.CreatedAt[1])
		if err != nil {
			return nil, 0, err
		}
		db = db.WhereBetween(dao.SysOss.Columns().CreatedAt, startTime, endTime.EndOfDay())
	}

	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}

	err = db.Page(pageInfo.Page, pageInfo.PageSize).Order("id", "desc").Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return
}
