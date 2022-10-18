package main

import (
	"fmt"
	"time"
)

func run() {
	arr := make([]int, 0)
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}
	// defer func() {
	// 	err := recover()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()
	for i := 0; i < len(arr); i++ {
		fmt.Println(i)
		i := i
		go func() {
			fmt.Println("idx", i)
		}()
	}
	// for i := 0; i < len(arr); i++ {
	// 	go func(i int) {
	// 		fmt.Println("idx:val", i, arr[i])
	// 	}(i)
	// }
	time.Sleep(time.Second * 5)
}

func main() {
	run()
}
