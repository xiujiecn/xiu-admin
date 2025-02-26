package system

import (
	"context"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/entity"
	"server/internal/model/request"
	"server/internal/service"
)

type sSysRole struct {
}

func NewSysRole() *sSysRole {
	return &sSysRole{}
}

func init() {
	service.RegisterSysRole(NewSysRole())
}

// 获取租户下角色列表
func (s *sSysRole) GetRoleList(ctx context.Context, model *model.SysRoleListParam, pageInfo *request.PageInfo) (res []*model.SysRoleListModel, total int, err error) {
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, 0, err
	}
	db := dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().TenantId, claims.TenantId)
	if model.RoleName != "" {
		db = db.WhereLike(dao.SysRole.Columns().RoleName, "%"+model.RoleName+"%")
	}
	if model.Status != "" {
		db = db.Where(dao.SysRole.Columns().Status, model.Status)
	}
	if model.CreatedAt != nil {
		db = db.Where(dao.SysRole.Columns().CreatedAt, model.CreatedAt)
	}
	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}
	err = db.Page(pageInfo.Page, pageInfo.PageSize).Scan(&res)
	return
}

// 获取角色详情
func (s *sSysRole) GetRoleView(ctx context.Context, id int64) (res *model.SysRoleViewModel, err error) {
	err = dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().RoleId, id).Scan(&res)
	return
}

// 获取角色菜单
func (s *sSysRole) GetRoleMenu(ctx context.Context, id int64) (res []*entity.SysMenu, err error) {
	err = dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.Columns().RoleId, id).Scan(&res)
	return
}

// 获取角色列表对应菜单
func (s *sSysRole) GetRoleListMenu(ctx context.Context, ids []int64) (res []*entity.SysRoleMenu, err error) {
	err = dao.SysRoleMenu.Ctx(ctx).WhereIn(dao.SysRoleMenu.Columns().RoleId, ids).Scan(&res)
	return
}
