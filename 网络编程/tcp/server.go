package main

import (
	"net"
	"time"
)

//TCP Socket
//func (c *TCPConn) Write(b []byte) (int, error)
//func (c *TCPConn) Read(b []byte) (int, error)

//fmt.Println(net.ParseIP("192.168.1.10")) // ip不合法会打印nil



func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime)) // don't care about return value
		conn.Close()                // we're finished with this client
	}
}
