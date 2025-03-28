// package genview
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package genview

import (
	"context"
	"strings"
	gendao "xiuadmin/internal/library/xgen/gen_dao"
	genmodel "xiuadmin/internal/library/xgen/gen_model"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

func getJsonTagFromCase(str, caseStr string) string {
	switch gstr.ToLower(caseStr) {
	case gstr.ToLower("Camel"):
		return gstr.CaseCamel(str)

	case gstr.ToLower("CamelLower"):
		return gstr.CaseCamelLower(str)

	case gstr.ToLower("Kebab"):
		return gstr.CaseKebab(str)

	case gstr.ToLower("KebabScreaming"):
		return gstr.CaseKebabScreaming(str)

	case gstr.ToLower("Snake"):
		return gstr.CaseSnake(str)

	case gstr.ToLower("SnakeFirstUpper"):
		return gstr.CaseSnakeFirstUpper(str)

	case gstr.ToLower("SnakeScreaming"):
		return gstr.CaseSnakeScreaming(str)
	}
	return str
}

func CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{}) (string, error) {
	var (
		typeName    string
		typePattern string
	)
	match, _ := gregex.MatchString(`(.+?)\((.+)\)`, fieldType)
	if len(match) == 3 {
		typeName = gstr.Trim(match[1])
		typePattern = gstr.Trim(match[2])
	} else {
		typeName = gstr.Split(fieldType, " ")[0]
	}
	typeName = strings.ToLower(typeName)
	switch typeName {
	case
		"binary",
		"varbinary",
		"blob",
		"tinyblob",
		"mediumblob",
		"longblob":
		return string(gdb.LocalTypeBytes), nil

	case
		"int",
		"tinyint",
		"small_int",
		"smallint",
		"medium_int",
		"mediumint",
		"serial":
		if gstr.ContainsI(fieldType, "unsigned") {
			return string(gdb.LocalTypeUint), nil
		}
		return string(gdb.LocalTypeInt), nil

	case
		"big_int",
		"bigint",
		"bigserial":
		if gstr.ContainsI(fieldType, "unsigned") {
			return string(gdb.LocalTypeUint64), nil
		}
		return string(gdb.LocalTypeInt64), nil

	case
		"real":
		return string(gdb.LocalTypeFloat32), nil

	case
		"float",
		"double",
		"decimal",
		"money",
		"numeric",
		"smallmoney":
		return string(gdb.LocalTypeFloat64), nil

	case
		"bit":
		// It is suggested using bit(1) as boolean.
		if typePattern == "1" {
			return string(gdb.LocalTypeBool), nil
		}
		s := gconv.String(fieldValue)
		// mssql is true|false string.
		if strings.EqualFold(s, "true") || strings.EqualFold(s, "false") {
			return string(gdb.LocalTypeBool), nil
		}
		if gstr.ContainsI(fieldType, "unsigned") {
			return string(gdb.LocalTypeUint64Bytes), nil
		}
		return string(gdb.LocalTypeInt64Bytes), nil

	case
		"bool":
		return string(gdb.LocalTypeBool), nil

	case
		"date":
		return string(gdb.LocalTypeDate), nil

	case
		"datetime",
		"timestamp",
		"timestamptz":
		return string(gdb.LocalTypeDatetime), nil

	case
		"json":
		return string(gdb.LocalTypeJson), nil

	case
		"jsonb":
		return string(gdb.LocalTypeJsonb), nil

	default:
		// Auto-detect field type, using key match.
		switch {
		case strings.Contains(typeName, "text") || strings.Contains(typeName, "char") || strings.Contains(typeName, "character"):
			return string(gdb.LocalTypeString), nil

		case strings.Contains(typeName, "float") || strings.Contains(typeName, "double") || strings.Contains(typeName, "numeric"):
			return string(gdb.LocalTypeFloat64), nil

		case strings.Contains(typeName, "bool"):
			return string(gdb.LocalTypeBool), nil

		case strings.Contains(typeName, "binary") || strings.Contains(typeName, "blob"):
			return string(gdb.LocalTypeBytes), nil

		case strings.Contains(typeName, "int"):
			if gstr.ContainsI(fieldType, "unsigned") {
				return string(gdb.LocalTypeUint), nil
			}
			return string(gdb.LocalTypeInt), nil

		case strings.Contains(typeName, "time"):
			return string(gdb.LocalTypeDatetime), nil

		case strings.Contains(typeName, "date"):
			return string(gdb.LocalTypeDatetime), nil

		default:
			return string(gdb.LocalTypeString), nil
		}
	}
}

