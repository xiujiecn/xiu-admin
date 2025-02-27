package convert

import "github.com/gogf/gf/v2/text/gstr"

// EscapeFieldsToSlice 将转义过的字段转换为字段集切片
func EscapeFieldsToSlice(s string) []string {
	return gstr.Explode(",", gstr.Replace(gstr.Replace(s, "`,`", ","), "`", ""))
}
