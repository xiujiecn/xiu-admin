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
	"strings"
	"xiuadmin/internal/dao"
	genmodel "xiuadmin/internal/library/xgen/gen_model"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
)

const (
	InputTypeListInp          = 1 // 列表输入
	InputTypeListModel        = 2 // 列表输出
	InputTypeExportModel      = 3 // 列表导出
	InputTypeEditInpValidator = 4 // 添加&编辑验证器
	InputTypeUpdateFields     = 5 // 编辑修改过滤字段
	InputTypeInsertFields     = 6 // 编辑新增过滤字段
	InputTypeTreeOptionFields = 7 // 关系树查询字段
	EditInpValidatorGenerally = "if err := g.Validator().Rules(\"%s\").Data(in.%s).Messages(\"%s\").Run(ctx); err != nil {\n\t\treturn err.Current()\n\t}\n"
)

func (l *gCurd) inputTplData(ctx context.Context, in *genmodel.CurdPreviewParam) (data g.Map, err error) {
	data = make(g.Map)
	data["listInpColumns"] = l.genInputListColumns(ctx, in, InputTypeListInp)
	data["listModelColumns"] = l.genInputListColumns(ctx, in, InputTypeListModel)
	data["exportModelColumns"] = l.genInputListColumns(ctx, in, InputTypeExportModel)
	data["editInpValidator"] = l.genInputListColumns(ctx, in, InputTypeEditInpValidator)
	data["updateFieldsColumns"] = l.genInputListColumns(ctx, in, InputTypeUpdateFields)
	data["insertFieldsColumns"] = l.genInputListColumns(ctx, in, InputTypeInsertFields)
	data["viewModelColumns"] = l.genInputViewColumns(ctx, in)
	if in.Options.Step.IsTreeTable {
		data["treeOptionFields"] = l.genInputListColumns(ctx, in, InputTypeTreeOptionFields)
	}
	return
}

func (l *gCurd) genInputViewColumns(ctx context.Context, in *genmodel.CurdPreviewParam) string {
	buffer := bytes.NewBuffer(nil)

	index := 0
	array := make([][]string, 1000)
	// 主表
	for _, field := range in.MasterFields {
		// 查询用户摘要
		if field.IsList && in.Options.Step.HasHookMemberSummary && IsMemberSummaryField(field.Name) {
			tagKey := "`"
			descriptionTag := gstr.Replace(formatComment(field.Dc)+"摘要信息", `"`, `\"`)
			result := []string{"    #" + field.GoName + "Summa"}
			result = append(result, " #*hook.MemberSumma")
			result = append(result, " #"+fmt.Sprintf(tagKey+`json:"%s"`, field.TsName+"Summa"))
			result = append(result, " #"+fmt.Sprintf(`dc:"%s"`+tagKey, descriptionTag))
			array[index] = result
			index++
		}
	}
	table := tablewriter.NewTable(buffer)

	table.Options(tablewriter.WithRendition(tw.Rendition{
		Borders: tw.Border{
			Top:    tw.Fail,
			Bottom: tw.Fail,
			Left:   tw.Fail,
			Right:  tw.Fail,
		},
		Settings: tw.Settings{
			Lines: tw.Lines{
				ShowTop:        tw.Fail,
				ShowBottom:     tw.Fail,
				ShowHeaderLine: tw.Fail,
				ShowFooterLine: tw.Fail,
			},
		},
	}))
	// tw.SetBorder(false)
	// tw.SetRowLine(false)
	// tw.SetAutoWrapText(false)
	// tw.SetColumnSeparator("")
	table.Bulk(array)
	table.Render()
	stContent := buffer.String()
	// Let's do this hack of table writer for indent!
	stContent = gstr.Replace(stContent, "  #", "")
	stContent = gstr.Replace(stContent, "` ", "`")
	stContent = gstr.Replace(stContent, "``", "")
	stContent = removeEndWrap(stContent)

	buffer.Reset()
	buffer.WriteString(stContent)
	return "\tentity." + in.In.DaoName + "\n" + buffer.String()
}

