package bcache

import (
	"context"
	"fmt"
	"xiujieadmin/internal/library/cache"
)

func GetDeptName(ctx context.Context, deptId int64) (string, error) {
	deptName, err := cache.Instance().Get(ctx, fmt.Sprintf("dept_name_%d", deptId))
	if err != nil {
		return "", err
	}
	return deptName.String(), nil
}
