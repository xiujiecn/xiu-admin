// package model
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package model

type GoRunInfo struct {
	GoName    string `json:"goName"`    //语言名
	GoOs      string `json:"goOs"`      //操作系统
	Arch      string `json:"arch"`      //系统架构
	GoVersion string `json:"goVersion"` //GO版本
	StartTime string `json:"startTime"` //系统开始时间
	RunTime   int64  `json:"runTime"`   //运行时长
	RootPath  string `json:"rootPath"`  //根路径
	Pwd       string `json:"pwd"`       //当前路径
	Goroutine string `json:"goroutine"` //协程数
	GoMem     string `json:"goMem"`     //内存使用
	GoSize    string `json:"goSize"`    //程序大小
}

type CpuInfo struct {
	Number      int       //cup个数
	Cores       int32     //核数
	UsedPercent []float64 //cpu使用率
	ModelName   string
}

// NetWorkInfo 网速信息
type NetWorkInfo struct {
	Name         string
	Receive      uint64
	Sent         uint64
	ReceiveSpeed uint64
	SentSpeed    uint64
}
