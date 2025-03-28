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
)

func init() {
	event.EventsInstance().Register(consts.EventKeySysDeptUpdate, func(ctx context.Context, eventKey string, args ...interface{}) {
		if len(args) == 0 {
			return
		}
		deptId := args[0].(int64)
		RemoveDeptName(ctx, deptId)
	})
}

func GetDeptName(ctx context.Context, deptId int64) (string, error) {
	deptName, err := Instance().Get(ctx, fmt.Sprintf(consts.MemCacheDeptName, deptId))
	if err != nil {
		dept, err := service.SysDept().GetDeptById(ctx, deptId)
		if err != nil {
			return "", err
		}
		Instance().Set(ctx, fmt.Sprintf(consts.MemCacheDeptName, deptId), dept.DeptName, time.Hour*24)
		return dept.DeptName, nil
	}
	return deptName.String(), nil
}

func SetDeptName(ctx context.Context, deptId int64, deptName string) error {
	return Instance().Set(ctx, fmt.Sprintf(consts.MemCacheDeptName, deptId), deptName, time.Hour*24)
}

func RemoveDeptName(ctx context.Context, deptId int64) error {
	_, err := Instance().Remove(ctx, fmt.Sprintf(consts.MemCacheDeptName, deptId))
	return err
}
