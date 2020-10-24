package syncmutex

import (
	"fmt"
	"testing"
)

// 生产消费
func Test_channel_01(t *testing.T) {
	// 生产者
	chanOwner := func() <-chan int {
		result := make(chan int, 5)
		go func() {
			defer close(result)
			for i := 0; i < 6; i++ {
				result <- i
			}
		}()
		return result
	}
	// 消费者
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Println(result)
		}
		fmt.Println("Done receiving")
	}

	results := chanOwner()
	consumer(results)
}

// 管道
func Test_channel_02(t *testing.T) {
	// 乘法
	multiply := func(values []int, mutiplier int) []int {
		mutiplierdValue := make([]int, len(values))
		for k, v := range values {
			mutiplierdValue[k] = v * mutiplier
		}
		return mutiplierdValue
	}
	// 加法
	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for k, v := range values {
			addedValues[k] = v + additive
		}
		return addedValues
	}
	// 合并
	ints := []int{1, 2, 3, 4}
	for _, v := range add(multiply(ints, 2), 1) {
		fmt.Println(v)
	}
}

// 构建管道的最佳实践
func Test_channel_03(t *testing.T) {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	multiply := func(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
		multiplieredStream := make(chan int)
		go func() {
			defer close(multiplieredStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multiplieredStream <- i * multiplier:
				}
			}
		}()
		return multiplieredStream
	}

	add := func(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case addedStream <- i + additive:
				}
			}
		}()
		return addedStream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeple := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for i := range pipeple {
		fmt.Println(i)
	}
}

// 便利的生成器
func Test_channel_04(t *testing.T) {
	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{}) // 无缓存区
		go func() {
			defer close(valueStream)
			for { // 持续发送
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
						fmt.Println("send to valueStream", v)
					}
				}
			}
		}()
		return valueStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
					fmt.Println("receive from valueStream")
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done) // 往done里面发消息，repeat take 均返回
	for num := range take(done, repeat(done, 1), 10) {
		fmt.Println(num)
	}
}

// 扇入扇出
// 多个函数可以同时从一个channel接收数据，直到channel关闭，这种情况被称作扇出
// 一个函数同时接收并处理多个channel输入并转化为一个输出channel，直到所有的输入channel都关闭后，关闭输出channel，这种情况就被称作扇入。


// or-done-channel
// 需要用select语句来封装我 们的读取操作和done通道


// tee-channel
// 分割来自通道的多个值，以便将它们发送到两个独立区域



// bridge-channel
// 自己想要使用一系列通道的值