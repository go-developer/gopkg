// Package mysql...
//
// Description : 数据定义
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-01 9:27 下午
package mysql

import (
	"github.com/go-developer/gopkg/logger"
	"go.uber.org/zap/zapcore"
)

// DBConfig 数据库连接的配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:32 下午 2021/3/1
type DBConfig struct {
	Host              string // 主机
	Port              uint   // 端口
	Database          string // 数据库
	Username          string // 账号
	Password          string // 密码
	Charset           string // 编码
	MaxOpenConnection uint   // 打开的最大连接数
	MaxIdleConnection uint   // 最大空闲连接数
}

// LogConfig 日志配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:51 下午 2021/3/1
type LogConfig struct {
	Level            zapcore.Level
	ConsoleOutput    bool
	Encoder          zapcore.Encoder
	SplitConfig      *logger.RotateLogConfig
	ExtractFieldList []string
	TraceFieldName   string
}

const (
	// defaultTraceFieldName 默认trace_id字段
	defaultTraceFieldName = "trace_id"
)
