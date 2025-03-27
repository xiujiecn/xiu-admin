package system

import (
	"context"

	v1 "xiuadmin/api/system/v1"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/response"
	"xiuadmin/internal/service"
)

func (c *ControllerV1) SysDictTypeList(ctx context.Context, req *v1.SysDictTypeListReq) (res *v1.SysDictTypeListRes, err error) {
	items, total, err := service.SysDictType().List(ctx, &req.SysDictTypeListParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysDictTypeListRes{
		Items: items,
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}, nil
}

func (c *ControllerV1) SysDictDataList(ctx context.Context, req *v1.SysDictDataListReq) (res *v1.SysDictDataListRes, err error) {
	items, total, err := service.SysDictData().List(ctx, &req.SysDictDataListParam)
	if err != nil {
		return nil, err
	}
	typeData, err := service.SysDictType().View(ctx, &model.SysDictTypeViewParam{
		DictId:   req.SysDictDataListParam.DictId,
		DictType: req.SysDictDataListParam.DictType,
	})
	if err != nil {
		return nil, err
	}

	return &v1.SysDictDataListRes{
		Items: items,
		Type:  typeData,
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}, nil
}
func (c *ControllerV1) SysDictTypeAdd(ctx context.Context, req *v1.SysDictTypeAddReq) (res *v1.SysDictTypeAddRes, err error) {
	output, err := service.SysDictType().Add(ctx, &req.SysDictTypeAddParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysDictTypeAddRes{
		SysDictTypeAddModel: *output,
	}, nil
}
func (c *ControllerV1) SysDictTypeEdit(ctx context.Context, req *v1.SysDictTypeEditReq) (res *v1.SysDictTypeEditRes, err error) {
	output, err := service.SysDictType().Edit(ctx, &req.SysDictTypeEditParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysDictTypeEditRes{
		SysDictTypeEditModel: *output,
	}, nil
}
func (c *ControllerV1) SysDictTypeDelete(ctx context.Context, req *v1.SysDictTypeDeleteReq) (res *v1.SysDictTypeDeleteRes, err error) {
	output, err := service.SysDictType().Delete(ctx, &req.SysDictTypeDeleteParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysDictTypeDeleteRes{
		SysDictTypeDeleteModel: *output,
	}, nil
}
func (c *ControllerV1) SysDictTypeView(ctx context.Context, req *v1.SysDictTypeViewReq) (res *v1.SysDictTypeViewRes, err error) {
	output, err := service.SysDictType().View(ctx, &req.SysDictTypeViewParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysDictTypeViewRes{
		SysDictTypeViewModel: *output,
	}, nil
}
func (c *ControllerV1) SysDictDataAdd(ctx context.Context, req *v1.SysDictDataAddReq) (res *v1.SysDictDataAddRes, err error) {
	output, err := service.SysDictData().Add(ctx, &req.SysDictDataAddParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysDictDataAddRes{
		SysDictDataAddModel: *output,
	}, nil
}
func (c *ControllerV1) SysDictDataEdit(ctx context.Context, req *v1.SysDictDataEditReq) (res *v1.SysDictDataEditRes, err error) {
	output, err := service.SysDictData().Edit(ctx, &req.SysDictDataEditParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysDictDataEditRes{
		SysDictDataEditModel: *output,
	}, nil
}
func (c *ControllerV1) SysDictDataDelete(ctx context.Context, req *v1.SysDictDataDeleteReq) (res *v1.SysDictDataDeleteRes, err error) {
	_, err = service.SysDictData().Delete(ctx, &req.SysDictDataDeleteParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysDictDataDeleteRes{
		SysDictDataDeleteModel: model.SysDictDataDeleteModel{
			DictCodes: req.SysDictDataDeleteParam.DictCodes,
		},
	}, nil
}
func (c *ControllerV1) SysDictDataView(ctx context.Context, req *v1.SysDictDataViewReq) (res *v1.SysDictDataViewRes, err error) {
	output, err := service.SysDictData().View(ctx, &req.SysDictDataViewParam)
	if err != nil {
		return nil, err
	}
	return &v1.SysDictDataViewRes{
		SysDictDataViewModel: *output,
	}, nil
}
