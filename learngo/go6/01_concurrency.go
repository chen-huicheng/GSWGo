package main

import (
	"fmt"
	"time"
)

func newTask(name string) {
	for {
		fmt.Println("this is a ", name)
		time.Sleep(time.Second)
	}
}

// goroutine// go 的协程
func main() { // <== 主goroutine

	go newTask("newTask") //新建一个携程

	for {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second)
		break
	}
	go newTask("test") //新建一个携程 主协程直接退出 子协程可能不执行
	time.Sleep(time.Second * 5)
	// 如果主协程退出，子协程也要退出
}
