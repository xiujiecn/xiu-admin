// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"xiuadmin/internal/model"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IMemoryDB interface {
		Init(ctx context.Context) error
		EventHandler(ctx context.Context, chg *model.BroadcastDbChgCache) error
		DB(ctx context.Context) gdb.DB
		Ctx(ctx context.Context) gdb.DB
	}
	IMemDBSysConfig interface {
		Init(ctx context.Context) error
		EventHandler(ctx context.Context) error
		TableName(ctx context.Context) string
		CreateTable(ctx context.Context) error
		LoadData(ctx context.Context, pk int64) error
		DeleteData(ctx context.Context, ids []int64) error
	}
	IMemDBSysDept interface {
		Init(ctx context.Context) error
		EventHandler(ctx context.Context) error
		TableName(ctx context.Context) string
		CreateTable(ctx context.Context) error
		LoadData(ctx context.Context, pk int64) error
		DeleteData(ctx context.Context, ids []int64) error
	}
	IMemDBSysDictData interface {
		Init(ctx context.Context) error
		EventHandler(ctx context.Context) error
		TableName(ctx context.Context) string
		CreateTable(ctx context.Context) error
		LoadData(ctx context.Context, pk int64) error
		DeleteData(ctx context.Context, ids []int64) error
	}
	IMemDBSysDictType interface {
		Init(ctx context.Context) error
		EventHandler(ctx context.Context) error
		TableName(ctx context.Context) string
		CreateTable(ctx context.Context) error
		LoadData(ctx context.Context, pk int64) error
		DeleteData(ctx context.Context, ids []int64) error
	}
	IMemDBSysMenu interface {
		Init(ctx context.Context) error
		EventHandler(ctx context.Context) error
		TableName(ctx context.Context) string
		CreateTable(ctx context.Context) error
		LoadData(ctx context.Context, pk int64) error
		DeleteData(ctx context.Context, ids []int64) error
	}
	IMemDBSysPost interface {
		Init(ctx context.Context) error
		EventHandler(ctx context.Context) error
		TableName(ctx context.Context) string
		CreateTable(ctx context.Context) error
		LoadData(ctx context.Context, pk int64) error
		DeleteData(ctx context.Context, ids []int64) error
	}
	IMemDBSysRole interface {
		Init(ctx context.Context) error
		EventHandler(ctx context.Context) error
		TableName(ctx context.Context) string
		CreateTable(ctx context.Context) error
		LoadData(ctx context.Context, pk int64) error
		DeleteData(ctx context.Context, ids []int64) error
	}
	IMemDBSysTenant interface {
		Init(ctx context.Context) error
		EventHandler(ctx context.Context) error
		DB(ctx context.Context) gdb.DB
		TableName(ctx context.Context) string
		CreateTable(ctx context.Context) error
		LoadData(ctx context.Context, pk int64) error
		DeleteData(ctx context.Context, ids []int64) error
	}
)

var (
	localMemoryDB         IMemoryDB
	localMemDBSysConfig   IMemDBSysConfig
	localMemDBSysDept     IMemDBSysDept
	localMemDBSysDictData IMemDBSysDictData
	localMemDBSysDictType IMemDBSysDictType
	localMemDBSysMenu     IMemDBSysMenu
	localMemDBSysPost     IMemDBSysPost
	localMemDBSysRole     IMemDBSysRole
	localMemDBSysTenant   IMemDBSysTenant
)

func MemoryDB() IMemoryDB {
	if localMemoryDB == nil {
		panic("implement not found for interface IMemoryDB, forgot register?")
	}
	return localMemoryDB
}

func RegisterMemoryDB(i IMemoryDB) {
	localMemoryDB = i
}

func MemDBSysConfig() IMemDBSysConfig {
	if localMemDBSysConfig == nil {
		panic("implement not found for interface IMemDBSysConfig, forgot register?")
	}
	return localMemDBSysConfig
}

func RegisterMemDBSysConfig(i IMemDBSysConfig) {
	localMemDBSysConfig = i
}

func MemDBSysDept() IMemDBSysDept {
	if localMemDBSysDept == nil {
		panic("implement not found for interface IMemDBSysDept, forgot register?")
	}
	return localMemDBSysDept
}

func RegisterMemDBSysDept(i IMemDBSysDept) {
	localMemDBSysDept = i
}

func MemDBSysDictData() IMemDBSysDictData {
	if localMemDBSysDictData == nil {
		panic("implement not found for interface IMemDBSysDictData, forgot register?")
	}
	return localMemDBSysDictData
}

func RegisterMemDBSysDictData(i IMemDBSysDictData) {
	localMemDBSysDictData = i
}

func MemDBSysDictType() IMemDBSysDictType {
	if localMemDBSysDictType == nil {
		panic("implement not found for interface IMemDBSysDictType, forgot register?")
	}
	return localMemDBSysDictType
}

func RegisterMemDBSysDictType(i IMemDBSysDictType) {
	localMemDBSysDictType = i
}

func MemDBSysMenu() IMemDBSysMenu {
	if localMemDBSysMenu == nil {
		panic("implement not found for interface IMemDBSysMenu, forgot register?")
	}
	return localMemDBSysMenu
}

func RegisterMemDBSysMenu(i IMemDBSysMenu) {
	localMemDBSysMenu = i
}

func MemDBSysPost() IMemDBSysPost {
	if localMemDBSysPost == nil {
		panic("implement not found for interface IMemDBSysPost, forgot register?")
	}
	return localMemDBSysPost
}

func RegisterMemDBSysPost(i IMemDBSysPost) {
	localMemDBSysPost = i
}

func MemDBSysRole() IMemDBSysRole {
	if localMemDBSysRole == nil {
		panic("implement not found for interface IMemDBSysRole, forgot register?")
	}
	return localMemDBSysRole
}

func RegisterMemDBSysRole(i IMemDBSysRole) {
	localMemDBSysRole = i
}

func MemDBSysTenant() IMemDBSysTenant {
	if localMemDBSysTenant == nil {
		panic("implement not found for interface IMemDBSysTenant, forgot register?")
	}
	return localMemDBSysTenant
}

func RegisterMemDBSysTenant(i IMemDBSysTenant) {
	localMemDBSysTenant = i
}
