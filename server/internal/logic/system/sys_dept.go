// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"fmt"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/event"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/do"
	"xiuadmin/internal/model/request"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysDept struct {
}

func NewSysDept() *sSysDept {
	return &sSysDept{}
}

func init() {
	service.RegisterSysDept(NewSysDept())
}

func (l *sSysDept) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.SysDept.Ctx(ctx), option...)
}

func (l *sSysDept) ModelQuery(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   true,
		})
	}
	return handler.Model(service.MemoryDB().DB(ctx).Ctx(ctx).Model(dao.SysDept.Table()), option...)
}

func (l *sSysDept) GetDeptList(ctx context.Context, query model.SysDeptListParam) (items []*model.SysDeptListModel, total int, err error) {
	m := l.ModelQuery(ctx)
	if query.DeptName != "" {
		m = m.WhereLike(dao.SysDept.Columns().DeptName, "%"+query.DeptName+"%")
	}
	if query.Status != "" {
		m = m.Where(dao.SysDept.Columns().Status, query.Status)
	}
	if query.DeptType != nil {
		m = m.Where(dao.SysDept.Columns().DeptType, *query.DeptType)
	}
	// 排序：按照显示顺序OrderNum从小到大排序，如果相同则按照部门ID从小到大排序
	err = m.Page(1, 5000).Order(dao.SysDept.Columns().OrderNum, "ASC").Order(dao.SysDept.Columns().DeptId, "ASC").Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	total = len(items)
	return
}

func (l *sSysDept) GetDeptById(ctx context.Context, id int64) (dept *model.SysDeptViewModel, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return nil, gerror.NewCode(consts.CodeLoginExpired, "登录已过期")
	}
	err = l.ModelQuery(ctx).Where(dao.SysDept.Columns().DeptId, id).Scan(&dept)
	if err != nil {
		return nil, err
	}
	return dept, nil
}

// 构建树结构
func (l *sSysDept) DeptTree(ctx context.Context, parentDept *model.SysDeptTreeModel, deptList []*model.SysDeptListModel, ancestors string) (data []*model.SysDeptTreeModel, err error) {
	data = make([]*model.SysDeptTreeModel, 0)
	for _, dept := range deptList {
		pId := int64(0)
		if parentDept != nil {
			pId = parentDept.DeptId
		}
		if dept.ParentId == pId {
			item := &model.SysDeptTreeModel{
				DeptId:    dept.DeptId,
				ParentId:  dept.ParentId,
				DeptName:  dept.DeptName,
				Key:       fmt.Sprintf("%d", dept.DeptId),
				Children:  nil,
				Ancestors: ancestors,
			}
			subDept, err := l.DeptTree(ctx, item, deptList, ancestors+","+fmt.Sprintf("%d", dept.DeptId))
			if err != nil {
				return nil, err
			}
			item.Children = subDept
			data = append(data, item)
		}
	}
	return
}

func (l *sSysDept) GetDeptTree(ctx context.Context, query model.SysDeptTreeParam) (items []*model.SysDeptTreeModel, err error) {
	depts, _, err := l.GetDeptList(ctx, model.SysDeptListParam{
		DeptType: query.DeptType,
	})
	if err != nil {
		return nil, err
	}
	if len(depts) == 0 {
		return make([]*model.SysDeptTreeModel, 0), nil
	}
	//找出列表中的顶级部门（可能包含多个）
	topDepts := make([]*model.SysDeptTreeModel, 0)
	deptMap := make(map[int64]*model.SysDeptTreeModel)
	for _, dept := range depts {
		deptMap[dept.DeptId] = &model.SysDeptTreeModel{
			DeptId:    dept.DeptId,
			ParentId:  dept.ParentId,
			DeptName:  dept.DeptName,
			Key:       fmt.Sprintf("%d", dept.DeptId),
			Children:  nil,
			Ancestors: "0",
		}
	}

	for _, dept := range depts {
		if _, ok := deptMap[dept.ParentId]; !ok {
			topDepts = append(topDepts, deptMap[dept.DeptId])
		}
	}

	for _, dept := range topDepts {
		children, err := l.DeptTree(ctx, dept, depts, fmt.Sprintf("%d", dept.DeptId))
		if err != nil {
			return nil, err
		}
		dept.Children = children
	}

	return topDepts, nil
}

// 递归构建结构
func (l *sSysDept) RecursionDeptIds(ctx context.Context, parentId int64, deptList []*model.SysDeptListModel, data *[]int64) (err error) {
	for _, dept := range deptList {
		if dept.ParentId == parentId {
			*data = append(*data, dept.DeptId)
			err := l.RecursionDeptIds(ctx, dept.DeptId, deptList, data)
			if err != nil {
				return err
			}
		}
	}
	return
}

