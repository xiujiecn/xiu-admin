// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"encoding/json"
	"errors"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/do"
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/model/request"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysNotice struct {
}

func NewSysNotice() *sSysNotice {
	return &sSysNotice{}
}

func init() {
	service.RegisterSysNotice(NewSysNotice())
}

func (l *sSysNotice) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.SysNotice.Ctx(ctx), option...)
}

func (s *sSysNotice) List(ctx context.Context, param *model.SysNoticeListParam) (items []*model.SysNoticeListModel, total int, err error) {
	db := s.Model(ctx)

	if param.NoticeTitle != "" {
		db = db.WhereLike(dao.SysNotice.Columns().NoticeTitle, "%"+param.NoticeTitle+"%")
	}

	if param.NoticeType != "" {
		db = db.Where(dao.SysNotice.Columns().NoticeType, param.NoticeType)
	}

	if param.CreatedBy != "" {
		db = db.Where(dao.SysNotice.Columns().CreatedBy, param.CreatedBy)
	}

	if len(param.CreatedAt) == 2 {
		createdAt1 := gtime.NewFromStr(param.CreatedAt[0])
		createdAt2 := gtime.NewFromStr(param.CreatedAt[1])

		db = db.WhereBetween(dao.SysNotice.Columns().CreatedAt, createdAt1, createdAt2.EndOfDay())
	}

	db = db.OrderDesc(dao.SysNotice.Columns().NoticeId)

	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}

	err = db.Page(param.Page, param.PageSize).WithAll().Scan(&items)
	if err != nil {
		return nil, 0, err
	}

	return
}

func (s *sSysNotice) insertNoticeUserList(ctx context.Context, userIds *gset.IntSet, param *entity.SysNotice) (err error) {

	sysNotifyUserList := make([]*model.SysNoticeUserEditParam, 0)
	for _, userId := range userIds.Slice() {
		sysNotifyUserList = append(sysNotifyUserList, &model.SysNoticeUserEditParam{
			NoticeId:      param.NoticeId,
			NoticeTitle:   param.NoticeTitle,
			NoticeType:    param.NoticeType,
			NoticeContent: string(param.NoticeContent),
			Status:        param.Status,
			UserId:        int64(userId),
			IsRead:        0,
			Remark:        param.Remark,
			TenantId:      contexts.GetTenantId(ctx),
			CreatedDept:   contexts.GetDeptId(ctx),
			CreatedBy:     contexts.GetUserId(ctx),
			CreatedAt:     gtime.Now(),
			UpdatedBy:     contexts.GetUserId(ctx),
			UpdatedAt:     gtime.Now(),
		})
	}

	_, err = g.DB().Model(dao.SysNoticeUser.Table()).Data(sysNotifyUserList).OmitNil().Insert()
	if err != nil {
		return err
	}

	return
}

func (s *sSysNotice) Add(ctx context.Context, param *model.SysNoticeAddParam) (err error) {
	if param == nil || param.NoticeTitle == "" {
		return errors.New("公告标题不能为空")
	}
	if param.NoticeType == "" {
		return errors.New("公告类型不能为空")
	}
	if param.NoticeContent == "" {
		return errors.New("公告内容不能为空")
	}
	if param.NoticeRange == 0 {
		return errors.New("公告范围不能为空")
	}

	data := do.SysNotice{}
	gconv.Struct(param, &data)
	if param.NoticeRange == consts.NoticeRangeDept {
		if len(param.DeptIdList) == 0 {
			return errors.New("通知组织ID列表不能为空")
		}
		deptIds, err := json.Marshal(param.DeptIdList)
		if err != nil {
			return err
		}
		data.DeptIds = string(deptIds)
	}

	if param.NoticeRange == consts.NoticeRangeUser {
		if len(param.UserIdList) == 0 {
			return errors.New("通知用户ID列表不能为空")
		}
		userIds, err := json.Marshal(param.UserIdList)
		if err != nil {
			return err
		}
		data.UserIds = string(userIds)
	}

	data.TenantId = contexts.GetTenantId(ctx)
	data.CreatedDept = contexts.GetDeptId(ctx)
	data.CreatedBy = contexts.GetUserId(ctx)
	data.CreatedAt = gtime.Now()
	data.UpdatedBy = contexts.GetUserId(ctx)
	data.UpdatedAt = gtime.Now()
	g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		id, err := s.Model(ctx).Data(data).OmitEmpty().InsertAndGetId()
		if err != nil {
			return err
		}
		if id <= 0 {
			return errors.New("新增公告失败")
		}
		data.NoticeId = id

		// getUserIds
		userIds := gset.NewIntSet()
		if param.NoticeRange == consts.NoticeRangeUser {
			userIds.Add(gconv.Ints(param.UserIdList)...)
			g.Log().Debugf(ctx, "userIds: %+v", userIds.Slice())
		}
		if param.NoticeRange == consts.NoticeRangeDept {
			userList, _, err := service.SysUser().List(ctx, &request.PageInfo{
				Page:     1,
				PageSize: 10000,
			}, &model.UserListParam{
				DeptIdList: param.DeptIdList,
			})
			if err != nil {
				return err
			}
			for _, user := range userList {
				userIds.Add(int(user.UserId))
			}
		}
		if param.NoticeRange == consts.NoticeRangeAll {
			userList, _, err := service.SysUser().List(ctx, &request.PageInfo{
				Page:     1,
				PageSize: 10000,
			}, &model.UserListParam{})
			if err != nil {
				return err
			}
			for _, user := range userList {
				userIds.Add(int(user.UserId))
			}
		}

		if userIds.Size() <= 0 {
			return
		}

		entityData := entity.SysNotice{}
		gconv.Struct(data, &entityData)
		err = s.insertNoticeUserList(ctx, userIds, &entityData)
		if err != nil {
			return err
		}

		return
	})

	return
}

