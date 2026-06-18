// package xgen
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package xgen

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"xiuadmin/internal/library/addons"
	genconsts "xiuadmin/internal/library/xgen/gen_consts"
	gendao "xiuadmin/internal/library/xgen/gen_dao"
	genmodel "xiuadmin/internal/library/xgen/gen_model"
	genview "xiuadmin/internal/library/xgen/gen_view"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/request"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

func GenTypeSelects(ctx context.Context) (output []*genmodel.SelectGenTypeModel, err error) {
	output = make([]*genmodel.SelectGenTypeModel, 0)
	for k, v := range genconsts.GenCodesTypeNameMap {
		item := &genmodel.SelectGenTypeModel{
			SelectItemModel: genmodel.SelectItemModel{
				Value:    k,
				Name:     v,
				Label:    v,
				Disabled: false,
			},
			Templates: make([]*genmodel.SelectGenTemplateModel, 0),
		}
		confName, ok := genconsts.GenCodesTypeConfMap[k]
		if !ok {
			continue
		}
		confTemplates := make([]*genmodel.GenCodesConfigCrudTemplate, 0)
		err = g.Cfg().MustGet(ctx, "xgen.application."+confName+".templates").Scan(&confTemplates)
		if err != nil {
			return
		}
		if len(confTemplates) == 0 && confName != "crud" {
			err = g.Cfg().MustGet(ctx, "xgen.application.crud.templates").Scan(&confTemplates)
			if err != nil {
				return
			}
		}

		for index, v := range confTemplates {
			item.Templates = append(item.Templates, &genmodel.SelectGenTemplateModel{
				SelectItemModel: genmodel.SelectItemModel{
					Value:    index,
					Label:    v.Group,
					Name:     v.Group,
					Disabled: false,
				},
				IsAddon: v.IsAddon,
			})
		}
		sort.Slice(item.Templates, func(i, j int) bool {
			return item.Templates[i].Value.(int) < item.Templates[j].Value.(int)
		})

		output = append(output, item)
	}
	return
}
func DbSelects(ctx context.Context) (output []*genmodel.SelectItemModel, err error) {
	output = make([]*genmodel.SelectItemModel, 0)
	dbs := g.Cfg().MustGet(ctx, "xgen.selectDbs")
	if len(dbs.Strings()) == 0 {
		return
	}

	for _, v := range dbs.Strings() {
		output = append(output, &genmodel.SelectItemModel{
			Value:    v,
			Label:    v,
			Name:     v,
			Disabled: false,
		})
	}
	return
}

func StatusSelects(ctx context.Context) (output []*genmodel.SelectItemModel, err error) {
	output = make([]*genmodel.SelectItemModel, 0)
	for k, v := range genconsts.GenCodesStatusNameMap {
		output = append(output, &genmodel.SelectItemModel{
			Value:    k,
			Label:    v,
			Name:     v,
			Disabled: false,
		})
	}
	sort.Slice(output, func(i, j int) bool {
		return output[i].Value.(string) < output[j].Value.(string)
	})
	return
}

func LinkModeSelects(ctx context.Context) (output []*genmodel.SelectItemModel, err error) {
	output = make([]*genmodel.SelectItemModel, 0)
	for k, v := range genconsts.GenCodesJoinNameMap {
		output = append(output, &genmodel.SelectItemModel{
			Value:    k,
			Label:    v,
			Name:     v,
			Disabled: false,
		})
	}
	sort.Slice(output, func(i, j int) bool {
		return output[i].Value.(int) < output[j].Value.(int)
	})
	return
}

