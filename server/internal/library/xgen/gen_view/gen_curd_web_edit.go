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

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

func (l *gCurd) webEditTplData(ctx context.Context, in *genmodel.CurdPreviewParam) (data g.Map, err error) {
	data = make(g.Map)
	data["formItem"] = l.genWebEditFormItem(ctx, in)
	data["script"] = l.genWebEditScript(ctx, in)
	return
}

func (l *gCurd) genWebEditFormItem(ctx context.Context, in *genmodel.CurdPreviewParam) string {
	buffer := bytes.NewBuffer(nil)
	for _, field := range in.MasterFields {
		if !field.IsEdit {
			continue
		}

		if IsIndexPK(field.Index) {
			continue
		}

		var (
			defaultComponent = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n          <n-input placeholder=\"请输入%s\" v-model:value=\"formValue.%s\" />\n          </n-form-item>", field.Dc, field.TsName, field.Dc, field.TsName)
			component        string
		)

		if in.Options.Step.IsTreeTable && IsPidName(field.Name) {
			field.FormMode = FMPidTreeSelect
		}

		switch field.FormMode {
		case FMInput:
			component = defaultComponent

		case FMInputNumber:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <n-input-number placeholder=\"请输入%s\" v-model:value=\"formValue.%s\" />\n          </n-form-item>", field.Dc, field.TsName, field.Dc, field.TsName)

		case FMInputTextarea:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <n-input type=\"textarea\" placeholder=\"%s\" v-model:value=\"formValue.%s\" />\n          </n-form-item>", field.Dc, field.TsName, field.Dc, field.TsName)

		case FMInputEditor:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <Editor style=\"height: 450px\" id=\"%s\" v-model:value=\"formValue.%s\" />\n          </n-form-item>", field.Dc, field.TsName, field.TsName, field.TsName)

		case FMInputDynamic:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <n-dynamic-input\n            v-model:value=\"formValue.%s\"\n            preset=\"pair\"\n            key-placeholder=\"键名\"\n            value-placeholder=\"键值\"\n          />\n          </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FMDate:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <DatePicker v-model:formValue=\"formValue.%s\" type=\"date\" />\n          </n-form-item>", field.Dc, field.TsName, field.TsName)

		// case FMDateRange:  // 必须要有两个字段，后面优化下

		case FMTime:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <DatePicker v-model:formValue=\"formValue.%s\" type=\"datetime\" />\n          </n-form-item>", field.Dc, field.TsName, field.TsName)

		// case FMTimeRange: // 必须要有两个字段，后面优化下

		case FMRadio:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <n-radio-group v-model:value=\"formValue.%s\" name=\"%s\">\n            <n-radio-button\n              v-for=\"%s in dict.getOptionUnRef('%s')\"\n              :key=\"%s.value\"\n              :value=\"%s.value\"\n              :label=\"%s.label\"\n            />\n          </n-radio-group>\n          </n-form-item>", field.Dc, field.TsName, field.TsName, field.TsName, field.TsName, in.Options.DictMap[field.TsName], field.TsName, field.TsName, field.TsName)

		case FMCheckbox:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <n-checkbox-group v-model:value=\"formValue.%s\">\n            <n-space>\n              <n-checkbox\n                v-for=\"item in dict.getOptionUnRef('%s')\"\n                :key=\"item.value\"\n                :value=\"item.value\"\n                :label=\"item.label\"\n              />\n            </n-space>\n          </n-checkbox-group>\n          </n-form-item>", field.Dc, field.TsName, field.TsName, in.Options.DictMap[field.TsName])

		case FMSelect:
			if in.Options.DictMap[field.TsName] != nil {
				component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <n-select v-model:value=\"formValue.%s\" :options=\"dict.getOptionUnRef('%s')\" />\n          </n-form-item>", field.Dc, field.TsName, field.TsName, in.Options.DictMap[field.TsName])
			} else {
				component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <n-select v-model:value=\"formValue.%s\" options=\"\" />\n          </n-form-item>", field.Dc, field.TsName, field.TsName)
			}

		case FMSelectMultiple:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <n-select multiple v-model:value=\"formValue.%s\" :options=\"dict.getOptionUnRef('%s')\" />\n          </n-form-item>", field.Dc, field.TsName, field.TsName, in.Options.DictMap[field.TsName])

		case FMUploadImage:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <UploadImage :maxNumber=\"1\" v-model:value=\"formValue.%s\" />\n          </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FMUploadImages:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <UploadImage :maxNumber=\"10\" v-model:value=\"formValue.%s\" />\n          </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FMUploadFile:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <UploadFile :maxNumber=\"1\" v-model:value=\"formValue.%s\" />\n          </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FMUploadFiles:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <UploadFile :maxNumber=\"10\" v-model:value=\"formValue.%s\" />\n          </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FMSwitch:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <n-switch :unchecked-value=\"2\" :checked-value=\"1\" v-model:value=\"formValue.%s\"\n        />\n          </n-form-item>", field.Dc, field.TsName, field.TsName)

		case FMRate:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <n-rate allow-half :default-value=\"formValue.%s\" :on-update:value=\"update%s\" />\n          </n-form-item>", field.Dc, field.TsName, field.TsName, field.GoName)

		case FMCitySelector:
			component = fmt.Sprintf("<n-form-item label=\"%s\" path=\"%s\">\n            <CitySelector v-model:value=\"formValue.%s\" />\n          </n-form-item>", field.Dc, field.TsName, field.TsName)
		case FMPidTreeSelect:
			component = fmt.Sprintf(`<n-form-item label="%v" path="pid">
              <n-tree-select
                :options="treeOption"
                v-model:value="formValue.pid"
                key-field="%v"
                label-field="%v"
                clearable
                filterable
                default-expand-all
                show-path
              />
            </n-form-item>`, field.Dc, in.Pk.TsName, in.Options.Tree.TitleField.TsName)
		case FMTreeSelect:
			component = fmt.Sprintf(`<n-form-item label="%v" path="%v">
              <n-tree-select
                placeholder="请选择%v"
                v-model:value="formValue.%v"
                :options="[{ label: 'AA', key: 1, children: [{ label: 'BB', key: 2 }] }]"
                clearable
                filterable
                default-expand-all
              />
            </n-form-item>`, field.Dc, field.TsName, field.Dc, field.TsName)
		case FMCascader:
			component = fmt.Sprintf(`<n-form-item label="%v" path="%v">
              <n-cascader
                placeholder="请选择%v"
                v-model:value="formValue.%v"
                :options="[{ label: 'AA', value: 1, children: [{ label: 'BB', value: 2 }] }]"
                clearable
                filterable
              />
            </n-form-item>`, field.Dc, field.TsName, field.Dc, field.TsName)
		default:
			component = defaultComponent
		}

		buffer.WriteString(fmt.Sprintf("<n-gi span=\"%v\">%v</n-gi>\n\n", field.FormGridSpan, component))
	}
	return buffer.String()
}

