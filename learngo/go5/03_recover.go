package main

import (
	"fmt"
)

type Exception struct {
	Code   int32
	Msg    string
	Except bool
}

func test1(x int) {

	defer func() { //正常执行
		//recover() //当发生 panic 错误时 恢复错误
		if panicData := recover(); panicData != nil {
			msg := panicData
			if e, ok := panicData.(Exception); ok {
				msg = e.Msg
			}
			fmt.Println(msg)
		}
	}()
	// var a [10]int
	// a[x] = 11
	fmt.Println("run successful")
	// fmt.Println(4 / x)
	exp := Exception{Code: 1, Msg: "a panic"}
	panic(exp)
}
func main() {
	// test()
	test1(0)
	//panic
	//"runtime error: index out of range [15] with length 10"

}
