// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"xiuadmin/internal/model"
)

type (
	IInfo interface {
		GetInfo() *model.InfoModel
		SetInfoName(Type string, Name string, Title string, Description string)
		SetInfoBuild(buildVersion string, buildTime string, commitID string)
		SetInfoData(data map[string]interface{})
		GetInfoData() map[string]interface{}
		AddInfoData(key string, value interface{})
		DelInfoData(key string)
		GetInfoDataByKey(key string) interface{}
	}
)

var (
	localInfo IInfo
)

func Info() IInfo {
	if localInfo == nil {
		panic("implement not found for interface IInfo, forgot register?")
	}
	return localInfo
}

func RegisterInfo(i IInfo) {
	localInfo = i
}
