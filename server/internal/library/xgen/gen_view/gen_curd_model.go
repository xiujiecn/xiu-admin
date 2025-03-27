// package genview
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package genview

import (
	"bytes"
	"context"
	"fmt"
	genmodel "xiuadmin/internal/library/xgen/gen_model"
	"xiuadmin/utility"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type StateItem struct {
	Name         string
	DefaultValue interface{}
	Dc           string
}

func (l *gCurd) webModelTplData(ctx context.Context, in *genmodel.CurdPreviewParam) (data g.Map, err error) {
	data = make(g.Map)
	data["stateItems"] = l.genWebModelStateItems(ctx, in)
	data["rules"] = l.genWebModelRules(ctx, in)
	data["formSchema"] = l.genWebModelFormSchema(ctx, in)
	if data["columns"], err = l.genWebModelColumns(ctx, in); err != nil {
		return nil, err
	}
	if data["columnsInterface"], err = l.genWebModelColumnsInterface(ctx, in); err != nil {
		return nil, err
	}
	data["viewSchema"] = l.genWebModelViewSchema(ctx, in)
	data["editSchema"] = l.genWebModelEditSchema(ctx, in)
	// 根据表单生成情况，按需导包
	data["import"], data["const"] = l.genWebModelImport(ctx, in)
	return
}

func (l *gCurd) genWebModelImport(ctx context.Context, in *genmodel.CurdPreviewParam) (string, string) {
	importBuffer := bytes.NewBuffer(nil)
	constBuffer := bytes.NewBuffer(nil)

	importBuffer.WriteString(`import { h, ref } from 'vue';
import { Tag } from 'ant-design-vue';
import type { VxeGridProps } from '#/adapter/vxe-table';
import { getPopupContainer } from '@vben/utils';
import type { DescItem } from '#/components/description';
`)

	// 导入基础组件
	if len(in.Options.Step.ImportModel.NaiveUI) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(in.Options.Step.ImportModel.NaiveUI) + " from 'ant-design-vue';\n")
	}

	importBuffer.WriteString("import { cloneDeep } from 'lodash-es';\n")

	// 导入表单搜索
	if in.Options.Step.HasSearchForm {
		importBuffer.WriteString("import type { VbenFormSchema } from '@vben/common-ui';\n")
	}

	// 导入工具类
	if len(in.Options.Step.ImportModel.UtilsIs) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(in.Options.Step.ImportModel.UtilsIs) + " from '@/utils/is';\n")
	}

	if len(in.Options.Step.ImportModel.UtilsUrl) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(in.Options.Step.ImportModel.UtilsUrl) + " from '@/utils/urlUtils';\n")
	}

	if len(in.Options.Step.ImportModel.UtilsDate) > 0 {
		// importBuffer.WriteString("import " + ImportWebMethod(in.Options.Step.ImportModel.UtilsDate) + " from '@/utils/dateUtil';\n")
		importBuffer.WriteString("import dayjs from 'dayjs';\n")
	}

	if in.Options.Step.HasRulesValidator {
		importBuffer.WriteString("import { z } from '@vben/common-ui';\n")
	}

	if len(in.Options.Step.ImportModel.UtilsIndex) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(in.Options.Step.ImportModel.UtilsIndex) + " from '#/utils';\n")
	}

	// 导入api
	var importApiMethod []string
	if in.Options.Step.HasSwitch {
		importApiMethod = append(importApiMethod, "Switch")
	}
	if in.Options.Step.IsTreeTable {
		importApiMethod = append(importApiMethod, "TreeOption")
	}
	if len(importApiMethod) > 0 {
		importBuffer.WriteString("import " + ImportWebMethod(importApiMethod) + " from '" + in.Options.ImportWebApi + "';\n")
	}

	// 导入字典选项
	if in.Options.DictOps.Has {
		importBuffer.WriteString(`import { DictEnum } from '@vben/constants';
import { getDictOptions } from '#/utils/dict';`)
		constBuffer.WriteString("\n")
	}

	if in.Options.Step.HasSwitch {
		importBuffer.WriteString("import { usePermission } from '@/hooks/web/usePermission';\n")
		constBuffer.WriteString("const { hasPermission } = usePermission();\n")
		constBuffer.WriteString("const $message = window['$message'];\n")
	}

	return importBuffer.String(), constBuffer.String()
}

