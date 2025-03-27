// Package gen
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
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
)

type sGenTestDemo struct{}

func NewGenTestDemo() *sGenTestDemo {
	return &sGenTestDemo{}
}

func init() {
	service.RegisterGenTestDemo(NewGenTestDemo())
}

// Model 测试单表ORM模型
func (s *sGenTestDemo) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.TestDemo.Ctx(ctx), option...)
}

// List 获取测试单表列表
func (s *sGenTestDemo) List(ctx context.Context, in *genin.TestDemoListParam) (list []*genin.TestDemoListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.FieldsPrefix(dao.TestDemo.Table(), genin.TestDemoListModel{})
	mod = mod.Fields(xgorm.JoinFields(ctx, genin.TestDemoListModel{}, &dao.SysDept, "dept"))

	// 关联表字段
	mod = mod.LeftJoinOnFields(dao.SysDept.Table(), dao.TestDemo.Columns().DeptId, "=", dao.SysDept.Columns().DeptId)

	// 查询主键
	if in.Id > 0 {
		mod = mod.Where(dao.TestDemo.Columns().Id, in.Id)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		in.CreatedAt[1] = in.CreatedAt[1].EndOfDay()
	}
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TestDemo.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 查询部门名称
	if in.DeptDeptName != "" {
		mod = mod.WherePrefixLike(dao.SysDept.Table(), dao.SysDept.Columns().DeptName, in.DeptDeptName)
	}

	// 分页
	mod = mod.Page(in.Page, in.PageSize)

	// 排序
	mod = mod.OrderDesc(dao.TestDemo.Table() + "." + dao.TestDemo.Columns().Id)

	// 操作人摘要信息
	mod = mod.Hook(hook.MemberSummary)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取测试单表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出测试单表
func (s *sGenTestDemo) Export(ctx context.Context, in *genin.TestDemoListParam) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(genin.TestDemoExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出测试单表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, (totalCount+in.PageSize-1)/in.PageSize, in.Page, len(list))
		exports   []genin.TestDemoExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增测试单表
func (s *sGenTestDemo) Edit(ctx context.Context, in *genin.TestDemoEditParam) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			in.UpdatedAt = gtime.Now()
			in.UpdatedBy = contexts.GetUserId(ctx)
			if _, err = s.Model(ctx).
				Fields(genin.TestDemoUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改测试单表失败，请稍后重试！")
			}
			return
		}

		// 新增
		in.TenantId = contexts.GetTenantId(ctx)
		in.CreatedDept = contexts.GetDeptId(ctx)
		in.CreatedAt = gtime.Now()
		in.CreatedBy = contexts.GetUserId(ctx)
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(genin.TestDemoInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增测试单表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除测试单表
func (s *sGenTestDemo) Delete(ctx context.Context, in *genin.TestDemoDeleteParam) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.TestDemo.Columns().DeletedBy: contexts.GetUserId(ctx),
		dao.TestDemo.Columns().DeletedAt: gtime.Now(),
	}).Unscoped().Update(); err != nil {
		err = gerror.Wrap(err, "删除测试单表失败，请稍后重试！")
		return
	}
	return
}

// View 获取测试单表指定信息
func (s *sGenTestDemo) View(ctx context.Context, in *genin.TestDemoViewParam) (res *genin.TestDemoViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Hook(hook.MemberSummary).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取测试单表信息，请稍后重试！")
		return
	}
	return
}
