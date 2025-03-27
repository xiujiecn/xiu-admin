// package request
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 XiuAdmin CLI
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package request

type PageInfo struct {
	Page     int    `json:"page" form:"page" dc:"页码" v:"min:1#页码最小值不能低于1" `                                 // 页码
	PageSize int    `json:"pageSize" form:"pageSize" v:"min:1|max:2000#每页数量最小值不能低于1|最大值不能大于2000" dc:"每页大小"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword" dc:"关键字"`                                                // 关键字
}

type Empty struct{}

// SwitchReq 更新开关状态
type SwitchReq struct {
	Key   string `json:"key" v:"required#测试ID不能为空" dc:"开关字段"`
	Value int    `json:"value" v:"in:0,1#输入的开关值是无效的" dc:"更新值"`
}
