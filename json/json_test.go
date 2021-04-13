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
	tree.SetValue("extra.height.value", 180, false)
	tree.SetValue("extra.height.unit.use", "cm", false)
	tree.SetValue("extra.height.unit.open", "1", false)
	tree.SetValue("name", "zhangdeman", false)
	tree.SetValue("good.name", "good", false)
	tree.SetValue("work", "111", false)
	tree.SetValue("good.price", "180", false)
	tree.SetValue("good.unit", "$", false)
	tree.SetValue("slice.[0].name", "zhang", false)
	tree.SetValue("slice.[1].name", "de", false)
	tree.SetValue("slice.[2].name", "man", false)
	tree.SetValue("slice.[3]", "zhangdeman", false)
	fmt.Println(tree.String())
	tree = NewDynamicJSON()
	tree.SetValue("[0]", "zhang", false)
	tree.SetValue("[1]", "de", false)
	tree.SetValue("[2]", "man", false)
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
		"map":   map[string]interface{}{"a": 1, "b": 2, "c": 4},
		"table": []map[string]interface{}{
			{"name": "alex", "age": 18, "number": 1},
			{"name": "bob", "age": 28, "number": 2},
		},
	}
	pathList := []string{"name", "extra.age", "slice", "map", "table.[].name|number|test"}
	r, e := NewParseJSONTree(source).Parse(pathList)
	fmt.Println(r.String(), e)
}
