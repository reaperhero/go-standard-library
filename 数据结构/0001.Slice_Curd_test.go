package structure

import (
	"fmt"
	"testing"
)

//slice名为切片，是Go中的可变长数组，是对底层数组的封装和引用。切片指向一个底层数组，并且包含长度和容量信息。未初始化切片的值为 nil。
//作用于切片的内建函数主要有四个，分别是make、len、cap与append。
//make用于创建切片，len获取切片的长度，cap获取切片的容量，append向切片追加元素。

func Test_slice01(t *testing.T) {
	//创建切片，make([]T, length, capacity)
	fib := make([]int, 0, 10)
	fmt.Println("len(fib)=", len(fib))
	fmt.Println("cap(fib)=", cap(fib))
	fib = append(fib, []int{1, 1, 2, 3, 5, 8, 13}...)
	fmt.Println("fib=", fib)
}

// 插入元素
func insert(slice *[]interface{}, index int, value interface{}) {
	if index > len(*slice) {
		return
	}
	//尾部追加元素，使用append函数
	if index == len(*slice) {
		*slice = append(*slice, value)
		return
	}
	*slice = append((*slice)[:index+1], (*slice)[index:]...)
	(*slice)[index] = value
}

func Test_slice_insert01(t *testing.T) {
	fib := make([]interface{}, 0, 10)
	//切片头部插入元素
	insert(&fib, 0, 1)
	fmt.Println("fib =", fib)
}

// 删除元素
func slicedelete(slice *[]interface{}, index int) {
	if index > len(*slice)-1 {
		return
	}
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
}

func Test_slice_delete01(t *testing.T) {
	fib := []interface{}{1, 1, 2, 3, 5, 8}
	slicedelete(&fib, len(fib)-1)
	fmt.Println("fib =", fib)
}

// 修改元素
func update(slice *[]interface{}, index int, value interface{}) {
	if index > len(*slice)-1 {
		return
	}
	(*slice)[index] = value
}

func Test_slice_update01(t *testing.T) {
	fib := []interface{}{1, 1, 2, 3, 5, 8}
	update(&fib, len(fib)-1, 9)
	fmt.Println("fib =", fib)
}

// 查找元素下标
func search(slice []interface{}, value interface{}) []int {
	var index []int
	for i, v := range slice {
		if v == value {
			index = append(index, i)
		}
	}
	return index
}

func Test_slice_search01(t *testing.T) {
	fib := []interface{}{1, 1, 2, 3, 5, 8}
	indexSlice := search(fib, 1)
	fmt.Println("indexSlice =", indexSlice)
}
