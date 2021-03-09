// Package gin ...
//
// Description : 结合gin框架的一些工具集
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-09 4:51 下午
package util

import (
	"net/http"
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
