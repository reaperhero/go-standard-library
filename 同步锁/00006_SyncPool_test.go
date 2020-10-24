package syncmutex

import (
	"fmt"
	"sync"
	"testing"
)


//实例化sync.Pool时，给它一个新元素，该元素应该是线程安全的。
//当你从Get获得一个实例时，不要假设你接收到的对象状态。
//当你从池中取得实例时，请务必不要忘记调用Put。否则池的优越性就体现不出来了。这通常用defer 来执行延迟操作。
//池中的元素必须大致上是均匀的。


// 当多个 goroutine 都需要创建同⼀个对象的时候，如果 goroutine 数过多，导致对象的创建数⽬剧增，进⽽导致 GC 压⼒增大。形成 “并发⼤－占⽤内存⼤－GC 缓慢－处理并发能⼒降低－并发更⼤”这样的恶性循环。
// 在这个时候，需要有⼀个对象池，每个 goroutine 不再⾃⼰单独创建对象，⽽是从对象池中获取出⼀个对象（如果池中已经有的话）。关键思想就是对象的复用，避免重复创建、销毁

// 一、

func Test_SyncPool_01(t *testing.T) {

	type Person struct {
		Name string
	}

	var pool = &sync.Pool{
		New: func() interface{} {
			return new(Person)
		},
	}

	p := pool.Get().(*Person)
	fmt.Println(p) // nil

	p.Name = "first"

	pool.Put(p)

	fmt.Println(pool.Get().(*Person)) // &{first}
	fmt.Println(pool.Get().(*Person)) // nil
}

// 二、

func Test_SyncPool_02(t *testing.T) {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated++
			mem := make([]byte, 1024) // 增加1K
			return &mem
		},
	}
	// 将池扩充到4KB
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte) // 断言成byte数组
			defer calcPool.Put(mem)
		}()
	}
	wg.Wait()
	fmt.Printf("%d",numCalcsCreated) // 12
}
