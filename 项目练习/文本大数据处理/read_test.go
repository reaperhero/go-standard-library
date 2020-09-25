package chank

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

//一次性将全部数据载入内存（不可取）
func Test_Encode_01(t *testing.T) {
	//一次性将全部数据载入内存（不可取）
	contentBytes, err := ioutil.ReadFile(`./utils.go`)
	if err != nil {
		fmt.Println("读入失败", err)
	}
	contentStr := string(contentBytes)

	//以换行符为定界符，将大文本炸碎为行的切片
	lineStrs := strings.Split(contentStr, "\n\r")

	//逐行打印
	for _, lineStr := range lineStrs {
		fmt.Println(lineStr)
	}
}



//缓冲式读取
func Test_Encode_02(t *testing.T) {

	//打开文件
	file, _ := os.Open(`./utils.go`)
	defer file.Close()

	//创建文件的缓冲读取器
	reader := bufio.NewReader(file)

	for {

		//逐行读取
		lineBytes, _, err := reader.ReadLine()

		//读到文件末尾时退出循环
		if err == io.EOF {
			break
		}

		fmt.Println(string(lineBytes))
	}

}