package utils

import (
	"strconv"
	"strings"
	"time"
)

// ParseDuration 支持“xdxhxmxs"和数字字符串转换成time.duration类型
func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d) //移除字符串两端的空白字符
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
