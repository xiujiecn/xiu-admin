package consts

// 服务器监控信息
const (
	MonitorServerHost    = "host"
	MonitorServerSysLoad = "sysLoad"
	MonitorServerCpu     = "cpu"
	MonitorServerMem     = "mem"
	MonitorServerDisk    = "disk"
	MonitorServerNet     = "net"
	MonitorServerGo      = "go"
)

// WEBSOCKET订阅Tag  服务器监控信息
const (
	WSTagMonitorServer   = "ws:tag:monitor:server"      // WEBSOCKET订阅Tag  服务器监控信息
	WSEventMonitorServer = "ws:event:monitor:server:%s" // WEBSOCKET订阅 服务器监控信息  %s in [host,sysLoad,cpu,disk,net,go]
)
