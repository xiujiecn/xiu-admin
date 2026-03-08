// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"slices"
	"strings"
	"text/template"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/event"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/do"
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/service"
	"xiuadmin/utility"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysTenant struct {
}

func NewSysTenant() *sSysTenant {
	return &sSysTenant{}
}

func init() {
	s := NewSysTenant()
	service.RegisterSysTenant(s)
}
func (l *sSysTenant) ModelQuery(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		option = append(option, &handler.Option{
			FilterTenant: false,
		})
	}
	return handler.Model(service.MemoryDB().DB(ctx).Ctx(ctx).Model(dao.SysTenant.Table()), option...)
}

// 获取租户信息
func (l *sSysTenant) View(ctx context.Context, param *model.SysTenantViewParam) (data *model.SysTenantViewModel, err error) {
	db := l.ModelQuery(ctx)
	if param.Id != 0 {
		db = db.Where(dao.SysTenant.Columns().Id, param.Id)
	} else if param.TenantId != "" {
		db = db.Where(dao.SysTenant.Columns().TenantId, param.TenantId)
	} else {
		return nil, errors.New("参数错误")
	}
	err = db.Scan(&data)
	return
}

// 获取租户列表
func (l *sSysTenant) List(ctx context.Context, param *model.SysTenantListParam) (data []*model.SysTenantListModel, total int, err error) {

	// db := dao.SysTenant.Ctx(ctx)
	// 内存数据库优化
	db := l.ModelQuery(ctx)
	if param.TenantId != "" {
		db = db.Where(dao.SysTenant.Columns().TenantId, param.TenantId)
	}
	if param.ContactUserName != "" {
		db = db.WhereLike(dao.SysTenant.Columns().ContactUserName, "%"+param.ContactUserName+"%")
	}
	if param.ContactPhone != "" {
		db = db.WhereLike(dao.SysTenant.Columns().ContactPhone, "%"+param.ContactPhone+"%")
	}
	if param.CompanyName != "" {
		db = db.WhereLike(dao.SysTenant.Columns().CompanyName, "%"+param.CompanyName+"%")
	}
	if param.LicenseNumber != "" {
		db = db.WhereLike(dao.SysTenant.Columns().LicenseNumber, "%"+param.LicenseNumber+"%")
	}
	if param.Status != "" {
		db = db.Where(dao.SysTenant.Columns().Status, param.Status)
	}
	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}

	err = db.Page(param.Page, param.PageSize).OrderAsc(dao.SysTenant.Columns().PackageId).Scan(&data)
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func (l *sSysTenant) Add(ctx context.Context, param *model.SysTenantAddParam) (output *model.SysTenantAddModel, err error) {
	data := &do.SysTenant{}
	gconv.Struct(param, &data)
	data.CreatedAt = gtime.Now()
	data.CreatedBy = contexts.GetUserId(ctx)
	data.CreatedDept = contexts.GetDeptId(ctx)
	// 查询租户套餐
	tenantPackage, err := service.SysTenantPackage().View(ctx, &model.SysTenantPackageViewParam{
		PackageId: param.PackageId,
	})
	if err != nil || tenantPackage == nil {
		g.Log().Errorf(ctx, "sSysTenant.Add SysTenantPackage().View err: %v, packageId: %d, param: %+v",
			err, param.PackageId, param)
		err = errors.New("租户套餐不存在")
		return nil, err
	}
	if tenantPackage.MenuIds == "" {
		err = errors.New("租户套餐菜单为空")
		return nil, err
	}
	menuIds := strings.Split(tenantPackage.MenuIds, ",")

	tenantId := ""
	deptId := int64(0)
	roleId := int64(0)
	userId := int64(0)
	sysConfigIds := make([]int64, 0)
	err = g.DB().Transaction(context.TODO(), func(ctx context.Context, tx gdb.TX) error {
		data.TenantId = "100000"
		id, err := tx.Ctx(ctx).Model(dao.SysTenant.Table()).Data(data).OmitNil().InsertAndGetId()
		if err != nil {
			g.Log().Errorf(ctx, "sSysTenant.Add SysTenant.Table().InsertAndGetId err: %v, data: %+v", err, data)
			return err
		}
		tenantId = fmt.Sprintf("%06d", 100000+id)

		// 创建部门
		dataDeptInsert := do.SysDept{}
		dataDeptInsert.TenantId = tenantId
		dataDeptInsert.DeptName = param.CompanyName
		dataDeptInsert.ParentId = 0
		dataDeptInsert.OrderNum = 0
		dataDeptInsert.Status = consts.SysDeptStatusNormal
		dataDeptInsert.CreatedBy = contexts.GetUserId(ctx)
		dataDeptInsert.CreatedAt = gtime.Now()
		dataDeptInsert.UpdatedBy = contexts.GetUserId(ctx)
		dataDeptInsert.UpdatedAt = gtime.Now()
		deptId, err = tx.Ctx(ctx).Model(dao.SysDept.Table()).Data(dataDeptInsert).OmitNil().InsertAndGetId()
		if err != nil {
			g.Log().Errorf(ctx, "sSysTenant.Add SysDept.Table().InsertAndGetId err: %v, data: %+v", err, dataDeptInsert)
			return err
		}

		// 创建角色
		dataRoleInsert := do.SysRole{}
		dataRoleInsert.TenantId = tenantId
		dataRoleInsert.RoleName = "租户管理员"
		dataRoleInsert.RoleKey = "tenantAdmin"
		dataRoleInsert.RoleSort = 1
		dataRoleInsert.Status = "0"
		dataRoleInsert.DataScope = consts.SysRoleDataScopeAll
		dataRoleInsert.MenuCheckStrictly = 0
		dataRoleInsert.DeptCheckStrictly = 0
		dataRoleInsert.CreatedBy = contexts.GetUserId(ctx)
		dataRoleInsert.CreatedAt = gtime.Now()
		dataRoleInsert.UpdatedBy = contexts.GetUserId(ctx)
		dataRoleInsert.UpdatedAt = gtime.Now()
		roleId, err = tx.Ctx(ctx).Model(dao.SysRole.Table()).Data(dataRoleInsert).OmitNil().InsertAndGetId()
		if err != nil {
			g.Log().Errorf(ctx, "sSysTenant.Add SysRole.Table().InsertAndGetId err: %v, data: %+v", err, dataRoleInsert)
			return err
		}

		// 角色关联菜单
		dataRoleMenuInserts := []*do.SysRoleMenu{}
		for _, menuId := range menuIds {
			dataRoleMenuInserts = append(dataRoleMenuInserts, &do.SysRoleMenu{
				RoleId: roleId,
				MenuId: menuId,
			})
		}
		_, err = tx.Ctx(ctx).Model(dao.SysRoleMenu.Table()).Data(dataRoleMenuInserts).OmitNil().Insert()
		if err != nil {
			g.Log().Errorf(ctx, "sSysTenant.Add SysRoleMenu.Table().Insert err: %v, data: %+v", err, dataRoleMenuInserts)
			return err
		}

		// 创建系统管理员
		salt := utility.RandomString(5)
		password := utility.PasswordEncrypt(param.Password, salt)

		dataUserInsert := do.SysUser{}
		dataUserInsert.TenantId = tenantId
		dataUserInsert.DeptId = deptId
		dataUserInsert.NickName = param.ContactUserName
		dataUserInsert.UserName = param.Username
		dataUserInsert.Password = password
		dataUserInsert.UserType = consts.SysUserTypeSys
		dataUserInsert.Salt = salt
		dataUserInsert.Phonenumber = param.ContactPhone
		dataUserInsert.Email = ""
		dataUserInsert.Sex = consts.SysUserSexUnknown
		dataUserInsert.Status = consts.SysUserStatusNormal
		dataUserInsert.CreatedDept = contexts.GetDeptId(ctx)
		dataUserInsert.CreatedBy = contexts.GetUserId(ctx)
		dataUserInsert.CreatedAt = gtime.Now()
		dataUserInsert.UpdatedBy = contexts.GetUserId(ctx)
		dataUserInsert.UpdatedAt = gtime.Now()

		userId, err = tx.Ctx(ctx).Model(dao.SysUser.Table()).Data(dataUserInsert).OmitNil().InsertAndGetId()
		if err != nil {
			g.Log().Errorf(ctx, "sSysTenant.Add SysUser.Table().InsertAndGetId err: %v, data: %+v", err, dataUserInsert)
			return err
		}
		// 修改租户信息
		_, err = tx.Ctx(ctx).Model(dao.SysTenant.Table()).Data(g.Map{
			dao.SysTenant.Columns().TenantId:    tenantId,
			dao.SysTenant.Columns().AdminRoleId: roleId,
			dao.SysTenant.Columns().AdminDeptId: deptId,
			dao.SysTenant.Columns().AdminUserId: userId,
		}).Where(dao.SysTenant.Columns().Id, id).Update()
		if err != nil {
			g.Log().Errorf(ctx, "sSysTenant.Add SysTenant.Table().Update err: %v, data: %+v", err, data)
			return err
		}
		// 修改部门负责人
		dataDeptUpdate := do.SysDept{}
		dataDeptUpdate.DeptId = deptId
		dataDeptUpdate.Leader = userId
		dataDeptUpdate.UpdatedBy = contexts.GetUserId(ctx)
		dataDeptUpdate.UpdatedAt = gtime.Now()
		_, err = tx.Ctx(ctx).Model(dao.SysDept.Table()).Data(dataDeptUpdate).OmitNil().Where(dao.SysDept.Columns().DeptId, deptId).Update()
		if err != nil {
			g.Log().Errorf(ctx, "sSysTenant.Add SysDept.Table().Update err: %v, data: %+v", err, dataDeptUpdate)
			return err
		}
		output = &model.SysTenantAddModel{
			Id: id,
		}
		// 角色关联用户
		dataUserRoleInserts := []*do.SysUserRole{}
		dataUserRoleInserts = append(dataUserRoleInserts, &do.SysUserRole{
			UserId: userId,
			RoleId: roleId,
		})
		_, err = tx.Ctx(ctx).Model(dao.SysUserRole.Table()).Data(dataUserRoleInserts).OmitNil().Insert()
		if err != nil {
			g.Log().Errorf(ctx, "sSysTenant.Add SysUserRole.Table().Insert err: %v, data: %+v", err, dataUserRoleInserts)
			return err
		}
		// 插入系统参数
		sysConfigIds, err = l.BatchInsertSysParam(ctx, tx, tenantId, userId, deptId)
		if err != nil {
			g.Log().Errorf(ctx, "sSysTenant.Add batchInsertSysParam err: %v, tenantId: %s, createdBy: %d, createdDept: %d", err, tenantId, contexts.GetUserId(ctx), contexts.GetDeptId(ctx))
			return err
		}
		return nil
	})
	if err != nil {
		g.Log().Errorf(ctx, "sSysTenant.Add err: %v", err)
		return nil, err
	}
	// 更新缓存
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysDeptCreate, deptId)
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysRoleCreate, roleId)
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysTenantCreate, output.Id)
	event.EventsInstance().Emit(ctx, consts.EventKeyUserCreate, userId)
	for _, sysConfigId := range sysConfigIds {
		event.EventsInstance().Emit(ctx, consts.EventKeyDBSysConfigCreate, sysConfigId)
	}
	return
}

