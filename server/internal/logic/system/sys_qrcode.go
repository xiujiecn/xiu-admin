package system

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"xiuadmin/internal/consts"
	"xiuadmin/internal/dao"
	"xiuadmin/internal/library/contexts"
	"xiuadmin/internal/model"
	"xiuadmin/internal/model/entity"
	"xiuadmin/internal/model/request"
	"xiuadmin/internal/service"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
)

type sSysQrcode struct {
}

func NewSysQrcode() *sSysQrcode {
	return &sSysQrcode{}
}

func init() {
	service.RegisterSysQrcode(NewSysQrcode())
	go func() {
		for {
			time.Sleep(time.Second * 60)
			g.Log().Debug(context.Background(), "检查二维码是否过期，过期则删除, 每分钟检查一次")
			checkQrcodeExpiredToRemove(context.Background())
		}
	}()
}

type SysQrcodeConfig struct {
	ThirdId      int    `json:"thirdId" dc:"第三方ID"`
	SceneIdLogin int    `json:"sceneIdLogin" dc:"场景ID-登录"`
	SceneIdBind  int    `json:"sceneIdBind" dc:"场景ID-绑定"`
	Token        string `json:"token" dc:"Token"`
	ServerApi    string `json:"serverApi" dc:"服务端API"`
}

// 检查二维码是否过期，过期则删除
func checkQrcodeExpiredToRemove(ctx context.Context) {
	var sysQrcodeConfig *SysQrcodeConfig
	g.TryCatch(ctx, func(ctx context.Context) {
		var err error
		sysQrcodeConfig, err = getSysQrcodeConfig(ctx)
		if err != nil {
			g.Log().Errorf(ctx, "获取系统二维码配置失败, 错误: %v", err)
			return
		}
	}, func(ctx context.Context, exception error) {
		g.Log().Errorf(ctx, "检查二维码是否过期, 获取系统二维码配置失败, 错误: %v", exception)
	})

	if sysQrcodeConfig == nil {
		return
	}

	g.TryCatch(ctx, func(ctx context.Context) {
		redisKey := fmt.Sprintf(consts.KeyLoginScanWechat, sysQrcodeConfig.SceneIdLogin)
		err := checkQrcodeExpiredToRemoveByKey(ctx, redisKey)
		if err != nil {
			g.Log().Errorf(ctx, "检查登录二维码是否过期失败, 错误: %v", err)
			return
		}
	}, func(ctx context.Context, exception error) {
		g.Log().Errorf(ctx, "检查登录二维码是否过期, 错误: %v", exception)
	})

	g.TryCatch(ctx, func(ctx context.Context) {
		redisKey := fmt.Sprintf(consts.KeyUserBindWechat, sysQrcodeConfig.SceneIdBind)
		err := checkQrcodeExpiredToRemoveByKey(ctx, redisKey)
		if err != nil {
			g.Log().Errorf(ctx, "检查绑定二维码是否过期失败, 错误: %v", err)
			return
		}
	}, func(ctx context.Context, exception error) {
		g.Log().Errorf(ctx, "检查绑定二维码是否过期, 错误: %v", exception)
	})
}

// 检查指定redisKey的Hash的Field是否过期，过期则删除对应的feild
func checkQrcodeExpiredToRemoveByKey(ctx context.Context, redisKey string) (err error) {
	val, err := g.Redis().HGetAll(ctx, redisKey)
	if err != nil {
		g.Log().Errorf(ctx, "获取登录二维码缓存信息失败 redisKey:%s, 错误: %v", redisKey, err)
		return err
	}
	mapData := val.Map()
	if len(mapData) == 0 {
		return nil
	}

	for field, v := range mapData {
		cacheModel := &model.QrcodeCacheModel{}
		err = gconv.Struct(v, cacheModel)
		if err != nil {
			// 删除redisKey
			g.Log().Debug(ctx, "数据转换失败, 删除对应的数据 redisKey:%s, field:%s", redisKey, field)
			g.Redis().Do(ctx, "HDEL", redisKey, field)
			continue
		}
		// 判断是否过期 一分钟
		if cacheModel.ExpireTime.Before(gtime.Now().Add(time.Minute * -1)) {
			// 删除redisKey
			g.Log().Debug(ctx, "删除过期数据, redisKey:%s, field:%s", redisKey, field)
			g.Redis().Do(ctx, "HDEL", redisKey, field)
		}
	}
	return nil
}

