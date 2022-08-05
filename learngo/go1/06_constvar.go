package main

import "fmt"

func main() {
	const a int = 10
	fmt.Println("a = ", a)
	const b = a //不需要使用   :=
	fmt.Println("b = ", b)
}
