// Package easymap ...
//
// Description : 普通的的map,增加锁支持
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-23 10:03 下午
package easymap

import (
	"sync"

	"github.com/go-developer/gopkg/convert"
)

// NewNormal 获取map实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:07 下午 2021/2/23
func NewNormal(withLock bool) EasyMap {
	em := &normal{
		data:     make(map[interface{}]interface{}),
		withLock: withLock,
	}
	if withLock {
		em.lock = &sync.RWMutex{}
	}
	return em
}

// normal 普通map,内部可以加锁
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:03 下午 2021/2/23
type normal struct {
	data     map[interface{}]interface{}
	withLock bool
	lock     *sync.RWMutex
}

func (n *normal) Get(key interface{}) (interface{}, error) {
	if !n.Exist(key) {
		return nil, keyNotFound(key)
	}
	n.RLock()
	defer n.RUnlock()
	return n.data[key], nil
}

func (n *normal) GetWithReceiver(key interface{}, dest interface{}) error {
	var (
		val interface{}
		err error
	)
	if val, err = n.Get(key); nil != err {
		return err
	}
	return convertFail(convert.ConvertAssign(dest, val))
}

func (n *normal) GetUint(key interface{}) (uint, error) {
	var (
		result uint
		err    error
	)
	if err = n.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (n *normal) GetUint8(key interface{}) (uint8, error) {
	var (
		result uint8
		err    error
	)
	if err = n.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (n *normal) GetUint16(key interface{}) (uint16, error) {
	var (
		result uint16
		err    error
	)
	if err = n.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (n *normal) GetUint32(key interface{}) (uint32, error) {
	var (
		result uint32
		err    error
	)
	if err = n.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (n *normal) GetUint64(key interface{}) (uint64, error) {
	var (
		result uint64
		err    error
	)
	if err = n.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (n *normal) GetInt(key interface{}) (int, error) {
	var (
		result int
		err    error
	)
	if err = n.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (n *normal) GetInt8(key interface{}) (int8, error) {
	var (
		result int8
		err    error
	)
	if err = n.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (n *normal) GetInt16(key interface{}) (int16, error) {
	var (
		result int16
		err    error
	)
	if err = n.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (n *normal) GetInt32(key interface{}) (int32, error) {
	var (
		result int32
		err    error
	)
	if err = n.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (n *normal) GetInt64(key interface{}) (int64, error) {
	var (
		result int64
		err    error
	)
	if err = n.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (n *normal) GetFloat32(key interface{}) (float32, error) {
	var (
		result float32
		err    error
	)
	if err = n.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (n *normal) GetFloat64(key interface{}) (float64, error) {
	var (
		result float64
		err    error
	)
	if err = n.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (n *normal) GetBool(key interface{}) (bool, error) {
	var (
		result bool
		err    error
	)
	if err = n.GetWithReceiver(key, &result); nil != err {
		return false, err
	}
	return result, nil
}

func (n *normal) GetString(key interface{}) (string, error) {
	var (
		result string
		err    error
	)
	if err = n.GetWithReceiver(key, &result); nil != err {
		return "", err
	}
	return result, nil
}

func (n *normal) Set(key interface{}, value interface{}) {
	n.Lock()
	defer n.Unlock()
	n.data[key] = value
}

func (n *normal) Del(key interface{}) {
	n.Lock()
	defer n.Unlock()
	delete(n.data, key)
}

func (n *normal) Exist(key interface{}) bool {
	n.RLock()
	defer n.RUnlock()
	_, exist := n.data[key]
	return exist
}

// GetAll 读取全部数据使用的是原始数据深拷贝,避免获取到全部数据之后,直接读取导致并发读写
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:36 下午 2021/2/23
func (n *normal) GetAll() map[interface{}]interface{} {
	result := make(map[interface{}]interface{})
	n.lock.RLock()
	defer n.lock.RUnlock()
	for key, val := range n.data {
		result[key] = val
	}
	return result
}

func (n *normal) RLock() {
	if n.withLock {
		n.lock.RLock()
	}
}

func (n *normal) RUnlock() {
	if n.withLock {
		n.lock.RUnlock()
	}
}

func (n *normal) Lock() {
	if n.withLock {
		n.lock.Lock()
	}
}

func (n *normal) Unlock() {
	if n.withLock {
		n.lock.Unlock()
	}
}
