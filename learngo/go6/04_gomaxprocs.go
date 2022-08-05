package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(1) //单核打印  1111111111..111000000000..001111... 轮到谁使用时间片谁打印
	// n := runtime.GOMAXPROCS(2) //双核打印  不用等直接打印

	fmt.Println("n = ", n)

	for {
		go fmt.Print(1)
		fmt.Print(0)
	}
}
