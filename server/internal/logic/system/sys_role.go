package system

import (
	"context"
	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/do"
	"xiujieadmin/internal/model/entity"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/service"
	"xiujieadmin/utility"

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
	if len(model.RoleIds) > 0 {
		db = db.WhereIn(dao.SysRole.Columns().RoleId, model.RoleIds)
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
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	err = dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().RoleId, id).Where(dao.SysRole.Columns().TenantId, claims.TenantId).Scan(&res)
	if err != nil {
		return nil, err
	}
	// 获取角色菜单
	roleMenus, err := s.GetRoleMenu(ctx, id)
	if err != nil {
		return nil, err
	}
	menuIdList := make([]int64, 0)
	for _, menu := range roleMenus {
		menuIdList = append(menuIdList, menu.MenuId)
	}
	res.MenuIds = menuIdList
	// 获取角色部门
	roleDepts, err := s.GetRoleDept(ctx, id)
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
	err = dao.SysRoleMenu.Ctx(ctx).WhereIn(dao.SysRoleMenu.Columns().RoleId, ids).Scan(&res)
	return
}

// 新增角色
func (s *sSysRole) AddRole(ctx context.Context, model *model.SysRoleAddParam) (err error) {
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return err
	}
	data := do.SysRole{}
	gconv.Struct(model, &data)
	data.DataScope = "1"
	data.CreatedBy = claims.BaseClaims.ID
	data.CreatedAt = gtime.Now()
	data.CreatedDept = claims.BaseClaims.DeptId
	data.TenantId = claims.TenantId
	data.UpdatedBy = claims.BaseClaims.ID
	data.UpdatedAt = gtime.Now()

	result, err := dao.SysRole.Ctx(ctx).Data(data).OmitNil().Insert()
	if err != nil {
		return err
	}
	roleId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	// 新增角色菜单
	err = s.RoleMenu(ctx, roleId, model.MenuIds)
	if err != nil {
		return err
	}
	return
}

// 编辑角色
func (s *sSysRole) EditRole(ctx context.Context, model *model.SysRoleEditParam) (err error) {
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return err
	}
	data := do.SysRole{}
	gconv.Struct(model, &data)
	data.UpdatedBy = claims.BaseClaims.ID
	data.UpdatedAt = gtime.Now()

	_, err = dao.SysRole.Ctx(ctx).Data(data).Where(dao.SysRole.Columns().RoleId, model.RoleId).Where(dao.SysRole.Columns().TenantId, claims.TenantId).OmitNil().Update()
	if err != nil {
		return err
	}
	// 更新角色菜单
	err = s.RoleMenu(ctx, model.RoleId, model.MenuIds)
	if err != nil {
		return err
	}
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
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return err
	}
	data := do.SysRole{}
	gconv.Struct(model, &data)
	data.UpdatedBy = claims.BaseClaims.ID
	data.UpdatedAt = gtime.Now()

	_, err = dao.SysRole.Ctx(ctx).Data(data).Where(dao.SysRole.Columns().RoleId, model.RoleId).Where(dao.SysRole.Columns().TenantId, claims.TenantId).OmitNil().Update()
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
