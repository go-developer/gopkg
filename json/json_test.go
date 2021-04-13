// Package json...
//
// Description : json...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-10 11:44 下午
package json

import (
	"fmt"
	"testing"
)

// TestJSON ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:58 下午 2021/3/14
func TestJSON(t *testing.T) {
	tree := NewDynamicJSON()
	fmt.Println(tree.extraSliceIndex("[200]"))
	tree.SetValue("extra.height.value", 180)
	tree.SetValue("extra.height.unit.use", "cm")
	tree.SetValue("extra.height.unit.open", "1")
	tree.SetValue("name", "zhangdeman")
	tree.SetValue("good.name", "good")
	tree.SetValue("work", "111")
	tree.SetValue("good.price", "180")
	tree.SetValue("good.unit", "$")
	tree.SetValue("slice.[0].name", "zhang")
	tree.SetValue("slice.[1].name", "de")
	tree.SetValue("slice.[2].name", "man")
	tree.SetValue("slice.[3]", "zhangdeman")
	fmt.Println(tree.String())
	tree = NewDynamicJSON()
	tree.SetValue("[0]", "zhang")
	tree.SetValue("[1]", "de")
	tree.SetValue("[2]", "man")
	fmt.Println(tree.String())
}

// TestType 判断数据类型断言
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:59 下午 2021/3/14
func TestType(t *testing.T) {

}

// TestSelect 测试动态选择字段
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:47 下午 2021/4/13
func TestSelect(t *testing.T) {
	source := map[string]interface{}{
		"name": "zhangdeman",
		"extra": map[string]interface{}{
			"age":    18,
			"height": 180,
			"slice":  []int{1, 2, 3},
		},
		"slice": []int{1, 2, 3},
	}
	pathList := []string{"name", "extra.age", "slice"}
	r, e := NewParseJSONTree(source).Parse(pathList)
	fmt.Println(r.String(), e)
}
