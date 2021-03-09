// Package middleware...
//
// Description : middleware...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-09 5:52 下午
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	commonUtil "github.com/go-developer/gopkg/util"
)

// InitRequest 初始化请求信息,统一设置请求时间/请求ID等信息
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:53 下午 2021/3/9
func InitRequest() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 设置请求开始时间
		ctx.Set("start_time", time.Now().Unix())
		// 设置请求trace_id
		ctx.Set("trace_id", time.Now().Format("20060102150405")+"+"+commonUtil.GetHostIP()+"-"+commonUtil.Md5(commonUtil.GenRandomString("", 16)))
		ctx.Next()
	}
}