func (l *gCurd) genWebModelStateItems(ctx context.Context, in *genmodel.CurdPreviewParam) (items []*StateItem) {
	for _, field := range in.MasterFields {
		var value = field.DefaultValue
		if value == nil {
			value = "null"
		}
		if value == "" {
			value = `''`
		}
		if value == "NULL" {
			value = "null"
		}

		// 选项组件默认值调整
		if gconv.Int(value) == 0 && IsSelectFormModel(field.FormMode) {
			value = "null"
		}

		if field.Name == "status" {
			value = 0
		}
		if field.FormMode == FMSwitch {
			value = 0
		}
		if field.FormMode == FMInputDynamic {
			value = "[]"
		}
		items = append(items, &StateItem{
			Name:         field.TsName,
			DefaultValue: value,
			Dc:           field.Dc,
		})

		// 查询用户摘要
		if field.IsList && in.Options.Step.HasHookMemberSummary && IsMemberSummaryField(field.Name) {
			items = append(items, &StateItem{
				Name:         field.TsName + "Summa?: null | MemberSumma",
				DefaultValue: "null",
				Dc:           field.Dc + "摘要信息",
			})
		}
	}
	return
}

func (l *gCurd) genWebModelDictOptions(ctx context.Context, in *genmodel.CurdPreviewParam) error {
	type DictType struct {
		DictType string `json:"dictType"`
		DictName string `json:"dictName"`
	}

	var (
		dictTypeIds         []string
		dictTypeList        []*DictType
		builtinDictTypeIds  []int64
		builtinDictTypeList []*DictType
	)

	for _, field := range in.MasterFields {
		if field.DictType != "" {
			dictTypeIds = append(dictTypeIds, field.DictType)
		}

		// if field.DictType < 0 {
		// 	builtinDictTypeIds = append(builtinDictTypeIds, field.DictType)
		// }
	}

	dictTypeIds = utility.UniqueSlice(dictTypeIds)
	builtinDictTypeIds = utility.UniqueSlice(builtinDictTypeIds)

	if len(dictTypeIds) == 0 && len(builtinDictTypeIds) == 0 {
		return nil
	}

	if len(dictTypeIds) > 0 {
		err := g.Model("sys_dict_type").Ctx(ctx).
			Fields("dict_type", "dict_name").
			WhereIn("dict_type", dictTypeIds).
			Scan(&dictTypeList)
		if err != nil {
			return err
		}
	}

	// if len(builtinDictTypeIds) > 0 {
	// 	for _, id := range builtinDictTypeIds {
	// 		typ, err := dict.GetTypeById(ctx, id)
	// 		if err != nil && !errors.Is(err, dict.NotExistKeyError) {
	// 			return err
	// 		}
	// 		if len(typ) > 0 {
	// 			row := new(DictType)
	// 			row.Id = id
	// 			row.Type = typ
	// 			builtinDictTypeList = append(builtinDictTypeList, row)
	// 		}
	// 	}
	// }

	if len(dictTypeList) == 0 && len(builtinDictTypeList) == 0 {
		return nil
	}

	if len(builtinDictTypeList) > 0 {
		dictTypeList = append(dictTypeList, builtinDictTypeList...)
	}

	in.Options.DictOps.Has = true

	for _, v := range dictTypeList {
		// 字段映射字典
		for _, field := range in.MasterFields {
			if field.DictType != "" && v.DictType == field.DictType {
				in.Options.DictMap[field.TsName] = v.DictType
				in.Options.DictOps.Schemas = append(in.Options.DictOps.Schemas, &genmodel.OptionsSchemasField{
					Field: field.TsName,
					Type:  v.DictType,
				})
			}
		}
		in.Options.DictOps.Types = append(in.Options.DictOps.Types, v.DictName)
	}
	return nil
}

