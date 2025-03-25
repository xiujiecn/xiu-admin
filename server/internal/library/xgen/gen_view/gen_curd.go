package genview

import (
	"context"
	"runtime"
	"slices"
	"strings"
	"xiujieadmin/internal/dao"
	genconsts "xiujieadmin/internal/library/xgen/gen_consts"
	genmodel "xiujieadmin/internal/library/xgen/gen_model"
	"xiujieadmin/internal/library/xgorm"
	"xiujieadmin/internal/service"
	"xiujieadmin/utility"
	"xiujieadmin/utility/tree"
	version "xiujieadmin/utility/version"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gutil"
)

var Curd = gCurd{}

type gCurd struct{}

func (l *gCurd) DoBuild(ctx context.Context, in *genmodel.CurdBuildParam) (err error) {
	st := gtime.Now()
	preview, err := l.DoPreview(ctx, in.PreviewIn)
	if err != nil {
		return
	}

	db, err := g.DB().Open(ParseDBConfigNodeLink(&gdb.ConfigNode{Link: in.PreviewIn.DaoConfig.Link}))
	if err != nil {
		err = gerror.Newf("连接数据库失败，请检查配置文件[server/hack/config.yaml]数据库配置是否正确！err:%v", err.Error())
		return err
	}

	defer db.Close()
	if err = db.Ping(); err != nil {
		err = gerror.Newf("数据库访问异常，请检查配置文件[server/hack/config.yaml]数据库配置是否正确！err:%v", err.Error())
		return
	}

	// 前置操作
	if len(in.BeforeEvent) > 0 {
		for name, f := range in.BeforeEvent {
			if gstr.InArray(in.PreviewIn.Options.AutoOps, name) {
				if err = f(ctx); err != nil {
					return gerror.Newf("in doBuild operation beforeEvent to '%s' failed:%v", name, err)
				}
			}
		}
	}

	// 处理sql文件
	handleSqlFile := func(vi *genmodel.GenFile) (err error) {
		// 无需生成
		if vi.Meth != genconsts.GenCodesBuildMethCreate && vi.Meth != genconsts.GenCodesBuildMethCover {
			return
		}

		if err = gfile.PutContents(vi.Path, strings.TrimSpace(vi.Content)); err != nil {
			return gerror.Newf("writing content to '%s' failed: %v", vi.Path, err)
		}

		// 导入失败，将sql文件删除
		if err = ImportSql(ctx, vi.Path); err != nil {
			_ = gfile.Remove(vi.Path)
		}
		return
	}

	// 将sql文件提取出来优先处理
	// sql执行过程出错是高概率事件，后期在执行前要进行预效验，尽量减少在执行过程中出错的可能性
	sqlGenFile, ok := preview.Views["source.sql"]
	if ok {
		delete(preview.Views, "source.sql")
		if err = handleSqlFile(sqlGenFile); err != nil {
			return
		}
	}

	for _, vi := range preview.Views {
		// 无需生成
		if vi.Meth != genconsts.GenCodesBuildMethCreate && vi.Meth != genconsts.GenCodesBuildMethCover {
			continue
		}

		if err = gfile.PutContents(vi.Path, strings.TrimSpace(vi.Content)); err != nil {
			return gerror.Newf("writing content to '%s' failed: %v", vi.Path, err)
		}
	}

	// 后置操作
	if len(in.AfterEvent) > 0 {
		for name, f := range in.AfterEvent {
			if gstr.InArray(in.PreviewIn.Options.AutoOps, name) {
				if err = f(ctx); err != nil {
					return gerror.Newf("in doBuild operation afterEvent to '%s' failed:%v", name, err)
				}
			}
		}
	}
	g.Log().Debugf(ctx, "generate code operation completed, %vms", gtime.Now().Sub(st).Milliseconds())
	return
}

