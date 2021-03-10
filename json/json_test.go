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
	tree.SetValue("extra.height.value", 180)
	tree.SetValue("extra.height.unit", "cm")
	fmt.Println(tree.root)
}
