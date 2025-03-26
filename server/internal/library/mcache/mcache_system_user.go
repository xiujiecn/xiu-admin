package mcache

import (
	"context"
	"fmt"
	"time"
	"xiujieadmin/internal/consts"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	RegisterUserChangeClearCache(consts.UserInfo, RemoveUserInfo)
}

func GetUserInfo(ctx context.Context, userId int64) (*model.SysUserMiniModel, error) {
	userInfo, err := Instance().Get(ctx, fmt.Sprintf(consts.UserInfo, userId))
	if err != nil {
		user, err := service.SysUser().GetUserById(ctx, userId)
		if err != nil {
			return nil, err
		}

		miniUser := &model.SysUserMiniModel{}
		gconv.Struct(miniUser, user)
		Instance().Set(ctx, fmt.Sprintf(consts.UserInfo, userId), miniUser, time.Hour*24)
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
	return Instance().Set(ctx, fmt.Sprintf(consts.UserInfo, userId), userInfo, time.Hour*24)
}

func RemoveUserInfo(ctx context.Context, userId int64) error {
	_, err := Instance().Remove(ctx, fmt.Sprintf(consts.UserInfo, userId))
	return err
}
