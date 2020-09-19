package stand

import (
	"fmt"
	"github.com/groupcache/lru"
	"testing"
)


// groupcache 中实现的 LRU Cache 并不是并发安全的，如果用于多个 Go 程并发的场景，需要加锁
// LRU Cache 基于 map 与 list，map 用于快速检索，list 用于实现 LRU
func Test_lRu_01(t *testing.T)  {
	cache := lru.New(2)
	cache.Add("bill", 20)
	cache.Add("dable", 19)
	v, ok := cache.Get("bill")
	if ok {
		fmt.Printf("bill's age is %v\n", v)
	}
	cache.Add("cat", "18")

	fmt.Printf("cache length is %d\n", cache.Len())
	_, ok = cache.Get("dable")
	if !ok {
		fmt.Printf("dable was evicted out\n")
	}
}
