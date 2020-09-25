package command

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 子协程死了相当于break
func Test_FunExit_01(t *testing.T) {
	go func() {
		for i := 0; i < 10; i++ {
			if i == 5 {  // 相当于break
				runtime.Goexit()
			}
			fmt.Printf("routing1:%d\n", i)
			time.Sleep(time.Second * 1)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Printf("routing2:%d\n", i)
		time.Sleep(time.Second * 1)
	}
	fmt.Println("main over")
}

// 主协程死了，那么子协程将不受控制
func Test_FunExit_02(t *testing.T) {
	go func() {
		for i := 0; i < 10; i++ {

			fmt.Printf("routing1:%d\n", i)
			time.Sleep(time.Second * 1)
		}
	}()

	for i := 0; i < 10; i++ {
		if i == 5 {  // 相当于break
			runtime.Goexit()
		}
		fmt.Printf("routing2:%d\n", i)
		time.Sleep(time.Second * 1)
	}
	fmt.Println("main over")
}