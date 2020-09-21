package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

var (
	chanQuit chan bool
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	go handleSend(conn)
	go handleReceive(conn)
	<-chanQuit
}

func handleReceive(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != io.EOF {
		fmt.Println(err)
		os.Exit(1)
	}
	msg := string(buffer[:n])
	fmt.Println(msg)
}

func handleSend(conn net.Conn) {
	// 读取标准输入
	reader := bufio.NewReader(os.Stdin)
	lineBytes, _, _ := reader.ReadLine()

	// 发送到服务端
	_, err := conn.Write(lineBytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
