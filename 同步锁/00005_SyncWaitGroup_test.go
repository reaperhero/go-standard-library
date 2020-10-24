package syncmutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.WaitGroup 用于阻塞等待一组 Go 程的结束。主 Go 程调用 Add() 来设置等待的 Go 程数，然后该组中的每个 Go 程都需要在运行结束时调用 Done()， 递减 WaitGroup 的 Go 程计数器 counter。当 counter 变为 0 时，主 Go 程被唤醒继续执行。

func Test_Wg_01(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("entry foo1")
		time.Sleep(2 * time.Second)
		fmt.Println("exit foo1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("entry foo2")
		time.Sleep(4 * time.Second)
		fmt.Println("exit foo2")
	}()

	wg.Wait()

	fmt.Println("exit main")
}



func Test_Wg_02(t *testing.T) {
	hello := func(wg *sync.WaitGroup,id int) {
		defer wg.Done()
		fmt.Println(id)
	}
	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters) // 必须保证Add在Wait之前就准备好了
	for i := 0; i < numGreeters; i++ {
		hello(&wg,i)
	}
	wg.Wait()
}
