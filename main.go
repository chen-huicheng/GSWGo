package main

import (
	"fmt"
	"math/rand"

	"github.com/chen-huicheng/GSWGo/stl"
)

func main() {
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Int() % 100
	}
	fmt.Println(arr)
	stl.MySort(arr)
	fmt.Println(arr)
	fmt.Println("hello GSWGO")
	fmt.Println("fix bug")
}