func (l *sSysTenant) Edit(ctx context.Context, param *model.SysTenantEditParam) (output *model.SysTenantEditModel, err error) {
	if param.Id == 0 {
		return nil, errors.New("参数错误")
	}
	data := &do.SysTenant{}
	gconv.Struct(param, &data)
	data.UpdatedAt = gtime.Now()
	data.UpdatedBy = contexts.GetUserId(ctx)
	db := dao.SysTenant.Ctx(ctx)
	db = db.Where(dao.SysTenant.Columns().Id, param.Id)
	_, err = db.Data(data).OmitNil().Update()
	if err != nil {
		return nil, err
	}
	output = &model.SysTenantEditModel{
		Id: param.Id,
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysTenantUpdate, param.Id)
	return
}

func (l *sSysTenant) Delete(ctx context.Context, param *model.SysTenantDeleteParam) (output *model.SysTenantDeleteModel, err error) {
	if len(param.Ids) == 0 {
		return nil, errors.New("参数错误")
	}
	if slices.Contains(param.Ids, 1) {
		return nil, errors.New("不能删除默认租户")
	}
	data := &do.SysTenant{}
	gconv.Struct(param, &data)
	data.DeletedAt = gtime.Now()
	data.DeletedBy = contexts.GetUserId(ctx)

	db := dao.SysTenant.Ctx(ctx)
	db = db.WhereIn(dao.SysTenant.Columns().Id, param.Ids)
	_, err = db.Data(data).OmitNil().Update()
	if err != nil {
		return nil, err
	}
	output = &model.SysTenantDeleteModel{
		Ids: param.Ids,
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysTenantDelete, param.Ids)
	return
}

