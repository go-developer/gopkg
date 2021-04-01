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
		lockCnt: &LockCnt{
			Write: 0,
			Read:  0,
		},
	}
	for i := 0; i < segmentCnt; i++ {
		s.lockTable[i] = NewLock()
	}
	return s, nil
}

type segment struct {
	lockTable  []EasyLock
	segmentCnt int
	lockCnt    *LockCnt
	base
}

func (s *segment) Lock(optionFuncList ...OptionFunc) error {
	defer s.AddRLockCnt(s.lockCnt)
	o := s.ParseOption(optionFuncList...)
	return s.lockTable[util.GetHashIDMod(o.flag, s.segmentCnt)].Lock()
}

func (s *segment) Unlock(optionFuncList ...OptionFunc) error {
	defer s.DecreaseLockCnt(s.lockCnt)
	o := s.ParseOption(optionFuncList...)
	return s.lockTable[util.GetHashIDMod(o.flag, s.segmentCnt)].Unlock()
}

func (s *segment) RLock(optionFuncList ...OptionFunc) error {
	defer s.AddRLockCnt(s.lockCnt)
	o := s.ParseOption(optionFuncList...)
	return s.lockTable[util.GetHashIDMod(o.flag, s.segmentCnt)].RLock()
}

func (s *segment) RUnlock(optionFuncList ...OptionFunc) error {
	defer s.DecreaseRLockCnt(s.lockCnt)
	o := s.ParseOption(optionFuncList...)
	return s.lockTable[util.GetHashIDMod(o.flag, s.segmentCnt)].RUnlock()
}

func (s *segment) GetLockCnt() *LockCnt {
	return s.lockCnt
}
