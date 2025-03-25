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

	fmt.Println("   _  __     __     ___        __            _      \n  | |/ /    / /    /   |  ____/ /____ ___   (_)____ \n  |   /__  / /    / /| | / __  // __ `__ \\ / // __ \\\n /   |/ /_/ /    / ___ |/ /_/ // / / / / // // / / /\n/_/|_|\\____/    /_/  |_|\\__,_//_/ /_/ /_//_//_/ /_/ ")
	fmt.Println("")

	fmt.Println("Version   ：", buildVersion)
	fmt.Println("BuildTime ：", buildTime)
	fmt.Println("CommitID  ：", commitID)
	fmt.Println("")
	fmt.Println("Copyright:", "XiuJieZhiLian Technology Co.,Ltd")
}

func AppName(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "system.appName", "XJAdmin").String()
}
