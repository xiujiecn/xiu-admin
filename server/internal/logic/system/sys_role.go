package system

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
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
func (s *sSysRole) GetRoleList(ctx context.Context, tenantId int64) (res []*entity.SysRole, err error) {
	err = dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().TenantId, tenantId).Order("sort").Scan(&res)
	return
}

// 获取角色详情
func (s *sSysRole) GetRoleDetail(ctx context.Context, id int64) (res *entity.SysRole, err error) {
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
