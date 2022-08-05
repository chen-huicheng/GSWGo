package main

import "fmt"

func main() {
	a, b := 10, 20 // 交换 两个变量

	fmt.Printf("a = %d, b = %d", a, b)
	a, b = b, a
	fmt.Printf("a = %d, b = %d", a, b)

	//匿名变量
	i, j := 1, 2
	tmp, _ := i, j
	fmt.Printf("tmp = %d", tmp)
}
