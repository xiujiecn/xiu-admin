package gendao

import (
	"context"
	genmodel "xiujieadmin/internal/library/xgen/gen_model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	daoConfigList           = make([]*genmodel.GenDaoConfig, 0)
	cliConfigPath           = `hack/config.yaml`
	cliConfig               *gcfg.Config
	defaultGenServiceConfig = genmodel.GenServiceConfig{
		SrcFolder:       "internal/logic",
		DstFolder:       "internal/service",
		DstFileNameCase: "Snake",
		StPattern:       `s([A-Z]\w+)`,
		Clear:           false,
	}
	serviceConfig = genmodel.GenServiceConfig{}
)

func GetDaoConfig(ctx context.Context, group string) (output *genmodel.GenDaoConfig, err error) {
	for _, v := range daoConfigList {
		if v == nil {
			continue
		}
		if v.Group == group {
			output = v
			return
		}
	}
	return
}

func Init(ctx context.Context) (err error) {
	adapter, err := gcfg.NewAdapterFile(cliConfigPath)
	if err != nil {
		g.Log().Fatalf(ctx, "get cli configuration file:%v, err:%+v", cliConfigPath, err)
		return
	}
	cliConfig = gcfg.NewWithAdapter(adapter)

	cliConfig.MustGet(ctx, "gfcli.gen.dao").Scan(&daoConfigList)
	cliConfig.MustGet(ctx, "gfcli.gen.service").Scan(&serviceConfig)

	return
}

func GetServiceConfig() *genmodel.GenServiceConfig {
	conf := &genmodel.GenServiceConfig{}
	_ = gconv.Scan(defaultGenServiceConfig, &conf)
	_ = gconv.Scan(serviceConfig, &conf)
	return conf
}
