// Package easylock...
//
// Description : 包装各种姿势的锁
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-24 6:26 下午
package easylock

import "sync"

func NewLock() EasyLock {
	return &lock{
		sc: &sync.RWMutex{},
	}
}

type lock struct {
	sc *sync.RWMutex
}

func (l *lock) Lock() error {
	l.sc.Lock()
	return nil
}

func (l *lock) Unlock() error {
	l.sc.Unlock()
	return nil
}

func (l *lock) RLock() error {
	l.sc.RLock()
	return nil
}

func (l *lock) RUnlock() error {
	l.sc.Unlock()
	return nil
}