// 获取系统二维码配置
func getSysQrcodeConfig(ctx context.Context) (*SysQrcodeConfig, error) {

	thirdId, err := g.Cfg().Get(ctx, "wx.qrcode.thirdId")
	if err != nil {
		return nil, err
	}
	sceneIdLogin, err := g.Cfg().Get(ctx, "wx.qrcode.sceneIdLogin")
	if err != nil {
		return nil, err
	}
	sceneIdBind, err := g.Cfg().Get(ctx, "wx.qrcode.sceneIdBind")
	if err != nil {
		return nil, err
	}
	token, err := g.Cfg().Get(ctx, "wx.qrcode.token")
	if err != nil {
		return nil, err
	}
	serverApi, err := g.Cfg().Get(ctx, "wx.qrcode.serverApi")
	if err != nil {
		return nil, err
	}

	return &SysQrcodeConfig{
		ThirdId:      thirdId.Int(),
		SceneIdLogin: sceneIdLogin.Int(),
		SceneIdBind:  sceneIdBind.Int(),
		Token:        token.String(),
		ServerApi:    serverApi.String(),
	}, nil
}

// 调用二维码API
func callQrcodeApi(ctx context.Context, sysQrcodeConfig *SysQrcodeConfig, isLogin bool) (res *model.ApiCallQrcodeModel, err error) {

	// 调用API
	sceneId := sysQrcodeConfig.SceneIdLogin
	if !isLogin {
		sceneId = sysQrcodeConfig.SceneIdBind
	}
	request := g.Client()
	request.SetHeader("Content-Type", "application/json")
	timeStamp, sign := tokenEncode(sysQrcodeConfig.Token)
	response, err := request.Get(ctx, sysQrcodeConfig.ServerApi, g.Map{
		"thirdId":   sysQrcodeConfig.ThirdId,
		"sceneId":   sceneId,
		"timeStamp": timeStamp,
		"sign":      sign,
	})
	if err != nil {
		g.Log().Errorf(ctx, "调用二维码API失败, 请求数据: %v, 错误: %v", request, err)
		return nil, err
	}
	if response.StatusCode != 200 {
		g.Log().Errorf(ctx, "调用二维码API失败, 请求数据: %v, 响应数据: %v", request, response)
		return nil, errors.New("调用二维码API失败")
	}

	responseData := response.ReadAllString()

	g.Log().Debugf(ctx, "调用二维码API成功, 请求数据: %v, 响应数据: %v", request, responseData)
	result := &model.ApiCallQrcodeResult{}
	err = json.Unmarshal([]byte(responseData), &result)
	if err != nil {
		g.Log().Errorf(ctx, "调用二维码API失败, 数据转换失败, 请求数据: %v, 响应数据: %v, 错误: %v", request, responseData, err)
		return nil, err
	}
	if result.Code != 0 {
		g.Log().Errorf(ctx, "调用二维码API失败, 错误信息: %v", result.Error)
		return nil, errors.New(result.Error)
	}
	return &model.ApiCallQrcodeModel{
		Ticket:     result.Data.Ticket,
		ExpireTime: result.Data.ExpireTime,
		QrUrl:      result.Data.QrUrl,
	}, nil
}