func (l *gCurd) genInputListColumns(ctx context.Context, in *genmodel.CurdPreviewParam, inputType int) string {
	buffer := bytes.NewBuffer(nil)
	index := 0
	array := make([][]string, 1000)
	// 主表
	for _, field := range in.MasterFields {
		row := l.genStructFieldDefinition(in, field, inputType, true)
		if row == nil {
			continue
		}
		array[index] = row
		index++

		switch inputType {
		case InputTypeListModel:
			// 查询用户摘要
			if field.IsList && in.Options.Step.HasHookMemberSummary && IsMemberSummaryField(field.Name) {
				tagKey := "`"
				descriptionTag := gstr.Replace(formatComment(field.Dc)+"摘要信息", `"`, `\"`)
				result := []string{"    #" + field.GoName + "Summa"}
				result = append(result, " #*hook.MemberSumma")
				result = append(result, " #"+fmt.Sprintf(tagKey+`json:"%s"`, field.TsName+"Summa"))
				result = append(result, " #"+fmt.Sprintf(`dc:"%s"`+tagKey, descriptionTag))
				array[index] = result
				index++
			}
		}
	}

	// 关联表
	if len(in.Options.Join) > 0 {
		for _, v := range in.Options.Join {
			if !isEffectiveJoin(v) {
				g.Log().Infof(ctx, "genInputListColumns isEffectiveJoin: %+v", v)
				continue
			}
			g.Log().Infof(ctx, "genInputListColumns v: %+v", v)
			for _, field := range v.Columns {
				row := l.genStructFieldDefinition(in, field, inputType, false)
				g.Log().Infof(ctx, "genInputListColumns field: %+v row: %+v", field, row)
				if row != nil {
					array[index] = row
					index++
				}
			}
		}
	}
	g.Log().Infof(ctx, "genInputListColumns in.Options.Join: %+v", in.Options.Join)

	table := tablewriter.NewWriter(buffer)
	table.Options(tablewriter.WithRendition(tw.Rendition{
		Borders: tw.Border{
			Top:    tw.Fail,
			Bottom: tw.Fail,
			Left:   tw.Fail,
			Right:  tw.Fail,
		},
		Settings: tw.Settings{
			Lines: tw.Lines{
				ShowTop:        tw.Fail,
				ShowBottom:     tw.Fail,
				ShowHeaderLine: tw.Fail,
				ShowFooterLine: tw.Fail,
			},
		},
	}))
	// tw.SetBorder(false)
	// tw.SetRowLine(false)
	// tw.SetAutoWrapText(false)
	// tw.SetColumnSeparator("")
	table.Bulk(array)
	table.Render()
	stContent := buffer.String()
	// Let's do this hack of table writer for indent!
	stContent = gstr.Replace(stContent, "  #", "")
	stContent = gstr.Replace(stContent, "` ", "`")
	stContent = gstr.Replace(stContent, "``", "")
	stContent = removeEndWrap(stContent)

	buffer.Reset()
	buffer.WriteString(stContent)
	return buffer.String()
}

// genStructFieldForModel gens and returns the attribute definition for specified field.
func (l *gCurd) genStructFieldDefinition(in *genmodel.CurdPreviewParam, field *genmodel.GenCodesColumnListModel, inputType int, isMaster bool) []string {
	var (
		tagKey         = "`"
		result         = []string{"    #" + field.GoName}
		descriptionTag = gstr.Replace(formatComment(field.Dc), `"`, `\"`)
	)

	addResult := func() []string {
		result = append(result, " #"+field.GoType)
		result = append(result, " #"+fmt.Sprintf(tagKey+`json:"%s"`, field.TsName))
		result = append(result, " #"+fmt.Sprintf(`dc:"%s"`+tagKey, descriptionTag))
		return result
	}

	isUpdateOrInsertInputFields := func() bool {
		if !isMaster {
			return field.IsEdit
		}

		if field.IsEdit {
			return true
		}

		// 树表内部维护字段
		if in.Options.Step.IsTreeTable && gstr.InArray(defaultTreeFields, field.Name) {
			return true
		}

		// 修改人
		if inputType == InputTypeUpdateFields && field.GoName == "UpdatedBy" {
			return true
		}
		if inputType == InputTypeUpdateFields && field.GoName == "UpdatedAt" {
			return true
		}
		// 创建人
		if inputType == InputTypeInsertFields && field.GoName == "CreatedBy" {
			return true
		}
		if inputType == InputTypeInsertFields && field.GoName == "CreatedAt" {
			return true
		}
		if inputType == InputTypeInsertFields && field.GoName == "CreatedDept" {
			return true
		}
		if inputType == InputTypeInsertFields && field.GoName == "TenantId" {
			return true
		}
		return false
	}

	isQuery := false

	switch inputType {
	case InputTypeListInp:
		if in.Options.Step.IsTreeTable && IsPidName(field.Name) {
			isQuery = true
			field.QueryWhere = WMEq
		}
		if !field.IsQuery && !isQuery {
			return nil
		}

		if field.QueryWhere == WMBetween {
			result = append(result, " #[]"+field.GoType)
		} else {
			// 查询用户摘要时，固定接收字符串类型
			if field.IsQuery && in.Options.Step.HasQueryMemberSummary && IsMemberSummaryField(field.Name) {
				result = append(result, " #string")
			} else {
				result = append(result, " #"+field.GoType)
			}
		}
		result = append(result, " #"+fmt.Sprintf(tagKey+`json:"%s"`, field.TsName))
		result = append(result, " #"+fmt.Sprintf(`dc:"%s"`+tagKey, descriptionTag))

	case InputTypeListModel:
		// 主表的主键
		if IsIndexPK(field.Index) && isMaster {
			addResult()
			// 树表的pid字段
		} else if in.Options.Step.IsTreeTable && IsPidName(field.Name) {
			addResult()
		} else if field.IsList {
			addResult()
		} else {
			return nil
		}
	case InputTypeExportModel:
		if !field.IsExport {
			return nil
		}
		addResult()
	case InputTypeEditInpValidator:
		if !field.IsEdit {
			return nil
		}
		if !field.Required && (field.FormRole == "none" || field.FormRole == "") {
			return nil
		}
		rule := "// 验证" + field.Dc + "\n"
		if field.Required && (field.FormRole == FRNone || field.FormRole == "") {
			field.FormRole = "required"
		}
		if err, s := makeValidatorFunc(field); err != nil {
			return nil
		} else {
			rule += s
		}
		result = []string{rule}
	case InputTypeUpdateFields:
		if !isUpdateOrInsertInputFields() {
			return nil
		}
		addResult()
	case InputTypeInsertFields:
		if !isUpdateOrInsertInputFields() {
			return nil
		}
		addResult()
	case InputTypeTreeOptionFields:
		if IsIndexPK(field.Index) {
			return addResult()
		}
		if IsPidName(field.Name) {
			return addResult()
		}
		if in.Options.Tree.TitleColumn == field.Name {
			return addResult()
		}
		return nil
	default:
		panic("inputType is invalid")
	}
	return result
}

