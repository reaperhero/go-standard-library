package syncmutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//（1）计数器，统计函数执行次数；
//（2）线程安全，保障在多G情况下，函数仍然只执行一次，比如锁。

var once sync.Once
var onceBody = func() {
	fmt.Println("Only once")
}

func Test_Once01(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			once.Do(onceBody)
			fmt.Println("i=", i)
		}(i)
	}
	time.Sleep(time.Second) //睡眠1s用于执行go程，注意睡眠时间不能太短
}
