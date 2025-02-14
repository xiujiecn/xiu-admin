package system

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"
	"server/utility"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
)

type sSysUser struct {
}

func NewSysUser() *sSysUser {
	return &sSysUser{}
}

func init() {
	service.RegisterSysUser(NewSysUser())
}

// 根据用户名获取用户信息
func (l *sSysUser) GetUserByUsername(ctx context.Context, username string) (user *entity.SysUser, err error) {
	var data *entity.SysUser
	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserName, username).Scan(&data)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gerror.NewCode(consts.CodeUserNotFound, "账号不存在")
		}
		return nil, err
	}
	if data == nil {
		return nil, gerror.NewCode(consts.CodeUserNotFound, "账号不存在")
	}
	if data.Status == consts.UserStatusDisabled {
		return nil, gerror.NewCode(consts.CodeUserDisabled, "账号已禁用")
	}
	if data.Status == consts.UserStatusLocked {
		return nil, gerror.NewCode(consts.CodeUserLocked, "账号已锁定")
	}
	if data.Status == consts.UserStatusExpired {
		return nil, gerror.NewCode(consts.CodeUserExpired, "账号已过期")
	}
	if data.Status == consts.UserStatusDeleted {
		return nil, gerror.NewCode(consts.CodeUserDeleted, "账号已删除")
	}
	return data, nil
}

// 根据用户名和密码获取用户信息
func (l *sSysUser) GetUserByUsernameAndPassword(ctx context.Context, username string, password string) (user *entity.SysUser, err error) {
	// 判断用户名是否存在
	user, err = l.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	// 判断密码是否正确
	password = utility.PasswordEncrypt(password, user.Salt)
	if password != user.Password {
		g.Log().Errorf(ctx, "密码错误: username:%s, password:%s, user:%+v", username, password, user)
		return nil, gerror.NewCode(consts.CodeUserPasswordError, "密码错误")
	}
	return user, nil
}

// 根据用户ID获取用户信息
func (l *sSysUser) GetUserById(ctx context.Context, id int64) (user *entity.SysUser, err error) {
	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserId, id).Scan(&user)
	return
}
