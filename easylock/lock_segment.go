// Package easylock...
//
// Description : 分段的锁
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-24 10:44 下午
package easylock

import "github.com/go-developer/gopkg/util"

// NewSegment 获取分段锁
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:20 下午 2021/2/24
func NewSegment(segmentCnt int) (EasyLock, error) {
	if segmentCnt <= 0 {
		return nil, segmentError()
	}
	s := &segment{
		lockTable:  make([]EasyLock, segmentCnt),
		segmentCnt: segmentCnt,
	}
	for i := 0; i < segmentCnt; i++ {
		s.lockTable[i] = NewLock()
	}
	return s, nil
}

type segment struct {
	lockTable  []EasyLock
	segmentCnt int
}

func (s *segment) Lock(flag string) error {
	return s.lockTable[util.GetHashIDMod(flag, s.segmentCnt)].Lock(flag)
}

func (s *segment) Unlock(flag string) error {
	return s.lockTable[util.GetHashIDMod(flag, s.segmentCnt)].Unlock(flag)
}

func (s *segment) RLock(flag string) error {
	return s.lockTable[util.GetHashIDMod(flag, s.segmentCnt)].RLock(flag)
}

func (s *segment) RUnlock(flag string) error {
	return s.lockTable[util.GetHashIDMod(flag, s.segmentCnt)].RUnlock(flag)
}
