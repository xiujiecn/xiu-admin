package system

import (
	"context"
	"fmt"
	"slices"
	"sort"
	"strings"
	v1 "xiujieadmin/api/system/v1"
	"xiujieadmin/internal/consts"
	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/library/contexts"
	"xiujieadmin/internal/library/xgorm/handler"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/entity"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sSysMenu struct {
}

func NewSysMenu() *sSysMenu {
	return &sSysMenu{}
}

func init() {
	service.RegisterSysMenu(NewSysMenu())
}

func (s *sSysMenu) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: false,
			FilterAuth:   false,
		})
	}
	return handler.Model(dao.SysMenu.Ctx(ctx), option...)
}

func (l *sSysMenu) List(ctx context.Context, param *model.SysMenuListParam) (data []*model.SysMenuListModel, total int, err error) {
	m := l.Model(ctx)
	if param.MenuIds != nil {
		m = m.WhereIn(dao.SysMenu.Columns().MenuId, param.MenuIds)
	}
	if param.MenuName != "" {
		m = m.WhereLike(dao.SysMenu.Columns().MenuName, "%"+param.MenuName+"%")
	}
	if param.Status != "" {
		m = m.Where(dao.SysMenu.Columns().Status, param.Status)
	}
	total, err = m.Count()
	if err != nil {
		return nil, 0, err
	}
	err = m.Order(dao.SysMenu.Columns().OrderNum).Page(0, 5000).Scan(&data)
	return
}

// 获取租户菜单列表， 系统租户返回所有菜单，其他租户返回当前租户菜单
func (l *sSysMenu) GetTenantMenu(ctx context.Context, query *model.SysMenuListParam) (data []*model.SysMenuListModel, total int, err error) {
	// 获取当前用户信息
	tenantId := contexts.GetTenantId(ctx)
	m := dao.SysMenu.Ctx(ctx)
	if tenantId != consts.DefaultSystemTenantCode {
		// 获取租户信息
		tenantInfo, err := service.SysTenant().View(ctx, &model.SysTenantViewParam{
			TenantId: tenantId,
		})
		if err != nil {
			return nil, 0, err
		}
		// 获取租户套餐
		tenantPackage, err := service.SysTenantPackage().View(ctx, &model.SysTenantPackageViewParam{
			PackageId: tenantInfo.PackageId,
		})
		if err != nil {
			return nil, 0, err
		}
		if tenantPackage == nil {
			g.Log().Errorf(ctx, "sSysMenu.GetTenantMenu not found tenantPackage. TenantId:%s PackageId:%d",
				tenantId, tenantInfo.PackageId)
			return nil, 0, gerror.NewCode(gcode.CodeInternalError, "系统内部错误")
		}
		if tenantPackage.MenuIds != "" {
			menuIds := strings.Split(tenantPackage.MenuIds, ",")
			m = m.WhereIn(dao.SysMenu.Columns().MenuId, menuIds)
		} else {
			return nil, 0, nil
		}
	}
	if query.MenuName != "" {
		m = m.WhereLike(dao.SysMenu.Columns().MenuName, "%"+query.MenuName+"%")
	}
	if query.Status != "" {
		m = m.Where(dao.SysMenu.Columns().Status, query.Status)
	}
	if query.MenuIds != nil {
		m = m.WhereIn(dao.SysMenu.Columns().MenuId, query.MenuIds)
	}

	total, err = m.Count()
	if err != nil {
		return nil, 0, err
	}
	err = m.Order(dao.SysMenu.Columns().OrderNum).Page(0, 5000).Scan(&data)
	return
}

// 构建树结构
func (l *sSysMenu) MenuTree(ctx context.Context, parentMenu *model.SysMenuTree, menuList []*entity.SysMenu) (data []*model.SysMenuTree, err error) {
	data = make([]*model.SysMenuTree, 0)
	for _, menu := range menuList {
		pId := int64(0)
		if parentMenu != nil {
			pId = parentMenu.MenuId
		}
		if menu.ParentId == pId {
			item := &model.SysMenuTree{
				SysMenu:  *menu,
				Children: nil,
			}
			subMenu, err := l.MenuTree(ctx, item, menuList)
			if err != nil {
				return nil, err
			}
			item.Children = subMenu
			data = append(data, item)
		}
	}
	return
}

// 获取用户动态路由列表
func (l *sSysMenu) GetUserMenu(ctx context.Context) (data []*entity.SysMenu, err error) {
	roles := contexts.GetRoles(ctx)
	roleIds := make([]int64, 0)
	for _, role := range roles {
		roleIds = append(roleIds, role.RoleId)
	}
	// 判断是否存在超级管理员角色
	if slices.Contains(roleIds, consts.SuperAdminRoleId) {
		// 获取所有菜单
		err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Status, consts.SysMenuStatusNormal).Order(dao.SysMenu.Columns().OrderNum).Scan(&data)
		return
	}
	// 获取角色菜单列表
	menuList, err := service.SysRole().GetRoleListMenu(ctx, roleIds)
	if err != nil {
		return nil, err
	}
	// g.Log().Infof(ctx, "sSysMenu.GetUserMenu menuList: %v", menuList)
	menuIds := make([]int64, 0)
	for _, menu := range menuList {
		menuIds = append(menuIds, menu.MenuId)
	}
	// 获取用户角色菜单
	err = dao.SysMenu.Ctx(ctx).WhereIn(dao.SysMenu.Columns().MenuId, menuIds).Where(dao.SysMenu.Columns().Status, consts.SysMenuStatusNormal).Order(dao.SysMenu.Columns().OrderNum).Scan(&data)
	return
}

