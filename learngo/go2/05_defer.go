package main

import "fmt"

func test(x int) {
	x = 100 / x
}
func main() {
	defer fmt.Println("first defer") //程序结束前调用 多个defer 前面的后执行
	fmt.Println("hello")

	// test(1)

	defer fmt.Println("second defer")

	a, b := 1, 2
	defer func() {
		fmt.Println("third defer")
		fmt.Printf("不传参 a=%d,b=%d\n", a, b)
	}()

	defer func(a, b int) {
		fmt.Printf("传参 a=%d,b=%d\n", a, b)
	}(a, b) //  a b 值已传递  但函数未执行

	a, b = 111, 222
	fmt.Printf("外部 a=%d,b=%d\n", a, b)

}
