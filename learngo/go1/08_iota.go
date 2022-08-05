package main

import "fmt"

func main() {
	//iota 常量自动生成器，每隔一行自动累加1
	// iota给常量赋值使用
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a = %d b = %d c = %d\n", a, b, c)
	// iota 遇到 const 重置为0
	const d = iota
	fmt.Printf("d = %d\n", d)

	const (
		a1         = iota
		b1, b2, b3 = iota, iota, iota
		c1         = iota
	)
	fmt.Println(a1, b1, b2, b3, c1)
}