func (l *gCurd) DoPreview(ctx context.Context, in *genmodel.CurdPreviewParam) (res *genmodel.GenCodesPreviewModel, err error) {
	// 初始化
	if err = l.initInput(ctx, in); err != nil {
		return nil, err
	}

	// 加载模板
	if err = l.loadView(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateApiContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateInputContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateControllerContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateLogicContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateRouterContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateWebApiContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateWebModelContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateWebIndexContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateWebEditContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateWebViewContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateSqlContent(ctx, in); err != nil {
		return nil, err
	}

	in.Content.Config = in.Config
	res = in.Content
	return
}

func (l *gCurd) initInput(ctx context.Context, in *genmodel.CurdPreviewParam) (err error) {
	in.Content = new(genmodel.GenCodesPreviewModel)
	in.Content.Views = make(map[string]*genmodel.GenFile)

	// 初始化生成选项
	if err = initOptions(in); err != nil {
		return err
	}

	// 初始化表字段配置
	if err = initTableField(ctx, in); err != nil {
		return err
	}

	// 初始化树表
	if err = initTableTree(in); err != nil {
		return err
	}

	// 初始化生成选型
	initStep(in)

	// 初始化方法字典
	if err = initFuncDict(in); err != nil {
		return err
	}

	// 初始化生成模板
	if err = initTemplate(in); err != nil {
		return err
	}
	return
}
func initOptions(in *genmodel.CurdPreviewParam) (err error) {
	if err = in.In.Options.Scan(&in.Options); err != nil {
		return
	}
	in.Options.DictMap = make(g.Map)
	return
}

func initTableField(ctx context.Context, in *genmodel.CurdPreviewParam) (err error) {
	// 加载主表配置
	if err = in.In.MasterColumns.Scan(&in.MasterFields); err != nil {
		return
	}

	if len(in.MasterFields) == 0 {
		if in.MasterFields, err = TableColumnList(ctx, &genmodel.GenCodesColumnListParam{DbGroup: in.In.DbName, TableName: in.In.TableName}); err != nil {
			return
		}
	}

	// 主键属性
	in.Pk = getPkField(in)
	if in.Pk == nil {
		return gerror.New("initInput no primary key is set in the table!")
	}

	in.MasterFields = ReviseFields(in.MasterFields)

	// 检查表命名
	var names = []string{in.In.DaoName}
	for _, v := range in.Options.Join {
		v.Columns = ReviseFields(v.Columns)
		names = append(names, v.DaoName)
		g.Log().Infof(ctx, "initTableField in.Options.Join: %+v", v)
	}
	if err = CheckIllegalName("数据库表名", names...); err != nil {
		return
	}

	if err = CheckIllegalName("实体命名", in.In.VarName); err != nil {
		return
	}
	return
}

// getPkField 获取主键
func getPkField(in *genmodel.CurdPreviewParam) *genmodel.GenCodesColumnListModel {
	if len(in.MasterFields) == 0 {
		panic("getPkField masterFields uninitialized.")
	}
	for _, field := range in.MasterFields {
		if IsIndexPK(field.Index) {
			return field
		}
	}
	return nil
}

func initTableTree(in *genmodel.CurdPreviewParam) (err error) {
	// 检查树表字段
	if in.In.GenType == genconsts.GenCodesTypeTree {
		if err = CheckTreeTableFields(in.MasterFields); err != nil {
			return err
		}

		// 解析选项树名称字段
		has := false
		for _, field := range in.MasterFields {
			if in.Options.Tree.TitleColumn == field.Name {
				in.Options.Tree.TitleField = field
				has = true
				break
			}
		}
		if !has {
			err = gerror.New("请选择一个有效的树名称字段")
			return
		}
	}
	return err
}

// CheckTreeTableFields 检查树表字段
func CheckTreeTableFields(columns []*genmodel.GenCodesColumnListModel) (err error) {
	var fields = gutil.Copy(defaultTreeFields).([]string)
	for _, v := range columns {
		if gstr.InArray(fields, v.Name) {
			fields = slices.DeleteFunc(fields, func(item string) bool {
				return item == v.Name
			})
		}
	}

	if len(fields) > 0 {
		err = gerror.Newf("树表必须包含[%v]字段", strings.Join(fields, "、"))
		return err
	}
	return
}

func initStep(in *genmodel.CurdPreviewParam) {
	in.Options.Step = new(genmodel.CurdStep)
	in.Options.Step.HasMaxSort = HasMaxSort(in.MasterFields)
	in.Options.Step.HasAdd = gstr.InArray(in.Options.HeadOps, "add")
	in.Options.Step.HasBatchDel = gstr.InArray(in.Options.HeadOps, "batchDel") && gstr.InArray(in.Options.ColumnOps, "check")
	in.Options.Step.HasExport = gstr.InArray(in.Options.HeadOps, "export")
	in.Options.Step.HasNotFilterAuth = gstr.InArray(in.Options.ColumnOps, "notFilterAuth")
	in.Options.Step.HasEdit = gstr.InArray(in.Options.ColumnOps, "edit")
	in.Options.Step.HasDel = gstr.InArray(in.Options.ColumnOps, "del")
	in.Options.Step.HasView = gstr.InArray(in.Options.ColumnOps, "view")
	in.Options.Step.HasStatus = HasStatus(in.Options.ColumnOps, in.MasterFields)
	in.Options.Step.HasSwitch = HasSwitch(in.MasterFields)
	in.Options.Step.HasCheck = gstr.InArray(in.Options.ColumnOps, "check")
	in.Options.Step.HasMenu = gstr.InArray(in.Options.AutoOps, "genMenuPermissions")
	in.Options.Step.HasQueryMemberSummary = HasQueryMemberSummary(in.MasterFields)
	in.Options.Step.HasHookMemberSummary = HasHookMemberSummary(in.MasterFields)
	in.Options.Step.IsTreeTable = in.In.GenType == genconsts.GenCodesTypeTree
	if in.Options.Step.IsTreeTable {
		in.Options.Step.IsOptionTreeTable = in.Options.Tree.StyleType == genconsts.GenCodesTreeStyleTypeOption
	}
	in.Options.Step.HasFuncDict = gstr.InArray(in.Options.AutoOps, "genFuncDict")
	in.Options.Step.IsAddon = in.Config.Application.Crud.Templates[in.In.GenTemplate].IsAddon
	if in.Options.PresetStep.FormGridCols < 1 {
		in.Options.PresetStep.FormGridCols = 1
	}
}

func initFuncDict(in *genmodel.CurdPreviewParam) (err error) {
	g.Log().Infof(context.Background(), "server/internal/library/xgen/gen_curd.go initFuncDict in.Options.Step.HasFuncDict:%v in.Options.FuncDict:%+v", in.Options.Step.HasFuncDict, in.Options.FuncDict)
	if !in.Options.Step.HasFuncDict || in.Options.FuncDict == nil {
		return
	}

	if len(in.Options.FuncDict.LabelColumn) == 0 || len(in.Options.FuncDict.ValueColumn) == 0 {
		err = gerror.New("生成字典选项必须设置选项值和选项名称")
		return err
	}

	for _, field := range in.MasterFields {
		if field.Name == in.Options.FuncDict.ValueColumn {
			in.Options.FuncDict.Value = field
		}

		if field.Name == in.Options.FuncDict.LabelColumn {
			in.Options.FuncDict.Label = field
		}
	}
	g.Log().Infof(context.Background(), "server/internal/library/xgen/gen_curd.go initFuncDict in.Options.Step.HasFuncDict:%v in.Options.FuncDict:%+v", in.Options.Step.HasFuncDict, in.Options.FuncDict)
	return
}

func initTemplate(in *genmodel.CurdPreviewParam) (err error) {
	if len(in.Config.Application.Crud.Templates)-1 < in.In.GenTemplate {
		return gerror.New("没有找到生成模板的配置，请检查！")
	}

	// api前缀
	apiPrefix := gstr.LcFirst(in.In.VarName)
	if in.Config.Application.Crud.Templates[in.In.GenTemplate].IsAddon {
		apiPrefix = in.In.AddonName + "/" + apiPrefix
	}
	in.Options.ApiPrefix = apiPrefix

	if err = checkCurdPath(in.Config.Application.Crud.Templates[in.In.GenTemplate], in.In.AddonName); err != nil {
		return
	}
	in.Options.TemplateGroup = in.Config.Application.Crud.Templates[in.In.GenTemplate].MasterPackage
	return
}

func checkCurdPath(temp *genmodel.GenCodesConfigCrudTemplate, addonName string) (err error) {
	if temp == nil {
		return gerror.New("生成模板配置不能为空")
	}

	if temp.IsAddon {
		temp.TemplatePath = gstr.Replace(temp.TemplatePath, "{$name}", addonName)
		temp.ApiPath = gstr.Replace(temp.ApiPath, "{$name}", addonName)
		temp.InputPath = gstr.Replace(temp.InputPath, "{$name}", addonName)
		temp.ControllerPath = gstr.Replace(temp.ControllerPath, "{$name}", addonName)
		temp.LogicPath = gstr.Replace(temp.LogicPath, "{$name}", addonName)
		temp.RouterPath = gstr.Replace(temp.RouterPath, "{$name}", addonName)
		temp.SqlPath = gstr.Replace(temp.SqlPath, "{$name}", addonName)
		temp.WebApiPath = gstr.Replace(temp.WebApiPath, "{$name}", addonName)
		temp.WebViewsPath = gstr.Replace(temp.WebViewsPath, "{$name}", addonName)
	}

	tip := `生成模板配置参数'%s'路径不存在，请先创建路径:%s`

	if !gfile.Exists(temp.TemplatePath) {
		return gerror.Newf(tip, "TemplatePath", temp.TemplatePath)
	}
	if !gfile.Exists(temp.ApiPath) {
		return gerror.Newf(tip, "ApiPath", temp.ApiPath)
	}
	if !gfile.Exists(temp.InputPath) {
		return gerror.Newf(tip, "InputPath", temp.InputPath)
	}
	if !gfile.Exists(temp.ControllerPath) {
		return gerror.Newf(tip, "ControllerPath", temp.ControllerPath)
	}
	if !gfile.Exists(temp.LogicPath) {
		return gerror.Newf(tip, "LogicPath", temp.LogicPath)
	}
	if !gfile.Exists(temp.RouterPath) {
		return gerror.Newf(tip, "RouterPath", temp.RouterPath)
	}
	if !gfile.Exists(temp.SqlPath) {
		return gerror.Newf(tip, "SqlPath", temp.SqlPath)
	}
	if !gfile.Exists(temp.WebApiPath) {
		return gerror.Newf(tip, "WebApiPath", temp.WebApiPath)
	}
	if !gfile.Exists(temp.WebViewsPath) {
		return gerror.Newf(tip, "WebViewsPath", temp.WebViewsPath)
	}
	return
}

func (l *gCurd) loadView(ctx context.Context, in *genmodel.CurdPreviewParam) (err error) {
	temp := in.Config.Application.Crud.Templates[in.In.GenTemplate]
	view := gview.New()
	err = view.SetConfigWithMap(g.Map{
		"Paths":      temp.TemplatePath,
		"Delimiters": in.Config.Delimiters,
	})
	if err != nil {
		return
	}

	now := gtime.Now()
	view.BindFuncMap(g.Map{
		"NowYear":       now.Year,        // 当前年
		"ToLower":       strings.ToLower, // 全部小写
		"LcFirst":       gstr.LcFirst,    // 首字母小写
		"UcFirst":       gstr.UcFirst,    // 首字母大写
		"ToTSArray":     ToTSArray,       // 转为ts数组格式
		"DictToTSArray": DictToTSArray,   // 转为ts数组格式
	})

	if err = l.generateWebModelDictOptions(ctx, in); err != nil {
		return
	}

	modName, err := GetModName(ctx)
	if err != nil {
		return
	}
	importApi := gstr.Replace(temp.ApiPath, "./", modName+"/") + "/" + strings.ToLower(in.In.VarName) + "/v1"
	importInput := gstr.Replace(temp.InputPath, "./", modName+"/")
	importController := gstr.Replace(temp.ControllerPath, "./", modName+"/")
	importService := "xiujieadmin/internal/service"
	if temp.IsAddon {
		importService = "xiujieadmin/addons/" + in.In.AddonName + "/service"
	}

	in.Options.ImportWebApi = "#/api/gen/" + gstr.LcFirst(in.In.VarName)
	if temp.IsAddon {
		in.Options.ImportWebApi = "#/api/gen/addons/" + in.In.AddonName + "/" + gstr.LcFirst(in.In.VarName)
	}

	componentPrefix := gstr.LcFirst(in.In.VarName)
	if temp.IsAddon {
		componentPrefix = "addons/" + in.In.AddonName + "/" + componentPrefix
	}

	nowTime := now.Format("Y-m-d H:i:s")
	view.Assigns(gview.Params{
		"templateGroup":    in.Options.TemplateGroup,                                    // 生成模板分组名称
		"servFunName":      l.parseServFunName(in.Options.TemplateGroup, in.In.VarName), // 业务服务名称
		"nowTime":          nowTime,                                                     // 当前时间
		"version":          runtime.Version(),                                           // GO 版本
		"xjVersion":        version.BuildVersion,                                        // XJ 版本
		"varName":          in.In.VarName,                                               // 实体名称
		"tableComment":     in.In.TableComment,                                          // 对外名称
		"daoName":          in.In.DaoName,                                               // ORM模型
		"masterFields":     in.MasterFields,                                             // 主表字段
		"pk":               in.Pk,                                                       // 主键属性
		"options":          in.Options,                                                  // 提交选项
		"dictOptions":      in.Options.DictOps,                                          // web字典选项
		"importApi":        importApi,                                                   // 导入goApi包
		"importInput":      importInput,                                                 // 导入input包
		"importController": importController,                                            // 导入控制器包
		"importService":    importService,                                               // 导入业务服务
		"importWebApi":     in.Options.ImportWebApi,                                     // 导入webApi
		"apiPrefix":        in.Options.ApiPrefix,                                        // api前缀
		"componentPrefix":  componentPrefix,                                             // vue子组件前缀
	})

	in.View = view
	return
}

func (l *gCurd) generateApiContent(ctx context.Context, in *genmodel.CurdPreviewParam) (err error) {
	var (
		name    = "api.go"
		tplData = g.Map{}
		genFile = new(genmodel.GenFile)
	)
	genFile.Content, err = in.View.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}

	genFile.Content, err = FormatGo(ctx, name, genFile.Content)
	if err != nil {
		return err
	}

	genFile.Path = utility.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].ApiPath, strings.ToLower(in.In.VarName), "v1", strings.ToLower(in.In.VarName)+".go")
	genFile.Meth = genconsts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = genconsts.GenCodesBuildMethSkip
	}

	genFile.Required = true

	if genFile.Meth == genconsts.GenCodesBuildMethSkip && gstr.InArray(in.Options.AutoOps, "forcedCover") {
		genFile.Meth = genconsts.GenCodesBuildMethCover
	}

	in.Content.Views[name] = genFile
	return
}