// 获取登录二维码
func (l *sSysQrcode) GetQrcodeLogin(ctx context.Context) (res *model.QrcodeLoginModel, err error) {
	sysQrcodeConfig, err := getSysQrcodeConfig(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取系统二维码配置失败, 错误: %v", err)
		return nil, gerror.New("获取系统二维码配置失败")
	}
	if sysQrcodeConfig.ServerApi == "" || sysQrcodeConfig.Token == "" || sysQrcodeConfig.ThirdId == 0 || sysQrcodeConfig.SceneIdLogin == 0 || sysQrcodeConfig.SceneIdBind == 0 {
		return nil, errors.New("系统二维码配置错误")
	}
	apiResult, err := callQrcodeApi(ctx, sysQrcodeConfig, true)
	if err != nil {
		g.Log().Errorf(ctx, "获取二维码失败, 错误: %v", err)
		return nil, gerror.New("获取二维码失败")
	}

	expireTime := gtime.Now().Add(time.Second * time.Duration(apiResult.ExpireTime))
	// 将结果保存到redis中
	tempUserId := uuid.New().String()
	redisKey := fmt.Sprintf(consts.KeyLoginScanWechat, sysQrcodeConfig.SceneIdLogin)
	cacheModel := &model.QrcodeCacheModel{
		Ticket:     apiResult.Ticket,
		ExpireTime: expireTime,
		QrUrl:      apiResult.QrUrl,
		Scanned:    false,
		OpenId:     "",
	}
	mapData := make(map[string]interface{})
	mapData[tempUserId] = cacheModel
	g.Redis().HSet(ctx, redisKey, mapData)

	return &model.QrcodeLoginModel{
		TempUserId:       tempUserId,
		QrcodeCacheModel: *cacheModel,
		Bound:            false,
		Expired:          false,
	}, nil
}

// 对openId进行加密
func openIdEncode(openId string) string {
	data := []byte(openId)
	hash := md5.Sum(data)
	md5String := hex.EncodeToString(hash[:])
	return base64.StdEncoding.EncodeToString([]byte(openId + "," + md5String))
}

// 对openId进行解密
func openIdDecode(openIdEncode string) (string, error) {
	decode, err := base64.StdEncoding.DecodeString(openIdEncode)
	if err != nil {
		return "", err
	}
	strs := strings.Split(string(decode), ",")
	if len(strs) != 2 {
		return "", errors.New("openId 格式错误")
	}

	openId := strs[0]
	md5String := strs[1]

	data := []byte(openId)
	hash := md5.Sum(data)
	md5String2 := hex.EncodeToString(hash[:])
	if md5String != md5String2 {
		return "", errors.New("openId 格式错误")
	}

	return openId, nil
}

// 对token进行加密
func tokenEncode(token string) (timeStamp int64, md5String string) {
	nowTime := time.Now().Unix()
	tokenStr := fmt.Sprintf("%d%s", nowTime, token)
	data := []byte(tokenStr)
	hash := md5.Sum(data)
	md5String = hex.EncodeToString(hash[:])
	return nowTime, md5String
}

func tokenTimeStampEncode(timeStamp int64, token string) (md5String string) {
	tokenStr := fmt.Sprintf("%d%s", timeStamp, token)
	data := []byte(tokenStr)
	hash := md5.Sum(data)
	md5String = hex.EncodeToString(hash[:])
	return md5String
}

