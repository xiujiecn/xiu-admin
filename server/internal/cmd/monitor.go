// Package cmd
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package cmd

import (
	"context"
	"fmt"
	"time"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/library/websocket"
	"xiuadmin/internal/service"
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
