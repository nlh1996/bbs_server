package utils

import (
	"time"
)

// GetTimeStr 将时间格式化为字符串 .
func GetTimeStr() (string) {
	//世界时间转北京时间
	duration,_ := time.ParseDuration("8h")
	now := time.Now().UTC().Add(duration)
	strTime := now.Format("2006-01-02 15:04:05")
	return strTime
}

// GetDateStr 获取当天日期
func GetDateStr() (string) {
	duration,_ := time.ParseDuration("8h")
	now := time.Now().UTC().Add(duration)
	strTime := now.Format("2006-01-02")
	return strTime
}
