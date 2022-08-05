package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(path string, conn net.Conn) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	buf := make([]byte, 1024*8)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("send success.")
			} else {
				fmt.Println(err)
			}
			return
		}
		//发送内容
		n, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

func main() {
	fmt.Print("please input filename:")
	var fileName string
	fmt.Scan(&fileName)

	info, err := os.Stat(fileName) //获取文件属性
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.Dial("tcp", "127.0.0.1:12345") //连接服务器
	if err != nil {
		fmt.Println("Dial", err)
		return
	}
	defer conn.Close()
	_, err = conn.Write([]byte(info.Name())) //发送文件名
	if err != nil {
		fmt.Println("Dial", err)
		return
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf) //接受回复
	if err != nil {
		fmt.Println("Dial", err)
		return
	}
	if "ok" == string(buf[:n]) { //如果OK 发送文件
		SendFile(fileName, conn)
	}
}