func (l *gCurd) genWebModelRules(ctx context.Context, in *genmodel.CurdPreviewParam) string {
	// buffer := bytes.NewBuffer(nil)
	// buffer.WriteString("export const rules = {\n")
	// for _, field := range in.MasterFields {
	// 	if !field.IsEdit || (!field.Required && (field.FormRole == "" || field.FormRole == FRNone)) {
	// 		continue
	// 	}

	// 	in.Options.Step.HasRules = true
	// 	if field.FormRole == "" || field.FormRole == FRNone || field.FormRole == "required" {
	// 		buffer.WriteString(fmt.Sprintf("  %s: {\n    required: %v,\n    trigger: ['blur', 'input'],\n    type: '%s',\n    message: '请输入%s',\n  },\n", field.TsName, field.Required, field.TsType, field.Dc))
	// 	} else {
	// 		in.Options.Step.HasRulesValidator = true
	// 		buffer.WriteString(fmt.Sprintf("  %s: {\n    required: %v,\n    trigger: ['blur', 'input'],\n    type: '%s',\n    validator: validate.%v,\n  },\n", field.TsName, field.Required, field.TsType, field.FormRole))
	// 	}
	// }
	// buffer.WriteString("};\n")
	// return buffer.String()
	return ""
}

func (l *gCurd) genWebModelFormSchema(ctx context.Context, in *genmodel.CurdPreviewParam) string {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("export const querySchema: VbenFormSchema[] = [\n")

	// 主表
	l.genWebModelFormSchemaEach(buffer, in.MasterFields, in, false)

	// 关联表
	if len(in.Options.Join) > 0 {
		for _, v := range in.Options.Join {
			if !isEffectiveJoin(v) {
				continue
			}
			l.genWebModelFormSchemaEach(buffer, v.Columns, in, false)
		}
	}

	buffer.WriteString("];\n")
	return buffer.String()
}

// genWebDictOption 生产字典选项
func (l *gCurd) genWebDictOption(typ any) string {
	return fmt.Sprintf(`getDictOptions('%v')`, typ)
}

