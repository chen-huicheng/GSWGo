package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("CCCCCC")
	runtime.Goexit() //结束该协程 main 中 BBBBBBB 不再打印
	// return // 只结束当前函数
	fmt.Println("DDDDDD")
}
func main() {

	go func() {
		fmt.Println("AAAAAA")
		test()
		fmt.Println("BBBBBB")
	}()
	for {

	}

}
