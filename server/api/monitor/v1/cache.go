// Package v1
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025	 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RedisInfoReq struct {
	g.Meta `path:"/cache" method:"get" tags:"监控-缓存监控" summary:"获取缓存监控信息" x-check-permission:"cpm:monitor:cache:list"`
}

type RedisInfoRes struct {
	DBSize       int               `json:"dbSize"`
	CommandStats []*CommandStats   `json:"commandStats"`
	Info         map[string]string `json:"info"`
}

type CommandStats struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
