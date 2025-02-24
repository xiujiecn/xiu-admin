package system

import (
	"context"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/request"
	"server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sSysDict struct{}

func NewSysDict() *sSysDict {
	return &sSysDict{}
}

func init() {
	service.RegisterSysDict(NewSysDict())
}

func (s *sSysDict) GetDictTypeList(ctx context.Context, req *model.SysDictTypeListQuery, pageInfo *request.PageInfo) (items []model.SysDictType, total int, err error) {
	// 获取当前用户租户编码
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, 0, err
	}
	tenantId := claims.TenantId
	m := dao.SysDictType.Ctx(ctx).Where(dao.SysDictType.Columns().TenantId, tenantId)
	if req.DictName != "" {
		m = m.WhereLike(dao.SysDictType.Columns().DictName, "%"+req.DictName+"%")
	}
	if req.DictType != "" {
		m = m.WhereLike(dao.SysDictType.Columns().DictType, "%"+req.DictType+"%")
	}
	total, err = m.Count()
	if err != nil {
		return nil, 0, err
	}
	err = m.Order().Page(pageInfo.Page, pageInfo.PageSize).Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return
}

func (s *sSysDict) GetDictDataList(ctx context.Context, query *model.SysDictDataListQuery, pageInfo *request.PageInfo) (items []model.SysDictData, total int, err error) {
	// 获取当前用户租户编码
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, 0, err
	}
	tenantId := claims.TenantId
	// 获取字典类型
	if query.DictType == "" && query.DictId != 0 {
		var data *model.SysDictType
		err = dao.SysDictType.Ctx(ctx).Where(dao.SysDictType.Columns().TenantId, tenantId).
			Where(dao.SysDictType.Columns().DictId, query.DictId).Scan(&data)
		if err != nil {
			return nil, 0, err
		}
		if data == nil {
			return nil, 0, gerror.New("字典类型不存在")
		}
		query.DictType = data.DictType
	}
	if query.DictType == "" {
		return nil, 0, gerror.New("字典类型不能为空")
	}
	// 获取字典数据
	items = make([]model.SysDictData, 0)
	m := dao.SysDictData.Ctx(ctx).Where(dao.SysDictData.Columns().TenantId, tenantId)
	m = m.Where(dao.SysDictData.Columns().DictType, query.DictType)
	total, err = m.Count()
	if err != nil {
		return nil, 0, err
	}
	err = m.Order(dao.SysDictData.Columns().DictSort, "ASC").Page(pageInfo.Page, pageInfo.PageSize).Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return
}
