// Package gin ...
//
// Description : 便捷的相关API处理
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-26 2:06 下午
package api

import "github.com/gin-gonic/gin"

// IApi 每一个接口的实现约束
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:08 下午 2021/3/26
type IApi interface {
	// GetMethod 接口请求方法
	GetMethod() string
	// GetURI 接口URI
	GetURI() string
	// GetMiddleWareList 使用的中间件列表
	GetMiddleWareList() []gin.HandlerFunc
	// 处理的handler
	GetHandler() gin.HandlerFunc
}
