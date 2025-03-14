package system

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
	"xiujieadmin/internal/consts"
	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/library/contexts"
	"xiujieadmin/internal/library/xgorm/handler"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/entity"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysOss struct {
}

func NewSysOss() *sSysOss {
	return &sSysOss{}
}

func init() {
	service.RegisterSysOss(NewSysOss())
}

func (s *sSysOss) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) > 0 {
		option = append(option, &handler.Option{
			FilterTenant: true,
			FilterAuth:   true,
		})
	}
	return handler.Model(dao.SysOss.Ctx(ctx), option...)
}

func (s *sSysOss) List(ctx context.Context, param *model.SysOssListParam, pageInfo *request.PageInfo) (items []*model.SysOssListModel, total int, err error) {

	db := s.Model(ctx)

	if param.FileName != "" {
		db = db.WhereLike(dao.SysOss.Columns().FileName, "%"+param.FileName+"%")
	}

	if param.OriginalName != "" {
		db = db.WhereLike(dao.SysOss.Columns().OriginalName, "%"+param.OriginalName+"%")
	}

	if param.FileSuffix != "" {
		db = db.Where(dao.SysOss.Columns().FileSuffix, param.FileSuffix)
	}

	if param.Service != "" {
		db = db.Where(dao.SysOss.Columns().Service, param.Service)
	}

	if len(param.CreatedAt) == 2 {
		startTime, err := gtime.StrToTime(param.CreatedAt[0])
		if err != nil {
			return nil, 0, err
		}
		endTime, err := gtime.StrToTime(param.CreatedAt[1])
		if err != nil {
			return nil, 0, err
		}
		db = db.WhereBetween(dao.SysOss.Columns().CreatedAt, startTime, endTime.EndOfDay())
	}

	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}

	err = db.Page(pageInfo.Page, pageInfo.PageSize).Order(dao.SysOss.Columns().OssId, "DESC").Scan(&items)
	if err != nil {
		return nil, 0, err
	}
	return
}

func (s *sSysOss) View(ctx context.Context, param *model.SysOssViewParam) (output *model.SysOssViewModel, err error) {
	if param.OssId == 0 {
		return nil, errors.New("未选择文件")
	}
	db := s.Model(ctx).Where(dao.SysOss.Columns().OssId, param.OssId)
	err = db.Scan(&output)
	if err != nil {
		return nil, err
	}
	return
}

func (s *sSysOss) Download(ctx context.Context, param *model.SysOssDownloadParam) (output *model.SysOssDownloadModel, err error) {
	if param.OssId == 0 {
		return nil, errors.New("未选择文件")
	}
	data := &entity.SysOss{}
	db := s.Model(ctx).Where(dao.SysOss.Columns().OssId, param.OssId)
	err = db.Scan(&data)
	if err != nil {
		return nil, err
	}
	if data.Service != "local" {
		err = gerror.New("非本地文件")
		return
	}
	output = &model.SysOssDownloadModel{
		SysOss: data,
	}
	err = s.DownloadLocal(ctx, data)
	if err != nil {
		return nil, err
	}
	return
}

func (s *sSysOss) Delete(ctx context.Context, param *model.SysOssDeleteParam) (output *model.SysOssDeleteModel, err error) {
	if len(param.OssIds) == 0 {
		return nil, errors.New("未选择文件")
	}
	items := make([]*entity.SysOss, 0)
	err = s.Model(ctx).WhereIn(dao.SysOss.Columns().OssId, param.OssIds).Scan(&items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, errors.New("文件不存在")
	}
	_, err = s.Model(ctx).WhereIn(dao.SysOss.Columns().OssId, param.OssIds).Delete()
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.Service == "local" {
			s.DeleteLocal(ctx, item)
		}
	}
	return
}

