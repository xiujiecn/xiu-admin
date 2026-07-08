// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "xiuadmin/api/system/v1"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/model/request"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ISysAuth interface {
		Login(ctx context.Context, param *model.LoginParams) (res *model.LoginUserOut, token string, err error)
		// 生成Token
		GenerateToken(ctx context.Context, user *model.LoginUserOut) (claims *model.CustomClaims, token string, err error)
		// 解析Token
		ParseToken(ctx context.Context, token string) (claims *model.CustomClaims, err error)
		// 删除Token
		DeleteToken(ctx context.Context, token string) (err error)
		// 获取token
		GetAccessToken(ctx context.Context) (token string, err error)
		// 根据Token获取当前登录用户信息
		GetCurrentUser(ctx context.Context) (claims *model.CustomClaims, err error)
		// 获取用户权限码
		GetUserAccessCodeList(ctx context.Context, userId int64) (accessCodeList []string, menuRoleAccessCodeList []string, err error)
		// 根据openId登录
		LoginByOpenId(ctx context.Context, social *model.SysSocialListModel) (res *model.LoginUserOut, token string, err error)
	}
	ISysCaptcha interface {
		// 生成验证码
		GenerateCaptcha(ctx context.Context) (key string, image string, err error)
		// 验证验证码
		VerifyCaptcha(ctx context.Context, key string, value string) (err error)
	}
	ISysClient interface {
		List(ctx context.Context, query *model.SysClientListParam) (items []*model.SysClientListModel, total int, err error)
		View(ctx context.Context, param *model.SysClientViewParam) (item *model.SysClientViewModel, err error)
		Add(ctx context.Context, param *model.SysClientAddParam) (output *model.SysClientAddModel, err error)
		Edit(ctx context.Context, param *model.SysClientEditParam) (output *model.SysClientEditModel, err error)
		Delete(ctx context.Context, param *model.SysClientDeleteParam) (output *model.SysClientDeleteModel, err error)
		Status(ctx context.Context, param *model.SysClientStatusParam) (output *model.SysClientStatusModel, err error)
	}
	ISysConfig interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		ModelQuery(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, param *model.SysConfigListParam) (items []*model.SysConfigListModel, total int, err error)
		Add(ctx context.Context, param *model.SysConfigAddParam) (output *model.SysConfigAddModel, err error)
		Edit(ctx context.Context, param *model.SysConfigEditParam) (output *model.SysConfigEditModel, err error)
		Delete(ctx context.Context, param *model.SysConfigDeleteParam) (output *model.SysConfigDeleteModel, err error)
		View(ctx context.Context, param *model.SysConfigViewParam) (output *model.SysConfigViewModel, err error)
		GetConfigByKey(ctx context.Context, configKey string) (config *entity.SysConfig, err error)
	}
	ISysDept interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		ModelQuery(ctx context.Context, option ...*handler.Option) *gdb.Model
		GetDeptList(ctx context.Context, query model.SysDeptListParam) (items []*model.SysDeptListModel, total int, err error)
		GetDeptById(ctx context.Context, id int64) (dept *model.SysDeptViewModel, err error)
		// 构建树结构
		DeptTree(ctx context.Context, parentDept *model.SysDeptTreeModel, deptList []*model.SysDeptListModel, ancestors string) (data []*model.SysDeptTreeModel, err error)
		GetDeptTree(ctx context.Context, query model.SysDeptTreeParam) (items []*model.SysDeptTreeModel, err error)
		// 递归构建结构
		RecursionDeptIds(ctx context.Context, parentId int64, deptList []*model.SysDeptListModel, data *[]int64) (err error)
		// 根据父部门id获取部门列表
		GetDeptIdsByParentId(ctx context.Context, parentId int64) (ids []int64, err error)
		AddDept(ctx context.Context, dept *model.SysDeptAddModel) (deptId int64, err error)
		EditDept(ctx context.Context, dept *model.SysDeptEditModel) (deptId int64, err error)
		DeleteDept(ctx context.Context, dept *model.SysDeptDeleteModel) (deptId int64, err error)
		// 刷新部门 ancestors
		RefreshDeptAncestors(ctx context.Context) (err error)
		GetParentIDAncestors(ctx context.Context, depts []*model.SysDeptViewModel, printId int64, ancestors *string) (err error)
		GetDeptListByIds(ctx context.Context, ids []int64) (depts []*model.SysDeptListModel, err error)
		// 验证上级是否是公司
		ValidateParentIsCompany(ctx context.Context, parentId int64) (isCompany bool, err error)
		// 获取组织公司Map
		GetDeptCompanyMap(ctx context.Context) (map[int64]string, error)
		// 获取上级公司组织信息
		GetParentCompanyInfo(ctx context.Context, deptId int64) *model.SysDeptListModel
	}
	ISysDictData interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		ModelQuery(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, param *model.SysDictDataListParam) (items []model.SysDictDataListModel, total int, err error)
		View(ctx context.Context, param *model.SysDictDataViewParam) (data *model.SysDictDataViewModel, err error)
		Add(ctx context.Context, param *model.SysDictDataAddParam) (output *model.SysDictDataAddModel, err error)
		Edit(ctx context.Context, param *model.SysDictDataEditParam) (output *model.SysDictDataEditModel, err error)
		Delete(ctx context.Context, param *model.SysDictDataDeleteParam) (output *model.SysDictDataDeleteModel, err error)
		GetDictLabel(ctx context.Context, dictType string, dictCode string) string
		GetDictListByTypes(ctx context.Context, dictTypes []string) (dictDataList []model.SysDictDataListModel, err error)
	}
	ISysDictType interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		ModelQuery(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, param *model.SysDictTypeListParam) (items []model.SysDictTypeListModel, total int, err error)
		View(ctx context.Context, param *model.SysDictTypeViewParam) (data *model.SysDictTypeViewModel, err error)
		Add(ctx context.Context, param *model.SysDictTypeAddParam) (output *model.SysDictTypeAddModel, err error)
		Edit(ctx context.Context, param *model.SysDictTypeEditParam) (output *model.SysDictTypeEditModel, err error)
		Delete(ctx context.Context, param *model.SysDictTypeDeleteParam) (output *model.SysDictTypeDeleteModel, err error)
	}
	ISysJob interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, query *model.SysJobListParam, pageInfo *request.PageInfo) (Data []*model.SysJobListModel, total int, err error)
		View(ctx context.Context, jobId int64) (Data *model.SysJobViewModel, err error)
		Add(ctx context.Context, jobAdd *model.SysJobAddModel) (LastInsertId int64, err error)
		Update(ctx context.Context, jobUpdate *model.SysJobUpdateModel) (RowsAffected int64, err error)
		UpdateStatus(ctx context.Context, jobUpdate *model.SysJobUpdateStatusModel) (RowsAffected int64, err error)
		Delete(ctx context.Context, jobDelete *model.SysJobDeleteModel) (RowsAffected int64, err error)
		Exec(ctx context.Context, jobId int64) error
		// 查询所有的状态正常的任务列表，并进行初始注册
		InitRegister() error
	}
	ISysLogininfor interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, param *model.SysLogininforListParam) (items []*model.SysLogininforListModel, total int, err error)
		AddLogininfor(ctx context.Context, logininfor *model.SysLogininforAddModel) (id int64, err error)
		Delete(ctx context.Context, param *model.SysLogininforDeleteParam) (output *model.SysLogininforDeleteModel, err error)
	}
	ISysMenu interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		ModelQuery(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, param *model.SysMenuListParam) (data []*model.SysMenuListModel, total int, err error)
		// 获取租户菜单列表， 系统租户返回所有菜单，其他租户返回当前租户菜单
		GetTenantMenu(ctx context.Context, query *model.SysMenuListParam) (data []*model.SysMenuListModel, total int, err error)
		// 构建树结构
		MenuTree(ctx context.Context, parentMenu *model.SysMenuTree, menuList []*entity.SysMenu) (data []*model.SysMenuTree, err error)
		// 获取用户动态路由列表
		GetUserMenu(ctx context.Context) (data []*entity.SysMenu, err error)
		// 构建用户动态路由树
		BuildUserMenuTree(ctx context.Context, parentMenu *v1.RouteMenu, menuList []*entity.SysMenu, allPath string) (data v1.MenuAllRes, err error)
		// 获取用户动态路由树
		GetUserMenuTree(ctx context.Context) (data v1.MenuAllRes, err error)
		GetSysMenuView(ctx context.Context, menuId int64) (data *model.SysMenuViewModel, err error)
		UpdateSysMenu(ctx context.Context, menu *model.SysMenuUpdateModel) (data *model.SysMenuViewModel, err error)
		AddSysMenu(ctx context.Context, menu *model.SysMenuAddModel) (data *model.SysMenuViewModel, err error)
		DeleteSysMenu(ctx context.Context, menuId int64) (err error)
		GetFastList(ctx context.Context) (res map[int64]*entity.SysMenu, err error)
	}
	ISysNotice interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, param *model.SysNoticeListParam) (items []*model.SysNoticeListModel, total int, err error)
		Add(ctx context.Context, param *model.SysNoticeAddParam) (err error)
		Edit(ctx context.Context, param *model.SysNoticeEditParam) (err error)
		Delete(ctx context.Context, param *model.SysNoticeDeleteParam) (err error)
		View(ctx context.Context, param *model.SysNoticeViewParam) (data *model.SysNoticeViewModel, err error)
	}
	ISysNoticeUser interface {
		// Model 用户通知公告表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取用户通知公告表列表
		List(ctx context.Context, in *model.SysNoticeUserListParam) (list []*model.SysNoticeUserListModel, totalCount int, err error)
		// Export 导出用户通知公告表
		Export(ctx context.Context, in *model.SysNoticeUserListParam) (err error)
		// Edit 修改/新增用户通知公告表
		Edit(ctx context.Context, in *model.SysNoticeUserEditParam) (err error)
		// Delete 删除用户通知公告表
		Delete(ctx context.Context, in *model.SysNoticeUserDeleteParam) (err error)
		// View 获取用户通知公告表指定信息
		View(ctx context.Context, in *model.SysNoticeUserViewParam) (res *model.SysNoticeUserViewModel, err error)
		// Status 更新用户通知公告表状态
		Status(ctx context.Context, in *model.SysNoticeUserStatusParam) (err error)
		// Read 已读
		Read(ctx context.Context, in *model.SysNoticeUserReadParam) (err error)
	}
	ISysOperLog interface {
		GetOperLogList(ctx context.Context, query *model.SysOperLogListParam, page *request.PageInfo) (items []*model.SysOperLogListModel, total int, err error)
		AnalysisLog(ctx context.Context) (data *model.SysOperLogAddParam, err error)
		ClearOperationLogByDays(ctx context.Context, days int) error
		RealWrite(ctx context.Context, data entity.SysOperLog) (err error)
	}
	ISysOss interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, param *model.SysOssListParam, pageInfo *request.PageInfo) (items []*model.SysOssListModel, total int, err error)
		View(ctx context.Context, param *model.SysOssViewParam) (output *model.SysOssViewModel, err error)
		Download(ctx context.Context, param *model.SysOssDownloadParam) (output *model.SysOssDownloadModel, err error)
		Delete(ctx context.Context, param *model.SysOssDeleteParam) (output *model.SysOssDeleteModel, err error)
		Upload(ctx context.Context, param *model.SysOssUploadParam) (output *model.SysOssUploadModel, err error)
		MoveFile(ctx context.Context, param *model.SysOssMoveFileParam) (output *model.SysOssMoveFileModel, err error)
		SaveContent(ctx context.Context, param *model.SysOssSaveContentParam) (output *model.SysOssSaveContentModel, err error)
		GetSaveFilePathConfig(ctx context.Context, newFileType int, notAddDate int) (string, error)
		MoveFileLocal(ctx context.Context, filePath string, newFileType int, useOriginalName int, subDirName string, isAddDate int) (result model.UploadResponse, err error)
		// 保存内容
		SaveContentLocal(ctx context.Context, content []byte, newFileType int, fileName string, subDirName string, isAddDate int) (result model.UploadResponse, err error)
		UploadLocal(ctx context.Context, file *ghttp.UploadFile, newFileType int, useOriginalName int, subDirName string, notAddDate int) (result model.UploadResponse, err error)
		DownloadLocal(ctx context.Context, file *entity.SysOss) (err error)
		DeleteLocal(ctx context.Context, file *entity.SysOss) (err error)
		CheckType(ctx context.Context, checkFileType string, file *ghttp.UploadFile) (err error)
		CheckSize(ctx context.Context, checkFileType string, file *ghttp.UploadFile) (err error)
		GetAllUrl(ctx context.Context, url string) (allUrl string, ossId int64, fileSize int64, OriginalName string, md5 string, hmac string, err error)
	}
	ISysOssConfig interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, param *model.SysOssConfigListParam) (items []*model.SysOssConfigListModel, total int, err error)
		View(ctx context.Context, param *model.SysOssConfigViewParam) (item *model.SysOssConfigViewModel, err error)
		Add(ctx context.Context, param *model.SysOssConfigAddParam) (item *model.SysOssConfigAddModel, err error)
		Edit(ctx context.Context, param *model.SysOssConfigEditParam) (item *model.SysOssConfigEditModel, err error)
		Delete(ctx context.Context, param *model.SysOssConfigDeleteParam) (item *model.SysOssConfigDeleteModel, err error)
		GetAllUrlByService(ctx context.Context, path string, tenantId string, ossService string) (allUrl string, err error)
	}
	ISysPost interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		ModelQuery(ctx context.Context, option ...*handler.Option) *gdb.Model
		List(ctx context.Context, query *model.SysPostListParam) (items []*model.SysPostListModel, total int, err error)
		View(ctx context.Context, param *model.SysPostViewParam) (post *model.SysPostViewModel, err error)
		Add(ctx context.Context, param *model.SysPostAddParam) (post *model.SysPostAddModel, err error)
		Edit(ctx context.Context, param *model.SysPostEditParam) (post *model.SysPostEditModel, err error)
		Delete(ctx context.Context, param *model.SysPostDeleteParam) (post *model.SysPostDeleteModel, err error)
		Export(ctx context.Context, param *model.SysPostExportParam) (post *model.SysPostExportModel, err error)
	}
	ISysQrcode interface {
		// 获取登录二维码
		GetQrcodeLogin(ctx context.Context) (res *model.QrcodeLoginModel, err error)
		// 获取登录二维码扫码结果
		GetQrcodeLoginStatus(ctx context.Context, tempUserId string) (res *model.QrcodeLoginModel, err error)
		// 获取绑定二维码
		GetQrcodeBind(ctx context.Context) (res *model.QrcodeCacheModel, err error)
		// 获取绑定二维码扫码结果
		GetQrcodeBindStatus(ctx context.Context) (res *model.QrcodeBindStatusModel, err error)
		// 扫码回调
		QrcodeScanCallback(ctx context.Context, param *model.QrcodeScanCallbackParam) (err error)
		// 扫码后选择用户登录
		QrcodeLoginSelectUserId(ctx context.Context, param *model.QrcodeLoginSelectUserIdParam) (res *model.QrcodeLoginAndBindOpenIdModel, err error)
		// 登录并绑定
		QrcodeLoginAndBindOpenId(ctx context.Context, param *model.QrcodeLoginAndBindOpenIdParam) (res *model.QrcodeLoginAndBindOpenIdModel, err error)
		// 注册并绑定
		QrcodeRegisterAndBindOpenId(ctx context.Context, param *model.QrcodeRegisterAndBindOpenIdParam) (res *model.QrcodeLoginAndBindOpenIdModel, err error)
	}
	ISysRole interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		ModelQuery(ctx context.Context, option ...*handler.Option) *gdb.Model
		// 获取租户下角色列表
		List(ctx context.Context, param *model.SysRoleListParam) (res []*model.SysRoleListModel, total int, err error)
		// 获取角色详情
		View(ctx context.Context, param *model.SysRoleViewParam) (res *model.SysRoleViewModel, err error)
		// 获取角色菜单
		GetRoleMenu(ctx context.Context, id int64) (res []*entity.SysMenu, err error)
		// 获取角色部门
		GetRoleDept(ctx context.Context, id int64) (res []*entity.SysDept, err error)
		// 获取角色列表对应菜单
		GetRoleListMenu(ctx context.Context, ids []int64) (res []*entity.SysRoleMenu, err error)
		// 新增角色
		Add(ctx context.Context, param *model.SysRoleAddParam) (role *model.SysRoleAddModel, err error)
		// 编辑角色
		Edit(ctx context.Context, param *model.SysRoleEditParam) (role *model.SysRoleEditModel, err error)
		// 删除角色
		Delete(ctx context.Context, param *model.SysRoleDeleteParam) (role *model.SysRoleDeleteModel, err error)
		RoleMenu(ctx context.Context, roleId int64, menuIds []int64, dataScopes map[int64]string) (err error)
		// 自定义角色部门数据权限
		RoleDept(ctx context.Context, roleId int64, deptIds []int64) (err error)
		// 编辑角色数据权限
		EditRoleDataScope(ctx context.Context, model *model.SysRoleDataScopeEditParam) (err error)
	}
	ISysSocial interface {
		// 查询社会化关系表
		List(ctx context.Context, query *model.SysSocialListParam, page *request.PageInfo) (items []*model.SysSocialListModel, total int, err error)
		// 删除社会化关系表
		Delete(ctx context.Context, id int64) (err error)
		Create(ctx context.Context, social *model.SysSocialSaveParam) (err error)
	}
	ISysTenant interface {
		ModelQuery(ctx context.Context, option ...*handler.Option) *gdb.Model
		// 获取租户信息
		View(ctx context.Context, param *model.SysTenantViewParam) (data *model.SysTenantViewModel, err error)
		// 获取租户列表
		List(ctx context.Context, param *model.SysTenantListParam) (data []*model.SysTenantListModel, total int, err error)
		Add(ctx context.Context, param *model.SysTenantAddParam) (output *model.SysTenantAddModel, err error)
		Edit(ctx context.Context, param *model.SysTenantEditParam) (output *model.SysTenantEditModel, err error)
		Delete(ctx context.Context, param *model.SysTenantDeleteParam) (output *model.SysTenantDeleteModel, err error)
		Status(ctx context.Context, param *model.SysTenantStatusParam) (output *model.SysTenantStatusModel, err error)
		// 新建租户插入系统参数表数据
		BatchInsertSysParam(ctx context.Context, tx gdb.TX, tenantId string, createdBy int64, createdDept int64) (ids []int64, err error)
		// 同步租户菜单
		SyncTenantMenu(ctx context.Context, tenantId string) (err error)
		// 查询用户ID是否是租户管理员
		IsTenantAdmin(ctx context.Context, userIds []int64) (isAdminUserIds []int64, err error)
	}
	ISysTenantPackage interface {
		// 获取租户套餐
		View(ctx context.Context, param *model.SysTenantPackageViewParam) (data *model.SysTenantPackageViewModel, err error)
		List(ctx context.Context, param *model.SysTenantPackageListParam) (data []*model.SysTenantPackageListModel, total int, err error)
		Add(ctx context.Context, param *model.SysTenantPackageAddParam) (output *model.SysTenantPackageAddModel, err error)
		Edit(ctx context.Context, param *model.SysTenantPackageEditParam) (output *model.SysTenantPackageEditModel, err error)
		Delete(ctx context.Context, param *model.SysTenantPackageDeleteParam) (output *model.SysTenantPackageDeleteModel, err error)
		Status(ctx context.Context, param *model.SysTenantPackageStatusParam) (output *model.SysTenantPackageStatusModel, err error)
	}
	ISysUser interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// 根据用户名获取用户信息，不验证当前租户
		GetUserByUsername(ctx context.Context, username string, tenantId string) (user *entity.SysUser, err error)
		// 根据邮箱获取用户信息，不验证当前租户
		GetUserByEmail(ctx context.Context, email string, tenantId string) (user *entity.SysUser, err error)
		// 根据手机号获取用户信息,不验证当前租户
		GetUserByPhone(ctx context.Context, phone string, tenantId string) (user *entity.SysUser, err error)
		// 根据用户名和密码获取用户信息
		GetUserByUsernameAndPassword(ctx context.Context, tenantId string, username string, password string) (user *entity.SysUser, err error)
		// 根据用户ID获取用户信息
		GetUserById(ctx context.Context, id int64) (user *model.SysUserViewModel, err error)
		// 获取用户列表
		List(ctx context.Context, page *request.PageInfo, query *model.UserListParam) (items []*model.SysUserListModel, total int, err error)
		// 新增用户
		AddUser(ctx context.Context, req *model.SysUserAddModel) (data *model.SysUserViewModel, err error)
		Profile(ctx context.Context) (user *model.UserProfileModel, err error)
		UpdateCurrentUser(ctx context.Context, req *model.UpdateCurrentUserModel) (user *model.SysUserViewModel, err error)
		UpdateCurrentUserPassword(ctx context.Context, req *model.UpdateCurrentUserPasswordModel) (err error)
		SendCurrentUserContactCode(ctx context.Context, req *model.UpdateCurrentUserContactCodeModel) (err error)
		UpdateCurrentUserContact(ctx context.Context, req *model.UpdateCurrentUserContactModel) (err error)
		UpdateUser(ctx context.Context, req *model.SysUserUpdateModel) (err error)
		DeleteUser(ctx context.Context, userIds []int64) (err error)
		ResetPassword(ctx context.Context, userId int64, password string) (err error)
		UpdateLoginInfo(ctx context.Context, userId int64, loginIp string) (err error)
		// 获取用户角色ID列表
		GetUserRoleIds(ctx context.Context, userId int64) (roleIds []int64, err error)
		// 获取用户岗位ID列表
		GetUserPostIds(ctx context.Context, userId int64) (postIds []int64, err error)
		Register(ctx context.Context, param *model.SysUserRegisterModel) (err error)
		SendRegisterCode(ctx context.Context, param *model.SysUserRegisterCodeModel) (err error)
		// Status 更新用户状态
		Status(ctx context.Context, param *model.SysUserStatusParam) (err error)
		// 批量查询用户迷你信息
		BatchGetUserMiniInfo(ctx context.Context, userIds []int64) (users []*model.SysUserMiniModel, err error)
	}
	ISysUserOnline interface {
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		Add(ctx context.Context, userOnline *model.SysUserOnlineAddModel) (err error)
		List(ctx context.Context, query *model.SysUserOnlineListParam, page *request.PageInfo) (items []*model.SysUserOnlineListModel, total int, err error)
		Delete(ctx context.Context, ids []int64) (err error)
		DeleteByToken(ctx context.Context, token string) (err error)
	}
)