func BuildMethodSelects(ctx context.Context) (output []*genmodel.SelectItemModel, err error) {
	output = make([]*genmodel.SelectItemModel, 0)
	for k, v := range genconsts.GenCodesBuildMethNameMap {
		output = append(output, &genmodel.SelectItemModel{
			Value:    k,
			Label:    v,
			Name:     v,
			Disabled: false,
		})
	}
	sort.Slice(output, func(i, j int) bool {
		return output[i].Value.(int) < output[j].Value.(int)
	})
	return
}
func FormModeSelects(ctx context.Context) (output []*genmodel.SelectItemModel, err error) {
	output = make([]*genmodel.SelectItemModel, 0)
	for _, v := range genview.FMs {
		output = append(output, &genmodel.SelectItemModel{
			Value:    v,
			Label:    genview.FMMap[v],
			Name:     genview.FMMap[v],
			Disabled: false,
		})
	}
	return
}
func FormRoleSelects(ctx context.Context) (output []*genmodel.SelectItemModel, err error) {
	output = make([]*genmodel.SelectItemModel, 0)
	for k, v := range genview.FRMap {
		output = append(output, &genmodel.SelectItemModel{
			Value:    k,
			Label:    v,
			Name:     v,
			Disabled: false,
		})
	}
	return
}

func DictModeSelects(ctx context.Context) (output []*genmodel.SelectItemModel, err error) {
	dictMode, _, err := service.SysDictType().List(ctx, &model.SysDictTypeListParam{
		PageInfo: request.PageInfo{
			Page:     1,
			PageSize: 2000,
		},
	})
	if err != nil {
		return
	}
	for _, v := range dictMode {
		output = append(output, &genmodel.SelectItemModel{
			Value: v.DictType,
			Label: v.DictName + "(" + v.DictType + ")",
			Name:  v.DictName,
		})
	}
	return
}

func WhereModeSelects(ctx context.Context) (output []*genmodel.SelectItemModel, err error) {
	output = make([]*genmodel.SelectItemModel, 0)
	for _, v := range genview.WMs {
		output = append(output, &genmodel.SelectItemModel{
			Value:    v,
			Label:    v,
			Name:     v,
			Disabled: false,
		})
	}
	return
}

func TableAlignSelects(ctx context.Context) (output []*genmodel.SelectItemModel, err error) {
	output = make([]*genmodel.SelectItemModel, 0)
	for _, v := range genview.TableAligns {
		output = append(output, &genmodel.SelectItemModel{
			Value:    v,
			Label:    genview.TableAlignMap[v],
			Name:     genview.TableAlignMap[v],
			Disabled: false,
		})
	}
	return
}

func TreeStyleTypeSelects(ctx context.Context) (output []*genmodel.SelectItemModel, err error) {
	output = make([]*genmodel.SelectItemModel, 0)
	for k, v := range genconsts.GenCodesTreeStyleTypeNameMap {
		output = append(output, &genmodel.SelectItemModel{
			Value:    k,
			Label:    v,
			Name:     v,
			Disabled: false,
		})
	}
	sort.Slice(output, func(i, j int) bool {
		return output[i].Value.(int) < output[j].Value.(int)
	})
	return
}

func GenCodesSelects(ctx context.Context) (output *genmodel.SelectsModel, err error) {
	output = &genmodel.SelectsModel{}
	output.GenType, err = GenTypeSelects(ctx)
	if err != nil {
		return
	}
	output.Db, err = DbSelects(ctx)
	if err != nil {
		return
	}
	output.Status, err = StatusSelects(ctx)
	if err != nil {
		return
	}
	output.LinkMode, err = LinkModeSelects(ctx)
	if err != nil {
		return
	}
	output.BuildMethod, err = BuildMethodSelects(ctx)
	if err != nil {
		return
	}
	output.FormMode, err = FormModeSelects(ctx)
	if err != nil {
		return
	}
	output.WhereMode, err = WhereModeSelects(ctx)
	if err != nil {
		return
	}
	output.FormRole, err = FormRoleSelects(ctx)
	if err != nil {
		return
	}
	output.DictMode, err = DictModeSelects(ctx)
	if err != nil {
		return
	}
	output.Addons, err = addons.ModuleSelects(ctx)
	if err != nil {
		return
	}
	output.TableAlign, err = TableAlignSelects(ctx)
	if err != nil {
		return
	}
	output.TreeStyleType, err = TreeStyleTypeSelects(ctx)
	if err != nil {
		return
	}
	return
}

