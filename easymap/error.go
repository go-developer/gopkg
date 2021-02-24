// Package easymap...
//
// Description : easymap...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-23 10:15 下午
package easymap

import "github.com/pkg/errors"

// keyNotFound key 不存在
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:17 下午 2021/2/23
func keyNotFound(key interface{}) error {
	return errors.Errorf("%v 未找到", key)
}

// convertFail 数据类型妆换失败
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:28 下午 2021/2/23
func convertFail(err error) error {
	return errors.Wrapf(err, "数据类型转换失败")
}

// segmentError ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 1:44 下午 2021/2/24
func segmentError() error {
	return errors.New("segment需要大于0")
}
