// package xgorm
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package xgorm

import (
	"context"
	"fmt"
	"slices"
	"xiuadmin/utility"

	"reflect"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	descTags  = []string{"description", "dc", "json"} // 实体描述标签
	fieldTags = []string{"json"}                      // 实体字段名称映射
)

// GetFieldsToSlice 获取dao实例中的所有字段
func GetFieldsToSlice(ctx context.Context, dao daoInstance) ([]string, error) {
	fields, err := dao.Ctx(ctx).TableFields(dao.Table())
	if err != nil {
		return nil, err
	}
	if len(fields) == 0 {
		return nil, gerror.New("field not found")
	}

	var keys []string
	for _, field := range fields {
		keys = append(keys, field.Name)
	}
	return keys, nil
}

// GetPkField 获取dao实例中的主键名称
func GetPkField(ctx context.Context, dao daoInstance) (string, error) {
	fields, err := dao.Ctx(ctx).TableFields(dao.Table())
	if err != nil {
		return "", err
	}
	if len(fields) == 0 {
		return "", gerror.New("field not found")
	}

	for _, field := range fields {
		if field.Key == "PRI" {
			return field.Name, nil
		}
	}
	return "", gerror.New("no primary key")
}

// IsUnique 是否唯一
func IsUnique(ctx context.Context, dao daoInstance, where g.Map, message string, pkId ...interface{}) error {
	if len(where) == 0 {
		return gerror.New("where condition cannot be empty")
	}

	m := dao.Ctx(ctx).Where(where)
	if len(pkId) > 0 {
		field, err := GetPkField(ctx, dao)
		if err != nil {
			return err
		}
		m = m.WhereNot(field, pkId[0])
	}

	count, err := m.Count(1)
	if err != nil {
		return err
	}

	if count > 0 {
		if message == "" {
			for k := range where {
				message = fmt.Sprintf("in the table：%s, %v not uniqued", dao.Table(), where[k])
				break
			}
		}
		return gerror.New(message)
	}
	return nil
}

func JoinFields(ctx context.Context, entity interface{}, dao daoInstance, as string) (fs string) {
	entityFs, err := GetEntityFieldTags(entity)
	if err != nil {
		g.Log().Errorf(ctx, "JoinFields err: %v", err)
		return
	}

	if len(entityFs) == 0 {
		g.Log().Errorf(ctx, "JoinFields entityFs is empty")
		return ""
	}

	fields, err := dao.Ctx(ctx).TableFields(dao.Table())
	if err != nil {
		g.Log().Errorf(ctx, "JoinFields err: %v", err)
		return
	}

	var columns []string
	for _, v := range entityFs {
		if !gstr.HasPrefix(v, as) {
			g.Log().Infof(ctx, "JoinFields gstr.HasPrefix: %+v, %+v", v, as)
			continue
		}
		g.Log().Infof(ctx, "JoinFields gstr.HasPrefix2: %+v, %+v", v, as)
		field := gstr.CaseSnakeFirstUpper(gstr.StrEx(v, as))
		if _, ok := fields[field]; ok {
			columns = append(columns, fmt.Sprintf("`%s`.`%s` as `%s`", dao.Table(), field, v))
		}
		g.Log().Infof(ctx, "JoinFields field: %+v fields: %+v", field, fields)
	}

	if len(columns) > 0 {
		g.Log().Infof(ctx, "JoinFields columns: %+v", columns)
		return gstr.Implode(",", utility.UniqueSlice(columns))
	}
	return
}

// GetEntityFieldTags 获取实体中的字段名称
func GetEntityFieldTags(entity interface{}) (tags []string, err error) {
	var formRef = reflect.TypeOf(entity)
	for i := 0; i < formRef.NumField(); i++ {
		field := formRef.Field(i)
		if field.Type.Kind() == reflect.Struct {
			addTags, err := reflectTag(field.Type, fieldTags, []string{})
			if err != nil {
				return nil, err
			}
			tags = append(tags, addTags...)
			continue
		}
		tags = append(tags, reflectTagName(field, fieldTags, true))
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