var (
	localSysAuth          ISysAuth
	localSysCaptcha       ISysCaptcha
	localSysClient        ISysClient
	localSysConfig        ISysConfig
	localSysDept          ISysDept
	localSysDictData      ISysDictData
	localSysDictType      ISysDictType
	localSysJob           ISysJob
	localSysLogininfor    ISysLogininfor
	localSysMenu          ISysMenu
	localSysNotice        ISysNotice
	localSysNoticeUser    ISysNoticeUser
	localSysOperLog       ISysOperLog
	localSysOss           ISysOss
	localSysOssConfig     ISysOssConfig
	localSysPost          ISysPost
	localSysQrcode        ISysQrcode
	localSysRole          ISysRole
	localSysSocial        ISysSocial
	localSysTenant        ISysTenant
	localSysTenantPackage ISysTenantPackage
	localSysUser          ISysUser
	localSysUserOnline    ISysUserOnline
)

func SysAuth() ISysAuth {
	if localSysAuth == nil {
		panic("implement not found for interface ISysAuth, forgot register?")
	}
	return localSysAuth
}

func RegisterSysAuth(i ISysAuth) {
	localSysAuth = i
}

func SysCaptcha() ISysCaptcha {
	if localSysCaptcha == nil {
		panic("implement not found for interface ISysCaptcha, forgot register?")
	}
	return localSysCaptcha
}