func makeValidatorFunc(field *genmodel.GenCodesColumnListModel) (err error, rule string) {
	if field.FormRole == "required" {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "required", field.GoName, field.Dc+"不能为空")
	} else if field.FormRole == FRIp {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "ip", field.GoName, field.Dc+"必须为IPV4或IPV6")
	} else if field.FormRole == FRPercentage {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "min:0|max:100", field.GoName, field.Dc+"必须0-100之间")
	} else if field.FormRole == FRTel {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "telephone", field.GoName, field.Dc+"不是座机号码")
	} else if field.FormRole == FRPhone {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "phone", field.GoName, field.Dc+"不是手机号码")
	} else if field.FormRole == FRQq {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "qq", field.GoName, field.Dc+"不是QQ号码")
	} else if field.FormRole == FREmail {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "email", field.GoName, field.Dc+"不是邮箱地址")
	} else if field.FormRole == FRIdCard {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "resident-id", field.GoName, field.Dc+"不是身份证号码")
	} else if field.FormRole == FRNum {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "min:1", field.GoName, field.Dc+"必须大于0")
	} else if field.FormRole == FRBankCard {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "bank-card", field.GoName, field.Dc+"不是银行卡号")
	} else if field.FormRole == FRWeibo {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "regex:^[0-9a-zA-Z\\u4e00-\\u9fa5_-]*$", field.GoName, field.Dc+"不是微博号")
	} else if field.FormRole == FRUserName {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "regex:^[0-9a-zA-Z]{6,16}$", field.GoName, field.Dc+"必须为6-16位由字母和数字组成")
	} else if field.FormRole == FRAccount {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "regex:^[\\w_\\d]{6,16}$", field.GoName, field.Dc+"必须为6-16位由字母、数字或下划线组成")
	} else if field.FormRole == FRPassword {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "regex:^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,18}$", field.GoName, field.Dc+"必须包含6-18为字母和数字")
	} else if field.FormRole == FRAmount {
		rule = fmt.Sprintf(EditInpValidatorGenerally, "regex:(^[0-9]{1,10}$)|(^[0-9]{1,10}[\\\\.]{1}[0-9]{1,2}$)", field.GoName, field.Dc+"最多允许输入10位整数及2位小数")
	} else {
		err = gerror.New("not support")
	}

	// 生成验证字典
	if field.DictType != "" {
		var (
			ctx         = context.Background()
			valueType   gdb.Value
			resultValue gdb.Result
		)
		valueType, err = dao.SysDictType.Ctx(ctx).Fields(dao.SysDictType.Columns().DictType).Where(dao.SysDictType.Columns().DictId, field.DictType).Value()
		if err != nil {
			return
		}
		resultValue, err = dao.SysDictData.Ctx(context.Background()).Fields(dao.SysDictData.Columns().DictValue).Where(dao.SysDictData.Columns().DictType, valueType.String()).All()
		if err != nil {
			return
		}
		if resultValue.Len() > 0 {
			names := make([]string, 0)
			for _, item := range resultValue {
				names = append(names, item["value"].String())
			}
			dictRule := "in:" + strings.Join(names, ",")
			rule += fmt.Sprintf(EditInpValidatorGenerally, dictRule, field.GoName, field.Dc+"值不正确")
			err = nil
		}
	}
	return
}
