package system

import (
	"context"
	"fmt"
	"xiujieadmin/internal/consts"
	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/library/contexts"
	"xiujieadmin/internal/library/xgorm/handler"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/do"
	"xiujieadmin/internal/service"

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

func (l *sSysDept) GetDeptList(ctx context.Context, query model.SysDeptListParam) (items []*model.SysDeptListModel, total int, err error) {
	m := l.Model(ctx)
	if query.DeptName != "" {
		m = m.WhereLike(dao.SysDept.Columns().DeptName, "%"+query.DeptName+"%")
	}
	if query.Status != "" {
		m = m.Where(dao.SysDept.Columns().Status, query.Status)
	}
	err = m.Page(1, 5000).Order(dao.SysDept.Columns().OrderNum, "ASC").Scan(&items)
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
	err = l.Model(ctx).Where(dao.SysDept.Columns().DeptId, id).Scan(&dept)
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

func (l *sSysDept) GetDeptTree(ctx context.Context) (items []*model.SysDeptTreeModel, err error) {
	depts, _, err := l.GetDeptList(ctx, model.SysDeptListParam{})
	if err != nil {
		return nil, err
	}
	items, err = l.DeptTree(ctx, nil, depts, "0")
	if err != nil {
		return nil, err
	}
	return items, nil
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
	data := do.SysDept{}
	gconv.Struct(dept, &data)
	data.UpdatedBy = user.ID
	data.UpdatedAt = gtime.Now()
	_, err = l.Model(ctx).Data(data).Where(dao.SysDept.Columns().DeptId, dept.DeptId).Update()
	if err != nil {
		return dept.DeptId, err
	}
	// 刷新部门祖先
	err = l.RefreshDeptAncestors(ctx)
	if err != nil {
		return 0, err
	}
	return dept.DeptId, nil
}

func (l *sSysDept) DeleteDept(ctx context.Context, dept *model.SysDeptDeleteModel) (deptId int64, err error) {
	count, err := l.Model(ctx).Where(dao.SysDept.Columns().ParentId, dept.DeptId).Count()
	if err != nil {
		return dept.DeptId, err
	}
	if count > 0 {
		return dept.DeptId, gerror.NewCode(consts.CodeDeptHasSub, "该部门下有子部门，不能删除")
	}
	_, err = l.Model(ctx).Where(dao.SysDept.Columns().DeptId, dept.DeptId).Delete()
	if err != nil {
		return dept.DeptId, err
	}
	return dept.DeptId, nil
}

// 刷新部门 ancestors
func (l *sSysDept) RefreshDeptAncestors(ctx context.Context) (err error) {
	var depts []*model.SysDeptViewModel
	err = l.Model(ctx, &handler.Option{
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
