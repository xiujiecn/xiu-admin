package genview

import (
	"bytes"
	"context"

	genmodel "xiujieadmin/internal/library/xgen/gen_model"

	"github.com/gogf/gf/v2/frame/g"
)

func (l *gCurd) webIndexTplData(ctx context.Context, in *genmodel.CurdPreviewParam) (g.Map, error) {
	var (
		data              = make(g.Map)
		importBuffer      = bytes.NewBuffer(nil)
		importVueMethod   = []string{"h", "reactive", "ref", "computed"}
		importApiMethod   = []string{"List"}
		importModelMethod = []string{"columns", "querySchema", "type RowType"}
		importUtilsMethod = []string{"adaTableScrollX"}
		importIcons       []string
		actionWidth       int64 = 72
	)

	// 添加 MdiPlus, MdiDelete,MdiExport
	if in.Options.Step.HasAdd {
		importIcons = append(importIcons, "MdiPlus")
	}

	// 编辑
	if in.Options.Step.HasEdit {
		in.Options.Step.ActionColumnWidth += actionWidth
		if in.Options.Step.IsTreeTable && !in.Options.Step.IsOptionTreeTable {
			in.Options.Step.ActionColumnWidth += actionWidth
		}
		if in.Options.Step.IsOptionTreeTable {
			importIcons = append(importIcons, "MdiEdit")
		}
	}

	// 导出
	if in.Options.Step.HasExport {
		importIcons = append(importIcons, "MdiExport")
		importApiMethod = append(importApiMethod, "Export")
	}

	// 删除
	if in.Options.Step.HasDel {
		importApiMethod = append(importApiMethod, "Delete")
		in.Options.Step.ActionColumnWidth += actionWidth
	}

	// 批量删除
	if in.Options.Step.HasBatchDel {
		importIcons = append(importIcons, "MdiDelete")
		importApiMethod = append(importApiMethod, "Delete")
	}

	// 修改状态
	if in.Options.Step.HasStatus {
		importApiMethod = append(importApiMethod, "Status")
		in.Options.Step.ActionColumnWidth += actionWidth
	}

	// 更多
	// 查看详情
	if in.Options.Step.HasView {
		in.Options.Step.ActionColumnWidth += actionWidth
	}

	// 展开树
	if in.Options.Step.IsTreeTable {
		// importIcons = append(importIcons, "MdiExport")
	}

	// 存在字典数据选项
	if in.Options.DictOps.Has {
		importVueMethod = append(importVueMethod, "onMounted")
		// importModelMethod = append(importModelMethod, "loadOptions")
	}

	// 普通树表
	if in.Options.Step.IsTreeTable && !in.Options.Step.IsOptionTreeTable {
		importUtilsMethod = append(importUtilsMethod, "convertListToTree")
	}

	// 选项式树表
	if in.Options.Step.IsOptionTreeTable {
		importVueMethod = append(importVueMethod, []string{"onMounted", "unref"}...)
		// importIcons = append(importIcons, []string{"FormOutlined", "SearchOutlined"}...)
		importApiMethod = append(importApiMethod, "TreeOption")
		importUtilsMethod = append(importUtilsMethod, "getTreeKeys")
		importModelMethod = append(importModelMethod, []string{"loadTreeOption", "treeOption", "State"}...)
	}

	// 操作按钮宽度最小值
	if in.Options.Step.ActionColumnWidth > 0 && in.Options.Step.ActionColumnWidth < actionWidth*2 {
		in.Options.Step.ActionColumnWidth = 100
	}

	// 导入基础包
	importBuffer.WriteString("  import " + ImportWebMethod(importVueMethod) + " from 'vue';\n")
	importBuffer.WriteString("  import { Button, message,Tag, Modal, Popconfirm,Switch } from 'ant-design-vue';\n")
	importBuffer.WriteString("  import type { VbenFormProps } from '#/adapter/form';\n")
	importBuffer.WriteString("  import type { VxeTableGridOptions, VxeGridListeners } from '#/adapter/vxe-table';\n")
	importBuffer.WriteString("  import type { DeepPartial } from '@vben/types';\n")
	importBuffer.WriteString("  import { getVxePopupContainer } from '@vben/utils';\n")
	importBuffer.WriteString("  import { Page, useVbenDrawer } from '@vben/common-ui';\n")
	importBuffer.WriteString("  import { useVbenVxeGrid } from '#/adapter/vxe-table';\n")
	if in.Options.Step.HasExport {
		importBuffer.WriteString("  import { commonDownloadExcel } from '#/utils/file/download';\n")
	}

	// 导入字典
	if in.Options.DictOps.Has {
		// importBuffer.WriteString("  import { useDictStore } from '@/store/modules/dict';\n")
	}

	// 导入api
	importBuffer.WriteString("  import " + ImportWebMethod(importApiMethod) + " from '" + in.Options.ImportWebApi + "';\n")

	// 导入icons
	if len(importIcons) > 0 {
		importBuffer.WriteString("  import " + ImportWebMethod(importIcons) + "  from '@vben/icons';\n")
	}

	// 导入model
	if in.Options.Step.IsTreeTable {
		importModelMethod = append(importModelMethod, "newState")
	}
	importBuffer.WriteString("  import " + ImportWebMethod(importModelMethod) + " from './model';\n")

	// 导入utils
	if len(importUtilsMethod) > 0 {
		// importBuffer.WriteString("  import " + ImportWebMethod(importUtilsMethod) + " from '@/utils/xjadmin';\n")
	}

	// 导入edit组件
	if in.Options.Step.HasEdit {
		importBuffer.WriteString("  import editDrawer from './edit.vue';\n")
	}

	// 导入view组件
	if in.Options.Step.HasView {
		importBuffer.WriteString("  import viewDrawer from './view.vue';\n")
	}

	// 没有需要查询的字段则隐藏搜索表单
	isSearchForm := false
	for _, field := range in.MasterFields {
		if field.IsQuery {
			isSearchForm = true
			break
		}
	}
	if !isSearchForm {
		if len(in.Options.Join) > 0 {
		LoopOut:
			for _, v := range in.Options.Join {
				for _, column := range v.Columns {
					if column.IsQuery {
						isSearchForm = true
						break LoopOut
					}
				}
			}
		}
	}
	data["isSearchForm"] = isSearchForm
	data["import"] = importBuffer.String()
	return data, nil
}
