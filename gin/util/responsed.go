// Package gin ...
//
// Description : 结合gin框架的一些工具集
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-09 4:51 下午
package util

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Response 向客户端响应数据
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:52 下午 2021/3/9
func Response(ctx *gin.Context, code interface{}, message string, data interface{}) {
	var responseData = gin.H{
		"code":     code,
		"message":  message,
		"data":     data,
		"trace_id": ctx.GetString("trace_id"),
		"cost":     time.Since(time.Unix(ctx.GetInt64("start_time"), 0)).Seconds(),
	}
	ctx.JSON(http.StatusOK, responseData)
}

// RegisterRouter 注册gin路由
//
// Author : go_developer@163.com<张德满>
//
// Date : 8:36 下午 2021/3/9
func RegisterRouter(router *gin.Engine, method string, uri string, handler gin.HandlerFunc, middlewareList []gin.HandlerFunc) error {
	if nil == middlewareList {
		middlewareList = make([]gin.HandlerFunc, 0)
	}
	switch strings.ToUpper(method) {
	case http.MethodGet:
		router.GET(uri, handler).Use(middlewareList...)
	case http.MethodPost:
		router.POST(uri, handler).Use(middlewareList...)
	case http.MethodDelete:
		router.DELETE(uri, handler).Use(middlewareList...)
	case http.MethodHead:
		router.HEAD(uri, handler).Use(middlewareList...)
	case http.MethodOptions:
		router.OPTIONS(uri, handler).Use(middlewareList...)
	case http.MethodPatch:
		router.PATCH(uri, handler).Use(middlewareList...)
	case http.MethodPut:
		router.PUT(uri, handler).Use(middlewareList...)
	case "ANY": // 一次性注册全部请求方法的路由
		router.Any(uri, handler).Use(middlewareList...)
	default:
		// 不是一个函数,数名method配置错误
		return fmt.Errorf("uri=%s method=%s 请求方法配置错误", uri, method)
	}
	return nil
}

// RegisterRouterGroup 注册gin路由
//
// Author : go_developer@163.com<张德满>
//
// Date : 8:36 下午 2021/3/9
func RegisterRouterGroup(router *gin.RouterGroup, method string, uri string, handler gin.HandlerFunc) error {
	switch strings.ToUpper(method) {
	case http.MethodGet:
		router.GET(uri, handler)
	case http.MethodPost:
		router.POST(uri, handler)
	case http.MethodDelete:
		router.DELETE(uri, handler)
	case http.MethodHead:
		router.HEAD(uri, handler)
	case http.MethodOptions:
		router.OPTIONS(uri, handler)
	case http.MethodPatch:
		router.PATCH(uri, handler)
	case http.MethodPut:
		router.PUT(uri, handler)
	case "ANY": // 一次性注册全部请求方法的路由
		router.Any(uri, handler)
	default:
		// 不是一个函数,数名method配置错误
		return fmt.Errorf("uri=%s method=%s 请求方法配置错误", uri, method)
	}
	return nil
}
