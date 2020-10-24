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

// 固定长度为2的队列，并且我们要将10个元素放入队列中。 我们希望一有空间就能放入，所以在队列中有空间时需要立刻通知:
func Test_Cond_02(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)
	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("remove from queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		if len(queue) == 2 {
			// Wait的调用不仅仅是阻塞，它暂停当前的goroutine， 允许其他goroutine在操作系统线程上运行。
			// 进入Wait后， Cond的变量Locker将调用Unlock，并在退出Wait时，Cond变量的Locker上会调用Lock
			c.Wait()
		}
		fmt.Println("add to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(time.Second * 1)
		c.L.Unlock()
	}
}

// Broadcast 通知
func Test_Cond_03(t *testing.T) {
	type button struct {
		click *sync.Cond
	}

	subscribe := func(c *sync.Cond, fn func()) {
		var tempwg sync.WaitGroup
		tempwg.Add(1)
		go func() {
			tempwg.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		tempwg.Wait()
	}
	var wg sync.WaitGroup
	buttoninstance := button{click: sync.NewCond(&sync.Mutex{})}
	wg.Add(3)
	subscribe(buttoninstance.click, func() {
		fmt.Println("maximzing windows")
		wg.Done()
	})
	subscribe(buttoninstance.click, func() {
		fmt.Println("display annoying  dialog box")
		wg.Done()
	})
	subscribe(buttoninstance.click, func() {
		fmt.Println("miuse leave")
		wg.Done()
	})
	buttoninstance.click.Broadcast()
	wg.Wait()

}