// 根据父部门id获取部门列表
func (l *sSysDept) GetDeptIdsByParentId(ctx context.Context, parentId int64) (ids []int64, err error) {
	depts, _, err := l.GetDeptList(ctx, model.SysDeptListParam{})
	if err != nil {
		return nil, err
	}
	ids = make([]int64, 0)
	err = l.RecursionDeptIds(ctx, parentId, depts, &ids)
	if err != nil {
		return nil, err
	}
	// g.Log().Infof(ctx, "sSysDept.GetDeptIdsByParentId 获取部门列表: ids:%+v", ids)
	return ids, nil
}

func (l *sSysDept) AddDept(ctx context.Context, dept *model.SysDeptAddModel) (deptId int64, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return 0, gerror.NewCode(consts.CodeLoginExpired, "登录已过期")
	}
	if dept.DeptType == 1 {
		isCompany, err := l.ValidateParentIsCompany(ctx, dept.ParentId)
		if err != nil {
			g.Log().Errorf(ctx, "sSysDept.AddDept ValidateParentIsCompany err: %v, parentId:%d, dept:%+v", err, dept.ParentId, dept)
			return 0, err
		}
		if !isCompany {
			return 0, gerror.New("上级部门必须是公司")
		}
	}

	data := do.SysDept{}
	gconv.Struct(dept, &data)
	data.CreatedDept = user.DeptId
	data.CreatedBy = user.ID
	data.CreatedAt = gtime.Now()
	data.UpdatedBy = user.ID
	data.UpdatedAt = gtime.Now()
	data.TenantId = user.TenantId

	deptId, err = l.Model(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return 0, err
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysDeptCreate, deptId)
	// 刷新部门祖先
	err = l.RefreshDeptAncestors(ctx)
	if err != nil {
		return 0, err
	}
	return deptId, nil
}

func (l *sSysDept) EditDept(ctx context.Context, dept *model.SysDeptEditModel) (deptId int64, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return dept.DeptId, gerror.NewCode(consts.CodeLoginExpired, "登录已过期")
	}
	// 查询现有数据
	existingDept, err := l.GetDeptById(ctx, dept.DeptId)
	if err != nil {
		g.Log().Errorf(ctx, "sSysDept.EditDept GetDeptById err: %v, deptId:%d, dept:%+v", err, dept.DeptId, gconv.String(dept))
		return 0, err
	}
	if existingDept == nil || existingDept.DeptId == 0 || existingDept.DeptId != dept.DeptId {
		g.Log().Warningf(ctx, "sSysDept.EditDept GetDeptById err: %v, deptId:%d, dept:%+v", err, dept.DeptId, gconv.String(dept))
		return 0, gerror.New("部门不存在")
	}
	newParentId := existingDept.ParentId
	if dept.ParentId != nil {
		newParentId = *dept.ParentId
	}
	if dept.DeptType != nil && *dept.DeptType == 1 {
		isCompany, err := l.ValidateParentIsCompany(ctx, newParentId)
		if err != nil {
			return 0, err
		}
		if !isCompany {
			return 0, gerror.New("上级组织必须是公司")
		}
	}

	if dept.DeptType != nil && *dept.DeptType != existingDept.DeptType {
		if existingDept.DeptType == 1 {
			return 0, gerror.New("组织类型不能修改")
		}
	}

	data := do.SysDept{}
	gconv.Struct(dept, &data)
	data.UpdatedBy = user.ID
	data.UpdatedAt = gtime.Now()
	_, err = l.Model(ctx).Data(data).Where(dao.SysDept.Columns().DeptId, dept.DeptId).Update()
	if err != nil {
		return dept.DeptId, err
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysDeptUpdate, dept.DeptId)
	// 刷新部门祖先
	err = l.RefreshDeptAncestors(ctx)
	if err != nil {
		return 0, err
	}
	return dept.DeptId, nil
}

func (l *sSysDept) DeleteDept(ctx context.Context, dept *model.SysDeptDeleteModel) (deptId int64, err error) {
	count, err := l.ModelQuery(ctx).Where(dao.SysDept.Columns().ParentId, dept.DeptId).Count()
	if err != nil {
		return dept.DeptId, err
	}
	if count > 0 {
		return dept.DeptId, gerror.NewCode(consts.CodeDeptHasSub, "该部门下有子部门，不能删除")
	}

	// 部门下有用户，不能删除
	userList, _, err := service.SysUser().List(ctx, &request.PageInfo{
		Page:     1,
		PageSize: 100,
	}, &model.UserListParam{
		DeptId: dept.DeptId,
	})
	if err != nil {
		return dept.DeptId, err
	}
	if len(userList) > 0 {
		return dept.DeptId, gerror.New("该部门下有用户，不能删除")
	}

	_, err = l.Model(ctx).Where(dao.SysDept.Columns().DeptId, dept.DeptId).Delete()
	if err != nil {
		return dept.DeptId, err
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysDeptDelete, dept.DeptId)
	return dept.DeptId, nil
}

