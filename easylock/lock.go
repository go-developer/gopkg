// Package easylock...
//
// Description : 包装各种姿势的锁
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-24 6:26 下午
package easylock

import "sync"

// NewLock获取普通锁实例,因为只有一把锁,flag没有意义,传空即可
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:42 下午 2021/2/24
func NewLock() EasyLock {
	return &lock{
		sc: &sync.RWMutex{},
	}
}

type lock struct {
	sc *sync.RWMutex
}

func (l *lock) Lock(optionFuncList ...OptionFunc) error {
	l.sc.Lock()
	return nil
}

func (l *lock) Unlock(optionFuncList ...OptionFunc) error {
	l.sc.Unlock()
	return nil
}

func (l *lock) RLock(optionFuncList ...OptionFunc) error {
	l.sc.RLock()
	return nil
}

func (l *lock) RUnlock(optionFuncList ...OptionFunc) error {
	l.sc.Unlock()
	return nil
}