func (l *gCurd) genWebModelFormSchemaEach(buffer *bytes.Buffer, fields []*genmodel.GenCodesColumnListModel, in *genmodel.CurdPreviewParam, isEdit bool) {
	for _, field := range fields {
		if isEdit {
			if !field.IsEdit && field.Name != in.Pk.Name {
				continue
			}
			if field.Name == in.Pk.Name {
				buffer.WriteString(fmt.Sprintf("  {\n    fieldName: '%s',\n    component: '%s',\n    label: '%s',\n    dependencies: {   show: () => false,    triggerFields: [''],   },\n    componentProps: {\n      placeholder: '',\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  },\n", field.TsName, "Input", field.Dc))
				continue
			}
		} else {
			if !field.IsQuery {
				continue
			}
		}
		rules := "null"
		if field.IsEdit {
			if field.Required {
				if field.FormMode == FMInputNumber {
					rules = fmt.Sprintf("z.number({required_error: '请输入%s', invalid_type_error: '无效数字'})", field.Dc)
					in.Options.Step.HasRulesValidator = true
				} else if field.FormMode == FMSelect || field.FormMode == FMSelectMultiple || field.FormMode == FMCheckbox || field.FormMode == FMDate || field.FormMode == FMDateRange || field.FormMode == FMTime || field.FormMode == FMTimeRange {
					rules = "'selectRequired'"
				} else {
					rules = "'required'"
				}
			}
		}
		if field.FormRole != "" {
			in.Options.Step.HasRulesValidator = true
			if field.FormRole == FRNum {
				rules = fmt.Sprintf("z.number({required_error: '请输入%s', invalid_type_error: '无效数字'}	)", field.Dc)
			}
			if field.FormRole == FRIp {
				rules = "z.string().ip({message: '无效IP格式'})"
			}
			if field.FormRole == FRPhone {
				rules = "z.string().regex(/^1[3-9]\\d{9}$/, '无效手机号格式')"
			}
			if field.FormRole == FRTel {
				rules = "z.string().regex(/^(?:\\+?\\d{1,3}[-.\\s]?)?(?:$?\\d{1,4}$?[-.\\s]?)?\\d{1,4}[-.\\s]?\\d{1,4}[-.\\s]?\\d{1,9}$/, '无效手机号格式')"
			}
			if field.FormRole == FRPercentage {
				rules = fmt.Sprintf("z.number({required_error: '请输入%s', invalid_type_error: '无效百分比'}).min(0).max(100)", field.Dc)
			}
			if field.FormRole == FRQq {
				rules = "z.string().regex(/^[1-9]\\d{4,}$/, '无效QQ号格式')"
			}
			if field.FormRole == FRAmount {
				rules = "z.number({required_error: '必填', invalid_type_error: '无效金额'}	).min(0)"
			}
			if field.FormRole == FRBankCard {
				rules = "z.string().regex(/^[1-9]\\d{4,}$/, '无效银行卡号格式')"
			}
			if field.FormRole == FRWeibo {
				rules = "z.string().regex(/^[1-9]\\d{4,}$/, '无效微博号格式')"
			}
			if field.FormRole == FRUserName {
				rules = "z.string().regex(/^[a-zA-Z0-9_]+$/, '无效用户名格式')"
			}
			if field.FormRole == FRPassword {
				rules = "z.string().regex(/^[a-zA-Z0-9_]+$/, '无效密码格式')"
			}
			if field.FormRole == FRAccount {
				rules = "z.string().regex(/^[a-zA-Z0-9_]+$/, '无效账号格式')"
			}
			if field.FormRole == FRIdCard {
				rules = "z.string().regex(/^[1-9]\\d{4,}$/, '无效身份证号格式')"
			}
			if field.FormRole == FREmail {
				rules = "z.string().email('无效邮箱格式')"
			}
		}

		in.Options.Step.HasSearchForm = true

		// 查询用户摘要
		if field.IsQuery && in.Options.Step.HasQueryMemberSummary && IsMemberSummaryField(field.Name) {
			buffer.WriteString(fmt.Sprintf("  {\n    fieldName: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      placeholder: '请输入ID|用户名|姓名|手机号',\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  rules:%v\n},\n", field.TsName, "Input", field.Dc, rules))
			continue
		}

		var (
			defaultComponent = fmt.Sprintf("  {\n    fieldName: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      placeholder: '请输入%s',\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  rules:%v\n},\n", field.TsName, "Input", field.Dc, field.Dc, rules)
			component        string
		)

		// 这里根据编辑表单组件来进行推断，如果没有则使用默认input，这可能会导致和查询条件所需参数不符的情况
		switch field.FormMode {
		case FMInput, FMInputTextarea, FMInputEditor:
			component = defaultComponent

		case FMInputNumber:
			component = fmt.Sprintf("  {\n    fieldName: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      placeholder: '请输入%s',\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  rules:%v\n},\n", field.TsName, "InputNumber", field.Dc, field.Dc, rules)

		case FMDate:
			component = fmt.Sprintf("  {\n    fieldName: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      type: '%s',\n      clearable: true,\n      shortcuts: %s,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  rules:%v\n},\n", field.TsName, "RangePicker", field.Dc, "date", "'FMDate'", rules)
			in.Options.Step.ImportModel.UtilsDate = append(in.Options.Step.ImportModel.UtilsDate, "defShortcuts")

		case FMDateRange:
			component = fmt.Sprintf("  {\n    fieldName: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      type: '%s',\n      clearable: true,\n      valueFormat: %s,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  rules:%v\n},\n", field.TsName, "RangePicker", field.Dc, "daterange", "'YYYY-MM-DD HH:mm:ss'", rules)
			// in.Options.Step.ImportModel.UtilsDate = append(in.Options.Step.ImportModel.UtilsDate, "defRangeShortcuts")

		case FMTime:
			component = fmt.Sprintf("  {\n    fieldName: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      type: '%s',\n      clearable: true,\n      shortcuts: %s,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  rules:%v\n},\n", field.TsName, "DatePicker", field.Dc, "datetime", "'FMTime'", rules)
			in.Options.Step.ImportModel.UtilsDate = append(in.Options.Step.ImportModel.UtilsDate, "defShortcuts")

		case FMTimeRange:
			component = fmt.Sprintf("  {\n    fieldName: '%s',\n    component: '%s',\n    label: '%s',\n    componentProps: {\n      type: '%s',\n      clearable: true,\n      shortcuts: %s,\n      onUpdateValue: (e: any) => {\n        console.log(e);\n      },\n    },\n  rules:%v\n},\n", field.TsName, "RangePicker", field.Dc, "datetimerange", "'FMTimeRange'", rules)
			in.Options.Step.ImportModel.UtilsDate = append(in.Options.Step.ImportModel.UtilsDate, "defRangeShortcuts")

		case FMSwitch:
			fallthrough
		case FMRadio:
			component = fmt.Sprintf("  {\n    fieldName: '%s',\n    component: '%s',\n    label: '%s',\n    giProps: {\n      //span: 24,\n    },\n    componentProps: {\n      options: %v,\n      onUpdateChecked: (e: any) => {\n        console.log(e);\n      },\n    },\n  rules:%v\n},\n", field.TsName, "RadioGroup", field.Dc, l.genWebDictOption(in.Options.DictMap[field.TsName]), rules)

		case FMCheckbox:
			component = fmt.Sprintf("  {\n    fieldName: '%s',\n    component: '%s',\n    label: '%s',\n    giProps: {\n      span: 1,\n    },\n    componentProps: {\n      placeholder: '请选择%s',\n      options: %v,\n      onUpdateChecked: (e: any) => {\n        console.log(e);\n      },\n    },\n  rules:%v\n},\n", field.TsName, "NCheckbox", field.Dc, field.Dc, l.genWebDictOption(in.Options.DictMap[field.TsName]), rules)

		case FMSelect:
			component = fmt.Sprintf(`  {    
			fieldName: '%s',    
			component: '%s',    
			label: '%s',    
			defaultValue: null,    
			componentProps: {    
				placeholder: '请选择%s',    
				options: %v,    
				onUpdateValue: (e: any) => {    
					console.log(e);    
				},  
			},
			rules:%v
		},
		`, field.TsName, "Select", field.Dc, field.Dc, l.genWebDictOption(in.Options.DictMap[field.TsName]), rules)

		case FMSelectMultiple:
			component = fmt.Sprintf(`  {    
			fieldName: '%s',    
			component: '%s',    
			label: '%s',    
			defaultValue: null,    
			componentProps: {    
				mode: 'multiple',
				placeholder: '请选择%s',    
				options: %v,    
				onUpdateValue: (e: any) => {    
					console.log(e);    
				},  	
			},
			rules:%v
		},
		`, field.TsName, "Select", field.Dc, field.Dc, l.genWebDictOption(in.Options.DictMap[field.TsName]), rules)

		default:
			component = defaultComponent
		}

		buffer.WriteString(component)
	}
}

