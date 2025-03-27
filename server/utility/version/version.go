// package version
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package version

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	BuildTime    string
	BuildVersion string
	CommitID     string
)

func ShowLogo(buildVersion, buildTime, commitID string) {
	BuildVersion = buildVersion
	BuildTime = buildTime
	CommitID = commitID
	//版本号

	fmt.Println(" __  __ _             _        _             _        \n \\ \\/ /(_) _   _     / \\    __| | _ __ ___  (_) _ __  \n  \\  / | || | | |   / _ \\  / _` || '_ ` _ \\ | || '_ \\ \n  /  \\ | || |_| |  / ___ \\| (_| || | | | | || || | | |\n /_/\\_\\|_| \\__,_| /_/   \\_\\\\__,_||_| |_| |_||_||_| |_|")
	fmt.Println("")

	fmt.Println("Version   ：", buildVersion)
	fmt.Println("BuildTime ：", buildTime)
	fmt.Println("CommitID  ：", commitID)
	fmt.Println("")
	fmt.Println("Copyright:", "XiuAdmin Technology Co.,Ltd")
}

func AppName(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "system.appName", "XJAdmin").String()
}
