// Package mysql...
//
// Description : 异常定义
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-01 9:42 下午
package mysql

import "github.com/pkg/errors"

// ConnectionOpenError 数据库连接失败
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:43 下午 2021/3/1
func ConnectionOpenError(err error) error {
	return errors.WithMessage(err, "数据库连接失败")
}

// CreateDBLogError 打开日志失败
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:53 下午 2021/3/1
func CreateDBLogError(err error) error {
	return errors.WithMessage(err, "数据库日志初始化失败")
}