func GenCodesTableSelect(ctx context.Context, params *genmodel.GenCodesTableSelectParam) (output []*genmodel.GenCodesTableSelectModel, err error) {
	if params == nil || params.DbGroup == "" {
		return
	}
	sql := "SELECT TABLE_NAME as value, TABLE_COMMENT as label FROM information_schema.`TABLES` WHERE TABLE_SCHEMA = ?"
	config := g.DB(params.DbGroup).GetConfig()
	disableTables := g.Cfg().MustGet(ctx, "xgen.disableTables").Strings()

	lists := make([]*genmodel.GenCodesTableSelectModel, 0)
	err = g.DB(params.DbGroup).Ctx(ctx).Raw(sql, config.Name).Scan(&lists)
	if err != nil {
		return
	}
	patternStr := `addon_(\w+)_`
	repStr := ``
	for _, v := range lists {
		if gstr.InArray(disableTables, v.Value) {
			continue
		}
		newValue := v.Value
		if config.Prefix != "" {
			newValue = gstr.SubStrFromEx(v.Value, config.Prefix)
		}
		// g.Log().Infof(ctx, "server/internal/library/xgen/xgen_curd.go GenCodesTableSelect newValue:%v config.Prefix:%v", newValue, config.Prefix)
		if newValue == "" {
			continue
		}
		// 插件移除掉插件表前缀
		bt, err := gregex.Replace(patternStr, []byte(repStr), []byte(newValue))
		if err != nil {
			continue
		}
		// g.Log().Infof(ctx, "server/internal/library/xgen/xgen_curd.go GenCodesTableSelect newValue:%v string(bt):%v", newValue, string(bt))

		row := v
		row.DefTableComment = v.Label
		row.DaoName = gstr.CaseCamel(newValue)
		row.DefVarName = gstr.CaseCamel(string(bt))
		row.DefAlias = gstr.CaseCamelLower(newValue)
		row.Name = fmt.Sprintf("%s (%s)", v.Value, v.Label)
		row.Label = row.Name
		output = append(output, row)
	}
	return
}

func GenCodesColumnSelect(ctx context.Context, params *genmodel.GenCodesColumnSelectParam) (output []*genmodel.GenCodesColumnSelectModel, err error) {
	sql := "SELECT COLUMN_NAME as value, COLUMN_COMMENT as label FROM information_schema.`COLUMNS` WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?"
	dbConfig := g.DB(params.DbGroup).GetConfig()

	lists := make([]*genmodel.GenCodesColumnSelectModel, 0)
	err = g.DB(params.DbGroup).Ctx(ctx).Raw(sql, dbConfig.Name, params.TableName).Scan(&lists)
	if err != nil {
		return
	}
	if len(lists) == 0 {
		err = gerror.New("表不存在")
		return
	}
	for _, v := range lists {
		output = append(output, &genmodel.GenCodesColumnSelectModel{
			Value: v.Value,
			Label: v.Label,
			Name:  fmt.Sprintf("%s (%s)", v.Value, v.Label),
		})
	}
	return
}

func GenCodesColumnList(ctx context.Context, params *genmodel.GenCodesColumnListParam) (output []*genmodel.GenCodesColumnListModel, err error) {
	output, err = genview.TableColumnList(ctx, params)
	return
}

// GetLoadGenerate 获取本地生成配置
func GetLoadGenerate(ctx context.Context) (conf *genmodel.GenCodesConfig, err error) {
	err = g.Cfg().MustGet(ctx, "xgen").Scan(&conf)
	return
}

