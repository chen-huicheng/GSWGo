package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("dial", err)
		return
	}
	defer conn.Close()

	requestBuf := "GET /go HTTP/1.1\r\nHost: 127.0.0.1:8888\r\nConnection: keep-alive\r\n\r\n"

	conn.Write([]byte(requestBuf))
	buf := make([]byte, 1024*8)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf[:n]))

}