// 刷新部门 ancestors
func (l *sSysDept) RefreshDeptAncestors(ctx context.Context) (err error) {
	var depts []*model.SysDeptViewModel
	err = l.ModelQuery(ctx, &handler.Option{
		FilterTenant: true,
		FilterAuth:   false,
	}).Scan(&depts)
	if err != nil {
		return err
	}
	for _, dept := range depts {
		ancestors := ""
		l.GetParentIDAncestors(ctx, depts, dept.ParentId, &ancestors)

		if ancestors != dept.Ancestors {
			_, err = l.Model(ctx).Data(do.SysDept{
				Ancestors: ancestors,
			}).Where(dao.SysDept.Columns().DeptId, dept.DeptId).Update()
			if err != nil {
				return err
			}
			event.EventsInstance().Emit(ctx, consts.EventKeyDBSysDeptUpdate, dept.DeptId)
			g.Log().Infof(ctx, "sSysDept.RefreshDeptAncestors 刷新部门祖先: deptId:%d, ancestors:%s", dept.DeptId, ancestors)
		}
	}
	return nil
}

func (l *sSysDept) GetParentIDAncestors(ctx context.Context, depts []*model.SysDeptViewModel, printId int64, ancestors *string) (err error) {
	for _, dept := range depts {
		if dept.DeptId == printId {
			*ancestors = fmt.Sprintf("%d", dept.DeptId) + "," + *ancestors
			l.GetParentIDAncestors(ctx, depts, dept.ParentId, ancestors)
			return nil
		}
	}
	*ancestors = "0," + *ancestors
	return nil
}

func (l *sSysDept) GetDeptListByIds(ctx context.Context, ids []int64) (depts []*model.SysDeptListModel, err error) {
	err = l.ModelQuery(ctx).WhereIn(dao.SysDept.Columns().DeptId, ids).Scan(&depts)
	if err != nil {
		return nil, err
	}
	return depts, nil
}

// 验证上级是否是公司
func (l *sSysDept) ValidateParentIsCompany(ctx context.Context, parentId int64) (isCompany bool, err error) {
	dept, err := l.GetDeptById(ctx, parentId)
	if err != nil {
		return false, err
	}
	return dept.DeptType == 1 || dept.ParentId == 0, nil
}

// 获取上级公司名称
func (l *sSysDept) getParentCompanyInfo(ctx context.Context, depts []*model.SysDeptListModel, deptId int64) *model.SysDeptListModel {
	var topDept *model.SysDeptListModel
	for _, dept := range depts {
		if dept.ParentId == 0 {
			topDept = dept
		}
		if dept.DeptId == deptId {
			if dept.ParentId == 0 {
				return dept
			}
			if dept.DeptType == 1 {
				return dept
			}
			return l.getParentCompanyInfo(ctx, depts, dept.ParentId)
		}
	}
	return topDept
}

// 获取组织公司Map
func (l *sSysDept) GetDeptCompanyMap(ctx context.Context) (map[int64]string, error) {
	depts, _, err := l.GetDeptList(ctx, model.SysDeptListParam{})
	if err != nil {
		return nil, err
	}
	companyMap := make(map[int64]string)
	for _, dept := range depts {
		if dept.ParentId == 0 {
			companyMap[0] = dept.DeptName
		}
		if dept.DeptType == 1 {
			companyMap[dept.DeptId] = dept.DeptName
		} else if dept.ParentId == 0 {
			companyMap[dept.DeptId] = dept.DeptName
		} else {
			parentDept := l.getParentCompanyInfo(ctx, depts, dept.DeptId)
			if parentDept != nil {
				companyMap[dept.DeptId] = parentDept.DeptName
			} else {
				companyMap[dept.DeptId] = ""
			}
		}
	}
	return companyMap, nil
}

// 获取上级公司组织信息
func (l *sSysDept) GetParentCompanyInfo(ctx context.Context, deptId int64) *model.SysDeptListModel {
	depts, _, err := l.GetDeptList(ctx, model.SysDeptListParam{})
	if err != nil {
		g.Log().Errorf(ctx, "sSysDept.GetParentCompanyInfo GetDeptList err: %v, deptId:%d", err, deptId)
		return nil
	}
	return l.getParentCompanyInfo(ctx, depts, deptId)
}
