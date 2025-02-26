package system

import (
	"context"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/entity"
	"server/internal/model/request"
	"server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
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
	if len(model.CreatedAt) == 2 {
		start, end := gtime.NewFromStr(model.CreatedAt[0]), gtime.NewFromStr(model.CreatedAt[1])
		db = db.WhereBetween(dao.SysRole.Columns().CreatedAt, start, end.EndOfDay())
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

// 新增角色
func (s *sSysRole) AddRole(ctx context.Context, model *model.SysRoleAddParam) (err error) {
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return err
	}
	model.CreatedBy = claims.BaseClaims.ID
	model.CreatedAt = gtime.Now()
	model.CreatedDept = claims.BaseClaims.DeptId
	model.TenantId = claims.TenantId
	model.UpdatedBy = claims.BaseClaims.ID
	model.UpdatedAt = gtime.Now()

	_, err = dao.SysRole.Ctx(ctx).Data(model).Insert()
	return
}

// 编辑角色
func (s *sSysRole) EditRole(ctx context.Context, model *model.SysRoleEditParam) (err error) {
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return err
	}
	model.UpdatedBy = claims.BaseClaims.ID
	model.UpdatedAt = gtime.Now()

	_, err = dao.SysRole.Ctx(ctx).Data(model).Where(dao.SysRole.Columns().RoleId, model.RoleId).Where(dao.SysRole.Columns().TenantId, claims.TenantId).Update()
	return
}

// 删除角色
func (s *sSysRole) DeleteRole(ctx context.Context, model *model.SysRoleDeleteParam) (err error) {
	roleIds := make([]int64, 0)
	if model.RoleId != 0 {
		roleIds = append(roleIds, model.RoleId)
	} else {
		roleIds = model.RoleIds
	}
	if len(roleIds) == 0 {
		return gerror.New("请选择要删除的角色")
	}
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return err
	}
	_, err = dao.SysRole.Ctx(ctx).WhereIn(dao.SysRole.Columns().RoleId, roleIds).Where(dao.SysRole.Columns().TenantId, claims.TenantId).Delete()
	return
}
