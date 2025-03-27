// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package common

import (
	"context"

	"xiuadmin/api/common/v1"
)

type ICommonV1 interface {
	SysOssUpload(ctx context.Context, req *v1.SysOssUploadReq) (res *v1.SysOssUploadRes, err error)
	SysOssDownload(ctx context.Context, req *v1.SysOssDownloadReq) (res *v1.SysOssDownloadRes, err error)
}
