package syncmutex

import (
	"fmt"
	"sync"
	"testing"
)

// 当多个 goroutine 都需要创建同⼀个对象的时候，如果 goroutine 数过多，导致对象的创建数⽬剧增，进⽽导致 GC 压⼒增大。形成 “并发⼤－占⽤内存⼤－GC 缓慢－处理并发能⼒降低－并发更⼤”这样的恶性循环。
// 在这个时候，需要有⼀个对象池，每个 goroutine 不再⾃⼰单独创建对象，⽽是从对象池中获取出⼀个对象（如果池中已经有的话）。关键思想就是对象的复用，避免重复创建、销毁

// 一、

type Person struct {
	Name string
}

var pool = &sync.Pool{
	New: func() interface{} {
		return new(Person)
	},
}

func Test_SyncPool_01(t *testing.T) {
	p := pool.Get().(*Person)
	fmt.Println(p) // nil

	p.Name = "first"

	pool.Put(p)

	fmt.Println(pool.Get().(*Person)) // &{first}
	fmt.Println(pool.Get().(*Person)) // nil
}
