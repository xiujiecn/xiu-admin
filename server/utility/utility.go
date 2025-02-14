package utility

import (
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
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
