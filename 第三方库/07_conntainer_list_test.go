package third

import (
	"container/list"
	"fmt"
	"testing"
)

// https://cloud.tencent.com/developer/section/1140666


// 双向链接列表
func Test_container_03(t *testing.T) {
	// 创建一个新列表并在其中添加一些数字。
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// 遍历列表并打印其内容。
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}


