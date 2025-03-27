// package xgorm
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package xgorm

import (
	"context"
	"slices"
	gormmodel "xiuadmin/internal/library/xgorm/gorm_model"
	"xiuadmin/utility/tree"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	idDefaulField    = "id"
	pidDefaulField   = "pid"
	levelDefaulField = "level"
	treeDefaulField  = "tree"
)

type TreeFiledOption struct {
	IdField    string
	PidField   string
	LevelField string
	TreeField  string
}

type daoInstance interface {
	Table() string
	Ctx(ctx context.Context) *gdb.Model
}

func CustomTreeFiledToDefault(old map[string]interface{}, option *TreeFiledOption) *gormmodel.DefaultTree {
	pd := &gormmodel.DefaultTree{}
	if id, ok := old[option.IdField]; ok {
		pd.Id = gconv.Int64(id)
	}
	if pid, ok := old[option.PidField]; ok {
		pd.Pid = gconv.Int64(pid)
	}
	if level, ok := old[option.LevelField]; ok {
		pd.Level = gconv.Int(level)
	}
	if tree, ok := old[option.TreeField]; ok {
		pd.Tree = gconv.String(tree)
	}
	return pd
}

// AutoUpdateTree 自动更新关系树
func AutoUpdateTree(ctx context.Context, dao daoInstance, id, pid int64, option *TreeFiledOption) (newPid int64, newLevel int, newTree string, err error) {
	if option == nil {
		option = &TreeFiledOption{
			IdField:    idDefaulField,
			PidField:   pidDefaulField,
			LevelField: levelDefaulField,
			TreeField:  treeDefaulField,
		}
	}
	if err = CheckTreeTable(ctx, dao, option); err != nil {
		return
	}
	if pid <= 0 {
		newPid = 0
		newLevel = 1
		newTree = ""
	} else {
		var pdMap map[string]interface{}
		if err = dao.Ctx(ctx).Fields(option.IdField, option.PidField, option.LevelField, option.TreeField).WherePri(pid).Scan(&pdMap); err != nil {
			return 0, 0, "", err
		}

		if len(pdMap) == 0 {
			return 0, 0, "", gerror.New("未查询到树表上级信息，请检查！")
		}
		pd := CustomTreeFiledToDefault(pdMap, option)

		if id > 0 && slices.Contains(tree.GetIds(pd.Tree), id) {
			return 0, 0, "", gerror.New("上级不能设为自己当前的子级！")
		}

		newPid = pid
		newLevel = pd.Level + 1
		newTree = tree.GenLabel(pd.Tree, pid)
	}

	if id > 0 {
		if pid == id {
			return 0, 0, "", gerror.New("上级不能是自己！")
		}

		var models *gormmodel.DefaultTree
		if err = dao.Ctx(ctx).WherePri(id).Scan(&models); err != nil {
			return 0, 0, "", err
		}

		if models == nil {
			return 0, 0, "", gerror.New("树表信息不存在，请检查！")
		}

		// 上级发生变化时，遍历修改其所有的下级关系树
		if models.Pid != pid {
			if err = updateChildrenTree(ctx, dao, models.Id, newLevel, newTree, option); err != nil {
				return
			}
		}
	}
	return
}

// CheckTreeTable 检查树表
func CheckTreeTable(ctx context.Context, dao daoInstance, option *TreeFiledOption) (err error) {
	fields, err := GetFieldsToSlice(ctx, dao)
	if err != nil {
		return err
	}

	if !slices.Contains(fields, option.PidField) {
		return gerror.New("树表必须包含`pid`字段")
	}

	if !slices.Contains(fields, option.LevelField) {
		return gerror.New("树表必须包含`level`字段")
	}

	if !slices.Contains(fields, option.TreeField) {
		return gerror.New("树表必须包含`tree`字段")
	}
	return
}

// updateChildrenTree 更新下级关系树
func updateChildrenTree(ctx context.Context, dao daoInstance, pid int64, pLevel int, pTree string, option *TreeFiledOption) (err error) {
	var list []*gormmodel.DefaultTree
	var pdMaps []map[string]interface{}

	if err = dao.Ctx(ctx).Fields(option.IdField, option.PidField, option.LevelField, option.TreeField).Where(option.PidField, pid).Scan(&pdMaps); err != nil {
		return
	}

	for _, v := range pdMaps {
		list = append(list, CustomTreeFiledToDefault(v, option))
	}

	if len(list) == 0 {
		return
	}

	newLevel := pLevel + 1
	newTree := tree.GenLabel(pTree, pid)

	var updateIds []int64
	for _, v := range list {
		updateIds = append(updateIds, v.Id)
		if err = updateChildrenTree(ctx, dao, v.Id, newLevel, newTree, option); err != nil {
			return
		}
	}

	if len(updateIds) > 0 {
		update := g.Map{
			option.LevelField: newLevel,
			option.TreeField:  newTree,
		}
		_, err = dao.Ctx(ctx).WhereIn(option.IdField, updateIds).Data(update).Update()
	}
	return
}
