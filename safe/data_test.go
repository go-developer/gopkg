// Package safe...
//
// Description : safe...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-03-10 10:07 下午
package safe

import (
	"fmt"
	"testing"

	"github.com/tidwall/gjson"
)

func TestFilter(t *testing.T) {
	source := []byte(`
{
"name":"zhangdeman",
"age":18,
"extra":{"height":180, "company":"com"}
}
`)
	result, err := Filter(source, []string{"name", "extra.company"})
	if nil != err {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(result))
	}
}

func TestGJSON(t *testing.T) {
	source := `
{
"name":"zhangdeman",
"age":18,
"extra":{"height":180, "company":"com"}
}
`
	fmt.Println(gjson.Get(source, "extra.height").Raw)
}