// 获取登录二维码扫码结果
func (l *sSysQrcode) GetQrcodeLoginStatus(ctx context.Context, tempUserId string) (res *model.QrcodeLoginModel, err error) {
	sysQrcodeConfig, err := getSysQrcodeConfig(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取系统二维码配置失败, 错误: %v", err)
		return nil, gerror.New("获取系统二维码配置失败")
	}

	res = &model.QrcodeLoginModel{
		TempUserId: tempUserId,
		QrcodeCacheModel: model.QrcodeCacheModel{
			Ticket:     "",
			ExpireTime: gtime.Now(),
			QrUrl:      "",
			Scanned:    false,
			OpenId:     "",
		},
		Bound:   false,
		Expired: false,
	}

	// 获取不到缓存，统一标记为过期
	redisKey := fmt.Sprintf(consts.KeyLoginScanWechat, sysQrcodeConfig.SceneIdLogin)
	cashValue, err := g.Redis().HGet(ctx, redisKey, tempUserId)
	if err != nil || cashValue == nil || cashValue.Map() == nil {
		g.Log().Errorf(ctx, "未获取到%s的二维码扫码结果, 错误: %v", tempUserId, err)
		res.Expired = true
		return res, nil
	}

	// 将val 的数据保存到cacheModel中'
	cacheModel := &model.QrcodeCacheModel{}
	g.Log().Debugf(ctx, "获取登录二维码扫码结果, 缓存数据: %v", cashValue.Map())
	mapData := cashValue.Map()
	err = gconv.Struct(mapData, cacheModel)
	if err != nil {
		g.Log().Errorf(ctx, "获取登录二维码扫码结果失败, 错误: %v", err)
		res.Expired = true
		return res, nil
	}

	// 将缓存结果赋值给返回结果
	res.Ticket = cacheModel.Ticket
	res.ExpireTime = cacheModel.ExpireTime
	res.QrUrl = cacheModel.QrUrl
	res.Scanned = cacheModel.Scanned
	res.OpenId = openIdEncode(cacheModel.OpenId)
	g.Log().Debugf(ctx, "获取登录二维码扫码结果, 缓存数据: %v", cacheModel)

	//先看扫码状态
	if !cacheModel.Scanned {
		// 未扫码 需要判断是否过期
		if cacheModel.ExpireTime.Before(gtime.Now()) {
			res.Expired = true
			// 删除redisKey
			g.Redis().Do(ctx, "HDEL", redisKey, tempUserId)
		}
		return res, nil
	}
	// 已扫码 需要判断是否绑定
	//根据openId和source查询数据库
	userBindList, _, err := service.SysSocial().List(ctx, &model.SysSocialListParam{
		OpenId: cacheModel.OpenId,
		Source: "wechat_mp",
	}, &request.PageInfo{Page: 1, PageSize: 100})
	if err != nil {
		g.Log().Errorf(ctx, "获取社会化关系表信息失败, 错误: %v", err)
		return nil, gerror.New("获取绑定信息失败")
	}
	if len(userBindList) == 0 {
		res.Bound = false
		// 删除redisKey
		g.Redis().Do(ctx, "HDEL", redisKey, tempUserId)
		return res, nil
	}
	validBindUserList := make([]*model.SysSocialListModel, 0)
	tenantIdList := gset.NewStrSet()
	for _, bindUser := range userBindList {
		user, err := service.SysUser().GetUserByUsername(ctx, bindUser.UserName, bindUser.TenantId)
		if err != nil || user == nil {
			g.Log().Errorf(ctx, "关联信息无效, Id:%d, openId:%s, source:%s, userName:%s, tenantId:%s", bindUser.Id, cacheModel.OpenId, "wechat_mp", bindUser.UserName, bindUser.TenantId)
			continue
		}
		bindUser.User = &model.SysUserViewModel{
			UserId:   user.UserId,
			UserName: user.UserName,
			NickName: user.NickName,
			Avatar:   user.Avatar,
			TenantId: user.TenantId,
		}
		tenantIdList.Add(bindUser.TenantId)
		validBindUserList = append(validBindUserList, bindUser)
	}
	//租户Idmap
	tenantMap := make(map[string]*model.SysTenantListModel)
	if tenantIdList.Size() > 0 {
		// 查询租户列表
		tenantList := make([]*model.SysTenantListModel, 0)
		err := dao.SysTenant.Ctx(ctx).WhereIn(dao.SysTenant.Columns().TenantId, tenantIdList.Slice()).Scan(&tenantList)
		if err != nil {
			g.Log().Errorf(ctx, "获取租户列表失败, 错误: %v", err)
			return nil, gerror.New("获取租户列表失败")
		}
		for _, tenant := range tenantList {
			tenantMap[tenant.TenantId] = tenant
		}
	}

	if len(validBindUserList) == 1 {
		//单账号绑定 执行登录操作
		bindUser := validBindUserList[0]
		res.Bound = true

		// 执行登录操作
		userOut, token, err := service.SysAuth().LoginByOpenId(ctx, bindUser)
		if err != nil {
			g.Log().Errorf(ctx, "登录操作失败, 错误: %v", err)
			return nil, gerror.New("登录失败, 请稍后重试")
		}
		res.LoginUserOut = userOut
		res.Token = token
	} else {
		// 多账号绑定，返回用户列表
		userList := make([]*model.SysUserSocialModel, 0)

		for _, bindUser := range validBindUserList {
			user := bindUser.User
			tenant := tenantMap[user.TenantId]
			tenantName := ""
			if tenant != nil {
				tenantName = tenant.CompanyName
			}
			userList = append(userList, &model.SysUserSocialModel{
				TenantId:    user.TenantId,
				UserId:      user.UserId,
				UserName:    user.UserName,
				NickName:    user.NickName,
				Phonenumber: user.Phonenumber,
				Avatar:      user.Avatar,
				TenantName:  tenantName,
			})
		}
		res.UserList = userList

	}
	// 删除redisKey
	g.Redis().Do(ctx, "HDEL", redisKey, tempUserId)
	return res, nil
}

