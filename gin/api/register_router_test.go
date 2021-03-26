// Package api ...
//
// Description : 路由注册单元测试
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-26 3:49 下午
package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

// TestRegisterRouter ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 3:50 下午 2021/3/26
func TestRegisterRouter(t *testing.T) {
	r := gin.Default()
	err := RegisterRouter(r, demoApiFunc(), &demoApi{}, &otherApi{}, nil)
	assert.Nil(t, err, "路由注册异常 : %v", err)
}

func demoApiFunc() RouterFunc {
	return func() (method string, uri string, handlerFunc gin.HandlerFunc, middlewareList []gin.HandlerFunc) {
		return http.MethodGet, "/api/func/test", func(context *gin.Context) {

		}, nil
	}
}

type demoApi struct {
}

func (d demoApi) GetMethod() string {
	return http.MethodGet
}

func (d demoApi) GetURI() string {
	return "/api/struct/test"
}

func (d demoApi) GetMiddleWareList() []gin.HandlerFunc {
	return nil
}

func (d demoApi) GetHandler() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

type otherApi struct {
}

func (oa *otherApi) DemoApiFunc() RouterFunc {
	return func() (method string, uri string, handlerFunc gin.HandlerFunc, middlewareList []gin.HandlerFunc) {
		return http.MethodGet, "/api/other/test", func(context *gin.Context) {

		}, nil
	}
}

func (oa *otherApi) Lala() {

}

func (oa *otherApi) SelfApi() (method string, uri string, handlerFunc gin.HandlerFunc, middlewareList []gin.HandlerFunc) {
	return http.MethodGet, "/api/other/self/test", func(context *gin.Context) {

	}, nil
}
