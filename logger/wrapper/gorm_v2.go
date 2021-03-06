// Package wrapper...
//
// Description : gorm v2 版本接口实现
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-01 9:52 下午
package wrapper

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"

	logger2 "github.com/go-developer/gopkg/logger"
	"gorm.io/gorm/logger"
)

// NewGormV2 获取日志实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:56 下午 2021/3/1
func NewGormV2(loggerLevel zapcore.Level, consoleOutput bool, encoder zapcore.Encoder, splitConfig *logger2.RotateLogConfig, traceIDField string, skip int) (logger.Interface, error) {
	logConfList := []logger2.SetLoggerOptionFunc{logger2.WithEncoder(encoder), logger2.WithCallerSkip(skip), logger2.WithCaller()}
	if consoleOutput {
		logConfList = append(logConfList, logger2.WithConsoleOutput())
	}
	logInstance, err := logger2.NewLogger(loggerLevel, splitConfig, logConfList...)
	if nil != err {
		return nil, err
	}
	if len(traceIDField) == 0 {
		traceIDField = "trace_id"
	}
	return &Gorm{
		instance:     logInstance,
		traceIDField: traceIDField,
	}, nil
}

// Gorm v2 版本库日志实现
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:55 下午 2021/3/1
type Gorm struct {
	instance     *zap.Logger // 日志实例
	traceIDField string      // 串联请求上下文的的ID
	flag         string      // 数据库标识
}

// LogMode ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:08 下午 2021/3/1
func (g *Gorm) LogMode(level logger.LogLevel) logger.Interface {
	return g
}

// Info 日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:18 下午 2021/3/1
func (g *Gorm) Info(ctx context.Context, s string, i ...interface{}) {
	g.instance.Info(
		"Info日志",
		zap.String(g.traceIDField, g.getTraceID(ctx)),
		zap.String("db_flag", g.flag),
		zap.String("message", fmt.Sprintf(s, i...)),
	)
}

// Warn ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:16 下午 2021/3/1
func (g *Gorm) Warn(ctx context.Context, s string, i ...interface{}) {
	g.instance.Warn(
		"SQL执行产生Warning",
		zap.String(g.traceIDField, g.getTraceID(ctx)),
		zap.String("db_flag", g.flag),
		zap.String("message", fmt.Sprintf(s, i...)),
	)
}

// Error 日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:18 下午 2021/3/1
func (g *Gorm) Error(ctx context.Context, s string, i ...interface{}) {
	g.instance.Warn(
		"SQL执行产生Error",
		zap.String(g.traceIDField, g.getTraceID(ctx)),
		zap.String("db_flag", g.flag),
		zap.String("message", fmt.Sprintf(s, i...)),
	)
}

// Trace Trace 记录
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:19 下午 2021/3/1
func (g *Gorm) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	start := begin.UnixNano()
	end := time.Now().UnixNano()
	sql := ""
	affectRows := int64(0)
	if nil != fc {
		sql, affectRows = fc()
	}

	g.instance.Info(
		"SQL执行记录",
		zap.String(g.traceIDField, g.getTraceID(ctx)),
		zap.String("db_flag", g.flag),
		zap.Int64("begin_time", start),
		zap.Int64("finish_time", end),
		zap.String("used_time", fmt.Sprintf("%fms", float64(end-start)/1e6)),
		zap.String("sql", sql),
		zap.Int64("affect_rows", affectRows),
		zap.Error(err),
	)
}

// getTraceID 获取traceID
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:11 下午 2021/3/1
func (g *Gorm) getTraceID(ctx context.Context) string {
	return fmt.Sprintf("%v", ctx.Value(g.traceIDField))
}

// GetGormSQL 获取tracefn
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:38 下午 2021/3/1
func GetGormSQL(dbClient *gorm.DB) func() (string, int64) {
	return func() (string, int64) {
		return dbClient.Dialector.Explain(dbClient.Statement.SQL.String(), dbClient.Statement.Vars...), dbClient.RowsAffected
	}
}