func (l *gCurd) generateInputContent(ctx context.Context, in *genmodel.CurdPreviewParam) (err error) {
	var (
		name    = "input.go"
		genFile = new(genmodel.GenFile)
	)

	tplData, err := l.inputTplData(ctx, in)
	if err != nil {
		return err
	}

	genFile.Content, err = in.View.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}

	genFile.Content, err = FormatGo(ctx, name, genFile.Content)
	if err != nil {
		return err
	}

	genFile.Path = utility.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].InputPath, CamelCaseToUnderline(in.In.VarName)+".go")
	genFile.Meth = genconsts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = genconsts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if genFile.Meth == genconsts.GenCodesBuildMethSkip && gstr.InArray(in.Options.AutoOps, "forcedCover") {
		genFile.Meth = genconsts.GenCodesBuildMethCover
	}

	in.Content.Views[name] = genFile
	return
}

func (l *gCurd) generateControllerContent(ctx context.Context, in *genmodel.CurdPreviewParam) (err error) {
	var (
		name    = "controller.go"
		tplData = g.Map{}
		genFile = new(genmodel.GenFile)
	)

	genFile.Content, err = in.View.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}

	genFile.Content, err = FormatGo(ctx, name, genFile.Content)
	if err != nil {
		return err
	}

	genFile.Path = utility.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].ControllerPath, CamelCaseToUnderline(in.In.VarName)+".go")
	genFile.Meth = genconsts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = genconsts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if genFile.Meth == genconsts.GenCodesBuildMethSkip && gstr.InArray(in.Options.AutoOps, "forcedCover") {
		genFile.Meth = genconsts.GenCodesBuildMethCover
	}

	in.Content.Views[name] = genFile
	return
}