// 获取绑定二维码
func (l *sSysQrcode) GetQrcodeBind(ctx context.Context) (res *model.QrcodeCacheModel, err error) {
	userId := contexts.GetUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("用户未登录")
	}

	sysQrcodeConfig, err := getSysQrcodeConfig(ctx)
	if err != nil {
		return nil, err
	}
	if sysQrcodeConfig.ServerApi == "" || sysQrcodeConfig.Token == "" || sysQrcodeConfig.ThirdId == 0 || sysQrcodeConfig.SceneIdLogin == 0 || sysQrcodeConfig.SceneIdBind == 0 {
		return nil, errors.New("系统二维码配置错误")
	}
	apiResult, err := callQrcodeApi(ctx, sysQrcodeConfig, false)
	if err != nil {
		g.Log().Errorf(ctx, "获取绑定二维码失败, 错误: %v", err)
		return nil, gerror.New("获取绑定二维码失败")
	}

	expireTime := gtime.Now().Add(time.Second * time.Duration(apiResult.ExpireTime))
	// 将结果保存到redis中
	res = &model.QrcodeCacheModel{
		Ticket:     apiResult.Ticket,
		ExpireTime: expireTime,
		QrUrl:      apiResult.QrUrl,
		Scanned:    false,
		OpenId:     "",
	}

	redisKey := fmt.Sprintf(consts.KeyUserBindWechat, sysQrcodeConfig.SceneIdBind)
	mapData := make(map[string]interface{})
	g.Log().Debugf(ctx, "获取绑定二维码, 缓存数据: %v", res)
	mapData[strconv.FormatInt(userId, 10)] = res
	g.Log().Debugf(ctx, "用户 %d 设置绑定二维码缓存: key: %s  field: %s  value: %v", userId, redisKey, strconv.FormatInt(userId, 10), res)
	g.Redis().HSet(ctx, redisKey, mapData)

	return res, nil
}

