package main

import "fmt"

func main() {
	a := 0
	var pi *int
	pi = &a
	fmt.Println(*pi)
	// fmt.Scan(&a)
	// fmt.Println(a)
}
