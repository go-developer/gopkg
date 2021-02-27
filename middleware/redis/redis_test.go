// Package redis...
//
// Description : redis...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-27 10:14 下午
package redis

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	redisInstance "github.com/go-redis/redis/v8"
)

// TestCommandProxy ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:22 下午 2021/2/27
func TestCommandProxy(t *testing.T) {
	instance, err := NewClient(map[string]Options{
		"test_redis": Options{
			Conf: &redisInstance.Options{
				Addr: "127.0.0.1:6379",
			},
			Logger: &LoggerConfig{
				LoggerPath:    "/tmp/test-log",
				LoggerFile:    "test-pkg-redis-client.log",
				LoggerLevel:   0,
				ConsoleOutput: true,
				Encoder:       nil,
				SplitConfig:   nil,
			},
			LoggerFieldConfig: nil,
		},
	}, nil)
	if nil != err {
		panic(err.Error())
	}
	r, cmdErr := instance.CommandProxy(nil, "test_redis", "set", "command_proxy", "hello world")
	assert.Nil(t, cmdErr, "命令执行成功")
	assert.Equal(t, "OK", fmt.Sprintf("%v", r))
}
