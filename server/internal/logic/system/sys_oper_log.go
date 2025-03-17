package system

import (
	"context"
	"strings"
	"xiujieadmin/internal/dao"
	"xiujieadmin/internal/library/contexts"
	"xiujieadmin/internal/library/mcache"
	"xiujieadmin/internal/model"
	"xiujieadmin/internal/model/entity"
	"xiujieadmin/internal/model/request"
	"xiujieadmin/internal/service"
	"xiujieadmin/utility"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysOperLog struct {
}

func NewSysOperLog() *sSysOperLog {
	return &sSysOperLog{}
}

func init() {
	service.RegisterSysOperLog(NewSysOperLog())
}

func (s *sSysOperLog) GetOperLogList(ctx context.Context, query *model.SysOperLogListParam, page *request.PageInfo) (items []*model.SysOperLogListModel, total int, err error) {
	// 获取当前用户租户编码
	claims, err := service.SysAuth().GetCurrentUser(ctx)
	if err != nil {
		return nil, 0, err
	}
	tenantId := claims.TenantId

	db := dao.SysOperLog.Ctx(ctx).Where(dao.SysOperLog.Columns().TenantId, tenantId)

	if query.Title != "" {
		db = db.WhereLike(dao.SysOperLog.Columns().Title, "%"+query.Title+"%")
	}

	if query.BusinessType != "" {
		db = db.Where(dao.SysOperLog.Columns().BusinessType, query.BusinessType)
	}

	if query.Method != "" {
		db = db.WhereLike(dao.SysOperLog.Columns().Method, "%"+query.Method+"%")
	}

	total, err = db.Count()
	if err != nil {
		return nil, 0, err
	}

	err = db.Page(page.Page, page.PageSize).OrderDesc(dao.SysOperLog.Columns().OperId).Scan(&items)
	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (s *sSysOperLog) AnalysisLog(ctx context.Context) (data *model.SysOperLogAddParam, err error) {
	ctxData := contexts.Get(ctx)
	request := ghttp.RequestFromCtx(ctx)
	handlerResponse := request.GetHandlerResponse() // 响应结果
	param := request.GetMap()                       // 请求参数
	path := request.URL.Path

	res := gconv.Map(handlerResponse)
	deptName := ""
	operName := ""
	businessType := 0
	operatorType := 0
	status := 0
	costTime := gtime.Now().Sub(gtime.New(request.EnterTime)).Milliseconds()
	message := ""
	if ctxData != nil && ctxData.User != nil {
		deptName, _ = mcache.GetDeptName(ctx, ctxData.User.DeptId)
		operName = ctxData.User.Username
	}
	jsonResult := ""
	var code gcode.Code = gcode.CodeOK
	if erro := request.GetError(); erro != nil {
		status = 1
		code = gerror.Code(erro)
		message = gjson.New(g.Map{
			"code":    code,
			"message": erro.Error(),
		}).String()
	} else {
		status = 0
		if b, err := gjson.Encode(res); err == nil {
			if len(b) > 1024 {
				jsonResult = "数据过大，未记录"
			} else {
				jsonResult = string(b)
			}
		}
	}
	if strings.Contains(path, "/add") {
		businessType = 1
	}
	if strings.Contains(path, "/edit") {
		businessType = 2
	}
	if strings.Contains(path, "/delete") {
		businessType = 3
	}
	url := request.URL.String()
	if len(url) > 100 {
		url = url[:100] + "..."
	}
	data = &model.SysOperLogAddParam{
		TenantId:      contexts.GetTenantId(ctx),
		Title:         "",
		BusinessType:  businessType,
		Method:        request.URL.Path,
		RequestMethod: request.Method,
		OperatorType:  operatorType,
		OperName:      operName,
		DeptName:      deptName,
		OperUrl:       url,
		OperIp:        utility.GetClientIp(ctx),
		OperLocation:  utility.GetCityByIp(utility.GetClientIp(ctx)),
		OperParam:     gconv.String(param),
		JsonResult:    jsonResult,
		Status:        status,
		ErrorMsg:      message,
		OperTime:      gtime.Now(),
		CostTime:      costTime,
	}
	return data, nil
}

func (s *sSysOperLog) ClearOperationLogByDays(ctx context.Context, days int) error {
	_, err := dao.SysOperLog.Ctx(ctx).Delete("to_days(now())-to_days(`oper_time`) > ?", days+1)
	return err
}

func (s *sSysOperLog) RealWrite(ctx context.Context, data entity.SysOperLog) (err error) {
	_, err = dao.SysOperLog.Ctx(ctx).Data(data).Save()
	if err != nil {
		g.Log().Error(ctx, "sSysOperLog.RealWrite error:%v,data:%+v", err, data)
	}
	return
}