func (l *gCurd) genWebEditScript(ctx context.Context, in *genmodel.CurdPreviewParam) g.Map {
	var (
		data         = make(g.Map)
		importBuffer = bytes.NewBuffer(nil)
		setupBuffer  = bytes.NewBuffer(nil)
	)

	importBuffer.WriteString("  import { ref, computed } from 'vue';\n")

	// 导入字典
	if in.Options.DictOps.Has {
		importBuffer.WriteString("  import { useDictStore } from '@/store/modules/dict';\n")
	}

	// 导入api
	var importApiMethod = []string{"Edit", "View"}
	if in.Options.Step.HasMaxSort {
		importApiMethod = append(importApiMethod, "MaxSort")
	}
	importBuffer.WriteString("  import " + ImportWebMethod(importApiMethod) + " from '" + in.Options.ImportWebApi + "';\n")

	// 导入model
	var importModelMethod = []string{"State", "newState"}
	if in.Options.Step.IsTreeTable {
		importModelMethod = append(importModelMethod, []string{"treeOption", "loadTreeOption"}...)
	}

	if in.Options.Step.HasRules {
		importModelMethod = append(importModelMethod, "rules")
	}
	importBuffer.WriteString("  import " + ImportWebMethod(importModelMethod) + " from './model';\n")

	for _, field := range in.MasterFields {
		if !field.IsEdit {
			continue
		}
		switch field.FormMode {
		case FMDate, FMDateRange, FMTime, FMTimeRange:
			if !gstr.Contains(importBuffer.String(), `import DatePicker`) {
				importBuffer.WriteString("  import DatePicker from '@/components/DatePicker/datePicker.vue';\n")
			}
		case FMInputEditor:
			if !gstr.Contains(importBuffer.String(), `import Editor`) {
				importBuffer.WriteString("  import Editor from '@/components/Editor/editor.vue';\n")
			}
		case FMUploadImage, FMUploadImages:
			if !gstr.Contains(importBuffer.String(), `import UploadImage`) {
				importBuffer.WriteString("  import UploadImage from '@/components/Upload/uploadImage.vue';\n")
			}
		case FMUploadFile, FMUploadFiles:
			if !gstr.Contains(importBuffer.String(), `import UploadFile`) {
				importBuffer.WriteString("  import UploadFile from '@/components/Upload/uploadFile.vue';\n")
			}
		case FMRate:
			setupBuffer.WriteString(fmt.Sprintf("  function update%s(num) {\n    formValue.value.%s = num;\n  }\n", field.GoName, field.TsName))
		case FMCitySelector:
			if !gstr.Contains(importBuffer.String(), `import CitySelector`) {
				importBuffer.WriteString("  import CitySelector from '@/components/CitySelector/citySelector.vue';\n")
			}
		}
	}

	data["import"] = importBuffer.String()
	data["setup"] = setupBuffer.String()
	return data
}
