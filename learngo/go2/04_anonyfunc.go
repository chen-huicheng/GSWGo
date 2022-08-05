package main

import "fmt"

func test() func() int {
	x := 0
	return func() int {
		x++
		fmt.Println("x = ", x)
		return x
	}
}

func main() {
	a := 10
	f1 := func() {
		fmt.Println("a = ", a)
	}
	f1()

	type fType func()
	var f2 fType
	f2 = f1
	f2()

	func() {
		fmt.Println("hello")
	}()

	b := 10
	str := "hello"
	fmt.Println(b, str)
	func() { //引用形式捕获变量
		b = 666
		str = "world"
		fmt.Println(b, str)
	}()
	fmt.Println(b, str)

	//闭包：说白了就是函数的嵌套，内层的函数可以使用外层函数的所有变量，即使外层函数已经执行完毕。
	f3 := test()
	f3()
	f3()
	f3()

}
