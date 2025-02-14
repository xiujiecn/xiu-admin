package model

import "server/internal/model/entity"

type SysMenuTree struct {
	entity.SysMenu
	Children []*SysMenuTree `json:"children"`
}
