// Package gen
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
// @AutoGenerate Version
package gen

import (
	"context"
	"fmt"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
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
)

type sGenTestCategory struct{}

func NewGenTestCategory() *sGenTestCategory {
	return &sGenTestCategory{}
}

func init() {
	service.RegisterGenTestCategory(NewGenTestCategory())
}

// Model 测试分类ORM模型
func (s *sGenTestCategory) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.TestCategory.Ctx(ctx), option...)
}

// List 获取测试分类列表
func (s *sGenTestCategory) List(ctx context.Context, in *genin.TestCategoryListParam) (list []*genin.TestCategoryListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(genin.TestCategoryListModel{})

	// 查询分类ID
	if in.Id > 0 {
		mod = mod.Where(dao.TestCategory.Columns().Id, in.Id)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.TestCategory.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		in.CreatedAt[1] = in.CreatedAt[1].EndOfDay()
	}
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TestCategory.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PageSize)

	// 排序
	mod = mod.OrderAsc(dao.TestCategory.Columns().Sort).OrderDesc(dao.TestCategory.Columns().Id)

	// 操作人摘要信息
	mod = mod.Hook(hook.MemberSummary)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取测试分类列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出测试分类
func (s *sGenTestCategory) Export(ctx context.Context, in *genin.TestCategoryListParam) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(genin.TestCategoryExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出测试分类-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, (totalCount+in.PageSize-1)/in.PageSize, in.Page, len(list))
		exports   []genin.TestCategoryExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增测试分类
func (s *sGenTestCategory) Edit(ctx context.Context, in *genin.TestCategoryEditParam) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			in.UpdatedAt = gtime.Now()
			in.UpdatedBy = contexts.GetUserId(ctx)
			if _, err = s.Model(ctx).
				Fields(genin.TestCategoryUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改测试分类失败，请稍后重试！")
			}
			return
		}

		// 新增
		in.CreatedDept = contexts.GetDeptId(ctx)
		in.CreatedAt = gtime.Now()
		in.CreatedBy = contexts.GetUserId(ctx)
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(genin.TestCategoryInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增测试分类失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除测试分类
func (s *sGenTestCategory) Delete(ctx context.Context, in *genin.TestCategoryDeleteParam) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.TestCategory.Columns().DeletedAt: gtime.Now(),
		dao.TestCategory.Columns().DeletedBy: contexts.GetUserId(ctx),
	}).Unscoped().Update(); err != nil {
		err = gerror.Wrap(err, "删除测试分类失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取测试分类最大排序
func (s *sGenTestCategory) MaxSort(ctx context.Context, in *genin.TestCategoryMaxSortParam) (res *genin.TestCategoryMaxSortModel, err error) {
	if err = dao.TestCategory.Ctx(ctx).Fields(dao.TestCategory.Columns().Sort).OrderDesc(dao.TestCategory.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取测试分类最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(genin.TestCategoryMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取测试分类指定信息
func (s *sGenTestCategory) View(ctx context.Context, in *genin.TestCategoryViewParam) (res *genin.TestCategoryViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Hook(hook.MemberSummary).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取测试分类信息，请稍后重试！")
		return
	}
	return
}

// Status 更新测试分类状态
func (s *sGenTestCategory) Status(ctx context.Context, in *genin.TestCategoryStatusParam) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.TestCategory.Columns().Status:    in.Status,
		dao.TestCategory.Columns().UpdatedBy: contexts.GetUserId(ctx),
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新测试分类状态失败，请稍后重试！")
		return
	}
	return
}