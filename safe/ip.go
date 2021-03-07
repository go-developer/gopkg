// Package safe...
//
// Description : 安全策略之,访问黑名单
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-07 11:34 下午
package safe

import (
	"github.com/go-developer/gopkg/easymap"
	"github.com/pkg/errors"
)

// IPBlack ip黑名单
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:35 下午 2021/3/7
var IPBlack *ipBlack

// InitIPBlack 初始化ip黑名单
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:36 下午 2021/3/7
func InitIPBlack(ipList []string) error {
	// 不分片等价于只分一片
	return InitIPBlackWithSeg(1, ipList)
}

// InitIPBlackWithSeg ip黑名单分片存储
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:37 下午 2021/3/7
func InitIPBlackWithSeg(seg int, ipList []string) error {
	var err error
	IPBlack = &ipBlack{}
	if IPBlack.blackIPTable, err = easymap.NewSegment(seg, true); nil != err {
		return errors.New("初始化IP黑名单表失败,失败原因 : " + err.Error())
	}
	for _, ip := range ipList {
		// 将黑名单IP添加到内存表, easymap.EasyMap 是并发安全的
		go IPBlack.blackIPTable.Set(ip, 1)
	}
	return nil
}

type ipBlack struct {
	blackIPTable easymap.EasyMap // 黑名单IP列表
}

// Add 添加黑名单IP
//
// Author : go_developer@163.com<张德满>
//
// Date : 12:05 上午 2021/3/8
func (ib *ipBlack) Add(ip string) {
	ib.blackIPTable.Set(ip, 1)
}

// Del 删除一个黑名单IP ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 12:06 上午 2021/3/8
func (ib *ipBlack) Del(ip string) {
	ib.blackIPTable.Del(ip)
}

// IsBlack 判断ip是否存在于和名单之中
//
// Author : go_developer@163.com<张德满>
//
// Date : 12:07 上午 2021/3/8
func (ib *ipBlack) IsBlack(ip string) bool {
	return ib.blackIPTable.Exist(ip)
}
