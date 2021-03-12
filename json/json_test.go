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
	tree.SetValue("slice.[0].name", "zhang")
	tree.SetValue("slice.[1].name", "de")
	tree.SetValue("slice.[2].name", "man")
	tree.SetValue("slice.[3]", "zhangdeman")
	fmt.Println(tree.String())
}
