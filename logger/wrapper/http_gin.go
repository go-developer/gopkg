// Package wrapper...
//
// Description : http_gin 使用gin框架时的，记录日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-01-03 3:43 下午
package wrapper

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/go-developer/gopkg/logger"
	logger2 "github.com/go-developer/gopkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewGinLogger 使用gin框架记录日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 3:45 下午 2021/1/3
func NewGinLogger(loggerLevel zapcore.Level, consoleOutput bool, encoder zapcore.Encoder, splitConfig *logger.RotateLogConfig, extractFieldList []string) (*Gin, error) {
	var (
		err error
		l   *zap.Logger
	)
	logConfList := []logger2.SetLoggerOptionFunc{logger2.WithEncoder(encoder), logger2.WithCaller(), logger2.WithCallerSkip(1)}
	if consoleOutput {
		logConfList = append(logConfList, logger2.WithConsoleOutput())
	}
	if l, err = logger.NewLogger(loggerLevel, splitConfig, logConfList...); nil != err {
		return nil, err
	}

	return &Gin{
		loggerInstance:   l,
		extractFieldList: extractFieldList,
	}, nil
}

// Gin 包装gin实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 3:59 下午 2021/1/3
type Gin struct {
	loggerInstance   *zap.Logger // zap 的日志实例
	extractFieldList []string    // 从gin中抽取的字段
}

// formatFieldList 格式化日志field列表
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:13 下午 2021/1/3
func (gw *Gin) formatFieldList(ginCtx *gin.Context, inputFieldList []zap.Field) []zap.Field {
	if nil == inputFieldList {
		inputFieldList = make([]zap.Field, 0)
	}
	if nil != ginCtx {
		// 自动扩充抽取字段,字段不存在的话,忽略掉
		for _, extractField := range gw.extractFieldList {
			if v, exist := ginCtx.Get(extractField); exist {
				byteData, _ := json.Marshal(v)
				inputFieldList = append(inputFieldList, zap.String(extractField, string(byteData)))
			}
		}
	}
	return inputFieldList
}

// Debug 日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:14 下午 2021/1/3
func (gw *Gin) Debug(ginCtx *gin.Context, msg string, field ...zap.Field) {
	fieldList := gw.formatFieldList(ginCtx, field)
	gw.loggerInstance.Debug(msg, fieldList...)
}

// Info 日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:28 下午 2021/1/3
func (gw *Gin) Info(ginCtx *gin.Context, msg string, field ...zap.Field) {
	fieldList := gw.formatFieldList(ginCtx, field)
	gw.loggerInstance.Info(msg, fieldList...)
}

// Warn 日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:29 下午 2021/1/3
func (gw *Gin) Warn(ginCtx *gin.Context, msg string, field ...zap.Field) {
	fieldList := gw.formatFieldList(ginCtx, field)
	gw.loggerInstance.Warn(msg, fieldList...)
}

// Error 日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:29 下午 2021/1/3
func (gw *Gin) Error(ginCtx *gin.Context, msg string, field ...zap.Field) {
	fieldList := gw.formatFieldList(ginCtx, field)
	gw.loggerInstance.Error(msg, fieldList...)
}

// Panic 日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:29 下午 2021/1/3
func (gw *Gin) Panic(ginCtx *gin.Context, msg string, field ...zap.Field) {
	fieldList := gw.formatFieldList(ginCtx, field)
	gw.loggerInstance.Panic(msg, fieldList...)
}

// DPanic 日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:30 下午 2021/1/3
func (gw *Gin) DPanic(ginCtx *gin.Context, msg string, field ...zap.Field) {
	fieldList := gw.formatFieldList(ginCtx, field)
	gw.loggerInstance.DPanic(msg, fieldList...)
}

// GetZapLoggerInstance 获取zap日志实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021/01/03 22:56:47
func (gw *Gin) GetZapLoggerInstance() *zap.Logger {
	return gw.loggerInstance
}
