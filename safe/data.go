// Package safe ...
//
// Description : 按需返回对外暴露的字段
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-10 6:38 下午
package safe

import (
	"strings"

	"github.com/buger/jsonparser"
)

// Filter 按需输出数据
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:40 下午 2021/3/10
func Filter(source []byte, filter []string) ([]byte, error) {
	var (
		bt     []byte
		setErr error
	)
	for _, item := range filter {
		fieldList := strings.Split(item, ",")
		val, _, _, err := jsonparser.Get(source, fieldList...)
		if nil != err {
			return nil, err
		}
		if bt, setErr = jsonparser.Set(bt, val, fieldList...); nil != setErr {
			return nil, setErr
		}
	}
	return bt, nil
}
