package consts

const (
	KeySysAuthToken       = "login:%d:%s"         // 系统认证token缓存    login:{userId}:{uuid}
	KeySysAuthTokenReject = "login_reject:%d:%s:" // 系统认证token拒绝缓存 login_reject:{userId}:{uuid}
)
