// Package util...
//
// Description : 文件相关工具
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-04-26 6:00 下午
package util

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
	yml "gopkg.in/yaml.v2"
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

// ReadYmlConfig 读取yml配置问价,并解析到指定的结构体中
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:35 下午 2021/4/26
func ReadYmlConfig(filePath string, result interface{}) error {
	if nil == result {
		return errors.New("接收读取结果的数据指针为NIL")
	}
	var (
		fileContent []byte
		err         error
	)
	if fileContent, err = ReadFileContent(filePath); nil != err {
		return err
	}
	return yml.Unmarshal(fileContent, result)
}

// ReadFileContent 读取文件内容
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:37 下午 2021/4/26
func ReadFileContent(filePath string) ([]byte, error) {
	if exist, isFile := IsFileExist(filePath); !exist || !isFile {
		//文件不存在或者是一个目录
		return nil, errors.New(filePath + " 文件不存在或者是一个目录!")
	}
	//打开文件
	var (
		f   *os.File
		err error
	)
	if f, err = os.Open(filePath); nil != err {
		return nil, err
	}

	return ioutil.ReadAll(f)
}

// IsFileExist 判断文件是否存在
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:37 下午 2021/4/26
func IsFileExist(filePath string) (bool, bool) {
	f, err := os.Stat(filePath)
	return nil == err || os.IsExist(err), (nil == err || os.IsExist(err)) && !f.IsDir()
}
