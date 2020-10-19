package stand

import (
	"fmt"
	"index/suffixarray"
	"testing"
)

func Test_suffixarray(t *testing.T)  {
	index := suffixarray.New([]byte("bananaanaanaana"))
	offsets := index.Lookup([]byte("ana"), -1) // 匹配所有开头索引
	fmt.Println(offsets)  // [12 9 6 3 1]

}
