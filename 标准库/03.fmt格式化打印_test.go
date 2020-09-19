package stand

import (
	. "fmt"
	"testing"
)

type Website struct {
	Name string
}

var site = Website{Name: "studygolang"}

func Test_print01(t *testing.T) {
	//普通占位符
	Printf("%v", site)  // 相应值的默认格式。 {studygolang}
	Printf("%+v", site) // 在打印结构体时，“加号”标记（%+v）会添加字段名 {Name:studygolang}
	Printf("%#v", site) // 相应值的Go语法表示         main.Website{Name:"studygolang"}
	Printf("%T", site)  // 相应值的类型的Go语法表示 main.Website

	//布尔占位符
	Printf("%t", true) //  true

	// 整数占位符
	Printf("%b", 5)      // 二进制表示  101
	Printf("%c", 0x4E2D) // 相应Unicode码点所表示的字符  中
	Printf("%d", 0x12)   // 十进制表示   18
	Printf("%d", 10)     // 八进制表示   12
	Printf("%q", 0x4E2D) // 单引号围绕的字符字面值，由Go语法安全地转义  '中'
	Printf("%x", 13)     //  十六进制表示，字母形式为小写 a-f       d
	Printf("%x", 13)     // 十六进制表示，字母形式为大写 A-F         D
	Printf("%U", 0x4E2D) // Unicode格式：U+1234，等同于 "U+%04X"        U+4E2D

	// 字符串与字节切片
	Printf("%s", []byte("Go语言中文网")) // 输出字符串表示（string类型或[]byte) Go语言中文网
	Printf("%q", "Go语言中文网")         //  双引号围绕的字符串，由Go语法安全地转义   "Go语言中文网"
	Printf("%x", "golang")          // 十六进制，小写字母，每字节两个字符   676f6c616e67
	Printf("%X", "golang")          // 十六进制，大写字母，每字节两个字符    676F6C616E67

	//指针
	Printf("%p", &site)  //十六进制表示，前缀 0x  0x4f57f0
	Printf("%#p", &site) //十六进制表示，不带0x的指针  4f57f0

}
