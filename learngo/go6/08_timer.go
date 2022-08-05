package main

import (
	"fmt"
	"time"
)

func stopTimer() {
	timer := time.NewTimer(2 * time.Second)
	timer.Reset(10 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("child")
	}()

	// timer.Stop()
	for {
	}
}
func main() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("new :", time.Now())

	t := <-timer.C //没有数据会阻塞 C是一个只读的通道
	fmt.Println("延时", t)

	// for {
	// 	<-timer.C //多次调用无效 会阻塞
	// 	fmt.Println("time")
	// }

	ch := time.After(2 * time.Second) //after 返回一个延时写入的通道
	fmt.Println("after", time.Now())
	t = <-ch
	fmt.Println("延时", t)

	stopTimer()
}
