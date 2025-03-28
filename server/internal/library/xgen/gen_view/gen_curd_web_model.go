// package genview
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package genview

import (
	"context"
	genmodel "xiuadmin/internal/library/xgen/gen_model"
	"xiuadmin/utility"

	"github.com/gogf/gf/v2/frame/g"
)

func (l *gCurd) generateWebModelDictOptions(ctx context.Context, in *genmodel.CurdPreviewParam) error {
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
			g.Log().Debugf(context.Background(), "curd_generate_web_model.go generateWebModelDictOptions in.masterFields:%+v", field)
		}

		// if field.DictType < 0 {
		// 	builtinDictTypeIds = append(builtinDictTypeIds, field.DictType)
		// }
	}
	g.Log().Debugf(context.Background(), "curd_generate_web_model.go generateWebModelDictOptions dictTypeIds:%+v", dictTypeIds)
	dictTypeIds = utility.UniqueSlice(dictTypeIds)
	builtinDictTypeIds = utility.UniqueSlice(builtinDictTypeIds)
	g.Log().Debugf(context.Background(), "curd_generate_web_model.go generateWebModelDictOptions dictTypeIds:%+v", dictTypeIds)
	// g.Log().Debugf(context.Background(), "curd_generate_web_model.go generateWebModelDictOptions builtinDictTypeIds:%+v", builtinDictTypeIds)
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
	g.Log().Debugf(context.Background(), "curd_generate_web_model.go generateWebModelDictOptions dictTypeList:%+v", dictTypeList)
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
