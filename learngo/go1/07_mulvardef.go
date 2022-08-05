package main

import "fmt"

func main() {
	// var a int
	// var b float32
	// var (
	// 	a int     = 10
	// 	b float32 = 2
	// )
	var (
		a = 10
		b = 3.14
	)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	// const (
	// 	i int     = 10
	// 	j float32 = 3.14
	// )
	const (
		i = 10
		j = 3.14
	)
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
}
