package main

import (
	"fmt"
	"time"
)

func fib(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		}
	}
}

//select 类似 io 复用
func main() {
	ch := make(chan int, 5)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch
			fmt.Println(num)
		}
		quit <- true
	}()
	fib(ch, quit)

	overtime()

}

func overtime() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println(num)
			case <-time.After(3 * time.Second):
				// time.Sleep()
				// t = time.NewTimer(time.Second)//新建一个定时器
				// <-t.C
				fmt.Println("overtime")
				quit <- true
				return
			}
		}
	}()
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	<-quit
	fmt.Println("over")
}
