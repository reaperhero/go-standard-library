package syncmutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Cond 实现了一个条件变量，在 Locker 的基础上增加的一个消息通知的功能，保存了一个通知列表，用来唤醒一个或所有因等待条件变量而阻塞的 Go 程，以此来实现多个 Go 程间的同步

// 在调用 Signal，Broadcast 之前，应确保目标 Go 程进入 Wait 阻塞状态

func Test_Cond_01(t *testing.T) {
	var locker = new(sync.Mutex)
	var cond = sync.NewCond(locker)

	for i := 0; i < 10; i++ {
		go func(x int) {
			cond.L.Lock()         //获取锁
			defer cond.L.Unlock() //释放锁
			cond.Wait()           //等待通知，阻塞当前goroutine
			fmt.Println(x)
		}(i)
	}
	time.Sleep(time.Second * 1) // 睡眠1秒，使所有goroutine进入 Wait 阻塞状态
	fmt.Println("Signal...")
	cond.Signal() // 1秒后下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal() // 1秒后下发下一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	cond.Broadcast() // 1秒后下发广播给所有等待的goroutine
	fmt.Println("Broadcast...")
	time.Sleep(time.Second * 1) // 睡眠1秒，等待所有goroutine执行完毕
}
