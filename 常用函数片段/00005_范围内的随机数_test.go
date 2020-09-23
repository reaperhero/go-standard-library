package command

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var randomMutex sync.Mutex

func Test_Rand_fund_01(t *testing.T) {
	start := 1
	end := 100
	randomMutex.Lock()    // 随机数互斥锁
	<-time.After(1 * time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := start + r.Intn(end-start+1)
	randomMutex.Unlock()
	fmt.Println(n)
}
