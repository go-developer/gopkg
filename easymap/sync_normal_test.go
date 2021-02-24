// Package easymap...
//
// Description : easymap...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-24 5:57 下午
package easymap

import (
	"fmt"
	"testing"
)

func TestSyncNormal(t *testing.T) {
	syncMap := NewSync()
	syncMap.Set("name", "zhangdeman")
	syncMap.Set("age", 25)
	syncMap.Set("height", 180)
	fmt.Println(syncMap.GetAll())
}
