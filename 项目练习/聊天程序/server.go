package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

var (
	clientsMap = make(map[string]net.Conn)
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	defer listen.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		for _, conn := range clientsMap {
			conn.Write([]byte("all:服务器进入维护状态"))
		}
		listen.Close()
	}()
	for {
		conn, err := listen.Accept()
		clientsMap[conn.RemoteAddr().String()] = conn
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		go ioWithConn(conn)
	}
}

func ioWithConn(conn net.Conn) {
	clientAddr := conn.RemoteAddr().String()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != io.EOF {
		fmt.Println(err)
		return
	}
	if n > 0 {
		msg := string(buffer[:n])
		fmt.Printf("%s:%s\n", clientAddr, msg)
		strs := strings.Split(msg, "#")
		if len(strs) > 1 {
			targetAddr := strs[0]
			targetMsg := strs[1]
			if targetAddr == "all" {
				// 群发
				for _, conn := range clientsMap {
					conn.Write([]byte(clientAddr + ":" + targetMsg))
				}
			} else {
				// 点对点
				for addr, conn := range clientsMap {
					if addr == targetAddr {
						conn.Write([]byte(clientAddr + ":" + targetMsg))
						break
					}
				}
			}
		} else {
			conn.Write([]byte(clientAddr + ":" ))
		}
	}
}
