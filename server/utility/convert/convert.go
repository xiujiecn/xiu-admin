// package convert
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package convert

import (
	"reflect"
	"slices"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	descTags  = []string{"description", "dc", "json"} // 实体描述标签
	fieldTags = []string{"json"}                      // 实体字段名称映射
)

// EscapeFieldsToSlice 将转义过的字段转换为字段集切片
func EscapeFieldsToSlice(s string) []string {
	return gstr.Explode(",", gstr.Replace(gstr.Replace(s, "`,`", ","), "`", ""))
}

// GetEntityDescTags 获取实体中的描述标签
func GetEntityDescTags(entity interface{}) (tags []string, err error) {
	var formRef = reflect.TypeOf(entity)
	for i := 0; i < formRef.NumField(); i++ {
		field := formRef.Field(i)
		if field.Type.Kind() == reflect.Struct {
			addTags, err := reflectTag(field.Type, descTags, []string{})
			if err != nil {
				return nil, err
			}
			tags = append(tags, addTags...)
			continue
		}
		tags = append(tags, reflectTagName(field, descTags, true))
	}
	return
}

// reflectTag 层级递增解析tag
func reflectTag(reflectType reflect.Type, filterTags []string, tags []string) ([]string, error) {
	if reflectType.Kind() == reflect.Ptr {
		return nil, gerror.Newf("reflect type do not support reflect.Ptr yet, reflectType:%+v", reflectType)
	}
	if reflectType.Kind() != reflect.Struct {
		return nil, nil
	}
	for i := 0; i < reflectType.NumField(); i++ {
		tag := reflectTagName(reflectType.Field(i), filterTags, false)
		if tag == "" {
			addTags, err := reflectTag(reflectType.Field(i).Type, filterTags, tags)
			if err != nil {
				return nil, err
			}
			tags = append(tags, addTags...)
			continue
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

// reflectTagName 解析实体中的描述标签，优先级：description > dc > json > Name
func reflectTagName(field reflect.StructField, filterTags []string, isDef bool) string {

	if slices.Contains(filterTags, "description") {
		if description, ok := field.Tag.Lookup("description"); ok && description != "" {
			return description
		}
	}

	if slices.Contains(filterTags, "dc") {
		if dc, ok := field.Tag.Lookup("dc"); ok && dc != "" {
			return dc
		}
	}

	if slices.Contains(filterTags, "json") {
		if jsonName, ok := field.Tag.Lookup("json"); ok && jsonName != "" {
			return jsonName
		}
	}

	if !isDef {
		return ""
	}
	return field.Name
}
