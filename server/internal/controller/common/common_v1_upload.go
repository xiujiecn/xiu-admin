package common

import (
	"context"

	v1 "xiuadmin/api/common/v1"
	"xiuadmin/internal/service"
)

func (c *ControllerV1) SysOssUpload(ctx context.Context, req *v1.SysOssUploadReq) (res *v1.SysOssUploadRes, err error) {

	info, err := service.SysOss().Upload(ctx, req.SysOssUploadParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysOssUploadRes{
		SysOssUploadModel: info,
	}
	return
}
func (c *ControllerV1) SysOssDownload(ctx context.Context, req *v1.SysOssDownloadReq) (res *v1.SysOssDownloadRes, err error) {
	output, err := service.SysOss().Download(ctx, req.SysOssDownloadParam)
	if err != nil {
		return nil, err
	}
	res = &v1.SysOssDownloadRes{
		SysOssDownloadModel: output,
	}
	return
}