func RegisterSysCaptcha(i ISysCaptcha) {
	localSysCaptcha = i
}

func SysClient() ISysClient {
	if localSysClient == nil {
		panic("implement not found for interface ISysClient, forgot register?")
	}
	return localSysClient
}

func RegisterSysClient(i ISysClient) {
	localSysClient = i
}

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}

func SysDept() ISysDept {
	if localSysDept == nil {
		panic("implement not found for interface ISysDept, forgot register?")
	}
	return localSysDept
}

func RegisterSysDept(i ISysDept) {
	localSysDept = i
}

func SysDictData() ISysDictData {
	if localSysDictData == nil {
		panic("implement not found for interface ISysDictData, forgot register?")
	}
	return localSysDictData
}

func RegisterSysDictData(i ISysDictData) {
	localSysDictData = i
}

func SysDictType() ISysDictType {
	if localSysDictType == nil {
		panic("implement not found for interface ISysDictType, forgot register?")
	}
	return localSysDictType
}

func RegisterSysDictType(i ISysDictType) {
	localSysDictType = i
}

func SysJob() ISysJob {
	if localSysJob == nil {
		panic("implement not found for interface ISysJob, forgot register?")
	}
	return localSysJob
}

func RegisterSysJob(i ISysJob) {
	localSysJob = i
}

