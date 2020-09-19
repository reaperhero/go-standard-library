package stand

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"testing"
)

func Test_File01(t *testing.T) {
	/*
		文件操作：
		1.路径：
			相对路径：relative
				ab.txt
				相对于当前工程
			绝对路径：absolute
				/Users/ruby/Documents/pro/a/aa.txt

			.当前目录
			..上一层
		2.创建文件夹，如果文件夹存在，创建失败
			os.MkDir()，创建一层
			os.MkDirAll()，可以创建多层

		3.创建文件，Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
			os.Create()，创建文件

		4.打开文件：让当前的程序，和指定的文件之间建立一个连接
			os.Open(filename)
			os.OpenFile(filename,mode,perm)

		5.关闭文件：程序和文件之间的链接断开。
			file.Close()

		5.删除文件或目录：慎用，慎用，再慎用
			os.Remove()，删除文件和空目录
			os.RemoveAll()，删除所有
	*/
	//1.路径
	fileName1 := "/tmp/aa.txt"
	fileName2 := "bb.txt"
	fmt.Println(filepath.IsAbs(fileName1)) //true
	fmt.Println(filepath.IsAbs(fileName2)) //false
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2))

	fmt.Println("获取父目录：", path.Join(fileName1, ".."))

	//2.创建目录
	//err := os.Mkdir("/Users/ruby/Documents/pro/a/bb",os.ModePerm)
	//if err != nil{
	//	fmt.Println("err:",err)
	//	return
	//}
	//fmt.Println("文件夹创建成功。。")
	//err :=os.MkdirAll("/Users/ruby/Documents/pro/a/cc/dd/ee",os.ModePerm)
	//if err != nil{
	//	fmt.Println("err:",err)
	//	return
	//}
	//fmt.Println("多层文件夹创建成功")

	//3.创建文件:Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
	//file1,err :=os.Create("/Users/ruby/Documents/pro/a/ab.txt")
	//if err != nil{
	//	fmt.Println("err：",err)
	//	return
	//}
	//fmt.Println(file1)

	//file2,err := os.Create(fileName2)//创建相对路径的文件，是以当前工程为参照的
	//if err != nil{
	//	fmt.Println("err :",err)
	//	return
	//}
	//fmt.Println(file2)

	//4.打开文件：
	//file3 ,err := os.Open(fileName1) //只读的
	//if err != nil{
	//	fmt.Println("err:",err)
	//	return
	//}
	//fmt.Println(file3)
	/*
		第一个参数：文件名称
		第二个参数：文件的打开方式
			const (
		// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
			O_RDONLY int = syscall.O_RDONLY // open the file read-only.
			O_WRONLY int = syscall.O_WRONLY // open the file write-only.
			O_RDWR   int = syscall.O_RDWR   // open the file read-write.
			// The remaining values may be or'ed in to control behavior.
			O_APPEND int = syscall.O_APPEND // append data to the file when writing.
			O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
			O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
			O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
			O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
		)
		第三个参数：文件的权限：文件不存在创建文件，需要指定权限
	*/
	//file4,err := os.OpenFile(fileName1,os.O_RDONLY|os.O_WRONLY,os.ModePerm)
	//if err != nil{
	//	fmt.Println("err:",err)
	//	return
	//}
	//fmt.Println(file4)

	//5关闭文件，
	//file4.Close()

	//6.删除文件或文件夹：
	//删除文件
	//err :=  os.Remove("/Users/ruby/Documents/pro/a/aa.txt")
	//if err != nil{
	//	fmt.Println("err:",err)
	//	return
	//}
	//fmt.Println("删除文件成功。。")
	//删除目录
	err := os.RemoveAll("/Users/ruby/Documents/pro/a/cc")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("删除目录成功。。")
}

func Test_File02(t *testing.T) {
	/*
		FileInfo：文件信息
			interface
				Name()，文件名
				Size()，文件大小，字节为单位
				IsDir()，是否是目录
				ModTime()，修改时间
				Mode()，权限
	*/
	fileInfo, err := os.Stat("/Users/chenqiangjun/gocode/src/go-standard-library/go.mod")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	fmt.Printf("%T\n", fileInfo)
	//文件名
	fmt.Println(fileInfo.Name())
	//文件大小
	fmt.Println(fileInfo.Size())
	//是否是目录
	fmt.Println(fileInfo.IsDir()) //IsDirectory
	//修改时间
	fmt.Println(fileInfo.ModTime())
	//权限
	fmt.Println(fileInfo.Mode()) //-rw-r--r--
}

func Test_readFile(t *testing.T) {
	/*
		读取数据：
			Reader接口：
				Read(p []byte)(n int, error)
	*/
	//读取本地aa.txt文件中的数据
	//step1：打开文件
	fileName := "/Users/chenqiangjun/gocode/src/go-standard-library/go.mod"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//step3：关闭文件
	defer file.Close()

	//step2：读取数据
	bs := make([]byte, 4, 4)
	/*
		 //第一次读取
		 n,err :=file.Read(bs)
		 fmt.Println(err) //<nil>
		 fmt.Println(n) //4
		 fmt.Println(bs) //[97 98 99 100]
		fmt.Println(string(bs)) //abcd

		//第二次读取
		n,err = file.Read(bs)
		fmt.Println(err)//<nil>
		fmt.Println(n)//4
		fmt.Println(bs) //[101 102 103 104]
		fmt.Println(string(bs)) //efgh

		//第三次读取
		n,err = file.Read(bs)
		fmt.Println(err) //<nil>
		fmt.Println(n) //2
		fmt.Println(bs) //[105 106 103 104]
		fmt.Println(string(bs)) //ijgh

		//第四次读取
		n,err = file.Read(bs)
		fmt.Println(err) //EOF
		fmt.Println(n) //0
	*/
	n := -1
	for {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("读取到了文件的末尾，结束读取操作。。")
			break
		}
		fmt.Println(string(bs[:n]))
	}
}
