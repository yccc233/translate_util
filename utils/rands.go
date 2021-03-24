package utils

import (
	"math/rand"
	"time"
)

/*
获取长度为参数的随机字符串
*/
func GetRand(length int) string {
	if length < 1 {
		return ""
	}
	var allc = "abcdefghijklmnopqrstuvwxyz0123456789"
	var str string

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		str +=  string(allc[r.Intn(len(allc))])
	}
	return str
}