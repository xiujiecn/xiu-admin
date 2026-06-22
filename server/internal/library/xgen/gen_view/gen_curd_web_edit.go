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
	genmodel "xiuadmin/internal/library/xgen/gen_model"
	"xiuadmin/utility"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

func (l *gCurd) webEditTplData(ctx context.Context, in *genmodel.CurdPreviewParam) (data g.Map, err error) {
	data = make(g.Map)
	data["formItem"] = l.genWebEditFormItem(ctx, in)
	data["script"] = l.genWebEditScript(ctx, in)
	data["isEditModal"] = in.Options.Step.IsEditModal
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
			defaultComponent = fmt.Sprintf(`<Col :span="%d">
          <FormItem label="%s" name="%s">
            <Input placeholder="请输入%s" v-model:value="formValue.%s" />
          </FormItem>
        </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.Dc, field.TsName)
			component        string
		)

		if in.Options.Step.IsTreeTable && IsPidName(field.Name) {
			field.FormMode = FMPidTreeSelect
		}

		switch field.FormMode {
		case FMInput:
			component = defaultComponent

		case FMInputNumber:
			component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <InputNumber class="w-full" placeholder="请输入%s" v-model:value="formValue.%s" />
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.Dc, field.TsName)

		case FMInputTextarea:
			component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <Textarea placeholder="请输入%s" v-model:value="formValue.%s" />
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.Dc, field.TsName)

		case FMInputEditor:
			component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <RichTextarea style="height: 450px" v-model="formValue.%s" />
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.TsName)

		case FMInputDynamic:
			component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <Space direction="vertical" class="w-full">
                <Space v-for="(item, index) in formValue.%s" :key="index">
                  <Input v-model:value="item.key" placeholder="键名" />
                  <Input v-model:value="item.value" placeholder="键值" />
                </Space>
              </Space>
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.TsName)

		case FMDate:
			component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <DatePicker v-model:value="formValue.%s" class="w-full" />
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.TsName)

		case FMTime:
			component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <DatePicker v-model:value="formValue.%s" show-time class="w-full" />
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.TsName)

		case FMRadio:
			component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <RadioGroup
                v-model:value="formValue.%s"
                :options="getDictOptions('%s')"
                option-type="button"
                button-style="solid"
              />
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.TsName, in.Options.DictMap[field.TsName])

		case FMCheckbox:
			component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <CheckboxGroup v-model:value="formValue.%s">
                <Space>
                  <Checkbox
                    v-for="item in getDictOptions('%s')"
                    :key="item.value"
                    :value="item.value"
                  >
                    {{ item.label }}
                  </Checkbox>
                </Space>
              </CheckboxGroup>
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.TsName, in.Options.DictMap[field.TsName])

		case FMSelect:
			if in.Options.DictMap[field.TsName] != nil {
				component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <Select v-model:value="formValue.%s" :options="getDictOptions('%s')" placeholder="请选择%s" class="w-full" />
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.TsName, in.Options.DictMap[field.TsName], field.Dc)
			} else {
				component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <Select v-model:value="formValue.%s" placeholder="请选择%s" class="w-full" />
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.TsName, field.Dc)
			}

		case FMSelectMultiple:
			component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <Select v-model:value="formValue.%s" mode="multiple" :options="getDictOptions('%s')" placeholder="请选择%s" class="w-full" />
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.TsName, in.Options.DictMap[field.TsName], field.Dc)

		case FMUploadImage:
			component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <ImageUpload :max-count="1" v-model:value="formValue.%s" />
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.TsName)

		case FMUploadImages:
			component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <ImageUpload :max-count="10" v-model:value="formValue.%s" />
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.TsName)

		case FMUploadFile:
			component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <FileUpload :max-count="1" v-model:value="formValue.%s" />
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.TsName)

		case FMUploadFiles:
			component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <FileUpload :max-count="10" v-model:value="formValue.%s" />
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.TsName)

		case FMSwitch:
			component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <Switch v-model:checked="formValue.%s" :checked-value="1" :un-checked-value="2" />
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.TsName)

		case FMRate:
			component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <Rate allow-half v-model:value="formValue.%s" />
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.TsName)

		case FMCitySelector:
			component = fmt.Sprintf(`<Col :span="%d">
            <FormItem label="%s" name="%s">
              <Cascader v-model:value="formValue.%s" placeholder="请选择%s" class="w-full" />
            </FormItem>
          </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.TsName, field.Dc)

		case FMPidTreeSelect:
			component = fmt.Sprintf(`<Col :span="%d">
              <FormItem label="%v" name="pid">
                <TreeSelect
                  v-model:value="formValue.pid"
                  :tree-data="treeOption"
                  :field-names="{ label: '%v', value: '%v', children: 'children' }"
                  allow-clear
                  show-search
                  tree-default-expand-all
                  :tree-node-filter-prop="'%v'"
                  class="w-full"
                />
              </FormItem>
            </Col>`, field.FormGridSpan, field.Dc, in.Options.Tree.TitleField.TsName, in.Pk.TsName, in.Options.Tree.TitleField.TsName)

		case FMTreeSelect:
			component = fmt.Sprintf(`<Col :span="%d">
              <FormItem label="%v" name="%v">
                <TreeSelect
                  placeholder="请选择%v"
                  v-model:value="formValue.%v"
                  :tree-data="[{ title: 'AA', value: 1, children: [{ title: 'BB', value: 2 }] }]"
                  allow-clear
                  show-search
                  tree-default-expand-all
                  class="w-full"
                />
              </FormItem>
            </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.Dc, field.TsName)

		case FMCascader:
			component = fmt.Sprintf(`<Col :span="%d">
              <FormItem label="%v" name="%v">
                <Cascader
                  placeholder="请选择%v"
                  v-model:value="formValue.%v"
                  :options="[{ label: 'AA', value: 1, children: [{ label: 'BB', value: 2 }] }]"
                  allow-clear
                  show-search
                  class="w-full"
                />
              </FormItem>
            </Col>`, field.FormGridSpan, field.Dc, field.TsName, field.Dc, field.TsName)

		default:
			component = defaultComponent
		}

		buffer.WriteString(component + "\n\n")
	}
	return buffer.String()
}

func (l *gCurd) genWebEditScript(ctx context.Context, in *genmodel.CurdPreviewParam) g.Map {
	var (
		data         = make(g.Map)
		importBuffer = bytes.NewBuffer(nil)
		antdUI       []string
	)

	importBuffer.WriteString("  import { ref, computed } from 'vue';\n")

	// 导入字典
	if in.Options.DictOps.Has {
		importBuffer.WriteString("  import { getDictOptions } from '#/utils/dict';\n")
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
		case FMInput, FMInputTextarea, FMInputDynamic:
			antdUI = append(antdUI, "Col", "FormItem", "Input", "Space")
		case FMInputNumber:
			antdUI = append(antdUI, "Col", "FormItem", "InputNumber")
		case FMDate, FMDateRange, FMTime, FMTimeRange:
			antdUI = append(antdUI, "Col", "FormItem", "DatePicker")
		case FMInputEditor:
			if !gstr.Contains(importBuffer.String(), `import { Tinymce as RichTextarea }`) {
				importBuffer.WriteString("  import { Tinymce as RichTextarea } from '#/components/tinymce';\n")
			}
			antdUI = append(antdUI, "Col", "FormItem")
		case FMUploadImage, FMUploadImages:
			if !gstr.Contains(importBuffer.String(), `import { ImageUpload }`) {
				importBuffer.WriteString("  import { ImageUpload } from '#/components/upload';\n")
			}
			antdUI = append(antdUI, "Col", "FormItem")
		case FMUploadFile, FMUploadFiles:
			if !gstr.Contains(importBuffer.String(), `import { FileUpload }`) {
				importBuffer.WriteString("  import { FileUpload } from '#/components/upload';\n")
			}
			antdUI = append(antdUI, "Col", "FormItem")
		case FMRadio:
			antdUI = append(antdUI, "Col", "FormItem", "RadioGroup")
		case FMCheckbox:
			antdUI = append(antdUI, "Col", "FormItem", "CheckboxGroup", "Checkbox", "Space")
		case FMSelect, FMSelectMultiple:
			antdUI = append(antdUI, "Col", "FormItem", "Select")
		case FMSwitch:
			antdUI = append(antdUI, "Col", "FormItem", "Switch")
		case FMRate:
			antdUI = append(antdUI, "Col", "FormItem", "Rate")
		case FMCitySelector, FMCascader:
			antdUI = append(antdUI, "Col", "FormItem", "Cascader")
		case FMPidTreeSelect, FMTreeSelect:
			antdUI = append(antdUI, "Col", "FormItem", "TreeSelect")
		}
	}

	antdUI = utility.UniqueSlice(antdUI)
	if len(antdUI) > 0 {
		importBuffer.WriteString("  import " + ImportWebMethod(antdUI) + " from 'ant-design-vue';\n")
	}

	data["import"] = importBuffer.String()
	data["setup"] = ""
	if in.Options.PresetStep != nil {
		data["formGridSpan"] = in.Options.PresetStep.FormGridCols
	}
	return data
}
