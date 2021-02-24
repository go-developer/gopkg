// Package easylock...
//
// Description : easylock...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-24 10:55 下午
package easylock

import "github.com/pkg/errors"

// segmentError ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 1:44 下午 2021/2/24
func segmentError() error {
	return errors.New("segment需要大于0")
}