// 构建用户动态路由树
func (l *sSysMenu) BuildUserMenuTree(ctx context.Context, parentMenu *v1.RouteMenu, menuList []*entity.SysMenu, allPath string) (data v1.MenuAllRes, err error) {
	data = make(v1.MenuAllRes, 0)
	for _, menu := range menuList {
		if menu.Status == consts.SysMenuStatusDisable {
			continue
		}
		if menu.MenuType == consts.SysMenuTypeBtn {
			continue
		}

		pId := int64(0)
		if parentMenu != nil {
			pId = parentMenu.Id
		}
		if menu.ParentId == pId {
			openInNewWindow := false
			var link *string
			if menu.IsFrame == consts.SysMenuIsFrameYes {
				openInNewWindow = true
				link = &menu.Path
			}

			query := make(map[string]string)
			if menu.QueryParam != "" {
				query2 := strings.Split(menu.QueryParam, "&")
				for _, q := range query2 {
					kv := strings.Split(q, "=")
					if len(kv) == 2 {
						query[kv[0]] = kv[1]
					}
				}
			}
			hideInMenu := false
			if menu.Visible == consts.SysMenuVisibleHide {
				hideInMenu = true
			}
			item := &v1.RouteMenu{
				Id:        menu.MenuId,
				Path:      menu.Path,
				Redirect:  nil,
				Name:      fmt.Sprintf("menu_%d", menu.MenuId),
				Component: menu.Component,
				Meta: &v1.RouteMeta{
					Title:           menu.MenuName,
					Icon:            &menu.Icon,
					Order:           &menu.OrderNum,
					OpenInNewWindow: &openInNewWindow,
					Query:           query,
					HideInMenu:      &hideInMenu,
					Authority:       strings.Split(menu.Perms, ","),
					Link:            link,
				},
			}
			if hideInMenu {
				activePath := allPath + "/"
				if strings.Contains(menu.Path, "-") {
					activePath = activePath + strings.Split(menu.Path, "-")[0]
				}
				item.Meta.ActivePath = &activePath
			}
			subMenu, err := l.BuildUserMenuTree(ctx, item, menuList, allPath+"/"+menu.Path)
			if err != nil {
				return nil, err
			}
			item.Children = subMenu
			if len(subMenu) == 0 && menu.MenuType == consts.SysMenuTypeDir {
				item.Children = []*v1.RouteMenu{}
			}
			data = append(data, item)
		}
	}
	// 排序
	sort.Slice(data, func(i, j int) bool {
		return *data[i].Meta.Order < *data[j].Meta.Order
	})
	return
}

// 获取用户动态路由树
func (l *sSysMenu) GetUserMenuTree(ctx context.Context) (data v1.MenuAllRes, err error) {
	// 获取用户动态路由列表
	menuList, err := l.GetUserMenu(ctx)
	if err != nil {
		return nil, err
	}
	// 构建树结构
	data, err = l.BuildUserMenuTree(ctx, nil, menuList, "")
	if err != nil {
		return nil, err
	}
	return
}

func (l *sSysMenu) GetSysMenuView(ctx context.Context, menuId int64) (data *model.SysMenuViewModel, err error) {
	menu := &entity.SysMenu{}
	err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().MenuId, menuId).Scan(menu)
	if err != nil {
		return nil, err
	}
	data = &model.SysMenuViewModel{
		SysMenu: *menu,
	}
	return
}

func (l *sSysMenu) UpdateSysMenu(ctx context.Context, menu *model.SysMenuUpdateModel) (data *model.SysMenuViewModel, err error) {
	menu.UpdatedAt = gtime.Now()
	menu.UpdatedBy = contexts.GetUserId(ctx)
	_, err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().MenuId, menu.MenuId).Data(menu).OmitNil().Update(menu)
	if err != nil {
		return nil, err
	}
	data, err = l.GetSysMenuView(ctx, menu.MenuId)
	if err != nil {
		return nil, err
	}
	return
}

func (l *sSysMenu) AddSysMenu(ctx context.Context, menu *model.SysMenuAddModel) (data *model.SysMenuViewModel, err error) {
	menu.CreatedDept = contexts.GetDeptId(ctx)
	menu.CreatedAt = gtime.Now()
	menu.CreatedBy = contexts.GetUserId(ctx)
	menu.UpdatedAt = gtime.Now()
	menu.UpdatedBy = contexts.GetUserId(ctx)
	re, err := dao.SysMenu.Ctx(ctx).Data(menu).Insert()
	if err != nil {
		return nil, err
	}
	menuId, err := re.LastInsertId()
	if err != nil {
		return nil, err
	}
	data, err = l.GetSysMenuView(ctx, menuId)
	if err != nil {
		return nil, err
	}
	return
}

func (l *sSysMenu) DeleteSysMenu(ctx context.Context, menuId int64) (err error) {
	_, err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().MenuId, menuId).Delete()
	return
}

func (s *sSysMenu) GetFastList(ctx context.Context) (res map[int64]*entity.SysMenu, err error) {
	var models []*entity.SysMenu
	if err = dao.SysMenu.Ctx(ctx).Scan(&models); err != nil {
		return
	}

	res = make(map[int64]*entity.SysMenu, len(models))
	for _, v := range models {
		res[v.MenuId] = v
	}
	return
}
