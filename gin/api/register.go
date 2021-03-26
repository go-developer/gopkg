// Package api ...
//
// Description : 注册路由
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-26 2:13 下午
package api

import (
	"log"
	"reflect"

	"github.com/go-developer/gopkg/gin/util"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

var (
	// DebugLogEnable 默认打开debug日志
	DebugLogEnable = true
)

// DisableDebugLog 禁用debug日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:17 下午 2021/3/26
func DisableDebugLog() {
	DebugLogEnable = false
}

// RegisterRouter 注册一个路由
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:14 下午 2021/3/26
func RegisterRouter(router *gin.Engine, apiInstanceList ...interface{}) error {
	for _, apiInstance := range apiInstanceList {
		if nil == apiInstance {
			continue
		}
		val := reflect.ValueOf(apiInstance)
		switch val.Kind() {
		case reflect.Struct:
			fallthrough
		case reflect.Ptr:
			api, ok := apiInstance.(IApi)
			if ok {
				if err := util.RegisterRouter(router, api.GetMethod(), api.GetURI(), api.GetHandler(), api.GetMiddleWareList()); nil != err {
					routerLog(err.Error())
					return err
				}
				continue
			}
			routerLog(val.String() + "结构体或者结构体指针, 自动识别函数是否包含RouterFunc")
			// 不是IApi接口,自动识别函数列表 RouterFunc 函数自动注册
			methodCnt := val.NumMethod()
			for i := 0; i < methodCnt; i++ {
				// TODO : 识别函数本身是不是 RouterFunc
				af, o := val.Method(i).Interface().(func() (string, string, gin.HandlerFunc, []gin.HandlerFunc))
				if o {
					method, uri, handler, middlewareList := af()
					if err := util.RegisterRouter(router, method, uri, handler, middlewareList); nil != err {
						routerLog(err.Error())
						return err
					}
					continue
				}
				apiFuncList := val.Method(i).Call(nil)
				for _, apiFuncVal := range apiFuncList {
					apiFunc, ok := apiFuncVal.Interface().(RouterFunc)
					if !ok {
						continue
					}
					method, uri, handler, middlewareList := apiFunc()
					if err := util.RegisterRouter(router, method, uri, handler, middlewareList); nil != err {
						routerLog(err.Error())
						return err
					}
				}

			}
		case reflect.Func:
			api, ok := apiInstance.(RouterFunc)
			if !ok {
				err := errors.New("函数方式注册路由必须是 RouterFunc")
				routerLog(err.Error())
				return err
			}
			method, uri, handler, middlewareList := api()
			if err := util.RegisterRouter(router, method, uri, handler, middlewareList); nil != err {
				routerLog(err.Error())
				return err
			}
		default:
			err := errors.New("注册的路由必须是 IApi 或者 RouterFunc 或者 包含 RouterFunc 的结构体")
			routerLog(err.Error())
			return err
		}
	}
	return nil
}

// routerLog 记录日志
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:28 下午 2021/3/26
func routerLog(msg string) {
	if !DebugLogEnable || len(msg) == 0 {
		return
	}
	log.Print(msg)
}
