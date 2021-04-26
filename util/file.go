// Package util...
//
// Description : 文件相关工具
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-04-26 6:00 下午
package util

import (
	"os"
	"strings"
)

// GetProjectPath 获取项目路径(可执行文件所在目录)
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:32 下午 2021/4/26
func GetProjectPath() (string, error) {
	rootPath, err := os.Getwd()
	if nil != err {
		return "", err
	}
	pathArr := strings.Split(rootPath, "/")
	if len(pathArr) > 0 {
		if pathArr[len(pathArr)-1] == "test" {
			rootPath = strings.Join(pathArr[0:len(pathArr)-1], "/")
		}
	}
	return rootPath, nil
}
