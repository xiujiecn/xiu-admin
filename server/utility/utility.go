package utility

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gcharset"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mileusna/useragent"
	"golang.org/x/exp/rand"
)

func Md5(str string) string {
	return gmd5.MustEncryptString(str)
}

// 密码加密
func PasswordEncrypt(password string, salt string) string {
	return Md5(password + salt)
}

// 解析时间
func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}

// 随机生成5位字符
func RandomString(n int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 获取客户端IP
func GetClientIp(ctx context.Context) string {
	return ghttp.RequestFromCtx(ctx).GetClientIp()
}

// 获取user-agent
func GetUserAgent(ctx context.Context) string {
	return ghttp.RequestFromCtx(ctx).Header.Get("User-Agent")
}

// 获取操作系统
func GetOs(ctx context.Context) string {
	userAgent := GetUserAgent(ctx)
	ua := useragent.Parse(userAgent)
	return ua.OS
}

// 获取浏览器
func GetBrowser(ctx context.Context) string {
	userAgent := GetUserAgent(ctx)
	ua := useragent.Parse(userAgent)
	return ua.Name
}

// 获取ip所属城市
func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}
	if ip == "[::1]" || ip == "127.0.0.1" {
		return "内网IP"
	}
	url := "https://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	bytes := g.Client().GetBytes(context.TODO(), url)
	src := string(bytes)
	srcCharset := "GBK"
	tmp, _ := gcharset.ToUTF8(srcCharset, src)
	json, err := gjson.DecodeToJson(tmp)
	if err != nil {
		return ""
	}
	if json.Get("code").Int() == 0 {
		city := ""
		if strings.EqualFold(json.Get("pro").String(), json.Get("city").String()) {
			city = fmt.Sprintf("%s", json.Get("pro").String())
		} else {
			city = fmt.Sprintf("%s %s", json.Get("pro").String(), json.Get("city").String())
		}

		return city
	} else {
		return ""
	}
}

// 时间戳转 yyyy-MM-dd HH:mm:ss
func TimeStampToDateTime(timeStamp int64) string {
	tm := gtime.NewFromTimeStamp(timeStamp)
	return tm.Format("Y-m-d H:i:s")
}

// 时间戳转 yyyy-MM-dd
func TimeStampToDate(timeStamp int64) string {
	tm := gtime.NewFromTimeStamp(timeStamp)
	return tm.Format("Y-m-d")
}

// 服务端ip
func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}