func SysLogininfor() ISysLogininfor {
	if localSysLogininfor == nil {
		panic("implement not found for interface ISysLogininfor, forgot register?")
	}
	return localSysLogininfor
}

func RegisterSysLogininfor(i ISysLogininfor) {
	localSysLogininfor = i
}

func SysMenu() ISysMenu {
	if localSysMenu == nil {
		panic("implement not found for interface ISysMenu, forgot register?")
	}
	return localSysMenu
}

func RegisterSysMenu(i ISysMenu) {
	localSysMenu = i
}

func SysNotice() ISysNotice {
	if localSysNotice == nil {
		panic("implement not found for interface ISysNotice, forgot register?")
	}
	return localSysNotice
}

func RegisterSysNotice(i ISysNotice) {
	localSysNotice = i
}

func SysNoticeUser() ISysNoticeUser {
	if localSysNoticeUser == nil {
		panic("implement not found for interface ISysNoticeUser, forgot register?")
	}
	return localSysNoticeUser
}

func RegisterSysNoticeUser(i ISysNoticeUser) {
	localSysNoticeUser = i
}

func SysOperLog() ISysOperLog {
	if localSysOperLog == nil {
		panic("implement not found for interface ISysOperLog, forgot register?")
	}
	return localSysOperLog
}