func (l *sSysTenant) Status(ctx context.Context, param *model.SysTenantStatusParam) (output *model.SysTenantStatusModel, err error) {
	if param.Id == 0 {
		return nil, errors.New("参数错误")
	}
	data := &do.SysTenant{}
	gconv.Struct(param, &data)
	data.UpdatedAt = gtime.Now()
	data.UpdatedBy = contexts.GetUserId(ctx)
	db := dao.SysTenant.Ctx(ctx)
	db = db.Where(dao.SysTenant.Columns().Id, param.Id)
	_, err = db.Data(data).OmitNil().Update()
	if err != nil {
		return nil, err
	}
	output = &model.SysTenantStatusModel{
		Id: param.Id,
	}
	event.EventsInstance().Emit(ctx, consts.EventKeyDBSysTenantUpdate, param.Id)
	return
}

// 新建租户插入系统参数表数据
func (l *sSysTenant) BatchInsertSysParam(ctx context.Context, tx gdb.TX, tenantId string, createdBy int64, createdDept int64) (ids []int64, err error) {
	sql := `
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '主框架页-默认皮肤样式名称', 'sys.index.skinName', 'skin-blue', 'Y', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow');
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '用户管理-账号初始密码', 'sys.user.initPassword', '123456', 'Y', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), '初始化密码 123456');
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '主框架页-侧边栏主题', 'sys.index.sideTheme', 'theme-dark', 'Y', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), '深色主题theme-dark，浅色主题theme-light');
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '账号自助-是否开启用户注册功能', 'sys.account.registerUser', 'false', 'Y', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), '是否开启注册用户功能（true开启，false关闭）');
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', 'OSS预览列表资源开关', 'sys.oss.previewListResource', 'true', 'Y', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), 'true:开启, false:关闭');
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '上传文件类型', 'sys.oss.fileType', 'txt,md,doc,docx,pdf,xls,xlsx,ppt,pptx,txt,jpg,jpeg,png,gif,mp4,avi,mov,mp3,wav,m4a,csv,json,application/json,bin,hex,fw,zip', 'Y', 103, 1, SYSDATE(), '');
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '上传图片类型', 'sys.oss.imgType', 'jpg,jpeg,png,gif,webp', 'Y', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), '');
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '上传文件大小', 'sys.oss.fileSize', '20M', 'Y', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), '');
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '上传图片大小', 'sys.oss.imgSize', '5M', 'Y', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), '');
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '上传文件Url路径', 'sys.oss.urlPath', 'resource/upload|/upload', 'Y', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), '\"http://ab.com/\" 或 \"/\" 或 \"resource/upload|http://ab.com/upload\"');
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '是否启动在线用户强制退出', 'sys.online.forceLogout', 'true', 'Y', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), 'true开启,false关闭');
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '系统启用的监控通道', 'iot.default.channel.gb28181', '0', 'Y', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), '');
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '报警资源匹配时间', 'iot.alarm.image.timesec', '10', 'N', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), NULL);
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '自主注册机构ID', 'iot.register.dept', '0', 'N', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), NULL);
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '自主注册角色ID', 'iot.register.role', '0', 'N', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), NULL);
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '系统默认MQTT渠道', 'iot.default.channel.mqtt', '0', 'N', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), NULL);
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '系统默认TCP渠道', 'iot.default.channel.tcp', '0', 'N', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), NULL);
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '系统默认UDP渠道', 'iot.default.channel.udp', '0', 'N', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), NULL);
INSERT INTO sys_config (tenant_id, config_name, config_key, config_value, config_type, created_dept, created_by, created_at, remark) VALUES ('{{.TenantId}}', '项目API默认地址', 'iot.project.api.default.addr', 'http://127.0.0.1:8510', 'Y', {{.CreatedDept}}, {{.CreatedBy}}, SYSDATE(), '');
`
	tmpl, err := template.New("example").Parse(sql)
	if err != nil {
		panic(err)
	}
	data := struct {
		TenantId    string
		CreatedBy   int64
		CreatedDept int64
	}{
		TenantId:    tenantId,
		CreatedBy:   createdBy,
		CreatedDept: createdDept,
	}
	var b strings.Builder
	err = tmpl.Execute(&b, data)
	if err != nil {
		return nil, err
	}
	sql = b.String()
	g.Log().Infof(ctx, "sSysTenant.batchInsertSysParam sql: %s", sql)

	sqlArr := strings.Split(sql, ";")
	ids = make([]int64, 0)
	for _, sql := range sqlArr {
		sql = strings.ReplaceAll(sql, "\n", "")
		sql = strings.TrimLeft(sql, " ")
		sql = strings.TrimRight(sql, " ")
		if sql == "" {
			continue
		}
		res, err := tx.Ctx(ctx).Exec(sql)
		if err != nil {
			g.Log().Errorf(ctx, "sSysTenant.batchInsertSysParam tx.Ctx(ctx).Exec err: %v, sql: %s", err, sql)
			return nil, err
		}
		lastInsertId, err := res.LastInsertId()
		if err != nil {
			g.Log().Errorf(ctx, "sSysTenant.batchInsertSysParam LastInsertId err: %v, sql: %s", err, sql)
			return nil, err
		}
		ids = append(ids, lastInsertId)
	}

	g.Log().Infof(ctx, "sSysTenant.batchInsertSysParam ids: %v", ids)
	return ids, nil
}

