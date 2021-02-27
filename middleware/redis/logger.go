// Package redis...
//
// Description : redis...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-27 5:26 下午
package redis

import (
	"github.com/go-developer/gopkg/logger"
	"go.uber.org/zap/zapcore"
)

// LoggerConfig 日志配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:26 下午 2021/2/27
type LoggerConfig struct {
	LoggerPath    string
	LoggerFile    string
	LoggerLevel   zapcore.Level
	ConsoleOutput bool
	Encoder       zapcore.Encoder
	SplitConfig   *logger.RotateLogConfig
}

// LogFieldConfig 日志字段配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:20 下午 2021/2/27
type LogFieldConfig struct {
	Message       string
	UsedTimeField string
	CommandField  string
	FlagField     string
}
