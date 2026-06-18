// Package gen
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2026 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
// @AutoGenerate Version
package gen

import (
	"context"
	"fmt"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/xgorm"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/library/xgorm/hook"
	"xiuadmin/internal/model/genin"
	"xiuadmin/internal/service"
	"xiuadmin/utility/convert"
	"xiuadmin/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"xiuadmin/utility/tree"
)

type sGenTestTree struct{}

func NewGenTestTree() *sGenTestTree {
	return &sGenTestTree{}
}

func init() {
	service.RegisterGenTestTree(NewGenTestTree())
}

// Model 测试树表ORM模型
func (s *sGenTestTree) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.TestTree.Ctx(ctx), option...)
}

// List 获取测试树表列表
func (s *sGenTestTree) List(ctx context.Context, in *genin.TestTreeListParam) (list []*genin.TestTreeListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(genin.TestTreeListModel{})

	// 查询主键
	if in.Id > 0 {
		mod = mod.Where(dao.TestTree.Columns().Id, in.Id)
	}

	// 查询父id
	if in.ParentId > 0 {
		mod = mod.Where(dao.TestTree.Columns().ParentId, in.ParentId)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		in.CreatedAt[1] = in.CreatedAt[1].EndOfDay()
	}
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TestTree.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 树形列表判断是否需要分页
	if in.Page > 0 && in.PageSize > 0 {
		mod = mod.Page(in.Page, in.PageSize)
	}

	// 排序
	mod = mod.OrderDesc(dao.TestTree.Columns().Id)

	// 操作人摘要信息
	mod = mod.Hook(hook.MemberSummary)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取测试树表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出测试树表
func (s *sGenTestTree) Export(ctx context.Context, in *genin.TestTreeListParam) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(genin.TestTreeExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出测试树表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, (totalCount+in.PageSize-1)/in.PageSize, in.Page, len(list))
		exports   []genin.TestTreeExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增测试树表
func (s *sGenTestTree) Edit(ctx context.Context, in *genin.TestTreeEditParam) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		in.ParentId, in.Level, in.Tree, err = xgorm.AutoUpdateTree(ctx, &dao.TestTree, in.Id, in.ParentId, &xgorm.TreeFiledOption{
			IdField:    dao.TestTree.Columns().Id,
			PidField:   dao.TestTree.Columns().ParentId,
			LevelField: dao.TestTree.Columns().Level,
			TreeField:  dao.TestTree.Columns().Tree,
		})
		if err != nil {
			return err
		}

		// 修改
		if in.Id > 0 {
			in.UpdatedAt = gtime.Now()
			updatedBy := contexts.GetUserId(ctx)
			in.UpdatedBy = &updatedBy
			if _, err = s.Model(ctx).
				Fields(genin.TestTreeUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改测试树表失败，请稍后重试！")
			}
			return
		}

		// 新增
		tenantId := contexts.GetTenantId(ctx)
		in.TenantId = &tenantId
		createdDept := contexts.GetDeptId(ctx)
		in.CreatedDept = &createdDept
		in.CreatedAt = gtime.Now()
		createdBy := contexts.GetUserId(ctx)
		in.CreatedBy = &createdBy
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(genin.TestTreeInsertFields{}).
			Data(in).Insert(); err != nil {
			err = gerror.Wrap(err, "新增测试树表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除测试树表
func (s *sGenTestTree) Delete(ctx context.Context, in *genin.TestTreeDeleteParam) (err error) {
	count, err := dao.TestTree.Ctx(ctx).Where(dao.TestTree.Columns().ParentId, in.Id).Count()
	if err != nil {
		err = gerror.Wrap(err, "查询测试树表下级失败，请稍后重试！")
		return err
	}
	if count > 0 {
		return gerror.New("请先删除该测试树表下的所有下级！")
	}

	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.TestTree.Columns().DeletedBy: contexts.GetUserId(ctx),
		dao.TestTree.Columns().DeletedAt: gtime.Now(),
	}).Unscoped().Update(); err != nil {
		err = gerror.Wrap(err, "删除测试树表失败，请稍后重试！")
		return
	}
	return
}

// View 获取测试树表指定信息
func (s *sGenTestTree) View(ctx context.Context, in *genin.TestTreeViewParam) (res *genin.TestTreeViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Hook(hook.MemberSummary).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取测试树表信息，请稍后重试！")
		return
	}
	return
}

// TreeOption 获取测试树表关系树选项
func (s *sGenTestTree) TreeOption(ctx context.Context) (nodes []tree.Node, err error) {
	var models []*genin.TestTreeTreeOption
	if err = s.Model(ctx).Fields(genin.TestTreeTreeOption{}).OrderAsc(dao.TestTree.Columns().ParentId).OrderDesc(dao.TestTree.Columns().Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, "获取测试树表关系树选项失败！")
		return
	}
	nodes = make([]tree.Node, len(models))
	for i, v := range models {
		nodes[i] = v
	}
	return tree.ListToTree(0, nodes), nil
}
