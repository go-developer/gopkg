// Package redis ...
//
// Description : redis 客户端
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-27 4:49 下午
package redis

import (
	"time"

	"github.com/go-developer/gopkg/convert"

	"github.com/go-developer/gopkg/logger"
	redisInstance "github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// Options 连接选项,百分之百兼容第三方包的选项
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:57 下午 2021/2/27
type Options struct {
	Conf              *redisInstance.Options // 第三方包的选项
	Logger            *LoggerConfig          // 日志的配置
	LoggerFieldConfig *LogFieldConfig        // 日志字段的配置
}

// RealClient 包装好的 redis client
type RealClient struct {
	Flag              string                // redis 标识
	Instance          *redisInstance.Client // redis 实例
	Logger            *zap.Logger           // 日志实例
	LoggerFieldConfig *LogFieldConfig       // 日志字段的配置
}

// NewClient 获取redis client实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:05 下午 2021/2/27
func NewClient(config map[string]Options) (*Client, error) {
	c := &Client{
		instanceTable: make(map[string]*RealClient),
		confTable:     config,
	}
	return c, c.init()
}

// Client 包装的redis client
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:52 下午 2021/2/27
type Client struct {
	instanceTable map[string]*RealClient // redis 实例
	confTable     map[string]Options     // redis 配置
}

// init 初始化redis连接
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:31 下午 2021/2/27
func (c *Client) init() error {
	var (
		err error
	)

	for flag, conf := range c.confTable {
		c.instanceTable[flag] = &RealClient{
			Flag:              flag,
			Instance:          redisInstance.NewClient(conf.Conf),
			Logger:            nil,
			LoggerFieldConfig: conf.LoggerFieldConfig,
		}
		if c.instanceTable[flag].Logger, err = c.getLogger(conf.Logger); nil != err {
			return LoggerInitFail(flag, err)
		}
		if nil == c.instanceTable[flag].LoggerFieldConfig {
			c.instanceTable[flag].LoggerFieldConfig = &LogFieldConfig{
				Message:       "",
				UsedTimeField: "",
				CommandField:  "",
				FlagField:     "",
			}
		}
		if len(c.instanceTable[flag].LoggerFieldConfig.Message) == 0 {
			c.instanceTable[flag].LoggerFieldConfig.Message = defaultMessage
		}
		if len(c.instanceTable[flag].LoggerFieldConfig.CommandField) == 0 {
			c.instanceTable[flag].LoggerFieldConfig.CommandField = defaultCommandField
		}
		if len(c.instanceTable[flag].LoggerFieldConfig.UsedTimeField) == 9 {
			c.instanceTable[flag].LoggerFieldConfig.UsedTimeField = defaultUsedTimeField
		}
		if len(c.instanceTable[flag].LoggerFieldConfig.FlagField) == 0 {
			c.instanceTable[flag].LoggerFieldConfig.FlagField = defaultFlagField
		}
	}
	return nil
}

// getLogger ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 7:07 下午 2021/2/27
func (c *Client) getLogger(conf *LoggerConfig) (*zap.Logger, error) {
	if nil == conf || nil == conf.SplitConfig {
		return nil, nil
	}
	return logger.NewLogger(
		conf.LoggerLevel,
		conf.ConsoleOutput,
		conf.Encoder,
		conf.SplitConfig,
	)
}

// GetRedisClient 获取redis实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:16 下午 2021/2/27
func (c *Client) GetRedisClient(flag string) (*RealClient, error) {
	redisClient, exist := c.instanceTable[flag]
	if !exist {
		return nil, FlagNotFound(flag)
	}
	return redisClient, nil
}

// log 记录redis请求日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 8:52 下午 2021/2/27
func (c *Client) log(ctx *Context, realClient *RealClient, cmdResult redisInstance.Cmder, startTime int64, finishTime int64) {
	if nil == realClient || nil == realClient.Logger {
		return
	}
	realClient.Logger.Info(
		"执行redis命令日志记录",
		zap.Any(ctx.RequestIDField, ctx.RequestID),                                                 // 上下文串联的requestID
		zap.String(realClient.LoggerFieldConfig.CommandField, cmdResult.String()),                  // 执行的命令
		zap.Float64(realClient.LoggerFieldConfig.UsedTimeField, float64(finishTime-startTime)/1e6), // 耗时,单位: ms
		zap.Error(cmdResult.Err()),                                                                 // 异常信息
	)
}

// CommandProxy 执行命令的代理
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:41 下午 2021/2/27
func (c *Client) CommandProxy(ctx *Context, flag string, cmd string, param ...interface{}) (interface{}, error) {
	var (
		realClient *RealClient
		err        error
	)
	if len(cmd) == 0 {
		return nil, EmptyCmd()
	}
	if realClient, err = c.GetRedisClient(ctx.Flag); nil != err {
		return nil, err
	}
	redisCmd := append([]interface{}{cmd}, param...)
	startTime := time.Now().Unix()
	cmdResult := realClient.Instance.Do(ctx.Ctx, redisCmd...)
	go c.log(ctx, realClient, cmdResult, startTime, time.Now().UnixNano())
	return cmdResult.Val(), cmdResult.Err()
}

// CommandProxyWithReceiver 执行命令,并解析结果
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:00 下午 2021/2/27
func (c *Client) CommandProxyWithReceiver(ctx *Context, flag string, receiver interface{}, cmd string, param ...interface{}) error {
	if nil == receiver {
		return ReceiverISNIL()
	}
	var (
		err    error
		result interface{}
	)

	if result, err = c.CommandProxy(ctx, flag, cmd, param); nil != err {
		return err
	}

	return ResultConvertFail(convert.ConvertAssign(receiver, result))
}

// Set set 命令
//
// Author : go_developer@163.com<张德满>
//
// Date : 8:18 下午 2021/2/27
func (c *Client) Set(ctx *Context, key string, value interface{}, expiration time.Duration) error {
	var (
		realClient *RealClient
		err        error
		statusCmd  *redisInstance.StatusCmd
	)
	if realClient, err = c.GetRedisClient(ctx.Flag); nil != err {
		return err
	}

	startTime := time.Now().UnixNano()
	statusCmd = realClient.Instance.Set(ctx.Ctx, key, value, expiration)
	go c.log(ctx, realClient, statusCmd, startTime, time.Now().UnixNano())
	return statusCmd.Err()
}
