package syncmutex

import (
	"fmt"
	"sync"
	"testing"
)

func Test_syncmutex_01(t *testing.T)  {
	var lock sync.Mutex
	var value int
	go func() {
		lock.Lock()
		value++
		lock.Unlock()
	}()
	lock.Lock()
	fmt.Println(value)
	lock.Unlock()
}