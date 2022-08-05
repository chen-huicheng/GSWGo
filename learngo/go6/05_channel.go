package main

import (
	"fmt"
	"time"
)

//  channel 主要用来同步
// goroutine 通过通信来共享内存，而不是通过共享内存来通信

var ch = make(chan int) //创建方式  make(chan Type)

// <-ch
func printer(data string) {
	for _, data := range data {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}
func person1() {
	printer("hello")
	ch <- 1
}
func person2() {
	<-ch
	printer("world")
}
func main() {

	go person2()
	go person1()

	for {

	}

}
