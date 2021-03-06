package command

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// 检查gorouting消耗的内存
func Test_mem_01(t *testing.T) {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c } // a goroutine which will never exit, because it's waiting for channel c all the time
	const numGoroutines = 1e5
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1024) // 2.063kb
}