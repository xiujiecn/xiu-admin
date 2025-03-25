package xgen

import (
	"context"
	gendao "xiujieadmin/internal/library/xgen/gen_dao"
)

func Init(ctx context.Context) (err error) {
	err = gendao.Init(ctx)
	if err != nil {
		return
	}
	return
}
