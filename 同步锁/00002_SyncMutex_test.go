package syncmutex

import (
	"fmt"
	"sync"
	"testing"
)

// Mutex（互斥锁）
//（1）使用Lock()加锁，Unlock()解锁；
//（2）对未解锁的Mutex使用Lock()会阻塞；
//（3）对未上锁的Mutex使用Unlock()会导致 panic 异常

func Test_Mutex_01(t *testing.T) {
	var intVar int
	var wg sync.WaitGroup
	var mutex sync.RWMutex
	go func() {
		defer wg.Done()
		mutex.Lock()
		intVar = 4
		mutex.Unlock()
		fmt.Printf("first goroutine, intVar=%d\n", intVar)
	}()

	go func() {
		defer wg.Done()
		mutex.Lock()
		intVar = 5
		mutex.Unlock()
		fmt.Printf("second goroutine, intVar=%d\n", intVar)
	}()

	wg.Add(2)
	wg.Wait()
	fmt.Println("end main goroutine")
}

func Test_Mutex_02(t *testing.T) {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	// Increment
	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	// Decrement
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
	}
	wg.Wait()
	fmt.Println("Arithmetic complete. Count: ",count)
}


// RWMutex（读写锁)
//（1）RWMutex是单写多读锁，该锁可以加多个读锁或者一个写锁；
//（2）读锁占用的情况下会阻止写，不会阻止读，多个 goroutine 可以同时获取读锁；
//（3）写锁会阻止其他 goroutine（无论读和写）进来，整个锁由该 goroutine 独占；
//（4）适用于读多写少的场景。
// Lock()与Unlock()
//（1）Lock() 加写锁，Unlock() 解写锁
//（2）如果在加写锁之前已经有其他的读锁和写锁，则 Lock() 会阻塞直到该锁可用，为确保该锁可用，已经阻塞的 Lock() 调用会从获得的锁中排除新的读取器，即写锁权限高于读锁，有写锁时优先进行写锁定；
//（3）在 Lock() 之前使用 Unlock() 会导致 panic 异常
// RLock() 和 RUnlock()
//（1）RLock() 加读锁，RUnlock() 解读锁；
//（2）RLock() 加读锁时，如果存在写锁，则无法加读锁；当只有读锁或者没有锁时，可以加读锁，读锁可以加多个；
//（3）RUnlock() 解读锁，RUnlock() 撤销单次 RLock() 调用，对于其他同时存在的读锁则没有效果；
//（4）在没有读锁的情况下调用 RUnlock()，会导致 panic 错误；
//（5）RUnlock() 的个数不得多于 RLock()，否则会导致 panic 错误。

func Test_RWMutex_01(t *testing.T) {
	mapNameAge := make(map[string]int) // map的读写是非原子操作
	var rwMutex sync.RWMutex
	var wg sync.WaitGroup
	go func() {
		defer wg.Done()
		rwMutex.Lock()
		mapNameAge["dablelv"] = 18
		rwMutex.Unlock()
		fmt.Println("first goroutine to write map end")
	}()

	go func() {
		defer wg.Done()
		rwMutex.RLock()
		age, _ := mapNameAge["dablelv"]
		rwMutex.RUnlock()
		fmt.Printf("second goroutine to read map end, map[dablelv]=%d\n", age)
	}()

	wg.Add(2)
	wg.Wait()
	fmt.Println("main goroutine end")
}


