// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"slices"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/do"
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/service"
	"xiuadmin/utility"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysRole struct {
}

func NewSysRole() *sSysRole {
	return &sSysRole{}
}

func init() {
	service.RegisterSysRole(NewSysRole())
}

func (l *sSysRole) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.SysRole.Ctx(ctx), option...)
}

// 获取租户下角色列表
func (s *sSysRole) List(ctx context.Context, param *model.SysRoleListParam) (res []*model.SysRoleListModel, total int, err error) {
	db := s.Model(ctx)
	if param.RoleName != "" {
		db = db.WhereLike(dao.SysRole.Columns().RoleName, "%"+param.RoleName+"%")
	}
	if param.Status != "" {
		db = db.Where(dao.SysRole.Columns().Status, param.Status)
	}
	if len(param.CreatedAt) == 2 {
		start, end := gtime.NewFromStr(param.CreatedAt[0]), gtime.NewFromStr(param.CreatedAt[1])
		db = db.WhereBetween(dao.SysRole.Columns().CreatedAt, start, end.EndOfDay())
	}
	if len(param.RoleIds) > 0 {
		db = db.WhereIn(dao.SysRole.Columns().RoleId, param.RoleIds)
	}
	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}
	err = db.Page(param.Page, param.PageSize).Scan(&res)
	return
}

// 获取角色详情
func (s *sSysRole) View(ctx context.Context, param *model.SysRoleViewParam) (res *model.SysRoleViewModel, err error) {
	err = s.Model(ctx).Where(dao.SysRole.Columns().RoleId, param.RoleId).Scan(&res)
	if err != nil {
		return nil, err
	}
	// 获取角色菜单
	roleMenus, err := s.GetRoleMenu(ctx, param.RoleId)
	if err != nil {
		return nil, err
	}
	menuIdList := make([]int64, 0)
	for _, menu := range roleMenus {
		menuIdList = append(menuIdList, menu.MenuId)
	}
	res.MenuIds = menuIdList
	// 获取角色部门
	roleDepts, err := s.GetRoleDept(ctx, param.RoleId)
	if err != nil {
		return nil, err
	}
	deptIdList := make([]int64, 0)
	for _, dept := range roleDepts {
		deptIdList = append(deptIdList, dept.DeptId)
	}
	res.DeptIds = deptIdList
	return
}

// 获取角色菜单
func (s *sSysRole) GetRoleMenu(ctx context.Context, id int64) (res []*entity.SysMenu, err error) {
	err = dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.Columns().RoleId, id).Scan(&res)
	return
}

// 获取角色部门
func (s *sSysRole) GetRoleDept(ctx context.Context, id int64) (res []*entity.SysDept, err error) {
	err = dao.SysRoleDept.Ctx(ctx).Where(dao.SysRoleDept.Columns().RoleId, id).Scan(&res)
	return
}

// 获取角色列表对应菜单
func (s *sSysRole) GetRoleListMenu(ctx context.Context, ids []int64) (res []*entity.SysRoleMenu, err error) {
	res = make([]*entity.SysRoleMenu, 0)
	if slices.Contains(ids, consts.SuperAdminRoleId) {
		menus, _, err := service.SysMenu().List(ctx, &model.SysMenuListParam{
			Status: consts.SysMenuStatusNormal,
		})
		if err != nil {
			return nil, err
		}
		for _, menu := range menus {
			res = append(res, &entity.SysRoleMenu{
				RoleId: consts.SuperAdminRoleId,
				MenuId: menu.MenuId,
			})
		}
		return res, nil
	}
	err = dao.SysRoleMenu.Ctx(ctx).WhereIn(dao.SysRoleMenu.Columns().RoleId, ids).Scan(&res)
	return
}

// 新增角色
func (s *sSysRole) Add(ctx context.Context, param *model.SysRoleAddParam) (role *model.SysRoleAddModel, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return nil, gerror.New("登录已过期")
	}
	data := do.SysRole{}
	gconv.Struct(param, &data)
	data.DataScope = "1"
	data.CreatedBy = user.ID
	data.CreatedAt = gtime.Now()
	data.CreatedDept = user.DeptId
	data.TenantId = user.TenantId
	data.UpdatedBy = user.ID
	data.UpdatedAt = gtime.Now()

	roleId, err := s.Model(ctx).Data(data).OmitNil().InsertAndGetId()
	if err != nil {
		return nil, err
	}

	role = &model.SysRoleAddModel{
		RoleId: roleId,
	}
	// 新增角色菜单
	err = s.RoleMenu(ctx, roleId, param.MenuIds)
	if err != nil {
		return role, err
	}
	return
}

