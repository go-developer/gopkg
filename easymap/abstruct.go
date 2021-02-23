// Package easymap...
//
// Description : easymap...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-23 9:55 下午
package easymap

// EasyMap 约束各种数据接口的实现
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:56 下午 2021/2/23
type EasyMap interface {
	Get(key interface{}) (interface{}, error)
	GetWithReceiver(key interface{}, dest interface{}) error
	GetUint(key interface{}) (uint, error)
	GetUint8(key interface{}) (uint8, error)
	GetUint16(key interface{}) (uint16, error)
	GetUint32(key interface{}) (uint32, error)
	GetUint64(key interface{}) (uint64, error)
	GetInt(key interface{}) (int, error)
	GetInt8(key interface{}) (int8, error)
	GetInt16(key interface{}) (int16, error)
	GetInt32(key interface{}) (int32, error)
	GetInt64(key interface{}) (int64, error)
	GetFloat32(key interface{}) (float32, error)
	GetFloat64(key interface{}) (float64, error)
	GetBool(key interface{}) (bool, error)
	GetString(key interface{}) (string, error)
	Set(key interface{}, value interface{})
	Del(key interface{})
	Exist(key interface{}) bool
	GetAll() map[interface{}]interface{}
}