func (l *gCurd) genWebModelColumns(ctx context.Context, in *genmodel.CurdPreviewParam) (string, error) {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("export const columns: VxeGridProps['columns'] = [\n")

	// 主表
	if err := l.genWebModelColumnsEach(buffer, in, in.MasterFields); err != nil {
		return "", err
	}

	// 关联表
	if len(in.Options.Join) > 0 {
		for _, v := range in.Options.Join {
			if !isEffectiveJoin(v) {
				continue
			}
			if err := l.genWebModelColumnsEach(buffer, in, v.Columns); err != nil {
				return "", err
			}
		}
	}
	buffer.WriteString("  { title: '操作', width: 120, slots: { default: 'action' } },\n")
	buffer.WriteString("];\n")
	return buffer.String(), nil
}

func (l *gCurd) genWebModelColumnsInterface(ctx context.Context, in *genmodel.CurdPreviewParam) (string, error) {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("export interface RowType {\n")

	// 主表
	for _, field := range in.MasterFields {
		buffer.WriteString(fmt.Sprintf("  %s: %s;\n", field.TsName, field.TsType))
	}

	// 关联表
	if len(in.Options.Join) > 0 {
		for _, v := range in.Options.Join {
			if !isEffectiveJoin(v) {
				continue
			}
			for _, field := range v.Columns {
				buffer.WriteString(fmt.Sprintf("  %s: %s;\n", field.TsName, field.TsType))
			}
		}
	}
	buffer.WriteString("};\n")
	return buffer.String(), nil
}

