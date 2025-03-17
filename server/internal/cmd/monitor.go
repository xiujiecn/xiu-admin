package cmd

import (
	"context"
	"fmt"
	"time"
	"xiujieadmin/internal/consts"
	"xiujieadmin/internal/library/websocket"
	"xiujieadmin/internal/service"
)

func InitMonitor() {
	go func() {
		for {
			RunMonitor()
			time.Sleep(time.Second * 5)
		}
	}()
	websocket.RegisterTagCallback(consts.WSTagMonitorServer, func(client *websocket.Client) {
		RunMonitor()
	})
}

func RunMonitor() {
	ctx := context.Background()
	hostInfo := service.MonitorServer().GetHostInfo(ctx)
	websocket.SendToTag(consts.WSTagMonitorServer, &websocket.WResponse{
		Event: fmt.Sprintf(consts.WSEventMonitorServer, consts.MonitorServerHost),
		Data:  hostInfo,
	})
	sysLoad := service.MonitorServer().GetSysLoad(ctx)
	websocket.SendToTag(consts.WSTagMonitorServer, &websocket.WResponse{
		Event: fmt.Sprintf(consts.WSEventMonitorServer, consts.MonitorServerSysLoad),
		Data:  sysLoad,
	})
	cpuInfo := service.MonitorServer().GetCpuInfo(ctx)
	websocket.SendToTag(consts.WSTagMonitorServer, &websocket.WResponse{
		Event: fmt.Sprintf(consts.WSEventMonitorServer, consts.MonitorServerCpu),
		Data:  cpuInfo,
	})
	memInfo := service.MonitorServer().GetMemInfo(ctx)
	websocket.SendToTag(consts.WSTagMonitorServer, &websocket.WResponse{
		Event: fmt.Sprintf(consts.WSEventMonitorServer, consts.MonitorServerMem),
		Data:  memInfo,
	})
	diskInfo := service.MonitorServer().GetDiskInfo(ctx)
	websocket.SendToTag(consts.WSTagMonitorServer, &websocket.WResponse{
		Event: fmt.Sprintf(consts.WSEventMonitorServer, consts.MonitorServerDisk),
		Data:  diskInfo,
	})
	netInfo := service.MonitorServer().GetNetStatusInfo(ctx)
	websocket.SendToTag(consts.WSTagMonitorServer, &websocket.WResponse{
		Event: fmt.Sprintf(consts.WSEventMonitorServer, consts.MonitorServerNet),
		Data:  netInfo,
	})
	goInfo := service.MonitorServer().GetGoInfo(ctx)
	websocket.SendToTag(consts.WSTagMonitorServer, &websocket.WResponse{
		Event: fmt.Sprintf(consts.WSEventMonitorServer, consts.MonitorServerGo),
		Data:  goInfo,
	})
}
