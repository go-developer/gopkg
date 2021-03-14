// Package util...
//
// Description : util ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-14 11:11 下午
package util

import "encoding/json"

// StructToMap 结构体转为map
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:12 下午 2021/3/14
func StructToMap(data interface{}) (map[string]interface{}, error) {
	var (
		byteData []byte
		err      error
		result   map[string]interface{}
	)
	if byteData, err = json.Marshal(data); nil != err {
		return nil, err
	}
	if err = json.Unmarshal(byteData, &result); nil != err {
		return nil, err
	}
	return result, nil
}