func RegisterSysOperLog(i ISysOperLog) {
	localSysOperLog = i
}

func SysOss() ISysOss {
	if localSysOss == nil {
		panic("implement not found for interface ISysOss, forgot register?")
	}
	return localSysOss
}

func RegisterSysOss(i ISysOss) {
	localSysOss = i
}

func SysOssConfig() ISysOssConfig {
	if localSysOssConfig == nil {
		panic("implement not found for interface ISysOssConfig, forgot register?")
	}
	return localSysOssConfig
}

func RegisterSysOssConfig(i ISysOssConfig) {
	localSysOssConfig = i
}

func SysPost() ISysPost {
	if localSysPost == nil {
		panic("implement not found for interface ISysPost, forgot register?")
	}
	return localSysPost
}

func RegisterSysPost(i ISysPost) {
	localSysPost = i
}

func SysQrcode() ISysQrcode {
	if localSysQrcode == nil {
		panic("implement not found for interface ISysQrcode, forgot register?")
	}
	return localSysQrcode
}

func RegisterSysQrcode(i ISysQrcode) {
	localSysQrcode = i
}

func SysRole() ISysRole {
	if localSysRole == nil {
		panic("implement not found for interface ISysRole, forgot register?")
	}
	return localSysRole
}

func RegisterSysRole(i ISysRole) {
	localSysRole = i
}

