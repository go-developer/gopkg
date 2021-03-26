// Package api ...
//
// Description : 注册路由
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-26 2:13 下午
package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-developer/gopkg/gin/util"
)

var (
	// DebugLogEnable 默认打开debug日志
	DebugLogEnable = true
)

// DisableDebugLog 禁用debug日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:17 下午 2021/3/26
func DisableDebugLog() {
	DebugLogEnable = false
}

// RegisterRouter 注册一个路由
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:14 下午 2021/3/26
func RegisterRouter(router *gin.Engine, apiInstanceList ...IApi) error {
	for _, apiInstance := range apiInstanceList {
		if err := util.RegisterRouter(router, apiInstance.GetMethod(), apiInstance.GetURI(), apiInstance.GetHandler(), apiInstance.GetMiddleWareList()); nil != err {
			routerLog(err)
			panic(err.Error())
		}
	}
	return nil
}

// routerLog 记录日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:28 下午 2021/3/26
func routerLog(err error) {
	if !DebugLogEnable || nil == err {
		return
	}
	log.Fatal(err.Error())
}
