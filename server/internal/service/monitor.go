// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"xiujieadmin/internal/model"
)

type (
	IMonitorServer interface {
		GetGoInfo(ctx context.Context) (res *model.GoRunInfo)
		// GetHostInfo 获取主机信息
		GetHostInfo(ctx context.Context) (data []byte)
		// GetSysLoad 获取系统负载信息
		GetSysLoad(ctx context.Context) (data []byte)
		// GetCpuInfo 获取CPU信息
		GetCpuInfo(ctx context.Context) (data []byte)
		// GetMemInfo 获取内存信息
		GetMemInfo(ctx context.Context) (data []byte)
		// GetDiskInfo 获取磁盘信息
		GetDiskInfo(ctx context.Context) (data []byte)
		// GetNetStatusInfo 获取网络信息
		GetNetStatusInfo(ctx context.Context) (data []byte)
	}
)

var (
	localMonitorServer IMonitorServer
)

func MonitorServer() IMonitorServer {
	if localMonitorServer == nil {
		panic("implement not found for interface IMonitorServer, forgot register?")
	}
	return localMonitorServer
}

func RegisterMonitorServer(i IMonitorServer) {
	localMonitorServer = i
}