func (l *gCurd) genWebModelViewSchema(ctx context.Context, in *genmodel.CurdPreviewParam) string {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("export const viewSchema: DescItem[] = [\n")

	// 主表
	for _, field := range in.MasterFields {
		l.genWebModelViewSchemaEach(buffer, []*genmodel.GenCodesColumnListModel{field}, in)
	}

	// 关联表
	if len(in.Options.Join) > 0 {
		for _, v := range in.Options.Join {
			if !isEffectiveJoin(v) {
				continue
			}
			for _, field := range v.Columns {
				l.genWebModelViewSchemaEach(buffer, []*genmodel.GenCodesColumnListModel{field}, in)
			}
		}
	}
	buffer.WriteString("];\n")
	return buffer.String()
}

func (l *gCurd) genWebModelEditSchema(ctx context.Context, in *genmodel.CurdPreviewParam) string {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("export const editSchema: VbenFormSchema[] = [\n")

	// 主表
	for _, field := range in.MasterFields {

		l.genWebModelFormSchemaEach(buffer, []*genmodel.GenCodesColumnListModel{field}, in, true)
	}

	// 关联表
	if len(in.Options.Join) > 0 {
		for _, v := range in.Options.Join {
			if !isEffectiveJoin(v) {
				continue
			}
			for _, field := range v.Columns {
				l.genWebModelFormSchemaEach(buffer, []*genmodel.GenCodesColumnListModel{field}, in, true)
			}
		}
	}
	buffer.WriteString("];\n")
	return buffer.String()
}

