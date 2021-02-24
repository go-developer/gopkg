// Package easymap...
//
// Description : 内置sync.Map + segment
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-24 2:08 下午
package easymap

import (
	"sync"

	"github.com/go-developer/gopkg/convert"
)

func NewSync() EasyMap {
	return &syncMap{}
}

type syncMap struct {
	data sync.Map
}

func (s *syncMap) Get(key interface{}) (interface{}, error) {
	val, exist := s.data.Load(key)
	if !exist {
		return nil, keyNotFound(key)
	}
	return val, nil
}

func (s *syncMap) GetWithReceiver(key interface{}, dest interface{}) error {
	var (
		val interface{}
		err error
	)
	if val, err = s.Get(key); nil != err {
		return err
	}
	if err = convert.ConvertAssign(dest, val); nil != err {
		return convertFail(err)
	}
	return nil
}

func (s *syncMap) GetUint(key interface{}) (uint, error) {
	var (
		result uint
		err    error
	)
	if err = s.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (s *syncMap) GetUint8(key interface{}) (uint8, error) {
	var (
		result uint8
		err    error
	)
	if err = s.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (s *syncMap) GetUint16(key interface{}) (uint16, error) {
	var (
		result uint16
		err    error
	)
	if err = s.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (s *syncMap) GetUint32(key interface{}) (uint32, error) {
	var (
		result uint32
		err    error
	)
	if err = s.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (s *syncMap) GetUint64(key interface{}) (uint64, error) {
	var (
		result uint64
		err    error
	)
	if err = s.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (s *syncMap) GetInt(key interface{}) (int, error) {
	var (
		result int
		err    error
	)
	if err = s.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (s *syncMap) GetInt8(key interface{}) (int8, error) {
	var (
		result int8
		err    error
	)
	if err = s.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (s *syncMap) GetInt16(key interface{}) (int16, error) {
	var (
		result int16
		err    error
	)
	if err = s.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (s *syncMap) GetInt32(key interface{}) (int32, error) {
	var (
		result int32
		err    error
	)
	if err = s.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (s *syncMap) GetInt64(key interface{}) (int64, error) {
	var (
		result int64
		err    error
	)
	if err = s.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (s *syncMap) GetFloat32(key interface{}) (float32, error) {
	var (
		result float32
		err    error
	)
	if err = s.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (s *syncMap) GetFloat64(key interface{}) (float64, error) {
	var (
		result float64
		err    error
	)
	if err = s.GetWithReceiver(key, &result); nil != err {
		return 0, err
	}
	return result, nil
}

func (s *syncMap) GetBool(key interface{}) (bool, error) {
	var (
		result bool
		err    error
	)
	if err = s.GetWithReceiver(key, &result); nil != err {
		return false, err
	}
	return result, nil
}

func (s *syncMap) GetString(key interface{}) (string, error) {
	var (
		result string
		err    error
	)
	if err = s.GetWithReceiver(key, &result); nil != err {
		return "", err
	}
	return result, nil
}

func (s *syncMap) Set(key interface{}, value interface{}) {
	s.data.Store(key, value)
}

func (s *syncMap) Del(key interface{}) {
	s.data.Delete(key)
}

func (s *syncMap) Exist(key interface{}) bool {
	_, exist := s.data.Load(key)
	return exist
}

func (s *syncMap) GetAll() map[interface{}]interface{} {
	result := make(map[interface{}]interface{})
	s.data.Range(func(key, value interface{}) bool {
		result[key] = value
		return true
	})
	return result
}
