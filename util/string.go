// Package util...
//
// Description : 字符串相关的工具
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-09 6:00 下午
package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// GenRandomString 获取随机长度的字符串
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:01 下午 2021/3/9
func GenRandomString(source string, length uint) string {
	if length == 0 {
		return ""
	}
	if len(source) == 0 {
		//字符串为空，默认字符源为如下:
		source = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	strByte := []byte(source)
	var genStrByte = make([]byte, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < int(length); i++ {
		genStrByte = append(genStrByte, strByte[r.Intn(len(strByte))])
	}
	return string(genStrByte)
}

// Md5 对字符串进行md5加密
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:01 下午 2021/3/9
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
