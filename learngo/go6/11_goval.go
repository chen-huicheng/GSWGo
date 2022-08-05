package main

import (
	"fmt"
	"runtime"
	"time"
)

var a = 10

func run() {
	fmt.Println("run:", a)
	a = 11
}
func main() {
	go run()
	runtime.Gosched()
	time.Sleep(time.Second)
	fmt.Println("main:", a)
}