// 获取绑定二维码扫码结果
func (l *sSysQrcode) GetQrcodeBindStatus(ctx context.Context) (res *model.QrcodeBindStatusModel, err error) {
	userId := contexts.GetUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("用户未登录")
	}
	sysQrcodeConfig, err := getSysQrcodeConfig(ctx)
	if err != nil {
		return nil, err
	}
	if sysQrcodeConfig.ServerApi == "" || sysQrcodeConfig.Token == "" || sysQrcodeConfig.ThirdId == 0 || sysQrcodeConfig.SceneIdLogin == 0 || sysQrcodeConfig.SceneIdBind == 0 {
		return nil, errors.New("系统二维码配置错误")
	}

	redisKey := fmt.Sprintf(consts.KeyUserBindWechat, sysQrcodeConfig.SceneIdBind)
	val, err := g.Redis().HGet(ctx, redisKey, strconv.FormatInt(userId, 10))
	g.Log().Debugf(ctx, "获取绑定二维码扫码结果, 用户 %d 的缓存数据: %v", userId, val.Map())
	if err != nil || val == nil || val.Map() == nil {
		return &model.QrcodeBindStatusModel{
			QrcodeCacheModel: model.QrcodeCacheModel{
				Ticket:     "",
				ExpireTime: gtime.Now(),
				QrUrl:      "",
			},
			Expired: true,
			Bound:   false,
		}, nil
	}

	cacheModel := &model.QrcodeCacheModel{}
	mapData := val.Map()
	err = gconv.Struct(mapData, cacheModel)
	if err != nil {
		return nil, gerror.New("获取绑定二维码扫码结果失败")
	}
	res = &model.QrcodeBindStatusModel{
		QrcodeCacheModel: *cacheModel,
		Bound:            false,
		Expired:          false,
	}
	if !cacheModel.Scanned {
		// TODO 判断是否过期
		if cacheModel.ExpireTime.Before(gtime.Now()) {
			res.Expired = true
			// 删除redisKey
			g.Redis().Do(ctx, "HDEL", redisKey, strconv.FormatInt(userId, 10))
		}
		return res, nil
	} else {
		// 已扫码 需要判断是否绑定
		// 判断是否已绑定
		userBindList, _, err := service.SysSocial().List(ctx, &model.SysSocialListParam{
			UserId: userId,
			OpenId: cacheModel.OpenId,
			Source: "wechat_mp",
		}, &request.PageInfo{Page: 1, PageSize: 1})
		if err != nil {
			return nil, gerror.New("获取绑定信息失败")
		}
		if len(userBindList) > 0 {
			res.Bound = true
		} else {
			res.Bound = false
		}
		// 删除redisKey
		g.Redis().Do(ctx, "HDEL", redisKey, strconv.FormatInt(userId, 10))
	}
	return res, nil
}

