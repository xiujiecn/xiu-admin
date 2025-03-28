// package monitor
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package monitor

import (
	"context"
	"encoding/json"
	"os"
	"runtime"
	"strconv"
	"time"
	"xiuadmin/internal/model"
	"xiuadmin/internal/service"
	"xiuadmin/utility"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type sMonitorServer struct {
}

func NewMonitorServer() *sMonitorServer {
	return &sMonitorServer{}
}

var SysStartTime *gtime.Time
var GoDiskSize string
var SysRunDir string
var LocalIP string
var PublicIP string

func init() {
	service.RegisterMonitorServer(NewMonitorServer())
	SysStartTime = gtime.Now()
	SysRunDir, _ = os.Getwd()
	GoDiskSize = utility.DirSize(SysRunDir)
	LocalIP, _ = utility.GetLocalIP()
	PublicIP, _ = utility.GetPublicIP()
}

func (l *sMonitorServer) GetGoInfo(ctx context.Context) (res *model.GoRunInfo) {
	SysRunDir, _ := os.Getwd()
	var gm runtime.MemStats
	runtime.ReadMemStats(&gm)
	res = &model.GoRunInfo{
		GoName:    "Golang",
		GoOs:      runtime.GOOS,
		Arch:      runtime.GOARCH,
		GoVersion: runtime.Version(),
		StartTime: SysStartTime.Format("2006-01-02 15:04:05"),
		RunTime:   gtime.Now().Timestamp() - SysStartTime.Timestamp(),
		RootPath:  runtime.GOROOT(),
		Pwd:       SysRunDir,
		Goroutine: strconv.Itoa(runtime.NumGoroutine()),
		GoMem:     utility.FileSize(int64(gm.Sys)),
		GoSize:    GoDiskSize,
	}
	return
}

// GetHostInfo 获取主机信息
func (l *sMonitorServer) GetHostInfo(ctx context.Context) (data []byte) {
	hostInfo, _ := host.Info()
	timestamp, _ := host.BootTime()
	t := time.Unix(int64(timestamp), 0)
	tmpData := gconv.Map(hostInfo)
	tmpData["bootTime"] = t

	tmpData["intranet_ip"] = LocalIP
	tmpData["public_ip"] = PublicIP
	data, _ = gjson.Encode(tmpData)
	return
}

// GetSysLoad 获取系统负载信息
func (l *sMonitorServer) GetSysLoad(ctx context.Context) (data []byte) {
	loadInfo, _ := load.Avg()
	data, _ = gjson.Encode(g.Map{
		"load1":  loadInfo.Load1,
		"load5":  loadInfo.Load5,
		"load15": loadInfo.Load15,
	})
	return
}

// GetCpuInfo 获取CPU信息
func (l *sMonitorServer) GetCpuInfo(ctx context.Context) (data []byte) {
	var CpuInfoData model.CpuInfo
	cpus, _ := cpu.Info()
	for _, c := range cpus {
		CpuInfoData.Cores = CpuInfoData.Cores + c.Cores
	}
	CpuInfoData.Number = len(cpus)
	percent, _ := cpu.Percent(time.Second, false) //获取CPU使用率
	CpuInfoData.UsedPercent = percent
	CpuInfoData.ModelName = cpus[0].ModelName //CPU型号
	data, _ = json.Marshal(CpuInfoData)
	return
}

// GetMemInfo 获取内存信息
func (l *sMonitorServer) GetMemInfo(ctx context.Context) (data []byte) {
	hostInfo, _ := mem.VirtualMemory()
	tmpData := gconv.Map(hostInfo)
	var gomem runtime.MemStats
	runtime.ReadMemStats(&gomem)
	if tmpData == nil {
		tmpData = make(map[string]interface{})
	}
	tmpData["goUsed"] = gomem.Sys
	data, _ = gjson.Encode(tmpData)
	return
}

// GetDiskInfo 获取磁盘信息
func (l *sMonitorServer) GetDiskInfo(ctx context.Context) (data []byte) {
	diskUsed, _ := disk.Usage("/")
	data, _ = gjson.Encode(diskUsed)
	return
}

// GetNetStatusInfo 获取网络信息
func (l *sMonitorServer) GetNetStatusInfo(ctx context.Context) (data []byte) {
	IOCountersStat, _ := net.IOCounters(true)
	netWorkInfo := make([]model.NetWorkInfo, len(IOCountersStat))
	for i, n := range IOCountersStat {
		netWorkInfo[i].Name = n.Name
		netWorkInfo[i].Receive = n.BytesRecv
		netWorkInfo[i].Sent = n.BytesSent
		if netWorkInfo != nil && len(netWorkInfo) > i {
			netWorkInfo[i].ReceiveSpeed = n.BytesRecv - netWorkInfo[i].Receive
			netWorkInfo[i].SentSpeed = n.BytesSent - netWorkInfo[i].Sent
		}
	}
	data, _ = gjson.Encode(netWorkInfo)
	return
}
