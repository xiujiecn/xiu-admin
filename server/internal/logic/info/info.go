package info

import (
	"context"
	"runtime"
	"xiuadmin/internal/model"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

type sInfo struct {
	Info model.InfoModel
}

func NewInfo() *sInfo {
	return &sInfo{
		Info: model.InfoModel{
			Data: make(map[string]interface{}),
		},
	}
}

func init() {
	service.RegisterInfo(NewInfo())
}
func toLogLevelStr(level int) string {
	if level == glog.LEVEL_NONE {
		return "NONE"
	}
	if level == glog.LEVEL_ALL {
		return "ALL"
	}
	if level == glog.LEVEL_DEV {
		return "DEV"
	}
	if level == glog.LEVEL_PROD {
		return "PROD"
	}
	str := ""
	if level&glog.LEVEL_DEBU != 0 {
		str += "DEBUG "
	}
	if level&glog.LEVEL_INFO != 0 {
		str += "INFO "
	}
	if level&glog.LEVEL_NOTI != 0 {
		str += "NOTI "
	}
	if level&glog.LEVEL_WARN != 0 {
		str += "WARN "
	}
	if level&glog.LEVEL_ERRO != 0 {
		str += "ERRO "
	}
	if level&glog.LEVEL_CRIT != 0 {
		str += "CRIT "
	}
	if level&glog.LEVEL_PANI != 0 {
		str += "PANI "
	}
	if level&glog.LEVEL_FATA != 0 {
		str += "FATA "
	}
	return str
}
func (s *sInfo) GetInfo() *model.InfoModel {
	ctx := context.Background()
	logLevel := g.Cfg().MustGet(ctx, "logger.level", "").String()
	if logLevel != "" {
		g.Log().SetLevelStr(logLevel)
	}
	s.AddInfoData("log_level_conf", logLevel)
	s.AddInfoData("log_level_real", toLogLevelStr(g.Log().GetLevel()))
	s.AddInfoData("Goroutines", runtime.NumGoroutine())
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	s.AddInfoData("MemAllocKB", m.Alloc/1024)
	s.AddInfoData("MemTotalAllocKB", m.TotalAlloc/1024)
	s.AddInfoData("MemSysKB", m.Sys/1024)
	s.AddInfoData("MemHeapAllocKB", m.HeapAlloc/1024)
	s.AddInfoData("MemHeapSysKB", m.HeapSys/1024)
	s.AddInfoData("MemHeapIdleKB", m.HeapIdle/1024)
	s.AddInfoData("MemHeapInuseKB", m.HeapInuse/1024)
	s.AddInfoData("MemHeapReleasedKB", m.HeapReleased/1024)
	s.AddInfoData("MemHeapObjects", m.HeapObjects)
	s.AddInfoData("MemStackInuseKB", m.StackInuse/1024)
	s.AddInfoData("MemStackSysKB", m.StackSys/1024)
	s.AddInfoData("MemNumGC", m.NumGC)
	return &s.Info
}

func (s *sInfo) SetInfoName(Type, Name, Title, Description string) {
	if s.Info.Name == "" {
		s.Info.Type = Type
		s.Info.Name = Name
		s.Info.Title = Title
		s.Info.Description = Description
	}
}

func (s *sInfo) SetInfoBuild(buildVersion, buildTime, commitID string) {
	s.Info.Version = buildVersion
	s.Info.Build = buildTime
	s.Info.Commit = commitID
}

func (s *sInfo) SetInfoData(data map[string]interface{}) {
	s.Info.Data = data
}

func (s *sInfo) GetInfoData() map[string]interface{} {
	return s.Info.Data
}

func (s *sInfo) AddInfoData(key string, value interface{}) {
	if s.Info.Data == nil {
		s.Info.Data = make(map[string]interface{})
	}
	s.Info.Data[key] = value
}

func (s *sInfo) DelInfoData(key string) {
	delete(s.Info.Data, key)
}

func (s *sInfo) GetInfoDataByKey(key string) interface{} {
	if s.Info.Data == nil {
		return nil
	}
	if _, ok := s.Info.Data[key]; !ok {
		return nil
	}
	return s.Info.Data[key]
}
