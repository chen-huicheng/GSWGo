package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func RecvFile(fileName string, conn net.Conn) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Create", err)
		return
	}
	buf := make([]byte, 8*1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("recv success.")
			} else {
				fmt.Println(err)
			}
			return
		}
		_, err = f.Write(buf[:n])
		if err != nil {
			fmt.Println("Write", err)
			return
		}
	}
}
func HandleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("conn success", conn.RemoteAddr().String())
	buf := make([]byte, 1024*8)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Read", err)
		return
	}
	fileName := string(buf[:n])
	conn.Write([]byte("ok"))
	RecvFile(fileName, conn)
}
func main() {
	// listen
	listener, err := net.Listen("tcp", "127.0.0.1:12345")
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
