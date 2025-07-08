// package system_test
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system_test

import (
	"context"
	"fmt"
	"testing"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/model"

	"xiuadmin/utility"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

func init() {
	// 数据库只读，不执行
	// g.DB().SetDryRun(true)
	g.DB().SetDebug(true)
	testing.Init()
}

func TestSysUser_UpdateUserPassword(t *testing.T) {
	// 设置上下文用户身份为用户
	ctx := context.WithValue(gctx.New(), consts.ContextKey, &model.Context{
		// 为了测试只设置了hook中需要用到的数据
		User: &model.Identity{
			BaseClaims: model.BaseClaims{
				ID:       1,
				TenantId: "000000",
				DeptId:   1,
			},
		},
	})
	newPassword := "123456"
	salt := "123456"
	password := utility.PasswordEncrypt(utility.Md5(newPassword), salt)
	_, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserId, 1).Data(map[string]any{
		dao.SysUser.Columns().Password:  password,
		dao.SysUser.Columns().Salt:      salt,
		dao.SysUser.Columns().UpdatedBy: 1,
		dao.SysUser.Columns().UpdatedAt: gtime.Now(),
	}).OmitEmpty().Update()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("UpdateUserPassword success")
}