func (l *gCurd) generateLogicContent(ctx context.Context, in *genmodel.CurdPreviewParam) (err error) {
	var (
		name    = "logic.go"
		genFile = new(genmodel.GenFile)
	)

	tplData, err := l.logicTplData(ctx, in)
	if err != nil {
		return err
	}

	genFile.Content, err = in.View.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}

	genFile.Content, err = FormatGo(ctx, name, genFile.Content)
	if err != nil {
		return err
	}

	genFile.Path = utility.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].LogicPath, CamelCaseToUnderline(in.In.VarName)+".go")
	genFile.Meth = genconsts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = genconsts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if genFile.Meth == genconsts.GenCodesBuildMethSkip && gstr.InArray(in.Options.AutoOps, "forcedCover") {
		genFile.Meth = genconsts.GenCodesBuildMethCover
	}

	in.Content.Views[name] = genFile
	return
}

func (l *gCurd) generateRouterContent(ctx context.Context, in *genmodel.CurdPreviewParam) (err error) {
	var (
		name    = "router.go"
		tplData = g.Map{}
		genFile = new(genmodel.GenFile)
	)
	genFile.Content, err = in.View.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}

	genFile.Content, err = FormatGo(ctx, name, genFile.Content)
	if err != nil {
		return err
	}

	genFile.Path = utility.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].RouterPath, CamelCaseToUnderline(in.In.VarName)+".go")
	genFile.Meth = genconsts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = genconsts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if genFile.Meth == genconsts.GenCodesBuildMethSkip && gstr.InArray(in.Options.AutoOps, "forcedCover") {
		genFile.Meth = genconsts.GenCodesBuildMethCover
	}

	in.Content.Views[name] = genFile
	return
}

