// Package json...
//
// Description : 动态构建json
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-10 10:26 下午
package json

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-developer/gopkg/easylock"
	"github.com/pkg/errors"
)

const (
	// PathSplit json 路径的分割符
	PathSplit = "."
)

// JSONode JSOM节点
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:33 下午 2021/3/10
type JSONode struct {
	Key              string      // json key
	Value            interface{} // 对应的值
	Child            []*JSONode  // 子节点
	IsRoot           bool        // 是否根节点
	IsHasLastBrother bool        // 在此之后是否有其他兄弟节点
	IsSlice          bool        // 是否是list
	IsIndexNode      bool        // 是否是slice的索引节点
	Sort             int         // 此属性用于　slice解析,保证最终排序是对的
}

// NewDynamicJSON 获取JSON实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:36 下午 2021/3/10
func NewDynamicJSON() *DynamicJSON {
	exp, _ := regexp.Compile(`\[(\d*?)]`)
	return &DynamicJSON{
		root: &JSONode{
			Key:    "",
			Value:  nil,
			Child:  nil,
			IsRoot: true,
		},
		nodeCnt:  0,
		lock:     easylock.NewLock(),
		sliceExp: exp,
	}
}

// DynamicJSON 动态json
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:03 下午 2021/3/10
type DynamicJSON struct {
	root     *JSONode          // 节点数
	nodeCnt  int               // 节点数量
	lock     easylock.EasyLock // 锁
	sliceExp *regexp.Regexp    // 抽取slice索引的正则
}

// SetValue 设置节点值,如果节点不存在,创建;如果已存在,更新, 多级key使用, value 必须是基础数据类型, 如果是结构体, 需要继续添加path,多级path使用.分割
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:45 下午 2021/3/10
func (dj *DynamicJSON) SetValue(path string, value interface{}) {
	pathList := strings.Split(path, PathSplit)
	searchRoot := dj.root
	parent := dj.root
	for keyIndex, key := range pathList {
		searchRoot = dj.search(searchRoot, key)
		if nil != searchRoot {
			searchRoot.Value = value // 查询到结果,更新值
			parent = searchRoot
		} else {
			var val interface{}
			if keyIndex == len(pathList)-1 {
				val = value
			}
			_ = dj.createNode(parent, key, val)
			if len(parent.Child) > 0 {
				searchRoot = parent.Child[len(parent.Child)-1]
				parent = parent.Child[len(parent.Child)-1]
			}
		}
	}
}

// String 获取字符串的格式JSON
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:16 下午 2021/3/11
func (dj *DynamicJSON) String() string {
	tplList := make([]string, 0)
	valList := make([]interface{}, 0)
	tplListResult, valListResult := dj.buildTpl(dj.root, &tplList, &valList)
	return fmt.Sprintf(strings.Join(*tplListResult, ""), *valListResult...)
}

// buildTpl 构建json模版与绑定数据
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:38 下午 2021/3/11
func (dj *DynamicJSON) buildTpl(root *JSONode, tplList *[]string, valList *[]interface{}) (*[]string, *[]interface{}) {
	if nil == root {
		*tplList = append(*tplList, "}")
		return tplList, valList
	}

	key := "\"" + root.Key + "\""
	if !root.IsIndexNode {
		if len(root.Child) > 0 {
			if root.IsRoot {
				*tplList = append(*tplList, "{")
			} else {
				if root.IsSlice {
					*tplList = append(*tplList, key+":[")
				} else {
					*tplList = append(*tplList, key+":{")
				}
			}
		} else {
			if root.IsHasLastBrother {
				*tplList = append(*tplList, key+":%v,")
			} else {
				*tplList = append(*tplList, key+":%v")
			}
			switch val := root.Value.(type) {
			case string:
				*valList = append(*valList, "\""+val+"\"")
			default:
				*valList = append(*valList, root.Value)
			}
			return tplList, valList

		}
	} else {
		if len(root.Child) == 0 {

			switch val := root.Value.(type) {
			case string:
				*valList = append(*valList, val)
				if root.IsHasLastBrother {
					*tplList = append(*tplList, "\"%v\",")
				} else {
					*tplList = append(*tplList, "\"%v\"")

				}
			default:
				if root.IsHasLastBrother {
					*tplList = append(*tplList, "%v,")
				} else {
					*tplList = append(*tplList, "%v")

				}
				*valList = append(*valList, root.Value)
			}
		} else {
			*tplList = append(*tplList, "{")
		}

	}
	for _, node := range root.Child {
		dj.buildTpl(node, tplList, valList)
	}
	if !root.IsIndexNode {
		if root.IsHasLastBrother {
			*tplList = append(*tplList, "},")
		} else {
			if root.IsSlice {
				*tplList = append(*tplList, "]")
			} else {
				*tplList = append(*tplList, "}")
			}
		}
	} else {
		if len(root.Child) > 0 {
			if root.IsHasLastBrother {
				*tplList = append(*tplList, "},")
			} else {
				*tplList = append(*tplList, "}")
			}
		}
	}

	return tplList, valList
}

// Search 搜索一个key TODO : 优化
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:19 下午 2021/3/10
func (dj *DynamicJSON) search(root *JSONode, key string) *JSONode {
	if root == nil {
		return nil
	}
	for _, node := range root.Child {
		if node == nil {
			continue
		}
		if node.Key == key {
			return node
		}
	}
	return nil
}

// createNode 创建新的节点
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:57 下午 2021/3/10
func (dj *DynamicJSON) createNode(parent *JSONode, key string, value interface{}) error {
	if nil == parent {
		return errors.New("create node error : parent id nil")
	}
	_ = dj.lock.Lock("")
	if parent.Child == nil {
		parent.Child = make([]*JSONode, 0)
	}
	if len(parent.Child) > 0 {
		// 存在子节点，设置当前子节点还有其他兄弟节点
		parent.Child[len(parent.Child)-1].IsHasLastBrother = true
	}

	newNode := &JSONode{
		Key:              key,
		Value:            value,
		Child:            make([]*JSONode, 0),
		IsRoot:           false,
		IsHasLastBrother: false,
	}
	parent.IsSlice, newNode.Sort = dj.extraSliceIndex(key)
	if parent.IsSlice {
		newNode.IsIndexNode = true
	}
	parent.Child = append(parent.Child, newNode)
	dj.nodeCnt++
	_ = dj.lock.Unlock("")
	return nil
}

// extraSliceIndex 抽取slice索引
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:37 下午 2021/3/11
func (dj *DynamicJSON) extraSliceIndex(key string) (bool, int) {
	if len(key) < 3 {
		// slice 至少是 [1] 格式
		return false, 0
	}

	if !strings.HasPrefix(key, "[") || !strings.HasSuffix(key, "]") {
		return false, 0
	}
	// 不用正则,直接字符串处理
	strByte := []byte(key)
	index, err := strconv.Atoi(string(strByte[1 : len(strByte)-1]))
	if nil != err {
		return false, 0
	}
	return true, index
}
