// package genview
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package genview

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/mcache"
	genmodel "xiuadmin/internal/library/xgen/gen_model"
	"xiuadmin/internal/library/xgen/gen_view/gohtml"
	"xiuadmin/internal/service"
	"xiuadmin/utility"
	version "xiuadmin/utility/version"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/tools/imports"
)

func replaceEmptyLinesWithSpace(input string) string {
	re := regexp.MustCompile(`\n\s*\n`)
	result := re.ReplaceAllString(input, "\n\n")
	return result
}

func FormatTs(code string) string {
	code = replaceEmptyLinesWithSpace(code)
	return code + "\n"
}

func ToTSArray(vs []string) string {
	formattedStrings := make([]string, len(vs))
	for i, str := range vs {
		formattedStrings[i] = fmt.Sprintf("'%s'", str)
	}
	return fmt.Sprintf("[%s]", strings.Join(formattedStrings, ", "))
}

func DictToTSArray(vs []string) string {
	formattedStrings := make([]string, len(vs))
	formattedObjStrings := make([]string, len(vs))
	for i, str := range vs {
		formattedStrings[i] = fmt.Sprintf("'%s'", str)
		formattedObjStrings[i] = fmt.Sprintf("dict_%s", str)
	}
	return fmt.Sprintf("[%s],[%s]", strings.Join(formattedStrings, ", "), strings.Join(formattedObjStrings, ", "))
}

func RemoveSlice[K comparable](src []K, sub K) []K {
	for k, v := range src {
		if v == sub {
			copy(src[k:], src[k+1:])
			return src[:len(src)-1]
		}
	}
	return src
}

// CheckIllegalName 检查命名是否合理
func CheckIllegalName(errPrefix string, names ...string) (err error) {
	reg, _ := regexp.Compile("^[a-z_][a-z0-9_]*$")
	for _, name := range names {
		name = strings.ToLower(name)
		match := reg.MatchString(name)
		if !match {
			err = gerror.Newf("%v存在格式不正确，必须全部小写且由字母、数字和下划线组成:%v", errPrefix, name)
			return
		}
		if strings.HasSuffix(name, "test") {
			err = gerror.Newf("%v当中不能以`test`结尾:%v", errPrefix, name)
			return
		}
		if StartsWithDigit(name) {
			err = gerror.Newf("%v当中不能以阿拉伯数字开头:%v", errPrefix, name)
			return
		}
	}
	return
}
func StartsWithDigit(s string) bool {
	r := []rune(s)
	if len(r) > 0 {
		return unicode.IsDigit(r[0])
	}
	return false
}

// GetModName 获取主包名
func GetModName(ctx context.Context) (modName string, err error) {
	if !gfile.Exists("go.mod") {
		err = gerror.New("go.mod does not exist in current working directory")
		return
	}

	var (
		goModContent = gfile.GetContents("go.mod")
		match, _     = gregex.MatchString(`^module\s+(.+)\s*`, goModContent)
	)

	if len(match) > 1 {
		modName = gstr.Trim(match[1])
	} else {
		err = gerror.New("module name does not found in go.mod")
		return
	}
	return
}

// parseServFunName 解析业务服务名称
func (l *gCurd) parseServFunName(templateGroup, varName string) string {
	templateGroup = gstr.UcFirst(templateGroup)
	if gstr.HasPrefix(varName, templateGroup) && varName != templateGroup {
		return varName
	}
	return templateGroup + varName
}

func GetTempGeneratePath(ctx context.Context) string {
	return gfile.Abs(gfile.Temp() + "/tmp_xjgen/" + version.AppName(ctx))
}

func FormatGo(ctx context.Context, name, code string) (string, error) {
	path := GetTempGeneratePath(ctx) + "/" + name
	if err := gfile.PutContents(path, code); err != nil {
		return "", err
	}
	res, err := imports.Process(path, []byte(code), nil)
	if err != nil {
		err = gerror.Newf(`FormatGo error format "%s" go files: %v`, path, err)
		return "", err
	}
	return string(res), nil
}
func formatComment(comment string) string {
	comment = gstr.ReplaceByArray(comment, g.SliceStr{
		"\n", " ",
		"\r", " ",
	})
	comment = gstr.Replace(comment, `\n`, " ")
	comment = gstr.Trim(comment)
	return comment
}

func removeEndWrap(comment string) string {
	if len(comment) > 2 && comment[len(comment)-2:] == " \n" {
		comment = comment[:len(comment)-2]
	}
	return comment
}
func isEffectiveJoin(join *genmodel.CurdOptionsJoin) bool {
	return join.Alias != "" && join.Field != "" && join.LinkTable != "" && join.MasterField != "" && join.DaoName != "" && join.LinkMode > 0
}

func IsPidName(name string) bool {
	return name == "pid"
}

// CamelCaseToUnderline 驼峰单词转下划线单词
func CamelCaseToUnderline(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
		} else {
			if unicode.IsUpper(r) {
				output = append(output, '_')
			}

			output = append(output, unicode.ToLower(r))
		}
	}
	return string(output)
}

