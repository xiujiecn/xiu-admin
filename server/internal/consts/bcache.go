// Package consts
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package consts

const (
	KeySysAuthToken       = "login:%d:%s"         // 系统认证token缓存    login:{userId}:{uuid}
	KeySysAuthTokenReject = "login_reject:%d:%s:" // 系统认证token拒绝缓存 login_reject:{userId}:{uuid}
)
