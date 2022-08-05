package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Dial", err)
		return
	}
	defer conn.Close()
	var msg string
	buf := make([]byte, 8*1024)
	go func() {
		for {
			fmt.Scanf("%s\n", &msg)
			// fmt.Println("send:", msg)
			conn.Write([]byte(msg))
		}
	}()
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn close")
			break
		}
		fmt.Println("get : ", string(buf[:n]))
	}

}
