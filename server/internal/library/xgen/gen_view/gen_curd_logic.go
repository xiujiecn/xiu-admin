// package genview
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package genview

import (
	"bytes"
	"context"
	"fmt"

	genconsts "xiuadmin/internal/library/xgen/gen_consts"
	genmodel "xiuadmin/internal/library/xgen/gen_model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

const (
	LogicWhereComments  = "\n\t// 查询%s\n"
	LogicWhereNoSupport = "\t// TODO 暂不支持生成[ %s ]查询方式，请自行补充此处代码！"
	LogicEditUpdate     = "\tif _, err = s.Model(ctx%s).\n\t\t\tFields(%sin.%sUpdateFields{}).\n\t\t\tWherePri(in.%s).Data(in).Update(); err != nil {\n\t\t\terr = gerror.Wrap(err, \"修改%s失败，请稍后重试！\")\n\t\t}\n\t\treturn"
	LogicEditInsert     = "\tif _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).\n\t\tFields(%sin.%sInsertFields{}).\n\t\tData(in).Insert(); err != nil {\n\t\terr = gerror.Wrap(err, \"新增%s失败，请稍后重试！\")\n\t}"
	LogicEditUnique     = "\t// 验证'%s'唯一\n\tif err = xgorm.IsUnique(ctx, &dao.%s, g.Map{dao.%s.Columns().%s: in.%s}, \"%s已存在\", in.%s); err != nil {\n\t\treturn\n\t}\n"
	LogicSwitchUpdate   = "g.Map{\n\t\tin.Key:                       in.Value,\n%s}"
	LogicStatusUpdate   = "g.Map{\n\t\tdao.%s.Columns().Status:    in.Status,\n%s}"
	LogicDeletedUpdate  = "g.Map{\n%s}"
)

func (l *gCurd) logicTplData(ctx context.Context, in *genmodel.CurdPreviewParam) (data g.Map, err error) {
	g.Log().Infof(ctx, "logicTplData in.Options.Join: %+v", in.Options.Join)
	data = make(g.Map)
	data["listWhere"] = l.genLogicListWhere(ctx, in)
	data["listJoin"] = l.genLogicListJoin(ctx, in)
	data["listFields"] = l.genLogicListFields(ctx, in)
	data["listOrder"] = l.genLogicListOrder(ctx, in)
	data["edit"] = l.genLogicEdit(ctx, in)
	data["switchFields"] = l.genLogicSwitchFields(ctx, in)
	data["switchUpdate"] = l.genLogicSwitchUpdate(ctx, in)
	data["statusUpdate"] = l.genLogicStatusUpdate(ctx, in)
	data["deletedUpdate"] = l.genLogicDeletedUpdate(ctx, in)
	return
}

func (l *gCurd) genLogicDeletedUpdate(ctx context.Context, in *genmodel.CurdPreviewParam) string {
	isDestroy := false
	var update string
	for _, field := range in.MasterFields {
		if field.GoName == "DeletedBy" {
			update += "\t\tdao." + in.In.DaoName + ".Columns().DeletedBy: contexts.GetUserId(ctx),\n"
			isDestroy = true
		}
		if field.GoName == "DeletedAt" {
			update += "\t\tdao." + in.In.DaoName + ".Columns().DeletedAt: gtime.Now(),\n"
		}
	}

	if !isDestroy {
		return ""
	}

	update += "\t"
	return fmt.Sprintf(LogicDeletedUpdate, update)
}

func (l *gCurd) genLogicStatusUpdate(ctx context.Context, in *genmodel.CurdPreviewParam) string {
	var update string
	for _, field := range in.MasterFields {
		if field.GoName == "UpdatedBy" {
			update += "\t\tdao." + in.In.DaoName + ".Columns().UpdatedBy: contexts.GetUserId(ctx),\n"
		}
	}

	update += "\t"
	return fmt.Sprintf(LogicStatusUpdate, in.In.DaoName, update)
}

func (l *gCurd) genLogicSwitchUpdate(ctx context.Context, in *genmodel.CurdPreviewParam) string {
	var update string
	for _, field := range in.MasterFields {
		if field.GoName == "UpdatedBy" {
			update += "\t\tdao." + in.In.DaoName + ".Columns().UpdatedBy: contexts.GetUserId(ctx),\n"
		}
	}

	update += "\t"
	return fmt.Sprintf(LogicSwitchUpdate, update)
}

func (l *gCurd) genLogicSwitchFields(ctx context.Context, in *genmodel.CurdPreviewParam) string {
	buffer := bytes.NewBuffer(nil)
	if in.Options.Step.HasSwitch {
		for _, field := range in.MasterFields {
			if field.FormMode == "Switch" {
				buffer.WriteString("\t\tdao." + in.In.DaoName + ".Columns()." + field.GoName + ",\n")
			}
		}
	}
	return buffer.String()
}

func (l *gCurd) genLogicEdit(ctx context.Context, in *genmodel.CurdPreviewParam) g.Map {
	var (
		data         = make(g.Map)
		updateBuffer = bytes.NewBuffer(nil)
		insertBuffer = bytes.NewBuffer(nil)
		uniqueBuffer = bytes.NewBuffer(nil)
	)

	for _, field := range in.MasterFields {
		if field.GoName == "TenantId" {
			insertBuffer.WriteString("\t\ttenantId := contexts.GetTenantId(ctx)\n\t\tin.TenantId = &tenantId\n")
		}
		if field.GoName == "CreatedDept" {
			insertBuffer.WriteString("\t\tin.CreatedDept = contexts.GetDeptId(ctx)\n")
		}
		if field.GoName == "CreatedAt" {
			insertBuffer.WriteString("\t\tin.CreatedAt = gtime.Now()\n")
		}
		if field.GoName == "CreatedBy" {
			insertBuffer.WriteString("\tin.CreatedBy = contexts.GetUserId(ctx)\n")
		}
		if field.GoName == "UpdatedBy" {
			updateBuffer.WriteString("\t\tin.UpdatedBy = contexts.GetUserId(ctx)\n")
		}
		if field.GoName == "UpdatedAt" {
			updateBuffer.WriteString("\t\tin.UpdatedAt = gtime.Now()\n")
		}

		if field.Unique {
			uniqueBuffer.WriteString(fmt.Sprintf(LogicEditUnique, field.GoName, in.In.DaoName, in.In.DaoName, field.GoName, field.GoName, field.Dc, in.Pk.GoName))
		}
	}

	notFilterAuth := ""
	if gstr.InArray(in.Options.ColumnOps, "notFilterAuth") {
		notFilterAuth = ", &handler.Option{FilterAuth: false}"
	}

	updateBuffer.WriteString(fmt.Sprintf(LogicEditUpdate, notFilterAuth, in.Options.TemplateGroup, in.In.VarName, in.Pk.GoName, in.In.TableComment))
	insertBuffer.WriteString(fmt.Sprintf(LogicEditInsert, in.Options.TemplateGroup, in.In.VarName, in.In.TableComment))

	data["update"] = updateBuffer.String()
	data["insert"] = insertBuffer.String()
	data["unique"] = uniqueBuffer.String()
	return data
}

func (l *gCurd) genLogicListOrder(ctx context.Context, in *genmodel.CurdPreviewParam) string {
	statement := ""
	if hasEffectiveJoins(in.Options.Join) {
		statement = "dao." + in.In.DaoName + ".Table() + \".\" +"
	}
	buffer := bytes.NewBuffer(nil)
	if in.Options.Step.HasMaxSort {
		buffer.WriteString("OrderAsc(" + statement + "dao." + in.In.DaoName + ".Columns().Sort).")
	}
	buffer.WriteString("OrderDesc(" + statement + "dao." + in.In.DaoName + ".Columns()." + in.Pk.GoName + ")")
	return buffer.String()
}

func (l *gCurd) genLogicListJoin(ctx context.Context, in *genmodel.CurdPreviewParam) (link string) {
	connector := `"="`
	if hasEffectiveJoins(in.Options.Join) {
		linkBuffer := bytes.NewBuffer(nil)
		for _, join := range in.Options.Join {
			if isEffectiveJoin(join) {
				linkBuffer.WriteString("\tmod = mod." + genconsts.GenCodesJoinLinkMap[join.LinkMode] + "OnFields(dao." + join.DaoName + ".Table(), dao." + in.In.DaoName + ".Columns()." + gstr.CaseCamel(join.MasterField) + "," + connector + ", dao." + join.DaoName + ".Columns()." + gstr.CaseCamel(join.Field) + ")\n")
			}
		}
		link = linkBuffer.String()
	}
	return
}

func (l *gCurd) genLogicListFields(ctx context.Context, in *genmodel.CurdPreviewParam) (fields string) {
	selectBuffer := bytes.NewBuffer(nil)
	if hasEffectiveJoins(in.Options.Join) {
		selectBuffer.WriteString("mod = mod.FieldsPrefix(dao." + in.In.DaoName + ".Table(), " + in.Options.TemplateGroup + "in." + in.In.VarName + "ListModel{})\n")
		for _, join := range in.Options.Join {
			if isEffectiveJoin(join) {
				selectBuffer.WriteString("mod = mod.Fields(xgorm.JoinFields(ctx, " + in.Options.TemplateGroup + "in." + in.In.VarName + "ListModel{}, &dao." + join.DaoName + ", \"" + join.Alias + "\"))\n")
			}
		}
		fields = selectBuffer.String()
	} else {
		fields = fmt.Sprintf("mod = mod.Fields(%sin.%sListModel{})", in.Options.TemplateGroup, in.In.VarName)
	}
	return
}

func (l *gCurd) genLogicListWhere(ctx context.Context, in *genmodel.CurdPreviewParam) string {
	buffer := bytes.NewBuffer(nil)

	// 主表
	l.genLogicListWhereEach(buffer, in, in.MasterFields, in.In.DaoName, "")
	// 关联表
	g.Log().Infof(ctx, "genLogicListWhere in.Options.Join: %+v", in.Options.Join)
	if hasEffectiveJoins(in.Options.Join) {
		for _, v := range in.Options.Join {
			if isEffectiveJoin(v) {
				l.genLogicListWhereEach(buffer, in, v.Columns, v.DaoName, v.Alias)
			}
		}
	}
	return buffer.String()
}

func (l *gCurd) genLogicListWhereEach(buffer *bytes.Buffer, in *genmodel.CurdPreviewParam, fields []*genmodel.GenCodesColumnListModel, daoName string, alias string) {
	isLink := false
	if alias != "" {
		alias = `"` + alias + `."+`
		isLink = true
	}

	tablePrefix := ""
	wherePrefix := "Where"
	if isLink {
		wherePrefix = "WherePrefix"
		tablePrefix = "dao." + daoName + ".Table(), "
	}

	for _, field := range fields {
		isQuery := false
		// 树表查询上级
		if in.Options.Step.IsTreeTable && field.Name == in.Options.Tree.PidColumn {
			isQuery = true
			field.QueryWhere = WMEq
		}

		if (!field.IsQuery && !isQuery) || field.QueryWhere == "" {
			continue
		}

		buffer.WriteString(fmt.Sprintf(LogicWhereComments, field.Dc))

		var (
			linkMode   string
			whereTag   string
			columnName string
		)

		// 查询用户摘要
		if field.IsQuery && in.Options.Step.HasQueryMemberSummary && IsMemberSummaryField(field.Name) {
			servicePackName := "service"
			if in.Options.Step.IsAddon {
				servicePackName = "isc"
			}
			buffer.WriteString(fmt.Sprintf("if in.%v != \"\" {\n\t\t\t\tids, err := %v.AdminMember().GetIdsByKeyword(ctx, in.%v)\n\t\t\t\tif err != nil {\n\t\t\t\t\treturn nil, 0, err\n\t\t\t\t}\n\t\t\t\tmod = mod.WhereIn(dao.%v.Columns().%v, ids)\n\t\t\t}\n", field.GoName, servicePackName, field.GoName, in.In.DaoName, field.GoName))
			continue
		}

		if IsNumberType(field.GoType) {
			linkMode = `in.` + field.GoName + ` > 0`
		} else if field.GoType == GoTypeGTime {
			linkMode = `in.` + field.GoName + ` != nil`
		} else if field.GoType == GoTypeJson {
			linkMode = `!in.` + field.GoName + `.IsNil()`
		} else {
			linkMode = `in.` + field.GoName + ` != ""`
		}

		if field.QueryWhere == WMBetween || field.QueryWhere == WMNotBetween {
			linkMode = `len(in.` + field.GoName + `) == 2`
		}

		// 如果是关联表重新转换字段
		columnName = field.GoName
		if isLink {
			columnName = gstr.CaseCamel(field.Name)
		}
		if field.FormMode == FMDateRange {
			tmp := "\tif " + linkMode + " {\n\t\tin." + field.GoName + "[1] = in." + field.GoName + "[1].EndOfDay()\n\t}\n"
			buffer.WriteString(tmp)
		}

		switch field.QueryWhere {
		case WMEq:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WMNeq:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "Not(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WMGt:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "GT(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WMGte:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "GTE(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WMLt:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "LT(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WMLte:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "LTE(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WMIn:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "In(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WMNotIn:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "NotIn(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WMBetween:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "Between(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + "[0], in." + field.GoName + "[1])\n\t}"
		case WMNotBetween:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "NotBetween(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + "[0], in." + field.GoName + "[1])\n\t}"
		case WMLike:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "Like(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WMLikeAll:
			val := `"%"+in.` + field.GoName + `+"%"`
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "Like(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", " + val + ")\n\t}"
		case WMNotLike:
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "NotLike(" + tablePrefix + "dao." + daoName + ".Columns()." + columnName + ", in." + field.GoName + ")\n\t}"
		case WMJsonContains:
			val := tablePrefix + `"JSON_CONTAINS("+dao.` + daoName + `.Columns().` + columnName + `+",?)", in.` + field.GoName
			whereTag = "\tif " + linkMode + " {\n\t\tmod = mod." + wherePrefix + "(" + val + ")\n\t}"

		default:
			buffer.WriteString(fmt.Sprintf(LogicWhereNoSupport, field.QueryWhere))
		}

		buffer.WriteString(whereTag + "\n")
	}
}