func (l *gCurd) genWebModelColumnsEach(buffer *bytes.Buffer, in *genmodel.CurdPreviewParam, fields []*genmodel.GenCodesColumnListModel) (err error) {
	for _, field := range fields {
		if !field.IsList {
			continue
		}

		if field.Name == in.Pk.Name {
			buffer.WriteString(fmt.Sprintf("  {\n    title: '%s',\n    field: '%s',\n    align: '%v',\n    width: %v,\n    type: 'checkbox',\n  },\n", field.Dc, field.TsName, field.Align, field.Width))
			continue
		}
		var (
			defaultComponent = fmt.Sprintf("  {\n    title: '%s',\n    field: '%s',\n    align: '%v',\n    width: %v,\n },\n", field.Dc, field.TsName, field.Align, field.Width)
			component        string
		)

		// 查询用户摘要
		if in.Options.Step.HasHookMemberSummary && IsMemberSummaryField(field.Name) {
			buffer.WriteString(fmt.Sprintf("  {\n    title: '%v',\n    field: '%v',\n    align: '%v',\n    width: %v,\n    slots: {\n      default: ({ row }) =>  {\n      return renderPopoverMemberSumma(row.%vSumma);\n    },\n },\n },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName))
			in.Options.Step.ImportModel.UtilsIndex = append(in.Options.Step.ImportModel.UtilsIndex, []string{"renderPopoverMemberSumma", "type MemberSumma"}...)
			continue
		}

		// 这里根据编辑表单组件来进行推断，如果没有则使用默认input，这可能会导致和查询条件所需参数不符的情况
		switch field.FormMode {
		case FMDate:
			component = fmt.Sprintf("  {\n    title: '%s',\n    field: '%s',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      return formatDateTime(row.%s);\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName)
			in.Options.Step.ImportModel.UtilsDate = append(in.Options.Step.ImportModel.UtilsDate, "formatToDate")

		case FMRadio:
			fallthrough
		case FMSelect:

			if g.IsEmpty(in.Options.DictMap[field.TsName]) {
				g.Log().Infof(context.Background(), "server/internal/library/xgen/gen_curd_model.go genWebModelColumnsEach in.Options.DictMap:%+v", in.Options.DictMap)
				err = gerror.Newf("设置单选下拉框选项时，必须选择字典类型，字段名称:%v", field.Name)
				return
			}

			component = fmt.Sprintf(` {    
				title: '%s',    field: '%s',    align: '%v',    width: %v, 
				slots: {
      				default: ({ row }) => {
						return renderDict(row.%v, '%v');
					}
				},
			},
			`,
				field.Dc, field.TsName, field.Align, field.Width, field.TsName, in.Options.DictMap[field.TsName])
			in.Options.Step.ImportModel.UtilsIndex = append(in.Options.Step.ImportModel.UtilsIndex, "renderDict")

		case FMSelectMultiple:
			if g.IsEmpty(in.Options.DictMap[field.TsName]) {
				err = gerror.Newf("设置多选下拉框选项时，必须选择字典类型，字段名称:%v", field.Name)
				return
			}
			component = fmt.Sprintf("  {\n    title: '%s',\n    field: '%s',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      if (isNullObject(row.%s) || !isArray(row.%s)) {\n        return ``;\n      }\n      return row.%s.map((tagKey) => {\n        return renderOptionTag('%s', row.tagKey)\n      });\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName, field.TsName, field.TsName, in.Options.DictMap[field.TsName])
			in.Options.Step.ImportModel.NaiveUI = append(in.Options.Step.ImportModel.NaiveUI, "NTag")
			in.Options.Step.ImportModel.UtilsIs = append(in.Options.Step.ImportModel.UtilsIs, "isNullObject")

		case FMUploadImage:
			component = fmt.Sprintf("  {\n    title: '%s',\n    field: '%s',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      return renderImage(row.%v);\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName)
			in.Options.Step.ImportModel.UtilsIndex = append(in.Options.Step.ImportModel.UtilsIndex, "renderImage")

		case FMUploadImages:
			component = fmt.Sprintf("  {\n    title: '%s',\n    field: '%s',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      return renderImageGroup(row.%v);\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName)
			in.Options.Step.ImportModel.UtilsIndex = append(in.Options.Step.ImportModel.UtilsIndex, "renderImageGroup")

		case FMUploadFile:
			component = fmt.Sprintf("  {\n    title: '%s',\n    field: '%s',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      return renderFile(row.%v);\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName)
			in.Options.Step.ImportModel.UtilsIndex = append(in.Options.Step.ImportModel.UtilsIndex, "renderFile")

		case FMUploadFiles:
			component = fmt.Sprintf("  {\n    title: '%s',\n    field: '%s',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      return renderFileGroup(row.%v);\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, field.TsName)
			in.Options.Step.ImportModel.UtilsIndex = append(in.Options.Step.ImportModel.UtilsIndex, "renderFileGroup")

		case FMSwitch:
			component = fmt.Sprintf("  {\n    title: '%s',\n    field: '%s',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      return h(%s, {\n        value: row.%s === 1,\n        checked: '开启',\n        unchecked: '关闭',\n        disabled: !hasPermission(['%s']),\n        onUpdateValue: function (e) {\n          console.log('onUpdateValue e:' + JSON.stringify(e));\n          row.%s = e ? 1 : 2;\n          Switch({ %s: row.%s, key: '%s', value: row.%s }).then((_res) => {\n            $message.success('操作成功');\n          });\n        },\n      });\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, "NSwitch", field.TsName, "/"+in.Options.ApiPrefix+"/switch", field.TsName, in.Pk.TsName, in.Pk.TsName, CamelCaseToUnderline(field.TsName), field.TsName)
			in.Options.Step.ImportModel.NaiveUI = append(in.Options.Step.ImportModel.NaiveUI, "NSwitch")

		case FMRate:
			component = fmt.Sprintf("  {\n    title: '%s',\n    field: '%s',\n    align: '%v',\n    width: %v,\n    render(row: State) {\n      return h(%s, {\n        allowHalf: true,\n        readonly: true,\n        defaultValue: row.%s,\n      });\n    },\n  },\n", field.Dc, field.TsName, field.Align, field.Width, "NRate", field.TsName)
			in.Options.Step.ImportModel.NaiveUI = append(in.Options.Step.ImportModel.NaiveUI, "NRate")

		default:
			component = defaultComponent
		}

		buffer.WriteString(component)
	}
	return
}

func (l *gCurd) genWebModelViewSchemaEach(buffer *bytes.Buffer, fields []*genmodel.GenCodesColumnListModel, in *genmodel.CurdPreviewParam) {
	for _, field := range fields {

		if field.Name == in.Pk.Name {
			tmp := fmt.Sprintf("  {  field: '%s',  label: '%s'},\n", field.TsName, field.Dc)
			buffer.WriteString(tmp)
			continue
		}

		// 查询用户摘要
		if field.IsQuery && in.Options.Step.HasQueryMemberSummary && IsMemberSummaryField(field.Name) {
			tmp := fmt.Sprintf("  {  field: '%s',  label: '%s'},\n", field.TsName, field.Dc)
			buffer.WriteString(tmp)
			continue
		}

		var (
			defaultComponent = fmt.Sprintf("  {  field: '%s',  label: '%s'},\n", field.TsName, field.Dc)
			component        string
		)

		// 这里根据编辑表单组件来进行推断，如果没有则使用默认input，这可能会导致和查询条件所需参数不符的情况
		switch field.FormMode {
		case FMInput, FMInputTextarea, FMInputEditor:
			component = defaultComponent

		case FMInputNumber:
			component = fmt.Sprintf("  {  field: '%s',  label: '%s'},\n", field.TsName, field.Dc)

		case FMDate:
			component = fmt.Sprintf("  {  field: '%s',  label: '%s'},\n", field.TsName, field.Dc)

		case FMDateRange:
			component = fmt.Sprintf("  {  field: '%s',  label: '%s'},\n", field.TsName, field.Dc)

		case FMTime:
			component = fmt.Sprintf("  {  field: '%s',  label: '%s'},\n", field.TsName, field.Dc)

		case FMTimeRange:
			component = fmt.Sprintf("  {  field: '%s',  label: '%s'},\n", field.TsName, field.Dc)

		case FMSwitch:
			fallthrough
		case FMRadio:
			component = fmt.Sprintf(`  {
				field: '%s',
				label: '%s',
				render(row: any) {
					return renderDict(row.%s, '%v');
				},
			},
			`, field.TsName, field.Dc, field.TsName, in.Options.DictMap[field.TsName])
			in.Options.Step.ImportModel.UtilsIndex = append(in.Options.Step.ImportModel.UtilsIndex, "renderDict")

		case FMCheckbox:
			component = fmt.Sprintf(`  {
				field: '%s',
				label: '%s',
				render(row: any) {
					return renderDictTags(row.%s, '%v');
				},
			},
			`, field.TsName, field.Dc, field.TsName, in.Options.DictMap[field.TsName])
			in.Options.Step.ImportModel.UtilsIndex = append(in.Options.Step.ImportModel.UtilsIndex, "renderDictTags")

		case FMSelect:
			component = fmt.Sprintf(`  {
				field: '%s',
				label: '%s',
				render(row: any) {
					return renderDict(row.%s, '%v');
				},
			},
			`, field.TsName, field.Dc, field.TsName, in.Options.DictMap[field.TsName])
			in.Options.Step.ImportModel.UtilsIndex = append(in.Options.Step.ImportModel.UtilsIndex, "renderDict")

		case FMSelectMultiple:
			component = fmt.Sprintf(`  {
				field: '%s',
				label: '%s',
				render(row: any) {
					return renderDictTags(row.%s, '%v');
				},
			},
			`, field.TsName, field.Dc, field.TsName, in.Options.DictMap[field.TsName])
			in.Options.Step.ImportModel.UtilsIndex = append(in.Options.Step.ImportModel.UtilsIndex, "renderDictTags")

		default:
			component = defaultComponent
		}

		buffer.WriteString(component)
	}
}
