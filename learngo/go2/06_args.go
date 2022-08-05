package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args

	n := len(list)
	fmt.Println(n)

	for i, x := range list {
		fmt.Println(i, x)
	}
}

// using
/*
	go build -o output file1.go file2.go
	./output
*/