func GenCodesPreview(ctx context.Context, params *genmodel.GenCodesPreviewParam) (output *genmodel.GenCodesPreviewModel, err error) {

	genConfig, err := GetLoadGenerate(ctx)
	if err != nil {
		return nil, err
	}
	daoConfig, err := gendao.GetDaoConfig(ctx, params.DbName)
	if err != nil {
		return nil, err
	}
	switch params.GenType {
	case genconsts.GenCodesTypeCurd, genconsts.GenCodesTypeTree:
		return genview.Curd.DoPreview(ctx, &genmodel.CurdPreviewParam{
			In:        params,
			DaoConfig: daoConfig,
			Config:    genConfig,
		})
	default:
		err = gerror.Newf("生成类型暂不支持！")
		return
	}
	return
}

// Build 提交生成
func GenCodesBuild(ctx context.Context, params *genmodel.GenCodesBuildParam) (output *genmodel.GenCodesBuildModel, err error) {
	output = &genmodel.GenCodesBuildModel{}
	genConfig, err := GetLoadGenerate(ctx)
	if err != nil {
		return nil, err
	}
	daoConfig, err := gendao.GetDaoConfig(ctx, params.DbName)
	if err != nil {
		return nil, err
	}
	switch params.GenType {
	case genconsts.GenCodesTypeCurd, genconsts.GenCodesTypeTree:
		pin := params

		return output, genview.Curd.DoBuild(ctx, &genmodel.CurdBuildParam{
			PreviewIn: &genmodel.CurdPreviewParam{
				In:        pin,
				DaoConfig: daoConfig,
				Config:    genConfig,
			},
			BeforeEvent: genmodel.CurdBuildEvent{"runDao": Dao},
			AfterEvent: genmodel.CurdBuildEvent{"runService": func(ctx context.Context) (err error) {
				cfg := gendao.GetServiceConfig()

				// 插件模块，切换到插件下运行gen service
				if genConfig.Application.Crud.Templates[pin.GenTemplate].IsAddon {
					// 依然使用配置中的参数，只是将生成路径指向插件模块路径
					cfg.SrcFolder = "addons/" + pin.AddonName + "/logic"
					cfg.DstFolder = "addons/" + pin.AddonName + "/service"
				}
				err = ServiceWithCfg(ctx, cfg)
				return
			}},
		})
	default:
		err = gerror.Newf("生成类型暂不支持！")
		return
	}
}

// ServiceWithCfg 生成业务接口
func ServiceWithCfg(ctx context.Context, cfg ...*genmodel.GenServiceConfig) (err error) {
	c := gendao.GetServiceConfig()
	if len(cfg) > 0 {
		c = cfg[0]
	}
	cmd := `gf gen service`

	if c.SrcFolder != "" {
		cmd += ` -s ` + c.SrcFolder
	}
	if c.DstFolder != "" {
		cmd += ` -d ` + c.DstFolder
	}
	if c.DstFileNameCase != "" {
		cmd += ` -f ` + c.DstFileNameCase
	}
	if c.WatchFile != "" {
		cmd += ` -w ` + c.WatchFile
	}
	if c.StPattern != "" {
		cmd += ` -a ` + c.StPattern
	}
	if c.Clear {
		cmd += ` -l`
	}
	if c.ImportPrefix != "" {
		cmd += ` -i ` + c.ImportPrefix
	}
	if c.Packages != nil {
		cmd += ` -p ` + strings.Join(c.Packages, ",")
	}
	r, err := gproc.ShellExec(gctx.New(), cmd)
	if err != nil {
		return err
	}
	g.Log().Info(ctx, "生成业务接口 Stdout", r)
	return
}

// Dao 生成数据库实体
func Dao(ctx context.Context) (err error) {
	cmd := `gf gen dao`
	r, err := gproc.ShellExec(gctx.New(), cmd)
	if err != nil {
		g.Log().Error(ctx, "生成数据库实体 Stderr", r)
		return err
	}
	g.Log().Info(ctx, "生成数据库实体 Stdout", r)
	return
}
