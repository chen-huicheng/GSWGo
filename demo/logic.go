package main

import "fmt"

func main() {
	for k := 0; k < 5; k++ {
		if k != 1 && k != 2 {
			fmt.Println(k)
		}
	}
}
