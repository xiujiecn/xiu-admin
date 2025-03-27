// Package bcache
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package bcache

import (
	"context"
	"fmt"
	"time"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/library/cache"
)

// 设置系统认证token缓存
func SetSysAuthToken(ctx context.Context, userId int64, uuid string, token string, timeout time.Duration) (err error) {
	key := fmt.Sprintf(consts.KeySysAuthToken, userId, uuid)
	err = cache.Instance().Set(ctx, key, token, timeout)
	if err != nil {
		return err
	}
	return nil
}

// 获取系统认证token缓存
func GetSysAuthToken(ctx context.Context, userId int64, uuid string) (token string, err error) {
	key := fmt.Sprintf(consts.KeySysAuthToken, userId, uuid)
	t, err := cache.Instance().Get(ctx, key)
	if err != nil {
		return "", err
	}
	return t.String(), nil
}

// 删除系统认证token缓存
func DelSysAuthToken(ctx context.Context, userId int64, uuid string) (err error) {
	key := fmt.Sprintf(consts.KeySysAuthToken, userId, uuid)
	_, err = cache.Instance().Remove(ctx, key)
	if err != nil {
		return err
	}
	return nil
}

// 设置系统认证token拒绝缓存
func SetSysAuthTokenReject(ctx context.Context, userId int64, uuid string, token string, timeout time.Duration) (err error) {
	key := fmt.Sprintf(consts.KeySysAuthTokenReject, userId, uuid)
	err = cache.Instance().Set(ctx, key, token, timeout)
	if err != nil {
		return err
	}
	return nil
}

// 获取系统认证token拒绝缓存
func GetSysAuthTokenReject(ctx context.Context, userId int64, uuid string) (token string, err error) {
	key := fmt.Sprintf(consts.KeySysAuthTokenReject, userId, uuid)
	t, err := cache.Instance().Get(ctx, key)
	if err != nil {
		return "", err
	}
	return t.String(), nil
}

// 删除系统认证token拒绝缓存
func DelSysAuthTokenReject(ctx context.Context, userId int64, uuid string) (err error) {
	key := fmt.Sprintf(consts.KeySysAuthTokenReject, userId, uuid)
	_, err = cache.Instance().Remove(ctx, key)
	if err != nil {
		return err
	}
	return nil
}