// 扫码回调
func (l *sSysQrcode) QrcodeScanCallback(ctx context.Context, param *model.QrcodeScanCallbackParam) (err error) {
	g.Log().Debug(ctx, "扫码回调, 参数: %v", param)
	if param.TimeStamp == 0 || param.Sign == "" {
		return gerror.New("签名或时间戳不能为空")
	}

	sysQrcodeConfig, err := getSysQrcodeConfig(ctx)
	if err != nil {
		return err
	}
	if sysQrcodeConfig.ServerApi == "" || sysQrcodeConfig.Token == "" || sysQrcodeConfig.ThirdId == 0 || sysQrcodeConfig.SceneIdLogin == 0 || sysQrcodeConfig.SceneIdBind == 0 {
		return errors.New("系统二维码配置错误")
	}

	//nowTime := gtime.Now().Unix()
	sign := tokenTimeStampEncode(param.TimeStamp, sysQrcodeConfig.Token)
	// if math.Abs(float64(param.TimeStamp-nowTime)) > 60 {
	// 	return gerror.New("时间戳无效")
	// }
	if param.Sign != sign {
		return gerror.New("签名无效")
	}

	if param.SceneId == sysQrcodeConfig.SceneIdLogin {
		redisKey := fmt.Sprintf(consts.KeyLoginScanWechat, param.SceneId)
		// 登录回调 更新一个扫码缓存状态为已扫码
		val, err := g.Redis().HGetAll(ctx, redisKey)
		if err != nil {
			return gerror.New("获取登录二维码扫码结果失败")
		}

		g.Log().Debugf(ctx, "登录回调, 获取登录二维码扫码结果: %v", val)
		mapData := val.Map()
		if len(mapData) == 0 {
			return gerror.New("获取登录二维码扫码结果失败")
		}

		for tempUserId, v := range mapData {
			cacheModel := &model.QrcodeCacheModel{}
			err = gconv.Struct(v, cacheModel)
			if err != nil {
				g.Log().Errorf(ctx, "缓存数据转换失败,数据: %v, 错误: %v", v, err)
				continue
			}
			if cacheModel.Scanned || cacheModel.Ticket != param.Ticket {
				continue
			}
			// 未扫码
			cacheModel.Scanned = true
			cacheModel.OpenId = param.OpenId
			updateMap := make(map[string]interface{})
			updateMap[tempUserId] = cacheModel
			g.Redis().HSet(ctx, redisKey, updateMap)
			break
		}
	}
	if param.SceneId == sysQrcodeConfig.SceneIdBind {
		// 绑定回调  保存绑定结果
		redisKey := fmt.Sprintf(consts.KeyUserBindWechat, param.SceneId)
		val, err := g.Redis().HGetAll(ctx, redisKey)
		if err != nil {
			return gerror.New("获取绑定二维码扫码结果失败 redisKey: " + redisKey)
		}
		mapData := val.Map()
		if len(mapData) == 0 {
			return gerror.New("获取绑定二维码扫码结果失败 redisKey: " + redisKey)
		}
		for userId, v := range mapData {
			cacheModel := &model.QrcodeCacheModel{}
			err = gconv.Struct(v, cacheModel)
			if err != nil {
				g.Log().Errorf(ctx, "缓存数据转换失败,数据: %v, 错误: %v", v, err)
				continue
			}
			if cacheModel.Scanned || cacheModel.Ticket != param.Ticket {
				continue
			}
			// 未扫码
			cacheModel.Scanned = true
			cacheModel.OpenId = param.OpenId
			updateMap := make(map[string]interface{})
			updateMap[userId] = cacheModel
			g.Redis().HSet(ctx, redisKey, updateMap)

			user := &entity.SysUser{}
			// 查询用户信息
			err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserId, gconv.Int64(userId)).Scan(&user)
			if err != nil {
				g.Log().Errorf(ctx, "查询用户信息失败, 错误: %v", err)
				return gerror.New("查询用户信息失败")
			}
			if user == nil {
				g.Log().Errorf(ctx, "用户不存在, 用户ID: %s", userId)
				return gerror.New("用户不存在")
			}

			// 保存绑定结果
			err = service.SysSocial().Create(ctx, &model.SysSocialSaveParam{
				UserId:      gconv.Int64(userId),
				AuthId:      fmt.Sprintf("%s-%s", consts.SourceWechatMp, param.OpenId),
				OpenId:      param.OpenId,
				Source:      consts.SourceWechatMp,
				UserName:    user.UserName,
				AccessToken: param.OpenId,
				TenantId:    user.TenantId,
				CreatedBy:   0, // 创建者为0 表示系统自动创建
				CreatedAt:   gtime.Now(),
			})
			if err != nil {
				g.Log().Errorf(ctx, "保存绑定结果失败, 错误: %v", err)
				break
			}
			break
		}
	}
	return nil
}

// 扫码后选择用户登录
func (l *sSysQrcode) QrcodeLoginSelectUserId(ctx context.Context, param *model.QrcodeLoginSelectUserIdParam) (res *model.QrcodeLoginAndBindOpenIdModel, err error) {
	openId, err := openIdDecode(param.OpenId)
	if err != nil {
		g.Log().Errorf(ctx, "openId 解密失败, 错误: %v", err)
		return nil, err
	}
	g.Log().Debugf(ctx, "扫码后选择用户登录, 参数: %v", param)
	user, err := service.SysUser().GetUserByUsername(ctx, param.UserName, param.TenantId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.New("用户不存在")
	}

	userOut, token, err := service.SysAuth().LoginByOpenId(ctx, &model.SysSocialListModel{
		UserId:    user.UserId,
		OpenId:    openId,
		Source:    consts.SourceWechatMp,
		TenantId:  user.TenantId,
		CreatedBy: 0, // 创建者为0 表示系统自动创建
		CreatedAt: gtime.Now(),
		User: &model.SysUserViewModel{
			UserId:   user.UserId,
			UserName: user.UserName,
			NickName: user.NickName,
			Avatar:   user.Avatar,
			TenantId: user.TenantId,
		},
	})
	return &model.QrcodeLoginAndBindOpenIdModel{
		LoginUserOut: userOut,
		Token:        token,
	}, nil
}

