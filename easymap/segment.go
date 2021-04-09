// Package easymap...
//
// Description : 分段存储的map，并发行更好,分段数量为 1， 将退化成普通的
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-23 10:47 下午
package easymap

import (
	"github.com/go-developer/gopkg/util"
)

// NewSegment 获取分段map实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:05 下午 2021/2/24
func NewSegment(segmentCnt int, withLock bool) (EasyMap, error) {
	if segmentCnt <= 0 {
		return nil, segmentError()
	}
	em := &segment{
		segment: segmentCnt,
	}
	em.dataTable = make([]EasyMap, segmentCnt)
	for i := 0; i < segmentCnt; i++ {
		em.dataTable[i] = NewNormal(withLock)
	}
	return em, nil
}

type segment struct {
	dataTable []EasyMap
	segment   int
}

func (s *segment) Get(key interface{}) (interface{}, error) {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	return s.dataTable[segmentIndex].Get(key)
}

func (s *segment) GetWithReceiver(key interface{}, dest interface{}) error {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	return s.dataTable[segmentIndex].GetWithReceiver(key, dest)
}

func (s *segment) GetUint(key interface{}) (uint, error) {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	return s.dataTable[segmentIndex].GetUint(key)
}

func (s *segment) GetUint8(key interface{}) (uint8, error) {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	return s.dataTable[segmentIndex].GetUint8(key)
}

func (s *segment) GetUint16(key interface{}) (uint16, error) {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	return s.dataTable[segmentIndex].GetUint16(key)
}

func (s *segment) GetUint32(key interface{}) (uint32, error) {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	return s.dataTable[segmentIndex].GetUint32(key)
}

func (s *segment) GetUint64(key interface{}) (uint64, error) {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	return s.dataTable[segmentIndex].GetUint64(key)
}

func (s *segment) GetInt(key interface{}) (int, error) {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	return s.dataTable[segmentIndex].GetInt(key)
}

func (s *segment) GetInt8(key interface{}) (int8, error) {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	return s.dataTable[segmentIndex].GetInt8(key)
}

func (s *segment) GetInt16(key interface{}) (int16, error) {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	return s.dataTable[segmentIndex].GetInt16(key)
}

func (s *segment) GetInt32(key interface{}) (int32, error) {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	return s.dataTable[segmentIndex].GetInt32(key)
}

func (s *segment) GetInt64(key interface{}) (int64, error) {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	return s.dataTable[segmentIndex].GetInt64(key)
}

func (s *segment) GetFloat32(key interface{}) (float32, error) {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	return s.dataTable[segmentIndex].GetFloat32(key)
}

func (s *segment) GetFloat64(key interface{}) (float64, error) {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	return s.dataTable[segmentIndex].GetFloat64(key)
}

func (s *segment) GetBool(key interface{}) (bool, error) {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	return s.dataTable[segmentIndex].GetBool(key)
}

func (s *segment) GetString(key interface{}) (string, error) {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	return s.dataTable[segmentIndex].GetString(key)
}

func (s *segment) Set(key interface{}, value interface{}) {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	s.dataTable[segmentIndex].Set(key, value)
}

func (s *segment) Del(key interface{}) {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	s.dataTable[segmentIndex].Del(key)
}

func (s *segment) Exist(key interface{}) bool {
	segmentIndex := util.GetHashIDMod(key, s.segment)
	return s.dataTable[segmentIndex].Exist(key)
}

func (s *segment) GetAll() map[interface{}]interface{} {
	result := make(map[interface{}]interface{})
	for i := 0; i < s.segment; i++ {
		for k, v := range s.dataTable[i].GetAll() {
			result[k] = v
		}
	}
	return result
}
