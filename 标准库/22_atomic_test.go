package stand

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func Test_acomic_01(t *testing.T) {
	var ops uint64 = 0
	for i := 0; i < 500; i++ {
		go func() {
			atomic.AddUint64(&ops, 1)
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println(opsFinal) // 500
}
