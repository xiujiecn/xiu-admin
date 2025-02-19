package utility

import (
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
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
