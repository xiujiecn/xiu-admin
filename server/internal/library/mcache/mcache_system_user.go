// Package mcache
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package mcache

import (
	"context"
	"fmt"
	"time"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/library/event"
	"xiuadmin/internal/model"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	event.EventsInstance().Register(consts.EventKeyUserUpdate, func(ctx context.Context, eventKey string, args ...interface{}) {
		if len(args) == 0 {
			return
		}
		userId := args[0].(int64)
		RemoveUserInfo(ctx, userId)
	})
}

func GetUserInfo(ctx context.Context, userId int64) (*model.SysUserMiniModel, error) {
	userInfo, err := Instance().Get(ctx, fmt.Sprintf(consts.MemCacheUserInfo, userId))
	if err != nil {
		user, err := service.SysUser().GetUserById(ctx, userId)
		if err != nil {
			return nil, err
		}

		miniUser := &model.SysUserMiniModel{}
		gconv.Struct(miniUser, user)
		Instance().Set(ctx, fmt.Sprintf(consts.MemCacheUserInfo, userId), miniUser, time.Hour*24)
		return miniUser, nil
	}
	miniUser := &model.SysUserMiniModel{}
	err = userInfo.Struct(miniUser)
	if err != nil {
		return nil, err
	}
	return miniUser, nil
}

func SetUserInfo(ctx context.Context, userId int64, userInfo *model.SysUserMiniModel) error {
	return Instance().Set(ctx, fmt.Sprintf(consts.MemCacheUserInfo, userId), userInfo, time.Hour*24)
}

func RemoveUserInfo(ctx context.Context, userId int64) error {
	_, err := Instance().Remove(ctx, fmt.Sprintf(consts.MemCacheUserInfo, userId))
	return err
}
