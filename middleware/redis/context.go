// Package redis...
//
// Description : redis...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-27 8:22 下午
package redis

import (
	"context"

	"github.com/go-developer/gopkg/easymap"

	"github.com/gin-gonic/gin"
)

const (
	// 默认的 request_id 字段名
	defaultRequestIDField = "request_id"
	// 默认的message
	defaultMessage = "执行redis命令日志记录"
	// 耗时字段
	defaultUsedTimeField = "used_field"
	// 默认的命令字段
	defaultCommandField = "command"
	// 默认记录 redis标识的字段
	defaultFlagField = "flag"
)

// Context 请求上下文
//
// Author : go_developer@163.com<张德满>
//
// Date : 8:25 下午 2021/2/27
type Context struct {
	Flag           string          // 哪个模块的上下文
	Ctx            context.Context // ctx
	GinCtx         *gin.Context    // http 请求绑定的gin.context
	RequestIDField string          // requestID 字段名
	RequestID      string          // requestID 此字段有值, 直接使用此值,无值, 去GinCtx 中读取 RequestIDField
	Extra          easymap.EasyMap // 扩展信息
}

// NewContext 生成一个上下文
//
// Author : go_developer@163.com<张德满>
//
// Date : 8:26 下午 2021/2/27
func NewContext(flag string, of ...SetContextFunc) *Context {
	ctx := &Context{
		Flag:           flag,
		Ctx:            nil,
		GinCtx:         nil,
		RequestIDField: "",
		RequestID:      "",
		Extra:          nil,
	}
	for _, f := range of {
		f(ctx)
	}
	if nil == ctx.Ctx {
		ctx.Ctx = context.Background()
	}
	if len(ctx.RequestIDField) == 0 {
		ctx.RequestIDField = defaultRequestIDField
	}
	if nil == ctx.Extra {
		ctx.Extra = easymap.NewNormal(true)
	}
	// requestID 填充
	if len(ctx.RequestID) == 0 {
		// 先从 gin 读
		if nil != ctx.Ctx {
			ctx.RequestID = ctx.GinCtx.GetString(ctx.RequestIDField)
		}
		// 再从extra读取
		if len(ctx.RequestID) == 0 {
			ctx.RequestID, _ = ctx.Extra.GetString(ctx.RequestID)
		}
	}
	return ctx
}

// SetContextFunc 设置context参数
type SetContextFunc func(rc *Context)

// WithCtx 设置context
//
// Author : go_developer@163.com<张德满>
//
// Date : 8:30 下午 2021/2/27
func WithCtx(ctx context.Context) SetContextFunc {
	return func(rc *Context) {
		rc.Ctx = ctx
	}
}

// WithGinCtx 设置gin上下文
//
// Author : go_developer@163.com<张德满>
//
// Date : 8:34 下午 2021/2/27
func WithGinCtx(ginCtx *gin.Context) SetContextFunc {
	return func(rc *Context) {
		rc.GinCtx = ginCtx
	}
}

// WithExtra 设置扩展信息
//
// Author : go_developer@163.com<张德满>
//
// Date : 8:36 下午 2021/2/27
func WithExtra(extra easymap.EasyMap) SetContextFunc {
	return func(rc *Context) {
		rc.Extra = extra
	}
}

// WithRequestIDField 设置request_id参数名
//
// Author : go_developer@163.com<张德满>
//
// Date : 8:41 下午 2021/2/27
func WithRequestIDField(requestIDField string) SetContextFunc {
	return func(rc *Context) {
		rc.RequestIDField = requestIDField
	}
}

// WithRequestID ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 8:42 下午 2021/2/27
func WithRequestID(requestID string) SetContextFunc {
	return func(rc *Context) {
		rc.RequestID = requestID
	}
}