// 同步租户菜单
func (l *sSysTenant) SyncTenantMenu(ctx context.Context, tenantId string) (err error) {
	tenantInfo, err := l.View(ctx, &model.SysTenantViewParam{
		TenantId: tenantId,
	})
	if err != nil {
		g.Log().Errorf(ctx, "sSysTenant.SyncTenantMenu SysTenant().View err: %v, tenantId: %s", err, tenantId)
		return err
	}
	tenantPackage, err := service.SysTenantPackage().View(ctx, &model.SysTenantPackageViewParam{
		PackageId: tenantInfo.PackageId,
	})
	if err != nil {
		g.Log().Errorf(ctx, "sSysTenant.SyncTenantMenu SysTenantPackage().View err: %v, tenantId: %s, packageId: %d", err, tenantId, tenantInfo.PackageId)
		return err
	}
	if tenantPackage == nil {
		return errors.New("租户套餐不存在")
	}
	menuIds := strings.Split(tenantPackage.MenuIds, ",")
	// 获取角色对应菜单
	roleMenuList, err := service.SysRole().GetRoleMenu(ctx, tenantInfo.AdminRoleId)
	if err != nil {
		g.Log().Errorf(ctx, "sSysTenant.SyncTenantMenu SysRole().GetRoleMenu err: %v, tenantId: %s, roleId: %d", err, tenantId, tenantInfo.AdminRoleId)
		return err
	}
	insertMenuIds := make([]do.SysRoleMenu, 0)
	for _, menuId := range menuIds {
		found := false
		for _, roleMenu := range roleMenuList {
			if roleMenu.MenuId == gconv.Int64(menuId) {
				found = true
				break
			}
		}
		if !found {
			insertMenuIds = append(insertMenuIds, do.SysRoleMenu{
				RoleId: tenantInfo.AdminRoleId,
				MenuId: gconv.Int64(menuId),
			})
		}
	}
	delMenuIds := make([]int64, 0)
	for _, roleMenu := range roleMenuList {
		if !slices.Contains(menuIds, gconv.String(roleMenu.MenuId)) {
			delMenuIds = append(delMenuIds, roleMenu.MenuId)
			continue
		}
	}
	if len(delMenuIds) > 0 {
		_, err = dao.SysRoleMenu.Ctx(ctx).WhereIn(dao.SysRoleMenu.Columns().RoleId, tenantInfo.AdminRoleId).
			WhereIn(dao.SysRoleMenu.Columns().MenuId, delMenuIds).Delete()
		if err != nil {
			g.Log().Errorf(ctx, "sSysTenant.SyncTenantMenu SysRoleMenu.Ctx(ctx).Delete err: %v, tenantId: %s, roleId: %d, delMenuIds: %v", err, tenantId, tenantInfo.AdminRoleId, delMenuIds)
			return err
		}
	}
	if len(insertMenuIds) > 0 {
		_, err = dao.SysRoleMenu.Ctx(ctx).Data(insertMenuIds).Insert()
		if err != nil {
			g.Log().Errorf(ctx, "sSysTenant.SyncTenantMenu SysRoleMenu.Ctx(ctx).Insert err: %v, tenantId: %s, roleId: %d, insertMenuIds: %v", err, tenantId, tenantInfo.AdminRoleId, insertMenuIds)
			return err
		}
	}
	return nil
}

// 查询用户ID是否是租户管理员
func (l *sSysTenant) IsTenantAdmin(ctx context.Context, userIds []int64) (isAdminUserIds []int64, err error) {
	entityList := make([]*entity.SysTenant, 0)
	err = dao.SysTenant.Ctx(ctx).WhereIn(dao.SysTenant.Columns().AdminUserId, userIds).Scan(&entityList)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	isAdminUserIds = make([]int64, 0)
	for _, entity := range entityList {
		isAdminUserIds = append(isAdminUserIds, entity.AdminUserId)
	}
	return isAdminUserIds, nil
}