func SysSocial() ISysSocial {
	if localSysSocial == nil {
		panic("implement not found for interface ISysSocial, forgot register?")
	}
	return localSysSocial
}

func RegisterSysSocial(i ISysSocial) {
	localSysSocial = i
}

func SysTenant() ISysTenant {
	if localSysTenant == nil {
		panic("implement not found for interface ISysTenant, forgot register?")
	}
	return localSysTenant
}

func RegisterSysTenant(i ISysTenant) {
	localSysTenant = i
}

func SysTenantPackage() ISysTenantPackage {
	if localSysTenantPackage == nil {
		panic("implement not found for interface ISysTenantPackage, forgot register?")
	}
	return localSysTenantPackage
}

func RegisterSysTenantPackage(i ISysTenantPackage) {
	localSysTenantPackage = i
}

func SysUser() ISysUser {
	if localSysUser == nil {
		panic("implement not found for interface ISysUser, forgot register?")
	}
	return localSysUser
}

func RegisterSysUser(i ISysUser) {
	localSysUser = i
}

func SysUserOnline() ISysUserOnline {
	if localSysUserOnline == nil {
		panic("implement not found for interface ISysUserOnline, forgot register?")
	}
	return localSysUserOnline
}

func RegisterSysUserOnline(i ISysUserOnline) {
	localSysUserOnline = i
}
