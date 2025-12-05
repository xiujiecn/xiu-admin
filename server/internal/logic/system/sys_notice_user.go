// Package gen
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
// @AutoGenerate Version
package system

import (
	"context"
	"fmt"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/library/xgorm/hook"
	"xiuadmin/internal/model"
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

type sSysNoticeUser struct{}

func NewSysNoticeUser() *sSysNoticeUser {
	return &sSysNoticeUser{}
}

func init() {
	service.RegisterSysNoticeUser(NewSysNoticeUser())
}

// Model 用户通知公告表ORM模型
func (s *sSysNoticeUser) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.SysNoticeUser.Ctx(ctx), option...)
}

// List 获取用户通知公告表列表
func (s *sSysNoticeUser) List(ctx context.Context, in *model.SysNoticeUserListParam) (list []*model.SysNoticeUserListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 根据当前用户查询
	mod = mod.Where(dao.SysNoticeUser.Columns().UserId, contexts.GetUserId(ctx))
	mod = mod.Where(dao.SysNoticeUser.Columns().Status, consts.SysStatusNormal)
	// 字段过滤
	mod = mod.Fields(model.SysNoticeUserListModel{})

	if in.IsRead != "" {
		mod = mod.Where(dao.SysNoticeUser.Columns().IsRead, in.IsRead)
	}

	if in.NoticeId > 0 {
		mod = mod.Where(dao.SysNoticeUser.Columns().NoticeId, in.NoticeId)
	}

	if in.NoticeTitle != "" {
		mod = mod.WhereLike(dao.SysNoticeUser.Columns().NoticeTitle, "%"+in.NoticeTitle+"%")
	}

	if in.NoticeType != "" {
		mod = mod.Where(dao.SysNoticeUser.Columns().NoticeType, in.NoticeType)
	}
	// 分页
	mod = mod.Page(in.Page, in.PageSize)

	// 排序
	mod = mod.OrderDesc(dao.SysNoticeUser.Columns().Id)

	// 操作人摘要信息
	mod = mod.Hook(hook.MemberSummary)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取用户通知公告表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出用户通知公告表
func (s *sSysNoticeUser) Export(ctx context.Context, in *model.SysNoticeUserListParam) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(model.SysNoticeUserExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出用户通知公告表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, (totalCount+in.PageSize-1)/in.PageSize, in.Page, len(list))
		exports   []model.SysNoticeUserExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增用户通知公告表
func (s *sSysNoticeUser) Edit(ctx context.Context, in *model.SysNoticeUserEditParam) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			in.UpdatedBy = contexts.GetUserId(ctx)
			in.UpdatedAt = gtime.Now()
			if _, err = s.Model(ctx).
				Fields(model.SysNoticeUserUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改用户通知公告表失败，请稍后重试！")
			}
			return
		}

		// 新增
		in.TenantId = contexts.GetTenantId(ctx)
		in.CreatedDept = contexts.GetDeptId(ctx)
		in.CreatedBy = contexts.GetUserId(ctx)
		in.CreatedAt = gtime.Now()
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(model.SysNoticeUserInsertFields{}).
			Data(in).Insert(); err != nil {
			err = gerror.Wrap(err, "新增用户通知公告表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除用户通知公告表
func (s *sSysNoticeUser) Delete(ctx context.Context, in *model.SysNoticeUserDeleteParam) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除用户通知公告表失败，请稍后重试！")
		return
	}
	return
}

// View 获取用户通知公告表指定信息
func (s *sSysNoticeUser) View(ctx context.Context, in *model.SysNoticeUserViewParam) (res *model.SysNoticeUserViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Hook(hook.MemberSummary).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取用户通知公告表信息，请稍后重试！")
		return
	}
	return
}

// Status 更新用户通知公告表状态
func (s *sSysNoticeUser) Status(ctx context.Context, in *model.SysNoticeUserStatusParam) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.SysNoticeUser.Columns().Status:    in.Status,
		dao.SysNoticeUser.Columns().UpdatedBy: contexts.GetUserId(ctx),
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新用户通知公告表状态失败，请稍后重试！")
		return
	}
	return
}

// Read 已读
func (s *sSysNoticeUser) Read(ctx context.Context, in *model.SysNoticeUserReadParam) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Ids).Data(g.Map{
		dao.SysNoticeUser.Columns().IsRead:    1,
		dao.SysNoticeUser.Columns().UpdatedBy: contexts.GetUserId(ctx),
		dao.SysNoticeUser.Columns().UpdatedAt: gtime.Now(),
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新用户通知公告表状态失败，请稍后重试！")
		return
	}
	return
}
