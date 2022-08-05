package main

import (
	"fmt"
	"net"
)

func HandleConn(conn net.Conn) {
	k := 8
	buf := make([]byte, 1024*k)
	defer conn.Close()
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%s\n", buf[:n])
		fmt.Printf("\n\nend\n\n")
		n, err = conn.Write([]byte("hello world!!"))
		if err != nil {
			break
		}
	}
}
func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go HandleConn(conn)
	}
}
