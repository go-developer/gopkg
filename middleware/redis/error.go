// Package redis...
//
// Description : redis...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-27 5:13 下午
package redis

import "github.com/pkg/errors"

// FlagNotFound flag不存在异常
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:15 下午 2021/2/27
func FlagNotFound(flag string) error {
	return errors.Errorf("标识为 %s 的redis未找到", flag)
}

// LoggerInitFail 日志初始化失败
//
// Author : go_developer@163.com<张德满>
//
// Date : 7:30 下午 2021/2/27
func LoggerInitFail(flag string, err error) error {
	return errors.Wrapf(err, "标识为 %s 的redis日志初始化失败", flag)
}

// EmptyCmd 未设置要执行的命令
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:46 下午 2021/2/27
func EmptyCmd() error {
	return errors.Errorf("未设置要执行的命令")
}

// CommandExecuteFail 命令执行失败
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:58 下午 2021/2/27
func CommandExecuteFail(err error) error {
	return errors.Wrapf(err, "命令执行异常")
}

// ReceiverISNIL 数据接收者是空指针
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:05 下午 2021/2/27
func ReceiverISNIL() error {
	return errors.Errorf("数据接收者指针为空")
}

// ResultConvertFail 数据结果解析失败
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:07 下午 2021/2/27
func ResultConvertFail(err error) error {
	return errors.Wrapf(err, "数据结果解析失败")
}