func (l *gCurd) generateWebApiContent(ctx context.Context, in *genmodel.CurdPreviewParam) (err error) {
	var (
		name    = "web.api.ts"
		tplData = g.Map{}
		genFile = new(genmodel.GenFile)
	)
	genFile.Content, err = in.View.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}

	genFile.Content = FormatTs(genFile.Content)

	genFile.Path = utility.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].WebApiPath, gstr.LcFirst(in.In.VarName), "index.ts")
	genFile.Meth = genconsts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = genconsts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if genFile.Meth == genconsts.GenCodesBuildMethSkip && gstr.InArray(in.Options.AutoOps, "forcedCover") {
		genFile.Meth = genconsts.GenCodesBuildMethCover
	}

	in.Content.Views[name] = genFile
	return
}

func (l *gCurd) generateWebModelContent(ctx context.Context, in *genmodel.CurdPreviewParam) (err error) {
	var (
		name    = "web.model.ts"
		genFile = new(genmodel.GenFile)
	)

	tplData, err := l.webModelTplData(ctx, in)
	if err != nil {
		return
	}

	genFile.Content, err = in.View.Parse(ctx, name+".template", tplData)
	if err != nil {
		return
	}

	genFile.Content = FormatTs(genFile.Content)

	genFile.Path = utility.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].WebViewsPath, gstr.LcFirst(in.In.VarName), "model.tsx")
	genFile.Meth = genconsts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = genconsts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if genFile.Meth == genconsts.GenCodesBuildMethSkip && gstr.InArray(in.Options.AutoOps, "forcedCover") {
		genFile.Meth = genconsts.GenCodesBuildMethCover
	}
	in.Content.Views[name] = genFile
	return
}