func (s *sSysOss) Upload(ctx context.Context, param *model.SysOssUploadParam) (output *model.SysOssUploadModel, err error) {
	r := g.RequestFromCtx(ctx)
	file := r.GetUploadFile("file")
	if file == nil {
		err = gerror.New("上传文件必须")
		return
	}
	// 检查文件类型
	err = s.CheckType(ctx, param.FileType, file)
	if err != nil {
		return
	}

	// 检查文件大小
	err = s.CheckSize(ctx, param.FileType, file)
	if err != nil {
		return
	}

	// TODO：待实现其他云存储
	result, err := s.UploadLocal(ctx, file)
	if err != nil {
		return nil, err
	}
	id, err := s.Model(ctx).Data(g.Map{
		dao.SysOss.Columns().TenantId:     contexts.GetTenantId(ctx),
		dao.SysOss.Columns().FileName:     result.Name,
		dao.SysOss.Columns().OriginalName: result.OriginalName,
		dao.SysOss.Columns().FileSuffix:   result.Ext,
		dao.SysOss.Columns().Path:         result.Path,
		dao.SysOss.Columns().Url:          result.FullPath,
		dao.SysOss.Columns().CreatedDept:  contexts.GetDeptId(ctx),
		dao.SysOss.Columns().CreatedAt:    gtime.Now(),
		dao.SysOss.Columns().CreatedBy:    contexts.GetUserId(ctx),
		dao.SysOss.Columns().UpdatedAt:    gtime.Now(),
		dao.SysOss.Columns().UpdatedBy:    contexts.GetUserId(ctx),
		dao.SysOss.Columns().Service:      "local",
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	output = &model.SysOssUploadModel{}
	err = s.Model(ctx).Where(dao.SysOss.Columns().OssId, id).Scan(&output)
	if err != nil {
		return nil, err
	}
	return
}

func (s *sSysOss) UploadLocal(ctx context.Context, file *ghttp.UploadFile) (result model.UploadResponse, err error) {
	if file == nil {
		err = errors.New("文件必须")
		return
	}
	uploadPath := g.Cfg().MustGet(ctx, "system.upload.local.path").String()
	if uploadPath == "" {
		err = errors.New("上传路径未配置")
		return
	}
	nowData := time.Now().Format("2006-01-02")
	fullDirPath := uploadPath + "/" + nowData
	if !gfile.Exists(fullDirPath) {
		err = gfile.Mkdir(fullDirPath)
		if err != nil {
			return
		}
	}
	fileName := gfile.Basename(file.Filename)
	ext := gfile.Ext(file.Filename)
	fileNewName, err := file.Save(fullDirPath, true)
	if err != nil {
		return
	}
	fullFilePath := fullDirPath + "/" + fileNewName
	url := fullFilePath
	configType, err := NewSysConfig().GetConfigByKey(ctx, consts.ConfigOssUrlPath)
	if err == nil && configType.ConfigValue != "" {
		if strings.Contains(configType.ConfigValue, "|") {
			split := strings.Split(configType.ConfigValue, "|")
			url = strings.Replace(url, split[0], split[1], 1)
		} else {
			url = configType.ConfigValue + fullFilePath
		}
	}

	result = model.UploadResponse{
		Size:         file.Size,
		Path:         fullFilePath,
		FullPath:     url,
		OriginalName: fileName,
		Name:         fileNewName,
		Type:         file.Header.Get("Content-type"),
		Ext:          ext,
	}
	return
}

func (s *sSysOss) DownloadLocal(ctx context.Context, file *entity.SysOss) (err error) {
	if file == nil {
		err = errors.New("文件必须")
		return
	}
	finalPath := file.Url
	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		err = gerror.New("ctx not http request")
		return
	}
	r.Response.ServeFileDownload(finalPath)
	r.Response.Status = http.StatusOK
	r.Response.Flush()
	return
}
func (s *sSysOss) DeleteLocal(ctx context.Context, file *entity.SysOss) (err error) {
	if file == nil {
		err = errors.New("文件不存在")
		return
	}
	filePath := file.Path
	if !gfile.Exists(filePath) {
		err = errors.New("文件不存在")
		return
	}
	err = gfile.Remove(filePath)
	if err != nil {
		return
	}
	return
}

func (s *sSysOss) CheckType(ctx context.Context, checkFileType string, file *ghttp.UploadFile) (err error) {

	var (
		configType *entity.SysConfig
	)

	if checkFileType == consts.SysUploadFileTypeFile {
		//获取上传类型配置
		configType, err = NewSysConfig().GetConfigByKey(ctx, consts.ConfigOssFileTypeKey)
		if err != nil {
			return
		}

	} else if checkFileType == consts.SysUploadFileTypeImg {
		//获取上传类型配置
		configType, err = NewSysConfig().GetConfigByKey(ctx, consts.ConfigOssImgTypeKey)
		if err != nil {
			return
		}
	} else {
		return errors.New(fmt.Sprintf("文件检查类型错误:%s|%s, 	文件类型:%s", consts.SysUploadFileTypeFile, consts.SysUploadFileTypeImg, checkFileType))
	}
	ext := gfile.ExtName(file.Filename)
	ext = strings.ToLower(ext)
	imageType := gstr.Split(configType.ConfigValue, ",")
	rightType := false
	for _, v := range imageType {
		if v == ext {
			rightType = true
			break
		}
	}
	if !rightType {
		err = gerror.New("上传文件类型错误，只能包含后缀为：" + configType.ConfigValue + "的文件。")
		return
	}
	return
}

func (s *sSysOss) CheckSize(ctx context.Context, checkFileType string, file *ghttp.UploadFile) (err error) {

	var (
		configType *entity.SysConfig
	)

	if checkFileType == consts.SysUploadFileTypeFile {
		//获取上传类型配置
		configType, err = NewSysConfig().GetConfigByKey(ctx, consts.ConfigOssFileSizeKey)
		if err != nil {
			return
		}

	} else if checkFileType == consts.SysUploadFileTypeImg {
		//获取上传类型配置
		configType, err = NewSysConfig().GetConfigByKey(ctx, consts.ConfigOssImgSizeKey)
		if err != nil {
			return
		}
	} else {
		return errors.New(fmt.Sprintf("文件检查类型错误:%s|%s", consts.SysUploadFileTypeFile, consts.SysUploadFileTypeImg))
	}
	match, err := gregex.MatchString(`^([0-9]+)(?i:([a-z]*))$`, configType.ConfigValue)
	if err != nil {
		return
	}
	if len(match) == 0 {
		err = gerror.New("上传文件大小未设置，请在后台配置，格式为（30M,30k,30MB）")
		return
	}
	var cfSize int64
	switch gstr.ToUpper(match[2]) {
	case "MB", "M":
		cfSize = gconv.Int64(match[1]) * 1024 * 1024
	case "KB", "K":
		cfSize = gconv.Int64(match[1]) * 1024
	case "":
		cfSize = gconv.Int64(match[1])
	}
	if cfSize == 0 {
		err = gerror.New("上传文件大小未设置，请在后台配置，格式为（30M,30k,30MB），最大单位为MB")
		return
	}
	if file.Size > cfSize {
		err = gerror.New("上传文件大小超过限制，最大为" + configType.ConfigValue)
		return
	}
	return
}
