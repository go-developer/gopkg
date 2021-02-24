// Package easymap...
//
// Description : 内置sync.Map + segment
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-24 2:08 下午
package easymap

import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

// NewSegmentSync 获取SegmentSync实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:02 下午 2021/2/24
func NewSegmentSync(segment int) (EasyMap, error) {
	if segment <= 0 {
		return nil, segmentError()
	}
	ss := &segmentSync{
		segment: segment,
	}
	ss.dataTable = make([]EasyMap, segment)
	for i := 0; i < segment; i++ {
		ss.dataTable[i] = NewSync()
	}
	return ss, nil
}

type segmentSync struct {
	dataTable []EasyMap
	segment   int
}

func (s *segmentSync) Get(key interface{}) (interface{}, error) {
	return s.dataTable[s.getSegment(key)].Get(key)
}

func (s *segmentSync) GetWithReceiver(key interface{}, dest interface{}) error {
	return s.dataTable[s.getSegment(key)].GetWithReceiver(key, dest)
}

func (s *segmentSync) GetUint(key interface{}) (uint, error) {
	return s.dataTable[s.getSegment(key)].GetUint(key)
}

func (s *segmentSync) GetUint8(key interface{}) (uint8, error) {
	return s.dataTable[s.getSegment(key)].GetUint8(key)
}

func (s *segmentSync) GetUint16(key interface{}) (uint16, error) {
	return s.dataTable[s.getSegment(key)].GetUint16(key)
}

func (s *segmentSync) GetUint32(key interface{}) (uint32, error) {
	return s.dataTable[s.getSegment(key)].GetUint32(key)
}

func (s *segmentSync) GetUint64(key interface{}) (uint64, error) {
	return s.dataTable[s.getSegment(key)].GetUint64(key)
}

func (s *segmentSync) GetInt(key interface{}) (int, error) {
	return s.dataTable[s.getSegment(key)].GetInt(key)
}

func (s *segmentSync) GetInt8(key interface{}) (int8, error) {
	return s.dataTable[s.getSegment(key)].GetInt8(key)
}

func (s *segmentSync) GetInt16(key interface{}) (int16, error) {
	return s.dataTable[s.getSegment(key)].GetInt16(key)
}

func (s *segmentSync) GetInt32(key interface{}) (int32, error) {
	return s.dataTable[s.getSegment(key)].GetInt32(key)
}

func (s *segmentSync) GetInt64(key interface{}) (int64, error) {
	return s.dataTable[s.getSegment(key)].GetInt64(key)
}

func (s *segmentSync) GetFloat32(key interface{}) (float32, error) {
	return s.dataTable[s.getSegment(key)].GetFloat32(key)
}

func (s *segmentSync) GetFloat64(key interface{}) (float64, error) {
	return s.dataTable[s.getSegment(key)].GetFloat64(key)
}

func (s *segmentSync) GetBool(key interface{}) (bool, error) {
	return s.dataTable[s.getSegment(key)].GetBool(key)
}

func (s *segmentSync) GetString(key interface{}) (string, error) {
	return s.dataTable[s.getSegment(key)].GetString(key)
}

func (s *segmentSync) Set(key interface{}, value interface{}) {
	s.dataTable[s.getSegment(key)].Set(key, value)
}

func (s *segmentSync) Del(key interface{}) {
	s.dataTable[s.getSegment(key)].Del(key)
}

func (s *segmentSync) Exist(key interface{}) bool {
	return s.dataTable[s.getSegment(key)].Exist(key)
}

func (s *segmentSync) GetAll() map[interface{}]interface{} {
	result := make(map[interface{}]interface{})
	for i := 0; i < s.segment; i++ {
		for k, v := range s.dataTable[i].GetAll() {
			result[k] = v
		}
	}
	return result
}

func (s *segmentSync) getSegment(key interface{}) int {
	return int(murmur3.Sum64([]byte(fmt.Sprintf("%v", key))) % uint64(s.segment))
}
