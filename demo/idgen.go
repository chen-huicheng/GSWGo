package main

import "fmt"

func IdGen() uint {
	v := int64(-1) ^ (int64(-1) << 4)
	return uint(v)
}
func printB(n int64) {
	for i := 0; i < 64; i++ {
		if n&(1<<(64-i)) != 0 {
			fmt.Printf("%d", 1)
		} else {
			fmt.Printf("%d", 0)
		}
	}
	fmt.Println()
}
func main() {
	fmt.Println(IdGen())
	fmt.Printf("%d\n", int64(-1))
	fmt.Printf("%b\n", int64(-21))
}