func (l *gCurd) generateWebIndexContent(ctx context.Context, in *genmodel.CurdPreviewParam) (err error) {
	var (
		name    = "web.index.vue"
		genFile = new(genmodel.GenFile)
	)

	tplData, err := l.webIndexTplData(ctx, in)
	if err != nil {
		return err
	}

	genFile.Content, err = in.View.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}

	genFile.Content = FormatVue(genFile.Content)

	genFile.Path = utility.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].WebViewsPath, gstr.LcFirst(in.In.VarName), "index.vue")
	genFile.Meth = genconsts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = genconsts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if genFile.Meth == genconsts.GenCodesBuildMethSkip && gstr.InArray(in.Options.AutoOps, "forcedCover") {
		genFile.Meth = genconsts.GenCodesBuildMethCover
	}
	in.Content.Views[name] = genFile
	return
}

func (l *gCurd) generateWebEditContent(ctx context.Context, in *genmodel.CurdPreviewParam) (err error) {
	var (
		name    = "web.edit.vue"
		genFile = new(genmodel.GenFile)
	)

	tplData, err := l.webEditTplData(ctx, in)
	if err != nil {
		return err
	}

	genFile.Content, err = in.View.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}

	genFile.Content = FormatVue(genFile.Content)

	genFile.Path = utility.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].WebViewsPath, gstr.LcFirst(in.In.VarName), "edit.vue")
	genFile.Meth = genconsts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = genconsts.GenCodesBuildMethSkip
	}
	genFile.Required = true
	if genFile.Meth == genconsts.GenCodesBuildMethSkip && gstr.InArray(in.Options.AutoOps, "forcedCover") {
		genFile.Meth = genconsts.GenCodesBuildMethCover
	}

	if !in.Options.Step.HasEdit {
		genFile.Meth = genconsts.GenCodesBuildIgnore
		genFile.Required = false
	}

	in.Content.Views[name] = genFile
	return
}

