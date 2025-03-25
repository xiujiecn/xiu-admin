package system

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "xiujieadmin/api/system/v1"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/response"
	"xiujieadmin/internal/service"
)

func (c *ControllerV1) SysJobList(ctx context.Context, req *v1.SysJobListReq) (res *v1.SysJobListRes, err error) {
	data, total, err := service.SysJob().List(ctx, &req.SysJobListParam, &req.PageInfo)
	if err != nil {
		return nil, err
	}

	return &v1.SysJobListRes{
		Data: data,
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}, nil

}
func (c *ControllerV1) SysJobView(ctx context.Context, req *v1.SysJobViewReq) (res *v1.SysJobViewRes, err error) {
	data, err := service.SysJob().View(ctx, req.JobId)
	if err != nil {
		return nil, err
	}
	return &v1.SysJobViewRes{
		SysJobViewModel: *data,
	}, nil

}
func (c *ControllerV1) SysJobAdd(ctx context.Context, req *v1.SysJobAddReq) (res *v1.SysJobAddRes, err error) {
	lastInserId, err := service.SysJob().Add(ctx, &req.SysJobAddModel)
	if err != nil {
		return nil, err
	}
	data, err := service.SysJob().View(ctx, lastInserId)

	if err != nil {
		return nil, err
	}
	return &v1.SysJobAddRes{
		SysJobViewModel: *data,
	}, nil

}
func (c *ControllerV1) SysJobUpdate(ctx context.Context, req *v1.SysJobUpdateReq) (res *v1.SysJobUpdateRes, err error) {
	effectedRow, err := service.SysJob().Update(ctx, &req.SysJobUpdateModel)
	if err != nil {
		return nil, err
	}

	if effectedRow == 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "要更新的数据不存在")
	}

	data, err := service.SysJob().View(ctx, req.JobId)
	if err != nil {
		return nil, err
	}
	return &v1.SysJobUpdateRes{
		SysJobViewModel: *data,
	}, nil
}

func (c *ControllerV1) SysJobUpdateStatus(ctx context.Context, req *v1.SysJobUpdateStatusReq) (res *v1.SysJobUpdateStatusRes, err error) {
	effectedRow, err := service.SysJob().UpdateStatus(ctx, &req.SysJobUpdateStatusModel)
	if err != nil {
		return nil, err
	}

	if effectedRow == 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "要更新的数据不存在")
	}

	data, err := service.SysJob().View(ctx, req.JobId)
	if err != nil {
		return nil, err
	}
	return &v1.SysJobUpdateStatusRes{
		SysJobViewModel: *data,
	}, nil
}

func (c *ControllerV1) SysJobDelete(ctx context.Context, req *v1.SysJobDeleteReq) (res *v1.SysJobDeleteRes, err error) {
	if len(req.JobIds) == 0 {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "参数不能为空")
	}

	_, err = service.SysJob().Delete(ctx, &model.SysJobDeleteModel{JobIds: req.JobIds})
	if err != nil {
		return nil, err
	}

	return &v1.SysJobDeleteRes{
		JobIds: req.JobIds,
	}, nil
}

func (c *ControllerV1) SysJobExec(ctx context.Context, req *v1.SysJobExecReq) (res *v1.SysJobExecRes, err error) {
	err = service.SysJob().Exec(ctx, req.JobId)
	if err != nil {
		return nil, err
	}

	return &v1.SysJobExecRes{
		JobId: req.JobId,
	}, nil
}
