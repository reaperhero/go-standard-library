package stand

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Io_01() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		fmt.Println(err)
	}
}
func Io_02() {
	r := strings.NewReader("some io.Reader stream to be read")

	if _, err := io.CopyN(os.Stdout, r, 5); err != nil { // 拷贝从 src 到 dst 的 n 个字节
		fmt.Println(err)
	}

}

func Io_03() {
	io.WriteString(os.Stdout, "Hello World")

}
