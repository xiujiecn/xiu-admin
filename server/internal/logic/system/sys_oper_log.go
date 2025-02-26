package system

import (
	"context"
	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/service"
)

type sSysOperLog struct {
}

func NewSysOperLog() *sSysOperLog {
	return &sSysOperLog{}
}

func init() {
	service.RegisterSysOperLog(NewSysOperLog())
}

func (s *sSysOperLog) GetOperLogList(ctx context.Context, query *model.SysOperLogListParam, page *request.PageInfo) (items []*model.SysOperLogListModel, total int, err error) {
	// 获取当前用户租户编码
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, 0, err
	}
	tenantId := claims.TenantId

	db := dao.SysOperLog.Ctx(ctx).Where(dao.SysOperLog.Columns().TenantId, tenantId)

	if query.Title != "" {
		db = db.WhereLike(dao.SysOperLog.Columns().Title, "%"+query.Title+"%")
	}

	if query.BusinessType != "" {
		db = db.Where(dao.SysOperLog.Columns().BusinessType, query.BusinessType)
	}

	if query.Method != "" {
		db = db.WhereLike(dao.SysOperLog.Columns().Method, "%"+query.Method+"%")
	}

	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}

	err = db.Page(page.Page, page.PageSize).Scan(&items)
	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}
