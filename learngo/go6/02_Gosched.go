package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	}()

	for i := 0; i < 2; i++ {
		//让出时间片，先让其他的执行   注释看差异
		runtime.Gosched()
		//如果不让出时间片  主线程可能直接执行结束，导致子线程不执行
		fmt.Println("hello")
	}
}
