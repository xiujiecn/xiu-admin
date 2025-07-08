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

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

func (l *gCurd) webEditTplData(ctx context.Context, in *genmodel.CurdPreviewParam) (data g.Map, err error) {
	data = make(g.Map)
	data["script"] = l.genWebEditScript(ctx, in)
	data["isEditModal"] = in.Options.Step.IsEditModal
	return
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
	if in.Options.PresetStep != nil {
		data["formGridSpan"] = in.Options.PresetStep.FormGridCols
	}
	return data
}
