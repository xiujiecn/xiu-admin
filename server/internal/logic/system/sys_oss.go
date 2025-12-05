// package system
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package system

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/library/xgorm/handler"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/model/request"
	"xiuadmin/internal/service"
	"xiuadmin/utility/hash"

	"github.com/gogf/gf/v2/crypto/gmd5"
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
	if len(option) == 0 {
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
	if param.NewFileType == 0 {
		if param.IsDevice {
			param.NewFileType = consts.SysUploadNewFileTypeDeviceFile
		} else if param.FileType == consts.SysUploadFileTypeImg {
			param.NewFileType = consts.SysUploadNewFileTypeUserImg
		} else {
			param.NewFileType = consts.SysUploadNewFileTypeUserFile
		}
	}
	result, err := s.UploadLocal(ctx, file, param.NewFileType, param.SaveOriginalName, param.SubDirName, param.NotAddDate)
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
		dao.SysOss.Columns().Md5:          result.Md5,
		dao.SysOss.Columns().FileSize:     result.Size,
		dao.SysOss.Columns().FileCrc16:    result.Crc16,
		dao.SysOss.Columns().FileSum:      result.Sum16,
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

func (s *sSysOss) MoveFile(ctx context.Context, param *model.SysOssMoveFileParam) (output *model.SysOssMoveFileModel, err error) {
	result, err := s.MoveFileLocal(ctx, param.FilePath, param.NewFileType, param.SaveOriginalName, param.SubDirName, param.NotAddDate)
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
		dao.SysOss.Columns().Md5:          result.Md5,
		dao.SysOss.Columns().FileSize:     result.Size,
		dao.SysOss.Columns().FileCrc16:    result.Crc16,
		dao.SysOss.Columns().FileSum:      result.Sum16,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	oss := &entity.SysOss{}
	err = s.Model(ctx).Where(dao.SysOss.Columns().OssId, id).Scan(&oss)
	if err != nil {
		return nil, err
	}
	output = &model.SysOssMoveFileModel{
		SysOss: oss,
	}
	return output, nil
}

func (s *sSysOss) SaveContent(ctx context.Context, param *model.SysOssSaveContentParam) (output *model.SysOssSaveContentModel, err error) {
	result, err := s.SaveContentLocal(ctx, param.Content, param.NewFileType, param.FileName, param.SubDirName, param.NotAddDate)
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
		dao.SysOss.Columns().Md5:          result.Md5,
		dao.SysOss.Columns().FileSize:     result.Size,
		dao.SysOss.Columns().FileCrc16:    result.Crc16,
		dao.SysOss.Columns().FileSum:      result.Sum16,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	oss := &entity.SysOss{}
	err = s.Model(ctx).Where(dao.SysOss.Columns().OssId, id).Scan(&oss)
	if err != nil {
		return nil, err
	}
	output = &model.SysOssSaveContentModel{
		SysOss: oss,
	}
	return output, nil
}

func (s *sSysOss) getNewFileTypePath(ctx context.Context, newFileType int) string {
	switch newFileType {
	case consts.SysUploadNewFileTypeUserFile:
		return consts.SysUploadNewFileTypePathUserFile
	case consts.SysUploadNewFileTypeUserImg:
		return consts.SysUploadNewFileTypePathUserImg
	case consts.SysUploadNewFileTypeDeviceFile:
		return consts.SysUploadNewFileTypePathDeviceFile
	default:
		key := fmt.Sprintf(consts.ConfigOssNewFileTypePath, newFileType)
		config, err := service.SysConfig().GetConfigByKey(ctx, key)
		if err == nil && config != nil {
			return config.ConfigValue
		}
	}
	return "other"
}

func (s *sSysOss) GetSaveFilePathConfig(ctx context.Context, newFileType int, notAddDate int) (string, error) {
	uploadPath := g.Cfg().MustGet(ctx, "system.upload.local.path").String()
	if uploadPath == "" {
		return "", errors.New("上传路径未配置")
	}
	tenantId := contexts.GetTenantId(ctx)
	if tenantId == "" {
		tenantId = "default"
	}
	fileTypePath := s.getNewFileTypePath(ctx, newFileType)
	fullDirPath := uploadPath + "/" + tenantId + "/" + fileTypePath
	if notAddDate == 0 {
		nowData := time.Now().Format("20060102")
		fullDirPath = fullDirPath + "/" + nowData
	}
	return fullDirPath, nil
}

func (s *sSysOss) MoveFileLocal(ctx context.Context, filePath string, newFileType int, useOriginalName int, subDirName string, isAddDate int) (result model.UploadResponse, err error) {
	if filePath == "" {
		err = errors.New("文件路径必须")
		return
	}
	fileName := gfile.Basename(filePath)
	fullDirPath, err := s.GetSaveFilePathConfig(ctx, newFileType, isAddDate)
	if err != nil {
		return
	}
	if subDirName != "" {
		fullDirPath = fullDirPath + "/" + subDirName
	}
	if !gfile.Exists(fullDirPath) {
		err = gfile.Mkdir(fullDirPath)
		if err != nil {
			return
		}
	}
	err = gfile.Move(filePath, fullDirPath)
	if err != nil {
		g.Log().Errorf(ctx, "sSysOss MoveFileLocal move file error: %+v, filePath: %s, fullDirPath: %s, newFileType: %d, useOriginalName: %d, subDirName: %s", err, filePath, fullDirPath, newFileType, useOriginalName, subDirName)
		return
	}
	fullFilePath := fullDirPath + "/" + fileName
	url := fullFilePath
	configType, err := service.SysConfig().GetConfigByKey(ctx, consts.ConfigOssUrlPath)
	if err == nil && configType.ConfigValue != "" {
		if strings.Contains(configType.ConfigValue, "|") {
			split := strings.Split(configType.ConfigValue, "|")
			url = strings.Replace(url, split[0], split[1], 1)
		} else {
			url = configType.ConfigValue + fullFilePath
		}
	}
	md5, _ := gmd5.EncryptFile(fullFilePath)
	crc16 := hash.Crc16ChecksumFile(fullFilePath)
	sum16 := hash.Sum16ChecksumFile(fullFilePath)
	result = model.UploadResponse{
		Size:         gfile.Size(fullFilePath),
		Path:         fullFilePath,
		FullPath:     url,
		OriginalName: fileName,
		Name:         fileName,
		Type:         gfile.ExtName(fileName),
		Ext:          gfile.ExtName(fileName),
		Md5:          md5,
		Crc16:        crc16,
		Sum16:        sum16,
	}
	return result, nil
}

// 保存内容
func (s *sSysOss) SaveContentLocal(ctx context.Context, content []byte, newFileType int, fileName string, subDirName string, isAddDate int) (result model.UploadResponse, err error) {
	fullDirPath, err := s.GetSaveFilePathConfig(ctx, newFileType, isAddDate)
	if err != nil {
		return
	}
	if subDirName != "" {
		fullDirPath = fullDirPath + "/" + subDirName
	}
	if !gfile.Exists(fullDirPath) {
		err = gfile.Mkdir(fullDirPath)
		if err != nil {
			return
		}
	}
	fullFilePath := fullDirPath + "/" + fileName
	err = gfile.PutBytes(fullFilePath, content)
	if err != nil {
		return
	}
	url := fullFilePath
	configType, err := service.SysConfig().GetConfigByKey(ctx, consts.ConfigOssUrlPath)
	if err == nil && configType.ConfigValue != "" {
		if strings.Contains(configType.ConfigValue, "|") {
			split := strings.Split(configType.ConfigValue, "|")
			url = strings.Replace(url, split[0], split[1], 1)
		} else {
			url = configType.ConfigValue + fullFilePath
		}
	}
	md5, _ := gmd5.EncryptFile(fullFilePath)
	crc16 := hash.Crc16ChecksumFile(fullFilePath)
	sum16 := hash.Sum16ChecksumFile(fullFilePath)
	result = model.UploadResponse{
		Size:         gfile.Size(fullFilePath),
		Path:         fullFilePath,
		FullPath:     url,
		OriginalName: fileName,
		Name:         fileName,
		Type:         "application/json",
		Ext:          gfile.ExtName(fileName),
		Md5:          md5,
		Crc16:        crc16,
		Sum16:        sum16,
	}
	return result, nil
}

func (s *sSysOss) UploadLocal(ctx context.Context, file *ghttp.UploadFile, newFileType int, useOriginalName int, subDirName string, notAddDate int) (result model.UploadResponse, err error) {
	if file == nil {
		err = errors.New("文件必须")
		return
	}

	fullDirPath, err := s.GetSaveFilePathConfig(ctx, newFileType, notAddDate)
	if err != nil {
		return
	}
	if subDirName != "" {
		fullDirPath = fullDirPath + "/" + subDirName
	}
	if !gfile.Exists(fullDirPath) {
		err = gfile.Mkdir(fullDirPath)
		if err != nil {
			return
		}
	}
	fileName := gfile.Basename(file.Filename)
	ext := gfile.Ext(file.Filename)
	fileNewName, err := file.Save(fullDirPath, useOriginalName == 0)
	if err != nil {
		return
	}
	fullFilePath := fullDirPath + "/" + fileNewName
	url := fullFilePath
	configType, err := service.SysConfig().GetConfigByKey(ctx, consts.ConfigOssUrlPath)
	if err == nil && configType.ConfigValue != "" {
		if strings.Contains(configType.ConfigValue, "|") {
			split := strings.Split(configType.ConfigValue, "|")
			url = strings.Replace(url, split[0], split[1], 1)
		} else {
			url = configType.ConfigValue + fullFilePath
		}
	}
	md5, _ := gmd5.EncryptFile(fullFilePath)
	crc16 := hash.Crc16ChecksumFile(fullFilePath)
	sum16 := hash.Sum16ChecksumFile(fullFilePath)
	result = model.UploadResponse{
		Size:         file.Size,
		Path:         fullFilePath,
		FullPath:     url,
		OriginalName: fileName,
		Name:         fileNewName,
		Type:         file.Header.Get("Content-type"),
		Ext:          ext,
		Md5:          md5,
		Crc16:        crc16,
		Sum16:        sum16,
	}
	return
}

func (s *sSysOss) DownloadLocal(ctx context.Context, file *entity.SysOss) (err error) {
	if file == nil {
		err = errors.New("文件必须")
		return
	}
	finalPath := file.Path
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
		configType, err = service.SysConfig().GetConfigByKey(ctx, consts.ConfigOssFileTypeKey)
		if err != nil {
			return
		}

	} else if checkFileType == consts.SysUploadFileTypeImg {
		//获取上传类型配置
		configType, err = service.SysConfig().GetConfigByKey(ctx, consts.ConfigOssImgTypeKey)
		if err != nil {
			return
		}
	} else {
		return fmt.Errorf("文件检查类型错误:%s|%s, 	文件类型:%s", consts.SysUploadFileTypeFile, consts.SysUploadFileTypeImg, checkFileType)
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
		configType, err = service.SysConfig().GetConfigByKey(ctx, consts.ConfigOssFileSizeKey)
		if err != nil {
			return
		}

	} else if checkFileType == consts.SysUploadFileTypeImg {
		//获取上传类型配置
		configType, err = service.SysConfig().GetConfigByKey(ctx, consts.ConfigOssImgSizeKey)
		if err != nil {
			return
		}
	} else {
		return fmt.Errorf("文件检查类型错误:%s|%s", consts.SysUploadFileTypeFile, consts.SysUploadFileTypeImg)
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

func (s *sSysOss) GetAllUrl(ctx context.Context, url string) (allUrl string, ossId int64, fileSize int64, OriginalName string, md5 string, hmac string, err error) {
	entity := &entity.SysOss{}
	s.Model(ctx).Where(dao.SysOss.Columns().Path, url).Scan(&entity)
	if entity.OssId <= 0 {
		return
	}
	ossId = entity.OssId
	ossConfig, err := service.SysOssConfig().View(ctx, &model.SysOssConfigViewParam{
		ConfigKey: entity.Service,
	})
	if err != nil {
		g.Log().Errorf(ctx, "sSysOss GetAllUrl error: %+v, url: %s", err, url)
		return "", 0, 0, "", "", "", err
	}
	if ossConfig.IsHttps == "Y" {
		allUrl = "https://" + ossConfig.Domain + entity.Path
	} else {
		allUrl = "http://" + ossConfig.Domain + entity.Path
	}
	fileSize = gfile.Size(entity.Path)
	OriginalName = entity.OriginalName
	md5, _ = gmd5.Encrypt(entity.Path)
	hmac = ""
	return
}
