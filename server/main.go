package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	_ "xiujieadmin/internal/packed"

	_ "xiujieadmin/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"xiujieadmin/internal/cmd"
	"xiujieadmin/utility/version"
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
	cmd.Main.Run(ctx)
}
