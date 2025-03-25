package system

import (
	"context"

	v1 "xiujieadmin/api/system/v1"
	"xiujieadmin/internal/model/response"
	"xiujieadmin/internal/service"
)

func (c *ControllerV1) PostList(ctx context.Context, req *v1.PostListReq) (res *v1.PostListRes, err error) {
	items, total, err := service.SysPost().List(ctx, req.SysPostListParam)
	if err != nil {
		return nil, err
	}
	return &v1.PostListRes{
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
		Data: items,
	}, nil
}
func (c *ControllerV1) PostAdd(ctx context.Context, req *v1.PostAddReq) (res *v1.PostAddRes, err error) {
	post, err := service.SysPost().Add(ctx, req.SysPostAddParam)
	if err != nil {
		return nil, err
	}
	return &v1.PostAddRes{
		SysPostAddModel: post,
	}, nil
}
func (c *ControllerV1) PostEdit(ctx context.Context, req *v1.PostEditReq) (res *v1.PostEditRes, err error) {
	post, err := service.SysPost().Edit(ctx, req.SysPostEditParam)
	if err != nil {
		return nil, err
	}
	return &v1.PostEditRes{
		SysPostEditModel: post,
	}, nil
}
func (c *ControllerV1) PostDelete(ctx context.Context, req *v1.PostDeleteReq) (res *v1.PostDeleteRes, err error) {
	post, err := service.SysPost().Delete(ctx, req.SysPostDeleteParam)
	if err != nil {
		return nil, err
	}
	return &v1.PostDeleteRes{
		SysPostDeleteModel: post,
	}, nil
}
func (c *ControllerV1) PostView(ctx context.Context, req *v1.PostViewReq) (res *v1.PostViewRes, err error) {
	post, err := service.SysPost().View(ctx, req.SysPostViewParam)
	if err != nil {
		return nil, err
	}
	return &v1.PostViewRes{
		SysPostViewModel: post,
	}, nil
}
func (c *ControllerV1) PostExport(ctx context.Context, req *v1.PostExportReq) (res *v1.PostExportRes, err error) {
	_, err = service.SysPost().Export(ctx, req.SysPostExportParam)
	if err != nil {
		return nil, err
	}
	return &v1.PostExportRes{}, nil
}