// hasEffectiveJoin 存在有效的关联表
func hasEffectiveJoins(joins []*genmodel.CurdOptionsJoin) bool {
	for _, join := range joins {
		if isEffectiveJoin(join) {
			return true
		}
	}
	return false
}

// ImportWebMethod 导入前端方法
func ImportWebMethod(vs []string) string {
	vs = utility.UniqueSlice(vs)
	str := "{ " + strings.Join(vs, ", ") + " }"
	str = strings.TrimSuffix(str, ", ")
	return str
}

func FormatVue(code string) string {
	endTag := `</template>`
	vueLen := gstr.PosR(code, endTag)
	vueCode := code[:vueLen+len(endTag)]
	tsCode := code[vueLen+len(endTag):]
	vueCode = gohtml.Format(vueCode)
	tsCode = FormatTs(tsCode)
	return vueCode + tsCode
}

// ImportSql 导出sql文件
func ImportSql(ctx context.Context, path string) error {
	rows, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	beforeMaxMenuId, _ := dao.SysMenu.Ctx(ctx).Max(dao.SysMenu.Columns().MenuId)
	sqlArr := strings.Split(string(rows), "\n")
	for _, sql := range sqlArr {
		sql = strings.TrimSpace(sql)
		if sql == "" || strings.HasPrefix(sql, "--") {
			continue
		}
		g.Log().Infof(ctx, "views.ImportSql sql:%v", sql)
		exec, err := g.DB().Exec(ctx, sql)
		g.Log().Infof(ctx, "views.ImportSql sql:%v, exec:%+v, err:%+v", sql, exec, err)
		if err != nil {
			return err
		}
	}
	if err = syncGeneratedMenus(ctx, gconv.Int64(beforeMaxMenuId)); err != nil {
		return err
	}
	return nil
}

func syncGeneratedMenus(ctx context.Context, beforeMaxMenuId int64) error {
	var menuIds []int64
	if err := dao.SysMenu.Ctx(ctx).
		Fields(dao.SysMenu.Columns().MenuId).
		WhereGT(dao.SysMenu.Columns().MenuId, beforeMaxMenuId).
		OrderAsc(dao.SysMenu.Columns().MenuId).
		Scan(&menuIds); err != nil {
		return err
	}
	if len(menuIds) == 0 {
		return nil
	}

	for _, menuId := range menuIds {
		if err := service.MemDBSysMenu().LoadData(ctx, menuId); err != nil {
			return err
		}
	}

	var roleIds []int64
	if err := dao.SysRoleMenu.Ctx(ctx).
		Fields(dao.SysRoleMenu.Columns().RoleId).
		Where(dao.SysRoleMenu.Columns().TenantId, contexts.GetTenantId(ctx)).
		WhereIn(dao.SysRoleMenu.Columns().MenuId, menuIds).
		Group(dao.SysRoleMenu.Columns().RoleId).
		Scan(&roleIds); err != nil {
		return err
	}
	if len(roleIds) == 0 {
		return nil
	}

	var userIds []int64
	if err := dao.SysUserRole.Ctx(ctx).
		Fields(dao.SysUserRole.Columns().UserId).
		Where(dao.SysUserRole.Columns().TenantId, contexts.GetTenantId(ctx)).
		WhereIn(dao.SysUserRole.Columns().RoleId, roleIds).
		Group(dao.SysUserRole.Columns().UserId).
		Scan(&userIds); err != nil {
		return err
	}
	for _, userId := range userIds {
		_ = mcache.RemoveUserAccessCodeList(ctx, userId)
		_ = mcache.RemoveUserRoleDataAccessCodeList(ctx, userId)
	}
	return nil
}

// ParseDBConfigNodeLink 解析数据库连接配置
func ParseDBConfigNodeLink(node *gdb.ConfigNode) *gdb.ConfigNode {
	const linkPattern = `(\w+):([\w\-\$]*):(.*?)@(\w+?)\((.+?)\)/{0,1}([^\?]*)\?{0,1}(.*)`
	const defaultCharset = `utf8`
	const defaultProtocol = `tcp`

	var match []string
	if node.Link != "" {
		match, _ = gregex.MatchString(linkPattern, node.Link)
		if len(match) > 5 {
			node.Type = match[1]
			node.User = match[2]
			node.Pass = match[3]
			node.Protocol = match[4]
			array := gstr.Split(match[5], ":")
			if len(array) == 2 && node.Protocol != "file" {
				node.Host = array[0]
				node.Port = array[1]
				node.Name = match[6]
			} else {
				node.Name = match[5]
			}
			if len(match) > 6 && match[7] != "" {
				node.Extra = match[7]
			}
			node.Link = ""
		}
	}
	if node.Extra != "" {
		if m, _ := gstr.Parse(node.Extra); len(m) > 0 {
			_ = gconv.Struct(m, &node)
		}
	}
	// Default value checks.
	if node.Charset == "" {
		node.Charset = defaultCharset
	}
	if node.Protocol == "" {
		node.Protocol = defaultProtocol
	}
	return node
}
