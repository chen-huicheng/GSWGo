package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("conn success", conn.RemoteAddr().String())
	buf := make([]byte, 1024*8)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read", err)
			return
		}
		if "exit" == string(buf[:n]) {
			fmt.Println("conn close", conn.RemoteAddr().String())
			return
		}
		fmt.Println("get : ", string(buf[:n]))
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	// listen
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("listen", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept", err)
			continue
		}
		//处理用户请求
		go HandleConn(conn)

	}
}