// 登录并绑定
func (l *sSysQrcode) QrcodeLoginAndBindOpenId(ctx context.Context, param *model.QrcodeLoginAndBindOpenIdParam) (res *model.QrcodeLoginAndBindOpenIdModel, err error) {
	openId, err := openIdDecode(param.OpenId)
	if err != nil {
		g.Log().Errorf(ctx, "openId 解密失败, 错误: %v", err)
		return nil, err
	}
	// 1. 执行登录
	userOut, token, err := service.SysAuth().Login(ctx, &model.LoginParams{
		Username:     param.Username,
		Password:     param.Password,
		CaptchaID:    param.CaptchaID,
		CaptchaValue: param.CaptchaValue,
		TenantId:     param.TenantId,
	})
	if err != nil {
		return nil, err
	}
	// 2. 执行绑定
	err = service.SysSocial().Create(ctx, &model.SysSocialSaveParam{
		UserId:      userOut.ID,
		AuthId:      fmt.Sprintf("%s-%s", consts.SourceWechatMp, openId),
		OpenId:      openId,
		Source:      consts.SourceWechatMp,
		UserName:    param.Username,
		AccessToken: openId,
		TenantId:    param.TenantId,
		CreatedBy:   0, // 创建者为0 表示系统自动创建
		CreatedAt:   gtime.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &model.QrcodeLoginAndBindOpenIdModel{
		LoginUserOut: userOut,
		Token:        token,
	}, nil
}

// 注册并绑定
func (l *sSysQrcode) QrcodeRegisterAndBindOpenId(ctx context.Context, param *model.QrcodeRegisterAndBindOpenIdParam) (res *model.QrcodeLoginAndBindOpenIdModel, err error) {
	openId, err := openIdDecode(param.OpenId)
	if err != nil || len(openId) == 0 {
		g.Log().Errorf(ctx, "openId 解密失败, 错误: %v", err)
		return nil, err
	}
	// 1. 执行注册
	err = service.SysUser().Register(ctx, &model.SysUserRegisterModel{
		UserName:     param.UserName,
		Password:     param.Password,
		TenantId:     param.TenantId,
		CaptchaID:    param.CaptchaID,
		CaptchaValue: param.CaptchaValue,
	})
	if err != nil {
		return nil, err
	}

	user, err := service.SysUser().GetUserByUsername(ctx, param.UserName, param.TenantId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.New("用户注册失败")
	}
	// 2. 执行绑定
	err = service.SysSocial().Create(ctx, &model.SysSocialSaveParam{
		UserId:      user.UserId,
		AuthId:      fmt.Sprintf("%s-%s", consts.SourceWechatMp, openId),
		OpenId:      openId,
		Source:      consts.SourceWechatMp,
		UserName:    param.UserName,
		AccessToken: openId,
		TenantId:    param.TenantId,
		CreatedBy:   0, // 创建者为0 表示系统自动创建
		CreatedAt:   gtime.Now(),
	})
	if err != nil {
		return nil, err
	}

	// 3. 执行登录
	userOut, token, err := service.SysAuth().LoginByOpenId(ctx, &model.SysSocialListModel{
		UserId:    user.UserId,
		OpenId:    openId,
		Source:    consts.SourceWechatMp,
		TenantId:  param.TenantId,
		CreatedBy: 0, // 创建者为0 表示系统自动创建
		CreatedAt: gtime.Now(),
		User: &model.SysUserViewModel{
			UserId:   user.UserId,
			UserName: user.UserName,
			NickName: user.NickName,
			Avatar:   user.Avatar,
			TenantId: user.TenantId,
		},
	})
	if err != nil {
		return nil, err
	}
	return &model.QrcodeLoginAndBindOpenIdModel{
		LoginUserOut: userOut,
		Token:        token,
	}, nil
}
