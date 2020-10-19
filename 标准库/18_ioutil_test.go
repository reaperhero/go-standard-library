package stand

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func IoUtil_01() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", b)

}

// 返回按文件名排序的目录条目列表
func IoUtil_02() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}


// 读文件
func IoUtil_03() {
	content, err := ioutil.ReadFile("testdata/hello")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("File contents: %s", content)

}
