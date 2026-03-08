// package utility
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package utility

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gcharset"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
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
	if ctx == nil {
		return ""
	}
	if ghttp.RequestFromCtx(ctx) == nil {
		return ""
	}
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
func GetPublicIP() (ip string, err error) {
	resp, err := http.Get("https://ifconfig.co/ip")
	if err != nil {
		fmt.Println("GetPublicIP", err)
		return
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println("GetPublicIP Body.Close", err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("GetPublicIP ReadAll", err)
		return
	}
	ip = string(body)
	// 去除空格
	ip = strings.Replace(ip, " ", "", -1)
	// 去除换行符
	ip = strings.Replace(ip, "\n", "", -1)

	return
}

// 获取两个字符串之间的字符串
func GetBetweenStr(str, start, end string) (result string) {
	startIndex := strings.Index(str, start)
	if startIndex == -1 {
		return ""
	}
	// 更新startIndex为开始字符串之后的位置
	startIndex += len(start)

	// 找到结束字符串的位置，从更新后的startIndex开始查找
	endIndex := strings.Index(str[startIndex:], end)
	if endIndex == -1 {
		return ""
	}

	// 切片操作获取中间的字符串
	result = str[startIndex : startIndex+endIndex]
	return
}

// A和B两个数组，A数组中存在，B数组中不存在
func ArrayRightDiff(a, b []int64) []int64 {
	diff := make([]int64, 0)
	for _, v := range a {
		if !slices.Contains(b, v) {
			diff = append(diff, v)
		}
	}
	return diff
}

type fileInfo struct {
	name string
	size int64
}

// WalkDir 获取目录下文件的名称和大小
func WalkDir(dirname string) ([]fileInfo, error) {
	var fileInfos []fileInfo
	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileInfos = append(fileInfos, fileInfo{name: path, size: info.Size()})
		}
		return nil
	})

	return fileInfos, err
}
func DirSize(dirname string) string {
	var (
		s        int64
		files, _ = WalkDir(dirname)
	)
	for _, n := range files {
		s += n.size
	}
	return FileSize(s)
}
func MergeAbs(path string, fileName ...string) string {
	var paths = []string{gfile.RealPath(path)}
	paths = append(paths, fileName...)
	return gfile.Join(paths...)
}

// FileSize 字节的单位转换 保留两位小数
func FileSize(fileSize int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "EB"}
	var size = float64(fileSize)
	var i int
	for i = 0; size > 1024; i++ {
		size /= 1024
	}
	return fmt.Sprintf("%.2f %s", size, units[i])
}

// UniqueSlice 切片去重
func UniqueSlice[K comparable](languages []K) []K {
	result := make([]K, 0, len(languages))
	temp := map[K]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
