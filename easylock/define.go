// Package easylock ...
//
// Description : easylock ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-04-01 4:13 下午
package easylock

type option struct {
	flag string // 锁的标识
}

// Option 设置option选项
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:15 下午 2021/4/1
type OptionFunc func(o *option)

// WithFlag 设置flag
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:17 下午 2021/4/1
func WithFlag(flag string) OptionFunc {
	return func(o *option) {
		o.flag = flag
	}
}

// LockCnt 计数lockCnt
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:41 下午 2021/4/1
type LockCnt struct {
	Write int64 `json:"write"`
	Read  int64 `json:"read"`
}

type base struct {
}

// ParseOption 解析option
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:24 下午 2021/4/1
func (b *base) ParseOption(optionFuncList ...OptionFunc) *option {
	o := &option{}
	for _, f := range optionFuncList {
		f(o)
	}
	return o
}

// AddLockCnt 锁数量加一
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:44 下午 2021/4/1
func (b *base) AddLockCnt(lockCnt *LockCnt) {
	lockCnt.Write++
}

// DecreaseLockCnt 锁数量
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:45 下午 2021/4/1
func (b *base) DecreaseLockCnt(lockCnt *LockCnt) {
	lockCnt.Write--
}

// AddRLockCnt 锁数量加一
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:44 下午 2021/4/1
func (b *base) AddRLockCnt(lockCnt *LockCnt) {
	lockCnt.Read++
}

// DecreaseRLockCnt 锁数量
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:45 下午 2021/4/1
func (b *base) DecreaseRLockCnt(lockCnt *LockCnt) {
	lockCnt.Read--
}