// 编辑角色
func (s *sSysRole) Edit(ctx context.Context, param *model.SysRoleEditParam) (role *model.SysRoleEditModel, err error) {
	role = &model.SysRoleEditModel{
		RoleId: param.RoleId,
	}
	user := contexts.GetUser(ctx)
	if user == nil {
		return role, gerror.New("登录已过期")
	}
	data := do.SysRole{}
	gconv.Struct(param, &data)
	data.UpdatedBy = user.ID
	data.UpdatedAt = gtime.Now()

	_, err = s.Model(ctx).Data(data).Where(dao.SysRole.Columns().RoleId, param.RoleId).OmitNil().Update()
	if err != nil {
		return role, err
	}
	// 更新角色菜单
	err = s.RoleMenu(ctx, param.RoleId, param.MenuIds)
	if err != nil {
		return role, err
	}

	return
}

// 删除角色
func (s *sSysRole) Delete(ctx context.Context, param *model.SysRoleDeleteParam) (role *model.SysRoleDeleteModel, err error) {
	role = &model.SysRoleDeleteModel{
		RoleId:  param.RoleId,
		RoleIds: param.RoleIds,
	}
	roleIds := make([]int64, 0)
	if param.RoleId != 0 {
		roleIds = append(roleIds, param.RoleId)
	} else {
		roleIds = param.RoleIds
	}
	if len(roleIds) == 0 {
		return role, gerror.New("请选择要删除的角色")
	}
	user := contexts.GetUser(ctx)
	if user == nil {
		return role, gerror.New("登录已过期")
	}
	data := do.SysRole{}
	data.DeletedBy = user.ID
	data.DeletedAt = gtime.Now()
	_, err = s.Model(ctx).WhereIn(dao.SysRole.Columns().RoleId, roleIds).Data(data).OmitNil().Update()
	return
}

func (s *sSysRole) RoleMenu(ctx context.Context, roleId int64, menuIds []int64) (err error) {
	// 查询角色菜单
	roleMenu := make([]*entity.SysRoleMenu, 0)
	err = dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.Columns().RoleId, roleId).Scan(&roleMenu)
	if err != nil {
		return err
	}
	currMenuIds := make([]int64, 0)
	for _, menu := range roleMenu {
		currMenuIds = append(currMenuIds, menu.MenuId)
	}

	delMenuIds := utility.ArrayRightDiff(currMenuIds, menuIds)
	addMenuIds := utility.ArrayRightDiff(menuIds, currMenuIds)

	// 删除角色菜单
	if len(delMenuIds) > 0 {
		_, err = dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.Columns().RoleId, roleId).WhereIn(dao.SysRoleMenu.Columns().MenuId, delMenuIds).Delete()
	}

	// 新增角色菜单
	if len(addMenuIds) > 0 {
		addData := make([]map[string]interface{}, 0)
		for _, menuId := range addMenuIds {
			addData = append(addData, map[string]interface{}{
				"role_id": roleId,
				"menu_id": menuId,
			})
		}
		_, err = dao.SysRoleMenu.Ctx(ctx).Data(addData).Insert()
	}
	return
}

// 自定义角色部门数据权限
func (s *sSysRole) RoleDept(ctx context.Context, roleId int64, deptIds []int64) (err error) {
	// 查询角色部门
	roleDept := make([]*entity.SysRoleDept, 0)
	err = dao.SysRoleDept.Ctx(ctx).Where(dao.SysRoleDept.Columns().RoleId, roleId).Scan(&roleDept)
	if err != nil {
		return err
	}
	currDeptIds := make([]int64, 0)
	for _, dept := range roleDept {
		currDeptIds = append(currDeptIds, dept.DeptId)
	}

	delDeptIds := utility.ArrayRightDiff(currDeptIds, deptIds)
	addDeptIds := utility.ArrayRightDiff(deptIds, currDeptIds)

	// 删除角色部门
	if len(delDeptIds) > 0 {
		_, err = dao.SysRoleDept.Ctx(ctx).Where(dao.SysRoleDept.Columns().RoleId, roleId).WhereIn(dao.SysRoleDept.Columns().DeptId, delDeptIds).Delete()
	}

	// 新增角色部门
	if len(addDeptIds) > 0 {
		addData := make([]map[string]interface{}, 0)
		for _, deptId := range addDeptIds {
			addData = append(addData, map[string]interface{}{
				"role_id": roleId,
				"dept_id": deptId,
			})
		}
		_, err = dao.SysRoleDept.Ctx(ctx).Data(addData).Insert()
	}
	return
}

// 编辑角色数据权限
func (s *sSysRole) EditRoleDataScope(ctx context.Context, model *model.SysRoleDataScopeEditParam) (err error) {

	data := do.SysRole{}
	gconv.Struct(model, &data)
	data.UpdatedBy = contexts.GetUserId(ctx)
	data.UpdatedAt = gtime.Now()

	_, err = s.Model(ctx).Data(data).Where(dao.SysRole.Columns().RoleId, model.RoleId).OmitNil().Update()
	if err != nil {
		return err
	}
	// 更新角色部门
	err = s.RoleDept(ctx, model.RoleId, model.DeptIds)
	if err != nil {
		return err
	}
	return
}
