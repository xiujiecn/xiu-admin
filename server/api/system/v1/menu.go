package v1

import (
	"xiujieadmin/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 路由菜单
type RouteMenu struct {
	Id        int64        `json:"id" dc:"路由ID"`
	ParentId  int64        `json:"parentId" dc:"父路由ID"`
	Path      string       `json:"path" dc:"路由路径"`
	Redirect  *string      `json:"redirect,omitempty" dc:"重定向路径"`
	Name      string       `json:"name" dc:"路由名称"` // 必须唯一
	Component string       `json:"component" dc:"组件路径"`
	Meta      *RouteMeta   `json:"meta,omitempty" dc:"路由元数据"`
	Children  []*RouteMenu `json:"children,omitempty" dc:"子路由"`
}

// 路由元数据
// 对应 vue/packages/@core/base/typings/src/vue-router.d.ts
type RouteMeta struct {
	Title                    string            `json:"title" dc:"菜单标题"`
	Icon                     *string           `json:"icon,omitempty" dc:"菜单图标"`
	Order                    *int              `json:"order,omitempty" dc:"菜单排序"`
	OpenInNewWindow          *bool             `json:"openInNewWindow,omitempty" dc:"是否在新窗口打开"`
	NoBasicLayout            *bool             `json:"noBasicLayout,omitempty" dc:"是否不使用基础布局"`
	HideInMenu               *bool             `json:"hideInMenu,omitempty" dc:"是否隐藏菜单"`
	HideInBreadcrumb         *bool             `json:"hideInBreadcrumb,omitempty" dc:"是否隐藏面包屑"`
	HideChildrenInMenu       *bool             `json:"hideChildrenInMenu,omitempty" dc:"是否隐藏子菜单"`
	HideInTab                *bool             `json:"hideInTab,omitempty" dc:"是否隐藏标签页"`
	MenuVisibleWithForbidden *bool             `json:"menuVisibleWithForbidden,omitempty" dc:"菜单可见但访问会被重定向到403"`
	IgnoreAccess             *bool             `json:"ignoreAccess,omitempty" dc:"忽略权限，直接可以访问"`
	KeepAlive                *bool             `json:"keepAlive,omitempty" dc:"开启KeepAlive缓存"`
	Link                     *string           `json:"link,omitempty" dc:"外链-跳转路径"`
	Query                    map[string]string `json:"query,omitempty" dc:"路由所携带的参数"`
	IframeSrc                *string           `json:"iframeSrc,omitempty" dc:"iframe地址"`
	MaxNumOfOpenTab          *int              `json:"maxNumOfOpenTab,omitempty" dc:"标签页最大打开数量"`
	Badge                    *string           `json:"badge,omitempty" dc:"徽标"`
	BadgeType                *string           `json:"badgeType,omitempty" dc:"徽标类型"`
	BadgeVariants            *string           `json:"badgeVariants,omitempty" dc:"徽标颜色"`
	Authority                []string          `json:"authority,omitempty" dc:"需要特定的角色标识才可以访问"`
	AffixTab                 *bool             `json:"affixTab,omitempty" dc:"是否固定标签页"`
	AffixTabOrder            *int              `json:"affixTabOrder,omitempty" dc:"固定标签页的顺序"`
	ActiveIcon               *string           `json:"activeIcon,omitempty" dc:"激活图标"`
	ActivePath               *string           `json:"activePath,omitempty" dc:"当前激活的菜单"`
}

// 获取用户所有菜单
type MenuAllReq struct {
	g.Meta `path:"/menu/all" method:"get" tags:"系统" summary:"获取用户所有菜单"`
}

//	type MenuAllRes struct {
//		Data []RouteMenu `json:"data" dc:"菜单列表"`
//	}
type MenuAllRes = []*RouteMenu

type MenuListReq struct {
	g.Meta `path:"/menu/list" method:"get" tags:"系统" summary:"获取菜单列表"`
	model.SysMenuListParam
}

type MenuListRes struct {
	Data  []*model.SysMenuListModel `json:"items" dc:"菜单列表"`
	Total int                       `json:"total" dc:"总数"`
}

type MenuViewReq struct {
	g.Meta `path:"/menu/view" method:"get" tags:"系统" summary:"获取菜单详情"`
	MenuId int64 `json:"menuId" dc:"菜单ID"`
}

type MenuViewRes struct {
	model.SysMenuViewModel
}

type MenuAddReq struct {
	g.Meta `path:"/menu/add" method:"post" tags:"系统" summary:"添加菜单"`
	model.SysMenuAddModel
}

type MenuAddRes struct {
	model.SysMenuViewModel
}

type MenuUpdateReq struct {
	g.Meta `path:"/menu/update" method:"post" tags:"系统" summary:"更新菜单"`
	model.SysMenuUpdateModel
}

type MenuUpdateRes struct {
	model.SysMenuViewModel
}

type MenuDeleteReq struct {
	g.Meta `path:"/menu/delete" method:"post" tags:"系统" summary:"删除菜单"`
	MenuId int64 `json:"menuId" dc:"菜单ID"`
}

type MenuDeleteRes struct {
}

type MenuTreeReq struct {
	g.Meta `path:"/menu/tree" method:"get" tags:"系统" summary:"获取菜单树"`
}

type MenuTreeRes struct {
	Data []*model.SysMenuTreeModel `json:"items" dc:"菜单树"`
}
