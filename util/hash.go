// Package util...
//
// Description : util...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-24 11:04 下午
package util

import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

// GetHashID ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:04 下午 2021/2/24
func GetHashID(key interface{}) uint64 {
	return murmur3.Sum64([]byte(fmt.Sprintf("%v", key)))
}

// GetHashIDMod 获取hashID并取模
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:07 下午 2021/2/24
func GetHashIDMod(key interface{}, shard int) int {
	return int(GetHashID(key) % uint64(shard))
}
