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
	fmt.Println(tree.String())
}
