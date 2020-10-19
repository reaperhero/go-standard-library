package third

import (
	"container/ring"
	"fmt"
	"testing"
)

func Test_RingFunc(t *testing.T) {
	r := ring.New(10) //初始长度10
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	for i := 0; i < r.Len(); i++ {
		fmt.Println(r.Value)
		r = r.Next()
	}

	r = r.Move(6)
	fmt.Println(r.Value) //6
	r1 := r.Unlink(19)   //删除链表中n % r.Len()个元素，从r.Next()开始删除。如果n % r.Len() == 0，不修改r。返回删除的元素构成的链表，r不能为空。19 % 10 = 9
	for i := 0; i < r1.Len(); i++ {
		fmt.Println(r1.Value)
		//7
		//8
		//9
		//0
		//1
		//2
		//3
		//4
		//5
		r1 = r1.Next()
	}
	fmt.Println(r.Len())  //10-9=1
	fmt.Println(r1.Len()) //9
}