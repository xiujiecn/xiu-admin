// Package mcache
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package mcache

import (
	"context"
	"fmt"
	"time"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/library/event"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	event.EventsInstance().Register(consts.EventKeyUserLogout, func(ctx context.Context, eventKey string, args ...interface{}) {
		if len(args) == 0 {
			return
		}
		userId := args[0].(int64)
		RemoveUserAccessCodeList(ctx, userId)
		g.Log().Infof(ctx, "mcache_access_list.go EventKeyUserLogout userId: %d", userId)
	})

}

func GetUserAccessCodeList(ctx context.Context, userId int64) ([]string, []string, error) {
	accessCodeList, err := Instance().Get(ctx, fmt.Sprintf(consts.MemCacheUserAccessCodeList, userId))
	menuRoleDataAccessCodeList, err := Instance().Get(ctx, fmt.Sprintf(consts.MemCacheUserRoleDataAccessCodeList, userId))
	if err != nil || accessCodeList == nil {
		accessCodeCurrList, menuRoleDataAccessCodeList, err := service.SysAuth().GetUserAccessCodeList(ctx, userId)
		if err != nil {
			return nil, nil, err
		}
		if len(accessCodeCurrList) == 0 {
			accessCodeCurrList = []string{"null"}
		}
		Instance().Set(ctx, fmt.Sprintf(consts.MemCacheUserAccessCodeList, userId), accessCodeCurrList, time.Hour*24)
		Instance().Set(ctx, fmt.Sprintf(consts.MemCacheUserRoleDataAccessCodeList, userId), menuRoleDataAccessCodeList, time.Hour*24)
		return accessCodeCurrList, menuRoleDataAccessCodeList, nil
	}
	return accessCodeList.Strings(), menuRoleDataAccessCodeList.Strings(), nil
}

func SetUserAccessCodeList(ctx context.Context, userId int64, accessCodeList []string) error {
	return Instance().Set(ctx, fmt.Sprintf(consts.MemCacheUserAccessCodeList, userId), accessCodeList, time.Hour*24)
}

func SetUserRoleDataAccessCodeList(ctx context.Context, userId int64, roleDataAccessCodeList []string) error {
	return Instance().Set(ctx, fmt.Sprintf(consts.MemCacheUserRoleDataAccessCodeList, userId), roleDataAccessCodeList, time.Hour*24)
}

func RemoveUserRoleDataAccessCodeList(ctx context.Context, userId int64) error {
	_, err := Instance().Remove(ctx, fmt.Sprintf(consts.MemCacheUserRoleDataAccessCodeList, userId))
	return err
}

func RemoveUserAccessCodeList(ctx context.Context, userId int64) error {
	_, err := Instance().Remove(ctx, fmt.Sprintf(consts.MemCacheUserAccessCodeList, userId))
	return err
}