// GenGotype 生成字段的go类型
func GenGotype(ctx context.Context, field *genmodel.GenCodesColumnListModel, in *genmodel.GenDaoConfig) (goName, typeName, tsName string, tsType string) {
	var err error
	tsName = getJsonTagFromCase(field.Name, in.JsonCase)
	goName = gstr.CaseCamel(field.Name)

	typeName, err = CheckLocalTypeForField(ctx, field.DataType, nil)
	if err != nil {
		panic(err)
	}

	switch gdb.LocalType(typeName) {
	case gdb.LocalTypeDate, gdb.LocalTypeDatetime:
		if in.StdTime {
			typeName = "time.Time"
		} else {
			typeName = "*gtime.Time"
		}

	case gdb.LocalTypeInt64Bytes:
		typeName = "int64"

	case gdb.LocalTypeUint64Bytes:
		typeName = "uint64"

	// Special type handle.
	case gdb.LocalTypeJson, gdb.LocalTypeJsonb:
		if in.GJsonSupport {
			typeName = "*gjson.Json"
		} else {
			typeName = "string"
		}
	}

	tsType = ShiftMap[typeName]
	return
}

func CustomLinkAttributes(ctx context.Context, alias string, field *genmodel.GenCodesColumnListModel, in *genmodel.GenDaoConfig) {
	field.GoName, field.GoType, field.TsName, field.TsType = GenGotype(ctx, field, in)

	field.GoName = gstr.UcFirst(alias + field.GoName)
	field.TsName = gstr.LcFirst(field.GoName)

	setDefaultQueryWhere(field)
	setDefaultValue(field)
}

func CustomAttributes(ctx context.Context, field *genmodel.GenCodesColumnListModel, in *genmodel.GenDaoConfig) {
	field.GoName, field.GoType, field.TsName, field.TsType = GenGotype(ctx, field, in)
	setDefault(field)
}

func TableColumnList(ctx context.Context, params *genmodel.GenCodesColumnListParam) (output []*genmodel.GenCodesColumnListModel, err error) {
	sql := "select ORDINAL_POSITION as `id`, COLUMN_NAME as `name`, COLUMN_COMMENT as `dc`, DATA_TYPE as `dataType`, COLUMN_TYPE as `sqlType`, CHARACTER_MAXIMUM_LENGTH as `length`, IS_NULLABLE as `isAllowNull`, COLUMN_DEFAULT as `defaultValue`, COLUMN_KEY as `index`, EXTRA as `extra` from information_schema.COLUMNS where TABLE_SCHEMA = ? and TABLE_NAME = ? ORDER BY `id` ASC"
	dbConfig := g.DB(params.DbGroup).GetConfig()
	daoConfig, err := gendao.GetDaoConfig(ctx, params.DbGroup)
	if err != nil {
		return
	}

	output = make([]*genmodel.GenCodesColumnListModel, 0)
	err = g.DB(params.DbGroup).Ctx(ctx).Raw(sql, dbConfig.Name, params.TableName).Scan(&output)
	if err != nil {
		return
	}
	if len(output) == 0 {
		return
	}
	for _, v := range output {
		if params.IsLink == 1 {
			CustomLinkAttributes(ctx, params.Alias, v, daoConfig)
		} else {
			CustomAttributes(ctx, v, daoConfig)
		}
	}
	return
}
