package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string //用于发送数据
	Name string      //用户名
	Addr string      //网络地址
}
type Message struct {
	data string
	user Client
}

var onlineMap map[string]Client

// onlineMap = make(map[string]Client)
var message = make(chan Message)

func Manager() { //广播
	onlineMap = make(map[string]Client)
	for {
		msg := <-message
		user := msg.user
		data := msg.data
		buf := "[" + user.Addr + "]" + user.Name + ": " + data
		for _, cli := range onlineMap {
			if cli.Addr == user.Addr {
				continue
			}
			cli.C <- buf
		}
	}
}
func MakeMsg(cli Client, msg string) {
	message <- Message{msg, cli}
}
func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func HandleConn(conn net.Conn) { //处理用户连接
	defer conn.Close()
	cliAddr := conn.RemoteAddr().String()
	fmt.Println("conn success", cliAddr)

	//添加到map
	cli := Client{make(chan string), cliAddr, cliAddr}
	onlineMap[cliAddr] = cli

	go WriteMsgToClient(cli, conn)

	// message <- "[" + cli.Addr + "]" + cli.Name + ": login"
	MakeMsg(cli, "login")
	isQuit := make(chan bool)
	hasData := make(chan bool)
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			hasData <- true
			if n == 0 {
				isQuit <- true
				fmt.Println("conn close:", err)
				break
			}
			msg := string(buf[:n])
			msg = strings.Trim(msg, " \n")
			if "who" == msg { //命令查询
				conn.Write([]byte("user list:\n"))
				for _, v := range onlineMap {
					data := v.Addr + ":" + v.Name + "\n"
					conn.Write([]byte(data))
				}
			} else if len(msg) >= 8 && "rename" == msg[:6] {
				cli.Name = strings.Split(msg, "|")[1]
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename ok\n"))
			} else { //转发消息
				MakeMsg(cli, msg)
			}

		}

	}()
	for {
		select {
		case <-isQuit:
			delete(onlineMap, cliAddr)
			MakeMsg(cli, "login out")
			return
		case <-hasData:

		case <-time.After(10 * time.Second):
			delete(onlineMap, cliAddr)
			MakeMsg(cli, "time out leave out")
			return
		}
	}

}

func main() {
	// listen
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("listen", err)
		return
	}
	defer listener.Close()

	//新开一个协程，用于消息转发
	go Manager()
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
