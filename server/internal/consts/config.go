// Package consts
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package consts

const (
	ConfigOssUrlPath         = "sys.oss.urlPath"            // 上传文件Url路径 "http://ab.com/" 或 "/" 或 "resource/upload|http://ab.com/upload"
	ConfigOssFileTypeKey     = "sys.oss.fileType"           // 上传文件类型
	ConfigOssImgTypeKey      = "sys.oss.imgType"            // 上传图片类型
	ConfigOssFileSizeKey     = "sys.oss.fileSize"           // 上传文件大小
	ConfigOssImgSizeKey      = "sys.oss.imgSize"            // 上传图片大小
	ConfigOnlineForceLogout  = "sys.online.forceLogout"     // 在线用户强制退出
	ConfigOssNewFileTypePath = "sys.oss.newFileTypePath.%d" // 上传文件类型路径 ${newFileType}
)

const (
	ConfigOnlineForceLogoutTrue  = "true"  // 在线用户强制退出
	ConfigOnlineForceLogoutFalse = "false" // 在线用户强制退出
)
