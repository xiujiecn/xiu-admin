package mcache

import (
	"context"
	"fmt"
	"time"
	"xiujieadmin/internal/consts"
	"xiujieadmin/internal/service"
)

func GetDeptName(ctx context.Context, deptId int64) (string, error) {
	deptName, err := Instance().Get(ctx, fmt.Sprintf(consts.DeptName, deptId))
	if err != nil {
		dept, err := service.SysDept().GetDeptById(ctx, deptId)
		if err != nil {
			return "", err
		}
		Instance().Set(ctx, fmt.Sprintf(consts.DeptName, deptId), dept.DeptName, time.Hour*24)
		return dept.DeptName, nil
	}
	return deptName.String(), nil
}

func SetDeptName(ctx context.Context, deptId int64, deptName string) error {
	return Instance().Set(ctx, fmt.Sprintf(consts.DeptName, deptId), deptName, time.Hour*24)
}

func RemoveDeptName(ctx context.Context, deptId int64) error {
	_, err := Instance().Remove(ctx, fmt.Sprintf(consts.DeptName, deptId))
	return err
}
