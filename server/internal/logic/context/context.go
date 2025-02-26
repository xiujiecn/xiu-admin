package context

import (
	"context"
	"slices"
	"xiujieadmin/internal/consts"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

type sContext struct {
}

func NewContext() *sContext {
	return &sContext{}
}

func init() {
	service.RegisterContext(NewContext())
}

// Init初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *sContext) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// Get 获取上下文变量；
func (s *sContext) Get(ctx context.Context) *model.Context {
	customCtx, ok := ctx.Value(consts.ContextKey).(*model.Context)
	if !ok {
		return nil
	}
	return customCtx
}

// Set 设置上下文变量
func (s *sContext) SetUser(ctx context.Context, user *model.Identity) {
	customCtx := s.Get(ctx)
	customCtx.User = user
}

// GetUser 获取上下文变量中的用户信息
func (s *sContext) GetUser(ctx context.Context) *model.Identity {
	customCtx := s.Get(ctx)
	if customCtx == nil {
		return nil
	}
	return customCtx.User
}

// GetUserId 获取上下文变量中的用户ID
func (s *sContext) GetUserId(ctx context.Context) int64 {
	user := s.GetUser(ctx)
	if user == nil {
		return 0
	}
	return user.ID
}

// GetUserName 获取上下文变量中的用户名
func (s *sContext) GetUserName(ctx context.Context) string {
	user := s.GetUser(ctx)
	if user == nil {
		return ""
	}
	return user.Username
}

// GetNickName 获取上下文变量中的用户昵称
func (s *sContext) GetNickName(ctx context.Context) string {
	user := s.GetUser(ctx)
	if user == nil {
		return ""
	}
	return user.NickName
}

// GetDeptId 获取上下文变量中的部门ID
func (s *sContext) GetDeptId(ctx context.Context) int64 {
	user := s.GetUser(ctx)
	if user == nil {
		return 0
	}
	return user.DeptId
}

// GetTenantId 获取上下文变量中的租户ID
func (s *sContext) GetTenantId(ctx context.Context) string {
	user := s.GetUser(ctx)
	if user == nil {
		return ""
	}
	return user.TenantId
}

// GetAuthorityIds 获取上下文变量中的角色ID
func (s *sContext) GetAuthorityIds(ctx context.Context) []int64 {
	user := s.GetUser(ctx)
	if user == nil {
		return nil
	}
	return user.AuthorityIds
}

// IsSystemTenant 判断是否是系统租户
func (s *sContext) IsSystemTenant(ctx context.Context) bool {
	tenantId := s.GetTenantId(ctx)
	return tenantId == consts.DefaultSystemTenantCode
}

// IsSuperAdmin 判断是否是超级管理员角色
func (s *sContext) IsSuperAdmin(ctx context.Context) bool {
	authorityIds := s.GetAuthorityIds(ctx)
	return slices.Contains(authorityIds, consts.SuperAdminRoleId)
}
