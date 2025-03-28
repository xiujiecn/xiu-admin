// package gohtml
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package gohtml

// An element represents an HTML element.
type element interface {
	isInline() bool
	write(*formattedBuffer, bool) bool
}