func (l *gCurd) generateWebViewContent(ctx context.Context, in *genmodel.CurdPreviewParam) (err error) {
	var (
		name    = "web.view.vue"
		genFile = new(genmodel.GenFile)
	)

	tplData, err := l.webViewTplData(ctx, in)
	if err != nil {
		return err
	}

	genFile.Content, err = in.View.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}

	genFile.Content = FormatVue(genFile.Content)

	genFile.Path = utility.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].WebViewsPath, gstr.LcFirst(in.In.VarName), "view.vue")
	genFile.Meth = genconsts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = genconsts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if genFile.Meth == genconsts.GenCodesBuildMethSkip && gstr.InArray(in.Options.AutoOps, "forcedCover") {
		genFile.Meth = genconsts.GenCodesBuildMethCover
	}

	if !in.Options.Step.HasView {
		genFile.Meth = genconsts.GenCodesBuildIgnore
		genFile.Required = false
	}

	in.Content.Views[name] = genFile
	return
}

func (l *gCurd) generateSqlContent(ctx context.Context, in *genmodel.CurdPreviewParam) (err error) {
	var (
		name    = "source.sql"
		config  = g.DB("default").GetConfig()
		tplData = g.Map{
			"dbName":        config.Name,
			"menuTable":     config.Prefix + "sys_menu",
			"mainComponent": "LAYOUT",
		}
		genFile = new(genmodel.GenFile)
	)

	menus, err := service.SysMenu().GetFastList(ctx)
	if err != nil {
		return err
	}

	tplData["dirPid"], tplData["dirLevel"], tplData["dirTree"], err = xgorm.AutoUpdateTree(ctx, &dao.SysMenu, 0, int64(in.Options.Menu.Pid), &xgorm.TreeFiledOption{
		IdField:    "menu_id",
		PidField:   "parent_id",
		LevelField: "level",
		TreeField:  "tree",
	})
	if err != nil {
		return err
	}

	tplData["listLevel"] = tplData["dirLevel"].(int) + 1
	tplData["btnLevel"] = tplData["dirLevel"].(int) + 2
	tplData["sortLevel"] = tplData["dirLevel"].(int) + 3

	pageRedirect := ""
	if in.Options.Menu.Pid > 0 {
		tplData["mainComponent"] = "ParentLayout"
		menu, ok := menus[int64(in.Options.Menu.Pid)]
		if !ok {
			err = gerror.New("选择的上级菜单不存在")
			return
		}
		for _, id := range tree.GetIds(menu.Tree) {
			if v, ok2 := menus[id]; ok2 {
				if !gstr.HasSuffix(pageRedirect, "/") && !gstr.HasPrefix(v.Path, "/") {
					pageRedirect += "/"
				}
				pageRedirect += v.Path
			}
		}

		if !gstr.HasSuffix(pageRedirect, "/") && !gstr.HasPrefix(menu.Path, "/") {
			pageRedirect += "/"
		}
		pageRedirect += menu.Path
	}
	pageRedirect += "/" + gstr.LcFirst(in.In.VarName) + "/index"
	tplData["pageRedirect"] = pageRedirect

	genFile.Path = utility.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].SqlPath, CamelCaseToUnderline(in.In.VarName)+"_menu.sql")
	genFile.Meth = genconsts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = genconsts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if !in.Options.Step.HasMenu {
		genFile.Meth = genconsts.GenCodesBuildIgnore
		genFile.Required = false
	}

	// 需要生成时，检查菜单命名是否存在
	if genFile.Meth == genconsts.GenCodesBuildMethCreate {
		menuNamePrefix := gstr.LcFirst(in.In.VarName)
		menuNames := []string{menuNamePrefix, menuNamePrefix + "Index"}
		if in.Options.Step.HasEdit {
			menuNames = append(menuNames, menuNamePrefix+"Edit")
			menuNames = append(menuNames, menuNamePrefix+"View")
		}
		if in.Options.Step.HasView {
			menuNames = append(menuNames, menuNamePrefix+"View")
		}
		if in.Options.Step.HasMaxSort {
			menuNames = append(menuNames, menuNamePrefix+"MaxSort")
		}
		if in.Options.Step.HasDel {
			menuNames = append(menuNames, menuNamePrefix+"Delete")
		}
		if in.Options.Step.HasStatus {
			menuNames = append(menuNames, menuNamePrefix+"Status")
		}
		if in.Options.Step.HasSwitch {
			menuNames = append(menuNames, menuNamePrefix+"Switch")
		}
		if in.Options.Step.HasExport {
			menuNames = append(menuNames, menuNamePrefix+"Export")
		}
		if in.Options.Step.IsTreeTable {
			menuNames = append(menuNames, menuNamePrefix+"TreeOption")
		}

		menuNames = utility.UniqueSlice(menuNames)

		hasMenus, err := service.SysMenu().Model(ctx).Fields("path").WhereIn("path", menuNames).Array()
		if err != nil {
			return err
		}

		if len(hasMenus) > 0 {
			err = gerror.Newf("要生成的菜单中有已存在的路由别名，请检查并删除:%v", strings.Join(gvar.New(hasMenus).Strings(), `、`))
			return err
		}
	}

	tplData["generatePath"] = genFile.Path
	genFile.Content, err = in.View.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}

	in.Content.Views[name] = genFile
	return
}
