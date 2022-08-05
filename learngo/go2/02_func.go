package main

import "fmt"

func MaxAndMin(a, b int) (max, min int) {
	if a > b {
		max, min = a, b
	} else {
		max, min = b, a
	}
	return
}

// 斐波那契
func fib(n int) (ret int) {
	if n <= 2 {
		return 1
	}
	ret = fib(n-1) + fib(n-2)
	// fmt.Println(ret)
	return
}

//实现累加
func sum(n int) (s int) {
	if n == 1 {
		s = 1
		return
	}
	s = n + sum(n-1)
	return
}
func main() {
	max, min := MaxAndMin(15, 213)
	fmt.Printf("max : %d ,min : %d\n", max, min)

	fibn := fib(15)
	fmt.Printf("fibn : %d\n", fibn)

	s := sum(100)
	fmt.Printf("sumn : %d\n", s)
}
