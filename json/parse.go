// Package json ...
//
// Description : 将复杂数据结构转化为 JSONNode 树
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-14 10:40 下午
package json

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/tidwall/gjson"

	"github.com/pkg/errors"

	"github.com/go-developer/gopkg/util"
)

// NewParseJSONTree 获取解析的实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:43 下午 2021/3/14
func NewParseJSONTree(data interface{}) *ParseJSONTree {
	return &ParseJSONTree{data: data}
}

// ParseJSONTree 解析json树
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:41 下午 2021/3/14
type ParseJSONTree struct {
	data interface{}
}

// Parse 解析数据
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:44 下午 2021/3/14
func (pjt *ParseJSONTree) Parse(pathList []string) (*DynamicJSON, error) {
	if !pjt.isLegalData() {
		return nil, errors.New("非法的数据,无法转换成json")
	}
	result := NewDynamicJSON()
	byteData, _ := json.Marshal(pjt.data)
	source := string(byteData)
	for _, path := range pathList {
		// 为数组的处理
		pathArr := strings.Split(path, "[]")
		for idx, item := range pathArr {
			val := gjson.Get(source, item)
			isComplexType := val.IsArray()
			if len(pathArr)-1 == idx {
				// 当前是最后一项,说明不是数组
				if isComplexType {
					fmt.Println("这是一个数组")
				}
				result.SetValue(item, val.Raw, isComplexType)
				continue
			}
		}
	}
	return result, nil
}

// isLegalData 判断是否能转换成json结构, 只有slice/map/struct/能转换成slice或map的[]byte是合法的
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:46 下午 2021/3/14
func (pjt *ParseJSONTree) isLegalData() bool {
	val := reflect.ValueOf(pjt.data)

	switch val.Kind() {
	case reflect.Slice:
		// slice 情况下,对字节数组进行特殊判断
		var (
			byteData []byte
			ok       bool
			err      error
		)
		if byteData, ok = pjt.data.([]byte); ok {
			// 字节数组转map或者slice
			if err = json.Unmarshal(byteData, &pjt.data); nil != err {
				return false
			}
			return true
		}
		return true
	case reflect.Map:
		return true
	case reflect.Struct:
		// 结构体转为字符串处理
		fallthrough
	case reflect.Ptr:
		// 指针
		var err error
		if pjt.data, err = util.StructToMap(pjt.data); nil != err {
			return false
		}
		return true
	default:
		return false
	}
}
