package model

type LoginUserOut struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
}
