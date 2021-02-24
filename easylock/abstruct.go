// Package easylock...
//
// Description : easylock...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-24 6:29 下午
package easylock

type EasyLock interface {
	Lock() error
	Unlock() error
	RLock() error
	RUnlock() error
}
