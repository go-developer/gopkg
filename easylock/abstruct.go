// Package easylock...
//
// Description : easylock...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-24 6:29 下午
package easylock

// EasyLock ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:18 下午 2021/4/1
type EasyLock interface {
	// Lock ...
	Lock(optionFuncList ...OptionFunc) error
	// Unlock ...
	Unlock(optionFuncList ...OptionFunc) error
	// RLock ...
	RLock(optionFuncList ...OptionFunc) error
	// RUnlock ...
	RUnlock(optionFuncList ...OptionFunc) error
}
