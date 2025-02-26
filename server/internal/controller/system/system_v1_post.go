package system

import (
	"context"

	v1 "server/api/system/v1"
	"server/internal/model/response"
	"server/internal/service"
)

func (c *ControllerV1) PostList(ctx context.Context, req *v1.PostListReq) (res *v1.PostListRes, err error) {
	items, total, err := service.SysPost().GetPostList(ctx, req.SysPostListParam, req.PageInfo)
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
