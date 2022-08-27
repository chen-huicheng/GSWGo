package main

import (
	"fmt"
	"sort"
)

type Number struct {
	A int
	B int
}

func (*Number) Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func (*Number) Add(a, b int) int {
	return a + b
}

type Inter struct {
	A int
}

func (*Inter) Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func (*Inter) Add(a, b int) int {
	return a + b
}

type Option interface {
	Abs(int) int
	Add(int, int) int
}

func Compute(opt ...Option) {
	fmt.Printf("%T", opt)
	sort.Slice(opt, func(i, j int) bool {
		_, isNumber := opt[i].(*Number)
		_, isNumber2 := opt[j].(*Number)
		return isNumber && !isNumber2
	})
	for _, n := range opt {
		fmt.Printf("%T\n", n)
	}
}

func main() {
	Compute(&Inter{1}, &Number{1, 2}, &Number{2, 3}, &Number{3, 4}, &Inter{5}, &Inter{6}, &Inter{7})
	arr := []int{1, 4, 3, 2, 8, 6}

	fmt.Printf("%T\n", arr)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	fmt.Println(arr)
}
