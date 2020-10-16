package data

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

//
func Test_channel_01(t *testing.T) {
	sum := func(a []int, c chan int) {
		total := 0
		for _, v := range a {
			total += v
		}
		c <- total // send total to c
	}

	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c) // -9 + 4 + 0 = -5
	go sum(a[len(a)/2:], c) // 7 + 2 + 8 = -17
	x, y := <-c, <-c
	fmt.Println(x, y, x+y) // -5
}

// Range和Close
func Test_channel_02(t *testing.T) {
	fibonacci := func(n int, c chan int) {
		x, y := 1, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

// Select
func Test_channel_03(t *testing.T) {
	fibonacci := func(c, quit chan int) {
		x, y := 1, 1
		for {
			select {
			case c <- x: // 写
				x, y = y, x+y
			case <-quit: // 读
				fmt.Println("quit")
				return
			}
		}
	}

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("读"+strconv.Itoa(<-c))
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}


// 超时
func Test_channel_04(t *testing.T) {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <- c: // 读阻塞
				println(v)
			case <- time.After(5 * time.Second): // 5秒都可以读取到然后break
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<- o
}

func Test_channel_05(t *testing.T) {

}
