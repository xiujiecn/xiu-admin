// package main
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	_ "xiuadmin/internal/packed"

	_ "xiuadmin/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"xiuadmin/internal/cmd"
	"xiuadmin/utility/version"
)

var (
	BuildVersion = "0.0"
	BuildTime    = ""
	CommitID     = ""
)

func main() {
	version.ShowLogo(BuildVersion, BuildTime, CommitID)
	ctx := gctx.GetInitCtx()

	cmd.SystemInit(ctx)
	err := cmd.InitSystemDeferFunc(ctx)
	if err != nil {
		panic(err)
	}
	cmd.Main.Run(ctx)
}
