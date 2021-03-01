// Package mysql ...
//
// Description : mysql客户端
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-01 9:20 下午
package mysql

import (
	"fmt"

	"github.com/go-developer/gopkg/logger/wrapper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GetDatabaseClient 获取日志实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:49 下午 2021/3/1
func GetDatabaseClient(conf *DBConfig, logConf *LogConfig) (*gorm.DB, error) {
	var (
		instance *gorm.DB
		err      error
	)

	if instance, err = gorm.Open(mysql.Open(buildConnectionDSN(conf)), &gorm.Config{}); nil != err {
		return nil, err
	}

	if len(logConf.TraceFieldName) == 0 {
		logConf.TraceFieldName = defaultTraceFieldName
	}

	if instance.Logger, err = wrapper.NewGormV2(
		logConf.Level,
		logConf.ConsoleOutput,
		logConf.Encoder,
		logConf.SplitConfig,
		logConf.TraceFieldName); nil != err {
		return nil, CreateDBLogError(err)
	}
	return instance, nil
}

// buildConnectionDSN 构建建立连接的DSN
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:34 下午 2021/3/1
func buildConnectionDSN(conf *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
		conf.Charset,
	)
}
