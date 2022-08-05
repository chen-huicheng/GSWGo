package main

import (
	"fmt"
	"time"
)

func test() {
	ch := make(chan string, 10)

}

func main() {
	ch := make(chan string, 0) //无缓存  多次写入也会阻塞
	defer fmt.Println("main done")
	go func() {
		// defer fmt.Println("child done") //可能执行不到
		for i := 0; i < 2; i++ {
			fmt.Println("i = ", i)
			time.Sleep(time.Second)
		}
		fmt.Println("child done")
		ch <- "done"
	}()
	str := <-ch
	fmt.Println(str)
}
