package syncmutex

import (
	"fmt"
	"sync"
	"testing"
)

// Golang 中的 map 在并发情况下，只读是线程安全的，并发读写线程不安全。为了解决这个问题，Golang 提供了语言层级的并发读写安全的 sync.Map
// sync.Map 开箱即用，无需像 map 调用 make() 进行显示初始化。此外，sync.Map 的 key 和 value 类型为空接口 interface{}，表示可存储任意类型的数据

func Test_SyncMap_01(t *testing.T) {
	var m sync.Map

	//写

	m.Store("dablelv", "27")
	m.Store("cat", "28")

	//读
	v, ok := m.Load("dablelv")
	fmt.Printf("Load: v, ok = %v, %v\n", v, ok)

	//删除
	m.Delete("dablelv")

	//读或写
	v, ok = m.LoadOrStore("dablelv", "18")
	fmt.Printf("LoadOrStore: v, ok = %v, %v\n", v, ok)

	//遍历
	//操作函数
	f := func(key, value interface{}) bool {
		fmt.Printf("Range: k, v = %v, %v\n", key, value)
		return true
	}
	m.Range(f)
}