func (s *sSysNotice) Edit(ctx context.Context, param *model.SysNoticeEditParam) (err error) {
	if param == nil || param.NoticeId == 0 {
		return errors.New("公告ID不能为空")
	}
	if param.NoticeTitle == nil || *param.NoticeTitle == "" {
		return errors.New("公告标题不能为空")
	}
	if param.NoticeType == nil || *param.NoticeType == "" {
		return errors.New("公告类型不能为空")
	}
	if param.NoticeContent == nil || len(*param.NoticeContent) == 0 {
		return errors.New("公告内容不能为空")
	}

	data := do.SysNotice{}
	gconv.Struct(param, &data)
	if param.NoticeRange == consts.NoticeRangeDept {
		if len(param.DeptIdList) == 0 {
			return errors.New("通知组织ID列表不能为空")
		}
		deptIds, err := json.Marshal(param.DeptIdList)
		if err != nil {
			return err
		}
		data.DeptIds = string(deptIds)
	}

	if param.NoticeRange == consts.NoticeRangeUser {
		if len(param.UserIdList) == 0 {
			return errors.New("通知用户ID列表不能为空")
		}
		userIds, err := json.Marshal(param.UserIdList)
		if err != nil {
			return err
		}
		data.UserIds = string(userIds)
	}
	data.TenantId = contexts.GetTenantId(ctx)
	data.NoticeContent = []byte(*param.NoticeContent)
	data.UpdatedBy = contexts.GetUserId(ctx)
	data.UpdatedAt = gtime.Now()

	g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		m := s.Model(ctx)

		m = m.Where(dao.SysNotice.Columns().NoticeId, param.NoticeId)
		g.Log().Debugf(ctx, "param: %+v", string(*param.NoticeContent))
		g.Log().Debugf(ctx, "data: %+v", string(data.NoticeContent))
		_, err = m.Data(data).OmitEmpty().Update()
		if err != nil {
			return err
		}

		// 根据 notifyId删除，添加新的列表
		userIds := gset.NewIntSet()
		if param.NoticeRange == consts.NoticeRangeUser {
			userIds.Add(gconv.Ints(param.UserIdList)...)
		}
		if param.NoticeRange == consts.NoticeRangeDept || param.NoticeRange == consts.NoticeRangeAll {
			userList, _, err := service.SysUser().List(ctx, &request.PageInfo{
				Page:     1,
				PageSize: 10000,
			}, &model.UserListParam{
				DeptIdList: param.DeptIdList,
			})
			if err != nil {
				return err
			}
			for _, user := range userList {
				userIds.Add(int(user.UserId))
			}
		}

		// 删除
		_, err = g.DB().Model(dao.SysNoticeUser.Table()).Where(dao.SysNoticeUser.Columns().NoticeId, param.NoticeId).Delete()
		if err != nil {
			return err
		}

		if userIds.Size() <= 0 {
			return
		}
		entityData := entity.SysNotice{}
		gconv.Struct(data, &entityData)
		err = s.insertNoticeUserList(ctx, userIds, &entityData)
		if err != nil {
			return err
		}
		return
	})

	return
}

func (s *sSysNotice) Delete(ctx context.Context, param *model.SysNoticeDeleteParam) (err error) {
	g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = s.Model(ctx).WhereIn(dao.SysNotice.Columns().NoticeId, param.NoticeIds).Delete()
		if err != nil {
			return err
		}

		// 删除SysNoticeUser
		_, err = g.DB().Model(dao.SysNoticeUser.Table()).WhereIn(dao.SysNoticeUser.Columns().NoticeId, param.NoticeIds).Delete()
		if err != nil {
			return err
		}

		return
	})
	return
}

func (s *sSysNotice) View(ctx context.Context, param *model.SysNoticeViewParam) (data *model.SysNoticeViewModel, err error) {
	if param == nil || param.NoticeId == 0 {
		return nil, errors.New("公告ID不能为空")
	}
	data = &model.SysNoticeViewModel{}
	err = s.Model(ctx).Where(dao.SysNotice.Columns().NoticeId, param.NoticeId).Scan(&data)
	g.Log().Debugf(ctx, "data: %+v", string(data.NoticeContent))
	if err != nil {
		return nil, err
	}
	if data.UserIds != "" {
		json.Unmarshal([]byte(data.UserIds), &data.UserIdList)
	}
	if data.DeptIds != "" {
		json.Unmarshal([]byte(data.DeptIds), &data.DeptIdList)
	}
	return
}
