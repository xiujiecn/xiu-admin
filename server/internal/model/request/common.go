package request

type PageInfo struct {
	Page     int    `json:"page" form:"page" dc:"页码" v:"min:1#页码最小值不能低于1" `                                 // 页码
	PageSize int    `json:"pageSize" form:"pageSize" v:"min:1|max:2000#每页数量最小值不能低于1|最大值不能大于2000" dc:"每页大小"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword" dc:"关键字"`                                                // 关键字
}

type Empty struct{}
