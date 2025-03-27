// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"fmt"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/do"
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

type sSysPost struct {
}

func NewSysPost() *sSysPost {
	return &sSysPost{}
}

func init() {
	service.RegisterSysPost(NewSysPost())
}

func (l *sSysPost) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.SysPost.Ctx(ctx), option...)
}

func (l *sSysPost) List(ctx context.Context, query *model.SysPostListParam) (items []*model.SysPostListModel, total int, err error) {
	m := l.Model(ctx)
	deptIds := make([]int64, 0)
	if query.DeptId != 0 {
		deptIds = append(deptIds, query.DeptId)
	} else if query.BelongDeptId != 0 {
		deptIds = append(deptIds, query.BelongDeptId)
		subDeptIds, err := service.SysDept().GetDeptIdsByParentId(ctx, query.BelongDeptId)
		if err != nil {
			return nil, 0, err
		}
		deptIds = append(deptIds, subDeptIds...)
	}
	if query.PostCode != "" {
		m = m.WhereLike(dao.SysPost.Columns().PostCode, "%"+query.PostCode+"%")
	}
	if query.PostName != "" {
		m = m.WhereLike(dao.SysPost.Columns().PostName, "%"+query.PostName+"%")
	}
	if len(deptIds) > 0 {
		m = m.WhereIn(dao.SysPost.Columns().DeptId, deptIds)
	}
	if query.Status != "" {
		m = m.Where(dao.SysPost.Columns().Status, query.Status)
	}
	if len(query.CreatedAt) == 2 {
		startTime := gtime.NewFromStr(query.CreatedAt[0])
		endTime := gtime.NewFromStr(query.CreatedAt[1])
		m = m.WhereBetween(dao.SysPost.Columns().CreatedAt, startTime, endTime.EndOfDay())
	}
	total, err = m.Count()
	if err != nil {
		return nil, 0, err
	}
	err = m.WithAll().Page(query.Page, query.PageSize).OrderAsc(dao.SysPost.Columns().PostSort).Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return
}

func (l *sSysPost) View(ctx context.Context, param *model.SysPostViewParam) (post *model.SysPostViewModel, err error) {
	err = l.Model(ctx).WithAll().Where(dao.SysPost.Columns().PostId, param.PostId).Scan(&post)
	if err != nil {
		return nil, err
	}
	return
}

func (l *sSysPost) Add(ctx context.Context, param *model.SysPostAddParam) (post *model.SysPostAddModel, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return nil, gerror.NewCode(consts.CodeLoginExpired, "登录已过期")
	}
	data := do.SysPost{}
	gconv.Struct(param, &data)
	data.TenantId = user.TenantId
	data.CreatedDept = user.DeptId
	data.CreatedBy = user.ID
	data.CreatedAt = gtime.Now()
	data.UpdatedBy = user.ID
	data.UpdatedAt = gtime.Now()
	postId, err := l.Model(ctx).Data(data).OmitNil().InsertAndGetId()
	if err != nil {
		return nil, err
	}
	post = &model.SysPostAddModel{
		PostId: postId,
	}
	return
}

func (l *sSysPost) Edit(ctx context.Context, param *model.SysPostEditParam) (post *model.SysPostEditModel, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return nil, gerror.NewCode(consts.CodeLoginExpired, "登录已过期")
	}
	data := do.SysPost{}
	gconv.Struct(param, &data)
	data.UpdatedBy = user.ID
	data.UpdatedAt = gtime.Now()
	_, err = l.Model(ctx).Data(data).OmitNil().Where(dao.SysPost.Columns().PostId, param.PostId).Update()
	if err != nil {
		return nil, err
	}
	post = &model.SysPostEditModel{
		PostId: param.PostId,
	}
	return
}

func (l *sSysPost) Delete(ctx context.Context, param *model.SysPostDeleteParam) (post *model.SysPostDeleteModel, err error) {
	user := contexts.GetUser(ctx)
	if user == nil {
		return nil, gerror.NewCode(consts.CodeLoginExpired, "登录已过期")
	}
	ids := make([]int64, 0)
	if param.PostId != 0 {
		ids = append(ids, param.PostId)
	}
	if len(param.PostIds) > 0 {
		ids = append(ids, param.PostIds...)
	}
	if len(ids) == 0 {
		return nil, gerror.New("请选择要删除的岗位")
	}
	data := do.SysPost{}
	data.DeletedBy = user.ID
	data.DeletedAt = gtime.Now()
	_, err = l.Model(ctx).Data(data).OmitNil().WhereIn(dao.SysPost.Columns().PostId, ids).Update()
	if err != nil {
		return nil, err
	}
	post = &model.SysPostDeleteModel{
		PostId: param.PostId,
	}
	return
}

func (l *sSysPost) Export(ctx context.Context, param *model.SysPostExportParam) (post *model.SysPostExportModel, err error) {
	if param.Page == 0 {
		param.Page = 1
	}
	if param.PageSize == 0 {
		param.PageSize = 1000
	}
	items, total, err := l.List(ctx, param.SysPostListParam)
	if err != nil {
		return
	}
	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(model.SysPostExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出CURD列表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", total, total+param.PageSize-1/param.PageSize, param.Page, len(items))
		exports   []model.SysPostExportModel
	)
	g.Log().Info(ctx, " sSysPost.Export total", total)
	if err = gconv.Scan(items, &exports); err != nil {
		return
	}
	g.Log().Info(ctx, " sSysPost.Export exports", len(exports))
	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	g.Log().Info(ctx, " sSysPost.Export err", err)
	post = &model.SysPostExportModel{}
	return
}
